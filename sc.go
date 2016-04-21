package main

import (
	"os"
	"github.com/codegangsta/cli"
	"io"
	"crypto/md5"
	"log"
)

//TODO checksum original
//TODO copy
//TODO checksum copy


func main() {
	app := cli.NewApp()
	app.Name = "sc"
	app.Usage = "copy files + checksum"
	app.Action = func(c *cli.Context) {

		fromPath := os.Args[1]
		toPath := os.Args[2]

		println("initsum")
		initSum, err := checksum(fromPath)
		check(err)

		println("copy")
		copyError := copy(fromPath, toPath)
		check(copyError)

		println("postsum")
		postSum, err := checksum(toPath)
		check(err)

		if (sliceEq(initSum, postSum)) {
			//log.Println("copied ok :D")
		} else {
			log.Fatal("did not copy ok :(")
		}
	}

	app.Run(os.Args)
}

func checksum(filePath string) ([]byte, error) {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}

	return hash.Sum(result), nil
}

func copy(src, dst  string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}

func sliceEq(a, b []byte) bool {

	if a == nil && b == nil {
		return true;
	}

	if a == nil || b == nil {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func check(e error) {
	if (e != nil) {
		log.Fatal(e)
	}
}