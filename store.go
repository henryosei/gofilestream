package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

type PathTransformFunc func(string) PathKey

func CASPathTransformFun(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])
	blockSize := 5
	sliceLen := len(hashString) / blockSize
	paths := make([]string, sliceLen)
	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, (i*blockSize)+blockSize
		paths[i] = hashString[from:to]
	}
	return PathKey{PathName: strings.Join(paths, "/"), Orignal: hashString}
}

var DefaultPathTransformfunc = func(key string) string {
	return key
}

type StoreOpts struct {
	PathTransformFunc PathTransformFunc
}

type Store struct {
	StoreOpts
}

func NewStore(opts StoreOpts) *Store {
	return &Store{StoreOpts: opts}
}

type PathKey struct {
	PathName string
	Orignal  string
}

func (s *Store) writeStream(key string, r io.Reader) error {
	pathkey := s.PathTransformFunc(key)
	if err := os.MkdirAll(pathkey.PathName, os.ModePerm); err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	io.Copy(buf, r)
	//filenameBytes := md5.Sum(buf.Bytes())
	//filename := hex.EncodeToString(filenameBytes[:])

	pathAndFinename := pathkey.PathName + "/" + pathkey.PathName

	f, err := os.Create(pathAndFinename)
	if err != nil {
		return err
	}
	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}
	log.Printf("writtend (%d) byte to disk: %s", n, pathAndFinename)
	return nil
}
