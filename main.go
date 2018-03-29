package main

import (
	"os"
	"path/filepath"
	"github.com/fatih/color"
	"runtime"
	"strings"
	"crypto/md5"
	"io"
	"flag"
	"fmt"
	"encoding/hex"
	"sort"
)

//github.com/OneOfOne/xxhash

type XFile struct {
	path string
	size int64
	sum  string
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
		xfiles = append(xfiles, &XFile{path: fp, size: info.Size(), sum: ""})
	}
	return nil
}

func main() {
	color.NoColor = false
	runtime.GOMAXPROCS(runtime.NumCPU())

	var hashes []string

	info("~~~~ Welcome to Super Fast Go Duplicates Finder ~~~~\n")

	if len(os.Args) < 3 {
		red("Error: Please enter the path for scanning\n")
		os.Exit(0)
	}

	thedir := flag.String("dir", "/home", "the path to the location you want to scan")
	xdelete := flag.String("delete", "no", "no = test mode, yes = it will DELETE them!")
	flag.Parse()

	switch *xdelete {
	case "no":
		info("This scripts runs in test mode!\n")
	case "yes":
		info("This scripts runs in deleting mode!\n")
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
	if er := filepath.Walk(dirScan, collectFiles); er != nil {
		red("error: %s", er)
		os.Exit(1)
	}

	fmt.Println()

	info("A total of %d files were found\n\n", len(xfiles))

	//brute force...is bad
	for _, xfile := range xfiles {

		yfile, err := os.Open(xfile.path)
		if err != nil {
			red("error: %s", err)
			os.Exit(1)
		}


		xhash := md5.New()

		if _, err := io.Copy(xhash, yfile);  err != nil {
			red("error: %s", err)
			os.Exit(1)
		}

		yfile.Close()

		xfile.sum = hex.EncodeToString(xhash.Sum(nil))

		sort.Strings(hashes)

		target := xfile.sum

		sort.Strings(hashes)
		i := sort.Search(len(hashes), func(i int) bool { return hashes[i] >= target })
		if i < len(hashes) && hashes[i] == target {
			if *xdelete == "yes" {
				if xerr := os.Remove(xfile.path); xerr != nil {
					red("error deleting file %s\n", xerr)
				}else{
					success("duplicate on %s is deleted!\n", xfile.path)
				}

			}else{
				warn("duplicate on %s\n", xfile.path)
			}
		} else {
			hashes = append(hashes, xfile.sum)
		}

	}

}
