package service

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func NewClassLoader(path string) *DirectoryLoader {
	loader := &DirectoryLoader{
		Path: path,
		Map:  make(map[string][]byte),
	}

	files, err := os.ReadDir(loader.Path)
	if err != nil {
		return nil
	}

	for _, file := range files {
		if !file.IsDir() {
			contents, err := os.ReadFile(filepath.Join(loader.Path, file.Name()))
			if err != nil {
				fmt.Println(err)
				continue
			}
			loader.Map[file.Name()] = contents
		}
	}

	return loader
}

type DirectoryLoader struct {
	Path string
	Map  map[string][]byte
}

func (l *DirectoryLoader) Load(internalName string) ([]byte, error) {
	originalName := internalName
	if !strings.HasSuffix(internalName, ".class") {
		originalName = internalName + ".class"
	}

	if contents, ok := l.Map[originalName]; ok {
		return contents, nil
	}

	contents, err := os.ReadFile(filepath.Join(l.Path, originalName))
	if err != nil {
		return nil, err
	}

	return contents, nil
}

func (l *DirectoryLoader) CanLoad(internalName string) bool {
	internalName = filepath.Join(l.Path, internalName+".class")
	if _, err := os.Stat(internalName); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
