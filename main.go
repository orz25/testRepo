package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"github.com/gen2brain/go-unarr"
	"github.com/mholt/archiver"
	"github.com/ulikunitz/xz"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//type entry struct {
//	Header *lha.Header
//	Size   int
//	Err    error
//}

func main() {

	testRAR("/Users/orz/Downloads/files_to_play/folder/xz/Automations/untitled folder/5eeea355389655.59822ff824b72.gif.Z")
	//if err != nil {
	//	log.Fatal(err)
	//}

}

func testRAR(src string) {
	w, err := os.Create("/Users/orz/Downloads/files_to_play/folder/xz/Automations/untitled folder/test.txt")
	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	read := bufio.NewReader(f)
	write := bufio.NewWriter(w)
	format := archiver.Lz4{}
	format.Decompress(read, write)

	//format := archiver.Lz4{}
	//err := format.Decompress(src)
	//if err != nil {
	//	panic(err)
	//}
	//for err == nil {
	//	_, err = format.Read()
	//}

}

func rarOpener(src string) {
	a, err := unarr.NewArchive(src)
	if err != nil {
		panic(err)
	}
	defer a.Close()
}

//func lhaDecopmress(src string) ([]*entry, error) {
//	f, err := os.Open(src)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer f.Close()
//	r := lha.NewReader(f)
//	var entries []*entry
//	for {
//		h, err := r.NextHeader()
//		if err != nil {
//			log.Fatalf("NextHeader failed: %s", err)
//		}
//		if h == nil {
//			return entries, err
//		}
//		n, err := r.Decode(ioutil.Discard)
//		entries = append(entries, &entry{
//			Header: h,
//			Size:   n,
//			Err:    err,
//		})
//	}
//}

func decompress(src string) ([]string, error) {
	f, err := os.Open(src)
	read := bufio.NewReader(f)
	data, _ := ioutil.ReadAll(read)
	f, _ = os.Create("/Users/orz/Downloads/files_to_play/folder/new.jar.xz")
	w, _ := xz.NewWriter(f)
	w.Write(data)
	w.Close()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fi, err := f.Stat()
	//if err != nil {
	//	log.Fatal(err, fi)
	//}
	//p := make([]byte, 1024)

	//Read data from reader and write to new file
	//fo, err := os.Create("/Users/orz/Downloads/files_to_play/folder/new.jar")
	//b := make([]byte, 1024)
	//for {
	//	n, err := rc.Read(b)
	//	if err != nil {
	//		log.Fatalf("NewReader error %v", n)
	//	}
	//	if n == 0 {
	//		break
	//	}
	//	if _, err := fo.Write(b[:n]); err != nil {
	//		panic(err)
	//	}
	//}
	return nil, err
}

// Unzip will decompress a zip archived file,
// copying all files and folders
// within the zip file (parameter 1)
// to an output directory (parameter 2).

func Unzip(src string, destination string) ([]string, error) {

	// a variable that will store any
	//file names available in a array of strings
	var filenames []string

	// OpenReader will open the Zip file
	// specified by name and return a ReadCloser
	// Readcloser closes the Zip file,
	// rendering it unusable for I/O
	// It returns two values:
	// 1. a pointer value to ReadCloser
	// 2. an error message (if any)
	r, err := zip.OpenReader(src)

	// if there is any error then
	// (err!=nill) becomes true
	if err != nil {
		// and this block will break the loop
		// and return filenames gathered so far
		// with an err message, and move
		// back to the main function

		return filenames, err
	}

	defer r.Close()
	// defer makes sure the file is closed
	// at the end of the program no matter what.

	for _, f := range r.File {

		// this loop will run until there are
		// files in the source directory & will
		// keep storing the filenames and then
		// extracts into destination folder until an err arises

		// Store "path/filename" for returning and using later on
		fpath := filepath.Join(destination, f.Name)

		// Checking for any invalid file paths
		if !strings.HasPrefix(fpath, filepath.Clean(destination)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s is an illegal filepath", fpath)
		}

		// the filename that is accessed is now appended
		// into the filenames string array with its path
		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Creating a new Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Creating the files in the target directory
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		// The created file will be stored in
		// outFile with permissions to write &/or truncate
		outFile, err := os.OpenFile(fpath,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			f.Mode())

		// again if there is any error this block
		// will be executed and process
		// will return to main function
		if err != nil {
			// with filenames gathered so far
			// and err message
			return filenames, err
		}

		rc, err := f.Open()

		// again if there is any error this block
		// will be executed and process
		// will return to main function
		if err != nil {
			// with filenames gathered so far
			// and err message back to main function
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer so that
		// it closes the outfile before the loop
		// moves to the next iteration. this kinda
		// saves an iteration of memory & time in
		// the worst case scenario.
		outFile.Close()
		rc.Close()

		// again if there is any error this block
		// will be executed and process
		// will return to main function
		if err != nil {
			// with filenames gathered so far
			// and err message back to main function
			return filenames, err
		}
	}

	// Finally after every file has been appended
	// into the filenames string[] and all the
	// files have been extracted into the
	// target directory, we return filenames
	// and nil as error value as the process executed
	// successfully without any errors*
	// *only if it reaches until here.
	return filenames, nil
}
