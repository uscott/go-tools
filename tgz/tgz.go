package tgz

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

const compressionLevel = gzip.BestSpeed

// Gzip applies gzip compression to the file located at the path argument
func Gzip(path string) (e error) {
	f, e := os.Open(path)
	defer f.Close()
	if e != nil {
		return e
	}
	gzf, e := os.OpenFile(path+".gz", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0754)
	defer gzf.Close()
	if e != nil {
		return e
	}
	gzw := gzip.NewWriter(gzf)
	_, e = io.Copy(gzw, f)
	if e != nil {
		return fmt.Errorf(
			"Could not copy the file %s data to the gzip, got error: %s",
			path, e.Error())
	}
	gzw.Close()
	return nil
}

// Tgzip puts the file in the path argument into a .tgz archive
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
	f, e := os.Create(tgzpath)
	defer f.Close()
	gzpw, e := gzip.NewWriterLevel(f, compressionLevel)
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

// Tgunzip extracts a .tgz archive
func Tgunzip(path string, names *[]string) (e error) {
	f, e := os.Open(path)
	if e != nil {
		f.Close()
		return fmt.Errorf("Could not open file %v: %v", path, e.Error())
	}
	gzf, e := gzip.NewReader(f)
	f.Close()
	if e != nil {
		return fmt.Errorf("Could not create gzip reader %+v: %v", gzf, e.Error())
	}
	var dir string
	n := strings.IndexRune(path, '/')
	if n >= 0 {
		dir = path[:n+1]
	}
	tr := tar.NewReader(gzf)
	for {
		header, e := tr.Next()
		if e == io.EOF {
			break
		}
		if e != nil {
			return fmt.Errorf("tar.Reader.Next() failed: %v", e.Error())
		}
		switch header.Typeflag {
		case tar.TypeDir:
			if e = os.Mkdir(header.Name, 0754); e != nil {
				return fmt.Errorf(
					"could not create directory %v: %v", header.Name, e.Error())
			}
		case tar.TypeReg:
			fname := dir + header.Name
			opf, e := os.Create(fname)
			if e != nil {
				opf.Close()
				return fmt.Errorf(
					"could not create file %v: %v", fname, e.Error())
			}
			if _, e = io.Copy(opf, tr); e != nil {
				opf.Close()
				return fmt.Errorf("io.Copy failed: %v", e.Error())
			}
			opf.Close()
			if names != nil {
				*names = append(*names, fname)
			}
		default:
			return fmt.Errorf("unknown type %v in %v", header.Typeflag, header.Name)
		}
	}
	return nil
}

// TargzToGz extracts files from tgz archive then gzip compresses them individually
func TargzToGz(path string) (e error) {
	const buflen int = 0
	names := make([]string, 0, buflen)
	e = Tgunzip(path, &names)
	if e != nil {
		return e
	}
	for _, x := range names {
		if e = Gzip(x); e != nil {
			return e
		}
	}
	return nil
}

// TargzToGzDir applies TgzToGz to all files with extension tgz in a directory
func TargzToGzDir(dir string) (e error) {
	f, e := os.Open(dir)
	if e != nil {
		f.Close()
		return fmt.Errorf("could not open %v: %v", dir, e.Error())
	}
	names, e := f.Readdirnames(0)
	f.Close()
	if e != nil {
		return fmt.Errorf("error returned from os.File.Readdirnames: %v", e.Error())
	}
	// We only apply this to files with extension tgz!
	for _, x := range names {
		n := strings.LastIndex(x, ".")
		var extension string
		if n >= 0 {
			extension = x[n+1:]
		}
		if extension != "tgz" {
			continue
		}
		if e = TargzToGz(dir + x); e != nil {
			return e
		}
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
