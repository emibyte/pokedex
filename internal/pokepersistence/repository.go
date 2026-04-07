package pokepersistence

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type Repository interface {
	Load() (state []byte, err error)
	Save(state []byte) error
}

type FileRepository struct {
	configPath string
}

func (f *FileRepository) Load() (state []byte, err error) {
	content, err := os.ReadFile(f.configPath)
	if err != nil && os.IsNotExist(err) {
		return []byte{}, nil
	} else if err != nil {
		return nil, err
	}
	return content, nil
}

func (f *FileRepository) Save(state []byte) error {
	// NOTE: creates file if there isn't one, overwrites file if it exists and is write only
	file, err := os.OpenFile(f.configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	bytesWritten, err := file.Write(state)
	if err != nil {
		return err
	}
	if bytesWritten != len(state) {
		return errors.New("The entire state of the program couldn't be written to file")
	}
	return nil
}

func getDefaultConfigPath() (string, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	path = fmt.Sprintf("%s/pokedex.json", path)
	return path, nil
}

func isFile(path *string) (bool, error) {
	if path == nil {
		return false, errors.New("path is nil")
	}
	infos, err := os.Stat(*path)
	if err != nil {
		return false, err
	}
	return !infos.IsDir(), nil
}

func InitFileRepository(cfgPath *string) (Repository, error) {
	f := FileRepository{}
	pathIsFile, err := isFile(cfgPath)
	if err != nil && cfgPath != nil {
		return &f, err
	}
	if cfgPath != nil && fs.ValidPath(*cfgPath) && pathIsFile {
		f.configPath = *cfgPath
		return &f, nil
	}
	defaultPath, err := getDefaultConfigPath()
	if err != nil {
		return &f, err
	}
	f.configPath = defaultPath
	return &f, nil
}
