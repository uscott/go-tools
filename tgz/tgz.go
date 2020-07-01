package tgz

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const compressionLevel = gzip.BestSpeed

// Gzip applies gzip compression to the file located at the path argument
func Gzip(path string) (e error) {
	file, e := os.Open(path)
	defer file.Close()
	if e != nil {
		return e
	}
	reader := bufio.NewReader(file)
	content, e := ioutil.ReadAll(reader)
	if e != nil {
		return e
	}
	gzf, e := os.OpenFile(path+".gz", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0754)
	defer gzf.Close()
	if e != nil {
		return e
	}
	gzw := gzip.NewWriter(gzf)
	gzw.Write(content)
	gzw.Flush()
	gzw.Close()
	return nil
}

// Tgzip puts the file in the path argument into a .tgz file
func Tgzip(path string, extension string) (e error) {
	var tgzpath string
	switch {
	case len(extension) == 0:
		n := strings.LastIndex(path, ".")
		if n >= 0 {
			extension = path[n+1:]
			tgzpath = strings.ReplaceAll(path, extension, "tgz")
		} else {
			tgzpath = fmt.Sprintf("%v.%v", path, "tgz")
		}
	default:
		tgzpath = strings.ReplaceAll(path, extension, "tgz")
	}
	file, e := os.Create(tgzpath)
	defer file.Close()
	gzpw, e := gzip.NewWriterLevel(file, compressionLevel)
	defer gzpw.Close()
	if e != nil {
		return fmt.Errorf("%v", e.Error())
	}
	tarw := tar.NewWriter(gzpw)
	defer tarw.Close()
	if e = addFileToTarWriter(path, tarw); e != nil {
		return fmt.Errorf(
			"Could not add file '%s', to tarball, got error '%s'", path, e.Error())
	}
	return nil
}

func addFileToTarWriter(path string, tw *tar.Writer) (e error) {
	file, e := os.Open(path)
	if e != nil {
		return fmt.Errorf(
			"Could not open file '%s', got eor '%s'", path, e.Error())
	}
	defer file.Close()
	stat, e := file.Stat()
	if e != nil {
		return fmt.Errorf(
			"Could not get stat for file '%s', got error '%s'", path, e.Error())
	}
	n := strings.IndexRune(path, '/')
	name := path[n+1:]
	header := &tar.Header{
		Name:    name,
		Size:    stat.Size(),
		Mode:    int64(stat.Mode()),
		ModTime: stat.ModTime(),
	}
	if e = tw.WriteHeader(header); e != nil {
		return fmt.Errorf(
			"Could not write header for file '%s', got error '%s'", path, e.Error())
	}
	_, e = io.Copy(tw, file)
	if e != nil {
		return fmt.Errorf(
			"Could not copy the file '%s' data to the tarball, got error '%s'",
			path, e.Error())
	}
	return nil
}
