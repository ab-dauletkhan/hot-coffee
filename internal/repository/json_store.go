package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/ab-dauletkhan/hot-coffee/internal/core"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

var ErrStorageOperation = errors.New("storage operation failed")

// JSONStorage represents a thread-safe JSON file storage
type JSONStorage struct {
	filePath string
	mu       sync.RWMutex
	schema   interface{} // cached schema for validation
}

// NewJSONStorage creates and initializes a new JSONStorage instance
func NewJSONStorage(filePath string) (*JSONStorage, error) {
	storage := &JSONStorage{
		filePath: filePath,
		schema:   determineSchema(filepath.Base(filePath)),
	}

	if storage.schema == nil {
		return nil, fmt.Errorf("unsupported file type: %s", filePath)
	}

	if err := storage.init(); err != nil {
		return nil, fmt.Errorf("storage initialization failed: %w", err)
	}

	return storage, nil
}

// init initializes the storage directory and file
func (s *JSONStorage) init() error {
	dir := filepath.Dir(s.filePath)
	if err := os.MkdirAll(dir, core.DirPerm); err != nil {
		return fmt.Errorf("directory creation failed: %w", err)
	}

	return s.initializeFile()
}

// initializeFile ensures the file exists and contains valid JSON
func (s *JSONStorage) initializeFile() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if exists, err := s.fileExists(); err != nil {
		return fmt.Errorf("file check failed: %w", err)
	} else if !exists {
		return s.writeEmptyJSONArray()
	}

	if err := s.validateContent(); err != nil {
		return fmt.Errorf("content validation failed: %w", err)
	}

	return nil
}

// validateContent checks if the file content is valid JSON
func (s *JSONStorage) validateContent() error {
	content, err := os.ReadFile(s.filePath)
	if err != nil {
		return fmt.Errorf("file read failed: %w", err)
	}

	if len(content) == 0 {
		return s.writeEmptyJSONArray()
	}

	// Create a new instance of the schema for validation
	validation := determineSchema(filepath.Base(s.filePath))
	if err := json.Unmarshal(content, validation); err != nil {
		return fmt.Errorf("invalid JSON content: %w", err)
	}

	return nil
}

// determineSchema returns the appropriate schema based on filename
func determineSchema(filename string) interface{} {
	switch filename {
	case core.MenuFile:
		return &[]models.MenuItem{}
	case core.InventoryFile:
		return &[]models.InventoryItem{}
	case core.OrderFile:
		return &[]models.Order{}
	default:
		return nil
	}
}

// Retrieve reads and unmarshals data from the JSON file
func (s *JSONStorage) Retrieve(v interface{}) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	content, err := os.ReadFile(s.filePath)
	if err != nil {
		return fmt.Errorf("file read failed: %w", err)
	}

	if err := json.Unmarshal(content, v); err != nil {
		return fmt.Errorf("JSON unmarshal failed: %w", err)
	}

	return nil
}

// Save atomically writes data to the JSON file
func (s *JSONStorage) Save(v interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.atomicWrite(v)
}

// atomicWrite performs atomic write operation using a temporary file
func (s *JSONStorage) atomicWrite(v interface{}) error {
	tempFile := s.filePath + ".tmp"

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("JSON marshal failed: %w", err)
	}

	if err := os.WriteFile(tempFile, data, core.FilePerm); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("temporary file write failed: %w", err)
	}

	if err := os.Rename(tempFile, s.filePath); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("atomic rename failed: %w", err)
	}

	return nil
}

// fileExists checks if the file exists
func (s *JSONStorage) fileExists() (bool, error) {
	_, err := os.Stat(s.filePath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Clear removes the storage file (useful for testing)
func (s *JSONStorage) Clear() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if exists, _ := s.fileExists(); exists {
		return os.Remove(s.filePath)
	}
	return nil
}

func (s *JSONStorage) writeEmptyJSONArray() error {
	return os.WriteFile(s.filePath, []byte("[]"), core.FilePerm)
}
