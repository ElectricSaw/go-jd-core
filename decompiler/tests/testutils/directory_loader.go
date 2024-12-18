package testutils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func NewDirectoryLoader(path string) *DirectoryLoader {
	loader := &DirectoryLoader{
		Path: path,
		Map:  make(map[string][]byte),
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
	}
	loader.Map = findFiles(absPath, path, loader.Map)

	return loader
}

func findFiles(root, path string, m map[string][]byte) map[string][]byte {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	for _, file := range files {
		if !file.IsDir() {
			contents, err := os.ReadFile(filepath.Join(path, file.Name()))
			if err != nil {
				fmt.Println(err)
				continue
			}

			absPath, err := filepath.Abs(filepath.Join(path, file.Name()))
			if err != nil {
				fmt.Println(err)
				continue
			}

			relPath, err := filepath.Rel(root, absPath)
			if err != nil {
				fmt.Println(err)
				continue
			}

			relPath = filepath.ToSlash(relPath)
			m[relPath] = contents
		} else {
			m = findFiles(root, filepath.Join(path, file.Name()), m)
		}
	}

	return m
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
