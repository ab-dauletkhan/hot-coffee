package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/ab-dauletkhan/hot-coffee/internal/core"
	"github.com/ab-dauletkhan/hot-coffee/models"
)

type JSONStorage struct {
	filePath string
	mu       sync.RWMutex
}

// NewJSONStorage initializes the directory and file structure as needed
func NewJSONStorage(filePath string) (*JSONStorage, error) {
	storage := &JSONStorage{
		filePath: filePath,
	}

	// Initialize directory and file, validating contents if necessary
	if err := storage.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize storage: %w", err)
	}

	return storage, nil
}

// Init initializes the storage directory and file, creating them if necessary and validating content
func (s *JSONStorage) Init() error {
	// Ensure the directory exists
	dir := filepath.Dir(s.filePath)
	if err := os.MkdirAll(dir, core.DirPerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Check and initialize the JSON file if it doesn't exist or is invalid
	if err := s.initFileIfNeeded(); err != nil {
		return fmt.Errorf("failed to initialize or validate file: %w", err)
	}
	return nil
}

// initFileIfNeeded checks if a file exists and initializes it if missing or invalid
func (s *JSONStorage) initFileIfNeeded() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	fileInfo, err := os.Stat(s.filePath)
	if os.IsNotExist(err) || (err == nil && fileInfo.Size() == 0) {
		// File doesn't exist or is empty, so initialize it
		return s.writeEmptyJSONArray()
	} else if err != nil {
		return err
	}

	// File exists; validate JSON content
	content, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}
	if !s.isValidJSON(content) {
		// Reinitialize if content is invalid
		return s.writeEmptyJSONArray()
	}

	return nil
}

// writeEmptyJSONArray initializes the JSON file with an empty array ("[]")
func (s *JSONStorage) writeEmptyJSONArray() error {
	return os.WriteFile(s.filePath, []byte("[]"), core.FilePerm)
}

// isValidJSON checks if the content is valid JSON and matches the expected structure for the file type
func (s *JSONStorage) isValidJSON(content []byte) bool {
	// Determine the expected structure based on the filename
	expected := s.fileToStruct()
	if expected == nil {
		return false
	}
	return json.Unmarshal(content, expected) == nil
}

func (s *JSONStorage) fileToStruct() interface{} {
	switch filepath.Base(s.filePath) {
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

// Retrieve reads data from the JSON file and returns it as an interface
func (s *JSONStorage) Retrieve(v interface{}) error {
	if err := s.Init(); err != nil {
		return err
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	content, err := os.ReadFile(s.filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	return json.Unmarshal(content, v)
}

// Save writes data from the provided interface to the JSON file
func (s *JSONStorage) Save(v interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tempFile := s.filePath + ".tmp"
	file, err := os.Create(tempFile)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(v); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	if err := file.Sync(); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to sync file: %w", err)
	}

	if err := file.Close(); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to close file: %w", err)
	}

	if err := os.Rename(tempFile, s.filePath); err != nil {
		os.Remove(tempFile)
		return fmt.Errorf("failed to save file: %w", err)
	}

	return nil
}

// exists checks if the file exists
func (s *JSONStorage) exists() bool {
	_, err := os.Stat(s.filePath)
	return !os.IsNotExist(err)
}

// Clear removes the storage file; useful for testing
func (s *JSONStorage) Clear() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.exists() {
		return os.Remove(s.filePath)
	}
	return nil
}
