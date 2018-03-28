package main

import (
	"os"
	"path/filepath"
	"github.com/fatih/color"
	"runtime"
	"strings"
	"hash"
	"sync"
	"crypto/md5"
	"strconv"
)

//github.com/OneOfOne/xxhash

type Group struct {
	files []*File
	fsize int64
}

const FIRST int64 = 4096 * 4
const MULT int64 = 8

var files = make([]*File, 0)
var groups = make([]*Group, 0)
var mutex = &sync.Mutex{}

type File struct {
	path string
	size int64
	sum  hash.Hash
}

var red = color.New(color.FgRed).PrintfFunc()
var info = color.New(color.Bold, color.FgBlue).PrintlnFunc()
var success = color.New(color.Bold, color.FgGreen).PrintlnFunc()

func collectFiles(fp string, info os.FileInfo, err error) error {
	if err != nil {
		red("error: %s",err)
		return nil
	}

	//no hidden files
	if strings.HasPrefix(info.Name(), ".") {
		if info.IsDir() {
			return filepath.SkipDir
		}
		return nil
	}

	//no directories
	if info.IsDir() {
		return nil
	}

	//no shortcuts
	if !info.Mode().IsRegular() {
		// ignore symlinks
		return nil
	}

	if info.Size() > 0 {
		files = append(files, &File{path: fp, size: info.Size(), sum: md5.New()})
	}
	return nil
}

func main() {
	color.NoColor = false
	runtime.GOMAXPROCS(runtime.NumCPU())


	info("~~~~ Welcome to Super Fast Go Duplicates Finder ~~~~\n")

	if len(os.Args) != 2 {
		red("Error: Please enter the path for scanning\n")
		os.Exit(0)
	}

	dirScan, err := filepath.Abs(os.Args[1])
	if err != nil {
		red("error: %s",err)
		os.Exit(1)
	}

	success("Scanning: " + dirScan)



	//get the files
	err = filepath.Walk(dirScan, collectFiles)

	if err != nil {
		red("error: %s",err)
		os.Exit(1)
	}

	info("A total of "+strconv.Itoa(len(files)) + " files were found")

}