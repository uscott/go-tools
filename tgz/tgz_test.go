package tgz

import (
	"fmt"
	"os"
	"strings"
	"testing"
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
	if err := TargzToGzDir("tmp/"); err != nil {
		t.Errorf("%v\n", err)
	}
}
