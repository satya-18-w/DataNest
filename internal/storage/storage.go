package storage

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type PathTransformFunc func(key string) PathKey

type PathKey struct {
	PathName string
	Original string
}



func ( p * PathKey) FileName()string {

	return fmt.Sprintf("%s/%s", p.PathName, p.Original)
}
type StoreOpts struct {
	PathTransformFunc
}
type Store struct {
	StoreOpts
}

// Real CAS PathTransform Function
func CASPathTransformFunc(key string) PathKey {
	// Md5 for the filename
	hash := sha1.Sum([]byte(key)) // [20]byte => []byte to slice then do [:]
	hashstr := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLen := len(hashstr) / blockSize

	paths := make([]string, sliceLen)
	for i := 0; i < sliceLen; i++ {
		paths[i] = hashstr[i*blockSize : (i+1)*blockSize]
	}

	return PathKey{
		PathName: strings.Join(paths, "/"),
		Original: hashstr,
	}

}

func DefaultPathTransformFunc(key string) string {
	transformKey := fmt.Sprintf("%s/DataNest", key)
	return transformKey
}

func NewStore(opts StoreOpts) *Store {
	return &Store{
		StoreOpts: opts,
	}
}

func (s *Store) WriteStream(key string, r io.Reader) error {
	var Pathkey PathKey
	if s.PathTransformFunc != nil {
		Pathkey = s.PathTransformFunc(key)
		fmt.Println("Transformed Path Name ", Pathkey.PathName)
	}
	fmt.Println("This is the pathKey : ", Pathkey)

	if err := os.MkdirAll(Pathkey.PathName, os.ModePerm); err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, r)

	fileNameBytes := md5.Sum(buf.Bytes())
	fileName := hex.EncodeToString(fileNameBytes[:])
	fullPath := fmt.Sprintf("%s/%s", Pathkey.PathName, fileName)
	fmt.Println(fullPath)

	f, err := os.Create(fullPath)
	if err != nil {
		return err

	}
	n, err := io.Copy(f, buf)
	if err != nil {
		return err
	}

	log.Printf("Bytes Written to the file %d", n)
	return nil
}
