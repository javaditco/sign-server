package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type UploadedFile struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Bytes  int64  `json:"bytes"`
	Sha256 string `json:"sha256"`
	Url    string `json:"url"`
}

func (u *UploadedFile) GenPath(path string) string {
	return filepath.Join(path, u.Name)
}

func (u *UploadedFile) GetSha256(path string) string {
	hasher := sha256.New()
	s, err := ioutil.ReadFile(filepath.Join(path, u.Name))
	hasher.Write(s)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

func (u *UploadedFile) GenURL() string {
	return fmt.Sprintf("http://%s/file/%s", Config_map_string["name"], u.Sha256)
}

func NewUploadedFile(fileName string, path string, bytes int64) UploadedFile {
	u := UploadedFile{
		Name:  fileName,
		Bytes: bytes,
	}
	u.Path = u.GenPath(path)
	u.Sha256 = u.GetSha256(path)
	u.Url = u.GenURL()
	return u
}
