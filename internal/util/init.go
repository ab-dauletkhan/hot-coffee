package util

import (
	"os"
	"path/filepath"

	"github.com/ab-dauletkhan/hot-coffee/internal/core"
)

func InitDir() error {
	err := os.MkdirAll(core.Dir, core.DirPerm)
	if err != nil {
		return err
	}
	err = InitFiles()
	if err != nil {
		return err
	}

	return nil
}

func InitFiles() error {
	err := os.WriteFile(filepath.Join(core.Dir, core.MenuFile), []byte("[]"), core.FilePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(core.Dir, core.InventoryFile), []byte("[]"), core.FilePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(core.Dir, core.OrderFile), []byte("[]"), core.FilePerm)
	if err != nil {
		return err
	}

	return nil
}
