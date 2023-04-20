package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func getSelf() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return GetFileName(ex)
}

// GetFiles 获取目录所有文件
func GetFiles(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && !strings.HasSuffix(path, ".fs256") && path != getSelf() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// GetHashFiles 获取目录所有.fs256文件
func GetHashFiles(path string) ([]string, error) {
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".fs256") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// GetFileSha256 获取文件sha256
func GetFileSha256(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}

// GetFileName 获取文件名
func GetFileName(path string) string {
	return filepath.Base(path)
}

// CreateFile 创建文件
func CreateFile(path string) error {
	_, err := os.Create(path)
	return err
}

// DeleteHashFiles 删除所有.fs256文件
func DeleteHashFiles(path string) error {
	files, err := GetHashFiles(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		err := os.Remove(file)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateHashFile 创建.fs256文件
func CreateHashFile(path string, hash string) error {
	filename := GetFileName(path)
	err := CreateFile(filename + ".fs256")
	if err != nil {
		return err
	}
	//写入sha256值
	file, err := os.OpenFile(filename+".fs256", os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(hash)
	return err
}

// GetFileNameWithoutHash 获取去除.fs256后缀的文件名
func GetFileNameWithoutHash(path string) string {
	return strings.TrimSuffix(GetFileName(path), ".fs256")
}
