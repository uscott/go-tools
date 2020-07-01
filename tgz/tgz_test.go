package tgz

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

const testfile string = "tmp/test.txt"

func TestFileCreate(t *testing.T) {
	file, err := os.OpenFile(testfile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0754)
	defer file.Close()
	if err != nil {
		t.Errorf("%v\n", err)
	}
	var b = []byte("hello world hello world hello world hello world\n")
	for i := 0; i < 1000; i++ {
		_, err = file.Write(b)
		if err != nil {
			t.Errorf("%v\n", err)
		}
	}
}

func TestGzip(t *testing.T) {
	if err := Gzip(testfile); err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestTgzip(t *testing.T) {
	if err := Tgzip(testfile, ""); err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestTgunzip(t *testing.T) {
	tgzf := strings.ReplaceAll(testfile, "txt", "tgz")
	var names []string
	if err := Tgunzip(tgzf, &names); err != nil {
		t.Errorf("%v\n", err)
	}
	for _, x := range names {
		fmt.Println(x)
	}
}

func TestTargzToGz(t *testing.T) {
	if err := os.Remove(testfile + ".gz"); err != nil {
		t.Errorf("%v\n", err)
	}
	if err := TargzToGz(strings.ReplaceAll(testfile, ".txt", ".tgz")); err != nil {
		t.Errorf("%v\n", err)
	}
}

func TestTargzToGzDir(t *testing.T) {
	var dir, fname, ext string
	fname = testfile
	n := strings.IndexRune(fname, '/')
	if n >= 0 {
		dir = testfile[:n+1]
		fname = testfile[n+1:]
	} else {
		dir = "."
	}
	n = strings.LastIndex(fname, ".")
	if n >= 0 {
		ext = fname[n+1:]
		fname = fname[:n]
	}
	for i := 0; i < 10; i++ {
		src, err := os.Open(testfile)
		if err != nil {
			src.Close()
			t.Errorf("%v\n", err)
		}
		newname := fmt.Sprintf("%v%v%d.%v", dir, fname, i, ext)
		dst, err := os.OpenFile(newname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0764)
		if err != nil {
			src.Close()
			dst.Close()
			t.Errorf("%v\n", err)
		}
		if _, err = io.Copy(dst, src); err != nil {
			src.Close()
			dst.Close()
			t.Errorf("%v\n", err)
		}
		src.Close()
		dst.Close()
		if err = Tgzip(newname, ""); err != nil {
			t.Errorf("%v\n", err)
		}
		time.Sleep(100 * time.Millisecond)
	}
	if err := TargzToGzDir(dir); err != nil {
		t.Errorf("%v\n", err)
	}
}
