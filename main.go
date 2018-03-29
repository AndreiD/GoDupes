package main

import (
	"os"
	"path/filepath"
	"github.com/fatih/color"
	"runtime"
	"strings"
	"hash"
	"crypto/md5"
	"strconv"
	"io"
	"flag"
)

//github.com/OneOfOne/xxhash

type XFile struct {
	path string
	size int64
	sum  hash.Hash
}

var xfiles = make([]*XFile, 0)



func collectFiles(fp string, info os.FileInfo, err error) error {
	if err != nil {
		red("error: %s", err)
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
		xfiles = append(xfiles, &XFile{path: fp, size: info.Size(), sum: md5.New()})
	}
	return nil
}

func main() {
	color.NoColor = false
	runtime.GOMAXPROCS(runtime.NumCPU())

	info("~~~~ Welcome to Super Fast Go Duplicates Finder ~~~~\n")

	if len(os.Args) < 3 {
		red("Error: Please enter the path for scanning\n")
		os.Exit(0)
	}

	thedir := flag.String("dir", "/home", "the path to the location you want to scan")
	delete := flag.String("delete", "no", "no = test mode, yes = it will delete them")
	flag.Parse()

	switch *delete {
	case "no":
		info("This scripts runs in test mode!")
	case "yes":
		info("This scripts runs in deleting mode!")
	default:
		red("please use -delete=no or -delete=yes as arguments")
		os.Exit(1)

	}

	dirScan, err := filepath.Abs(*thedir)
	if err != nil {
		red("error: %s", err)
		os.Exit(1)
	}

	success("Scanning: " + dirScan)

	//get the files
	err = filepath.Walk(dirScan, collectFiles)

	if err != nil {
		red("error: %s", err)
		os.Exit(1)
	}

	info("A total of " + strconv.Itoa(len(xfiles)) + " files were found")

	//brute force...is bad
	for _, file := range xfiles {

		yfile, err := os.Open(file.path)
		if err != nil {
			red("error: %s", err)
			os.Exit(1)
		}
		defer yfile.Close()

		xhash := md5.New()
		_, err = io.Copy(xhash, yfile)
		if err != nil {
			red("error: %s", err)
			os.Exit(1)
		}

		file.sum = xhash

	}

}
