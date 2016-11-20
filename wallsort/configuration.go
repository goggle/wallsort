package wallsort

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

func Initialize() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.config/wallsort")
	viper.AddConfigPath("$HOME/.wallsort")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	return err
}

func ReadConfiguration() error {
	err := viper.Unmarshal(&Config)
	// adjustConfiguration(&Config)
	// fmt.Println(Config)
	return err
}

// Check, if the specified directory exists, if the user has appropriate
// access rights, and create the the specified subfolders if necessary.
func InitDirectory() error {
	if Config.Directory == "" {
		// No directory specified, return error.
		err := errors.New("No directory specified.")
		return err
	}

	fi, err := os.Stat(Config.Directory)
	if err != nil {
		if os.IsNotExist(err) {
			// The specified directory does not exist.
			message := fmt.Sprintf("The specified directory %s does not exist!", Config.Directory)
			errNotExists := errors.New(message)
			return errNotExists
		}
		return err
	}
	if !fi.IsDir() {
		// The specified directory exits, but is a regular file, not a directory.
		message := fmt.Sprintf("The specified directory %s is a regular file, not a directory!", Config.Directory)
		errNoDir := errors.New(message)
		return errNoDir
	}

	for _, cat := range Config.Categories {
		if cat.Name == "" {
			// No name for the subfolder specified, return error.
			errNoName := errors.New("There is a category in the configuration without a name.")
			return errNoName
		}
		path := path.Join(Config.Directory, cat.Name)
		fi, err := os.Stat(path)
		if err != nil {
			if os.IsNotExist(err) {
				// The subfolder does not exist. Create it.
				errMkdir := os.Mkdir(path, 0755)
				if errMkdir != nil {
					// Could not create subfolder. Return error.
					return errMkdir
				}
			} else {
				return err
			}
		} else if !fi.IsDir() {
			// The specified subfolder exists, but is not a directory, return error.
			message := fmt.Sprintf("The category %s is a regular file on the filesystem, but it should be a directory!", cat.Name)
			errNoDir := errors.New(message)
			return errNoDir
		}
	}
	return nil
}

func SetBaseDirectory(basedir string) {
	Config.Directory = basedir
}

func SetDefaultConfiguration(config *Configuration, basedir string) {
	config.Directory = basedir
	config.Categories = make([]Category, 0)
	cat_1080p := Category{
		Name:      "1080p",
		Heights:   make([]int, 0),
		Widths:    make([]int, 0),
		Height:    0,
		Width:     1920,
		MinHeight: 1080,
		MinWidth:  0,
		MaxHeight: 1200,
		MaxWidth:  0,
		MinPixels: 0,
		MaxPixels: 0,
	}
	cat_1440p := Category{
		Name:      "1440p",
		Heights:   make([]int, 0),
		Widths:    make([]int, 0),
		Height:    0,
		Width:     2560,
		MinHeight: 1440,
		MinWidth:  0,
		MaxHeight: 1600,
		MaxWidth:  0,
		MinPixels: 0,
		MaxPixels: 0,
	}
	cat_qhdplus := Category{
		Name:      "QHDplus",
		Heights:   make([]int, 0),
		Widths:    make([]int, 0),
		Height:    1800,
		Width:     3200,
		MinHeight: 0,
		MinWidth:  0,
		MaxHeight: 0,
		MaxWidth:  0,
		MinPixels: 0,
		MaxPixels: 0,
	}
	cat_uhd := Category{
		Name:      "UHD",
		Heights:   make([]int, 0),
		Widths:    make([]int, 0),
		Height:    2160,
		Width:     3840,
		MinHeight: 0,
		MinWidth:  0,
		MaxHeight: 0,
		MaxWidth:  0,
		MinPixels: 0,
		MaxPixels: 0,
	}
	cat_lowres := Category{
		Name:      "lowres",
		Heights:   make([]int, 0),
		Widths:    make([]int, 0),
		Height:    0,
		Width:     0,
		MinHeight: 0,
		MinWidth:  0,
		MaxHeight: 1080,
		MaxWidth:  1920,
		MinPixels: 0,
		MaxPixels: 2073599,
	}
	cat_highres := Category{
		Name:      "highres",
		Heights:   make([]int, 0),
		Widths:    make([]int, 0),
		Height:    0,
		Width:     0,
		MinHeight: 2160,
		MinWidth:  3840,
		MaxHeight: 0,
		MaxWidth:  0,
		MinPixels: 8294401,
		MaxPixels: 0,
	}
	cat_others := Category{
		Name:      "others",
		Heights:   make([]int, 0),
		Widths:    make([]int, 0),
		Height:    0,
		Width:     0,
		MinHeight: 0,
		MinWidth:  0,
		MaxHeight: 0,
		MaxWidth:  0,
		MinPixels: 0,
		MaxPixels: 0,
	}

	config.Categories = append(config.Categories, cat_1080p)
	config.Categories = append(config.Categories, cat_1440p)
	config.Categories = append(config.Categories, cat_qhdplus)
	config.Categories = append(config.Categories, cat_uhd)
	config.Categories = append(config.Categories, cat_lowres)
	config.Categories = append(config.Categories, cat_highres)
	config.Categories = append(config.Categories, cat_others)

}
