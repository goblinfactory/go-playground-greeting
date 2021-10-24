package compression

// import (
// 	"archive/zip"
// 	"io"
// 	"io/ioutil"
// 	"log"
// )

// // TextFileZipped ...
// type TextFileZipped struct {
// 	Name  string
// 	Bytes []byte
// }

// // DemoReadingSimpleZipTextFile read a zip file of multiple text files, return content as []string
// func DemoReadingSimpleZipTextFile(zipfile string) []TextFileZipped {
// 	z, err := zip.OpenReader(zipfile)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	numFiles := len(z.File)
// 	files := make([]TextFileZipped, numFiles)

// 	for _, file := range z.File {

// 		f, err := file.Open()
// 		defer f.Close()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		b, err := ioutil.ReadAll(f)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 	}

// }

// type myReadCloser interface {
// 	io.ReadCloser
// }

// func (rc myReadCloser) readAllExitOnError() []byte {
// 	b, e := ioutil.ReadAll(rc)
// 	if e != nil {
// 		log.Fatal(e)
// 	}
// 	return b
// }

// func readContent(file *zip.File) []byte {
// 	zf, err := file.Open()

// 	defer func() {
// 		err := file.Close()
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	bytes, err := ioutil.ReadAll(zf)

// 	return content
// }
