package tgz

import (
	"bufio"
	"os"
	"testing"
)

const testfile string = "test.txt"

func TestFileCreate(t *testing.T) {
	file, err := os.OpenFile(testfile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0754)
	defer file.Close()
	if err != nil {
		t.Errorf("%v\n", err)
	}
	w := bufio.NewWriter(file)
	w.WriteString("hello world hello world hello world")
	err = w.Flush()
	if err != nil {
		t.Errorf("%v\n", err)
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
