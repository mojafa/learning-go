// command shasum -a 256 file
// 7c6c8937b2a120af15849db05c9f46326761e0eec852a2e973b1e0b6acd59a01  SNOWY-EVENING.txt
// 890845889270f185bfefeed33ac31f338b93d0e698323b8a9df7ac5de29b76c4  SNOWY-EVENING.txt

package main

// import (
// 	"crypto/sha1"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"strings"
// )

// func main() {
// 	signature, err := Sha1Sum("SNOWY-EVENING.txt")
// 	if err != nil {
// 		log.Fatalf("error: %s", err)
// 	}
// 	fmt.Println(signature)

// 	signature, err = Sha1Sum("main.go")
// 	if err != nil {
// 		log.Fatalf("error: %s", err)
// 	}
// 	fmt.Println(signature)
// }

// // SNOWY-EVENING.txt
// func Sha1Sum(filename string) (string, error) {

// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	var r io.Reader = file
// 	if strings.HasSuffix(filename, ".txt") {
// 		File, err := txt.NewReader(file)
// 		if err != nil {
// 			return "", err
// 		}
// 		defer File.Close()

// 		r = file
// 	}

// 	w := sha1.New()
// 	// io.CopyN(os.Stdout, r, 100)
// 	if _, err := io.Copy(w, r); err != nil {
// 		return "", err
// 	}

// 	sig := w.Sum(nil)
// 	return fmt.Sprintf("%x", sig), nil

// }
