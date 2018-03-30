package main

import (
	"os"
	"path/filepath"
	"github.com/fatih/color"
	"runtime"
	"strings"
	"flag"
	"fmt"
	"sort"
	"time"
)


type XFile struct {
	path string
	size int64
	hash  uint64
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
		xfiles = append(xfiles, &XFile{path: fp, size: info.Size()})
	}
	return nil
}

func main() {
	color.NoColor = false
	runtime.GOMAXPROCS(runtime.NumCPU())
	start := time.Now()

	info("Welcome to Super Fast Go Duplicates Finder\n")

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

	success("Scanning: " + dirScan +"\n")

	//get the files
	if er := filepath.Walk(dirScan, collectFiles); er != nil {
		red("error: %s", er)
		os.Exit(1)
	}

	fmt.Println()

	info("A total of %d files were found\n\n", len(xfiles))

	//sort them by size
	sort.Slice(xfiles, func(i, j int) bool {
		return xfiles[i].size > xfiles[j].size
	})

	for i := 0; i < len(xfiles)-1; i++ {
		currFile := xfiles[i]
		nextFile := xfiles[i+1]

		//only if they have equal size...
		if currFile.size == nextFile.size {
			//...compare their hash
			currFile.hash = HashXXHash(currFile.path)
			nextFile.hash = HashXXHash(nextFile.path)

			if currFile.hash == nextFile.hash {
				success("duplicates on %s | %s\n", currFile.path, nextFile.path)

				if *xdelete == "yes" {
					//if it contains copy in the filename, delete it
					fileToDelete := nextFile
					if strings.Contains(strings.ToLower(currFile.path), "copy") {
						fileToDelete = currFile
					}
					if xerr := os.Remove(fileToDelete.path); xerr != nil {
						red("error deleting file %s\n", xerr)
					} else {
						success("duplicate on %s is deleted!\n", fileToDelete.path)
					}
				}
			}

		}
	}

	info("\nGoDupes finished in %s", time.Since(start))

}
