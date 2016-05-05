package main

import (
	"os"
	"github.com/codegangsta/cli"
	"io"
	"crypto/md5"
	"log"
	"github.com/fatih/color"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "sc"
	app.Usage = "copy files + checksum"
	app.ActionFunc = func(c *cli.Context) {

		if (len(os.Args) > 0) {

			fromPath := os.Args[1]
			toPath := os.Args[2]

			if _, err := os.Stat(fromPath); os.IsNotExist(err) {
				// fromPath does not exist, BAD!
				fail(fromPath, "does not exist")
			}

			if _, err := os.Stat(toPath); err == nil {
				// toPath exists, BAD!
				fail(toPath, "already exists")
			}

			initSum, err := checksum(fromPath)
			check(err)

			copyError := copy(fromPath, toPath)
			check(copyError)

			postSum, err := checksum(toPath)
			check(err)

			if (sliceEq(initSum, postSum)) {
				//log.Println("copied ok :D")
				PrintGreen("copied ok")
			} else {
				fail("did not copy ok :(")
			}
		} else {
			app.ShowAppHelp(c);
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
		os.Exit(1)
	}
}

func fail(text ...string) {
	PrintRed(text)
	os.Exit(1)
}

func PrintGreen(s []string) {
	color.Set(color.FgGreen)
	PrintArray(s)
	color.Unset()
}

func PrintRed(s []string) {
	color.Set(color.FgRed)
	PrintArray(s)
	color.Unset()
}

func PrintArray(fs []string) {
	for i, v := range fs {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}