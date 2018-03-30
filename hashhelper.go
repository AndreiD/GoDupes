package main

//testing speed on various hashing algorithms

import (
	"fmt"
	"github.com/OneOfOne/xxhash"
	"io"
	"os"
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
)

func HashXXHash(path string) uint64 {
	h := xxhash.New64()
	r, err := os.Open(path)
	if err != nil {
		fmt.Println("error: ", err)
		return 0
	}
	io.Copy(h, r)
	return h.Sum64()
}

func HashMd5(path string) []byte {
	xhash := md5.New()
	xfile, _ := os.Open(path)
	io.Copy(xhash, xfile)
	xfile.Close()
	return xhash.Sum(nil)
}

func HashSha1(path string) string {
	xhash := sha1.New()
	xfile, _ := os.Open(path)
	io.Copy(xhash, xfile)
	xfile.Close()
	return hex.EncodeToString(xhash.Sum(nil))
}
