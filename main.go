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

func tree(root, indent string) error {

	s, err := os.Stat(root)

	if err != nil {
		return fmt.Errorf("no dir or files found")
	}

	if s.IsDir() {
		dirColor.Println(s.Name())

		dir++
	} else {

		fileColor.Println(s.Name())

		file++
	}

	if !s.IsDir() {
		return nil
	}

	dir, err := ioutil.ReadDir(root)

	if err != nil {
		return fmt.Errorf("could not read directory")
	}

	for i, fi := range dir {

		if fi.Name()[0] == '.' {
			continue
		}

		add := "￨ "

		if i == len(dir)-1 {
			fmt.Printf(indent + "|__")
			add = " "

		} else {
			fmt.Printf(indent + "￨--")

		}

		if err := tree(filepath.Join(root, fi.Name()), indent+add); err != nil {
			return err
		}
	}

	return nil
}

func main() {

	args := []string{"."}

	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	for _, arg := range args {
		err := tree(arg, "")

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	infoColor.Printf("%v diresctories / %v files", dir, file)
}
