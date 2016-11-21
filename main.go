package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/goggle/wallsort/wallsort"
)

func AskYesNo(question string, defaultYes bool) bool {
	if defaultYes {
		fmt.Printf("%s [Y/n] ", question)
	} else {
		fmt.Printf("%s [y/N] ", question)
	}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, " \n")
	if text == "Y" || text == "y" || text == "yes" || text == "Yes" || text == "YES" || text == "1" {
		return true
	} else if text == "N" || text == "n" || text == "no" || text == "No" || text == "NO" || text == "0" {
		return false
	}
	return defaultYes
}

func main() {
	// fmt.Println("Hello World")
	errConfig := wallsort.Initialize()
	if errConfig != nil {
		fmt.Println("No configuration file found. Aborting...")
		os.Exit(1)
	}
	errParse := wallsort.ReadConfiguration()
	if errParse != nil {
		fmt.Println("Could not parse configuration file!")
		fmt.Printf("%v\n", errParse)
		os.Exit(1)
	}
	errInit := wallsort.InitDirectory()
	if errInit != nil {
		fmt.Printf("%v\n", errInit)
		os.Exit(1)
	}

	imageList, errImgList := wallsort.GenerateImageList()
	// TODO: In some cases we might be able to continue the program:
	if errImgList != nil {
		fmt.Printf("%v\n", errImgList)
		os.Exit(1)
	}

	wallsort.SortImages(imageList)
	errMove := wallsort.MoveImages()
	if errMove != nil {
		// TODO: Printf
		os.Exit(1)
	}

	errConfWrite := wallsort.WriteConfiguration("/home/alex/test/config.toml")
	if errConfig != nil {
		fmt.Printf("%v\n", errConfWrite)
	}

}
