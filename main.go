package main

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path"
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
		fmt.Println("No configuration file found.")
		ans := AskYesNo("Do you want to write a default configuration file?", true)
		if !ans {
			fmt.Println("You can manually create a configuration file and restart wallsort.")
			os.Exit(0)
		} else {
			fmt.Printf("Enter the path to the image directory: ")
			reader := bufio.NewReader(os.Stdin)
			imagePath, _ := reader.ReadString('\n')
			imagePath = strings.Trim(imagePath, " \n")
			wallsort.SetBaseDirectory(imagePath)
			wallsort.SetDefaultConfiguration(&wallsort.Config, imagePath)
			usr, errUser := user.Current()
			if errUser != nil {
				fmt.Println("Could not retrieve system user. No configuration file written. Try to manually create a configuration file.")
				os.Exit(1)
			}
			homeDir := usr.HomeDir
			configFile := path.Join(homeDir, ".config/wallsort/config.toml")
			// configFile := path.Join(homeDir, "test/config.toml")
			dirCreate := path.Dir(configFile)
			errCreate := os.MkdirAll(dirCreate, 0755)
			if errCreate != nil {
				fmt.Println("Could not create configuration file directory. No configuration file written. Try to manually create a configuration file.")
				os.Exit(1)
			}
			err := wallsort.WriteConfiguration(configFile)
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
			fmt.Printf("Successfully written configuration file to %s.\n", configFile)
			fmt.Println("Edit the configuration file and rerun wallsort to sort the images.")
			os.Exit(0)
		}
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
