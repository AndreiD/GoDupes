package main

import (
	"testing"
)


func benchmarkHashing(b *testing.B, algo string, path string) {
	switch algo {
	case "xxhash":
		for i := 0; i < b.N; i++ {
			HashXXHash(path)
		}
	case "md5":
		for i := 0; i < b.N; i++ {
			HashMd5(path)
		}
	case "sha1":
		for i := 0; i < b.N; i++ {
			HashSha1(path)
		}


	}
}

func BenchmarkXXHashSmall(b *testing.B) { benchmarkHashing(b,"xxhash","D:\\FINISHED\\file.mp4") }
func BenchmarkXXHashMedium(b *testing.B)   { benchmarkHashing(b,"xxhash","D:\\FINISHED\\photoshop.exe") }
func BenchmarkXXHashBig(b *testing.B)    { benchmarkHashing(b,"xxhash","D:\\FINISHED\\mint.iso") }

func BenchmarkMd5Small(b *testing.B) { benchmarkHashing(b,"md5","D:\\FINISHED\\file.mp4") }
func BenchmarkMd5Medium(b *testing.B)   { benchmarkHashing(b,"md5","D:\\FINISHED\\photoshop.exe") }
func BenchmarkMd5Big(b *testing.B)    { benchmarkHashing(b,"md5","D:\\FINISHED\\mint.iso") }

func BenchmarkSha1Small(b *testing.B) { benchmarkHashing(b,"sha1","D:\\FINISHED\\file.mp4") }
func BenchmarkSha1Medium(b *testing.B)   { benchmarkHashing(b,"sha1","D:\\FINISHED\\photoshop.exe") }
func BenchmarkSha1Big(b *testing.B)    { benchmarkHashing(b,"sha1","D:\\FINISHED\\mint.iso") }



