package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

var (
	dir       int
	file      int
	dirColor  *color.Color = color.New(color.FgRed, color.Bold)
	fileColor *color.Color = color.New(color.FgCyan)
	infoColor *color.Color = color.New(color.FgBlue)
)

func main() {

	args := []string{"."}

	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		err := tree(arg, "")

		if err != nil {
			log.Fatal(err)
		}
	}

	infoColor.Printf("%v diresctories / %v files", dir, file)
}

func tree(root, indent string) error {

	fi, err := os.Stat(root)

	if err != nil {
		return fmt.Errorf("nothing found")
	}

	if fi.IsDir() {
		dirColor.Println(fi.Name())

		dir++
	} else {

		fileColor.Println(fi.Name())

		file++
	}

	if !fi.IsDir() {
		return nil
	}

	fis, err := ioutil.ReadDir(root)

	if err != nil {
		return fmt.Errorf("could not read directory")
	}

	for i, fi := range fis {

		if fi.Name()[0] == '.' {
			continue
		}

		add := "| "

		if i == len(fis)-1 {
			fmt.Printf(indent + "↳")
			add = " "

		} else {
			fmt.Printf(indent + "|--")

		}

		if err := tree(filepath.Join(root, fi.Name()), indent+add); err != nil {
			return err
		}
	}

	return nil
}
