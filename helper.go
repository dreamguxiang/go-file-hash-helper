package main

import (
	"os"
	"sync"
)

func Generate(oripath string, wg *sync.WaitGroup) {
	files, err := GetFiles(oripath)
	if err != nil {
		return
	}
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			hash, err := GetFileSha256(file)
			if err != nil {
				wg.Done()
				return
			}
			err = CreateHashFile(file, hash)
			if err != nil {
				wg.Done()
				return
			}
			wg.Done()
		}(file)
	}
}

func VerifyHashFile(path string) (bool, error) {
	//获取.fs256文件中的hash值
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()
	hash := make([]byte, 64)
	_, err = file.Read(hash)
	if err != nil {
		return false, err
	}
	//获取文件hash
	filehash, err := GetFileSha256(GetFileNameWithoutHash(path))
	if err != nil {
		return false, err
	}
	return string(hash) == filehash, nil
}

func Verify(oripath string, wg *sync.WaitGroup, isError *bool) {
	files, err := GetHashFiles(oripath)
	if err != nil {
		return
	}
	for _, orifile := range files {
		wg.Add(1)
		go func(orifile string) {
			result, err := VerifyHashFile(orifile)
			if err != nil {
				wg.Done()
				return
			}
			if result {
				println("Verify Success: " + GetFileNameWithoutHash(orifile))
			} else {
				*isError = true
				println("Verify Failed: " + GetFileNameWithoutHash(orifile))
			}
			wg.Done()
		}(orifile)
	}
}
