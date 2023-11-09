package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	sig, err := sha1Sum("http.log.gz")
	if err != nil {
		log.Println(err)
	}
	log.Println(sig)

	sig, err = sha1Sum("sha1.go")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	log.Println(sig)
}

/*
if file names end with gz

	$ cat http.log.gz| gunzip | sha1sum

else
$ cat http.log.gz | sha1sum
*/
func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", nil
	}
	defer file.Close()

	// several defers in a function they will run in reverse order like a stack (LIFO)
	var r io.Reader = file
	if strings.HasSuffix(fileName, ".gz") {

		file, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer file.Close()
		r = file
	}
	// io.CopyN(os.Stdout, r, 100) // <-- uncompressed reader
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}
	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}
