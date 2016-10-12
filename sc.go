package main

import (
	"os"
	"github.com/codegangsta/cli"
	"io"
	"crypto/md5"
	"log"
	"github.com/fatih/color"
	"fmt"
	"encoding/hex"
)

func main() {
	app := cli.NewApp()
	app.Name = "sc"
	app.Usage = "copy files + checksum"
	app.Action = func(c *cli.Context) error {

		//PrintGreen(string(len(os.Args)))

		if (len(os.Args) == 3) {

			fromPath := os.Args[1]
			toPath := os.Args[2]

			if _, err := os.Stat(fromPath); os.IsNotExist(err) {
				// fromPath does not exist, BAD!
				fail(fromPath + " DOES NOT EXIST")
				return nil
			}

			if _, err := os.Stat(toPath); err == nil {
				// toPath exists, BAD!
				fail(toPath + " ALREADY EXISTS")
				return nil
			}

			initSum, err := checksum(fromPath)
			check(err)

			copyError := copyFile(fromPath, toPath)
			check(copyError)

			postSum, err := checksum(toPath)
			check(err)

			if (sliceEq(initSum, postSum)) {
				//log.Println("copied ok :D")
				PrintGreen("COPIED OK (" + hex.EncodeToString(initSum) + " === " + hex.EncodeToString(postSum) + ")")
				return nil
			} else {
				fail("DID NOT COPY OK! :(")
				//TODO delete bad copy
				return nil
			}
		} else {
			cli.ShowAppHelp(c);
			return nil
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

func copyFile(src, dst  string) error {
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

func fail(text string) {
	PrintRed(text)
	os.Exit(1)
}

func PrintGreen(s string) {
	color.Set(color.FgGreen)
	fmt.Println(s)
	color.Unset()
}

func PrintRed(s string) {
	color.Set(color.FgRed)
	fmt.Println(s)
	color.Unset()
}