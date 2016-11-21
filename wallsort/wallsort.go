package wallsort

import (
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path"

	_ "image/jpeg"
	_ "image/png"
)

type Configuration struct {
	Directory  string     `mapstructure:"directory"`
	Categories []Category `mapstructure:"category"`
}

type Category struct {
	Name      string `mapstructure:"name"`
	Heights   []int  `mapstructure:"heights"`
	Widths    []int  `mapstructure:"widths"`
	Height    int    `mapstructure:"height"`
	Width     int    `mapstructure:"width"`
	MinHeight int    `mapstructure:"min_height"`
	MinWidth  int    `mapstructure:"min_width"`
	MaxHeight int    `mapstructure:"max_height"`
	MaxWidth  int    `mapstructure:"max_width"`
	MinPixels int    `mapstructure:"min_pixels"`
	MaxPixels int    `mapstructure:"max_pixels"`

	// The filenames in the base directory, which belong into this category.
	Filenames []string
}

type Image struct {
	Filename string
	Height   int
	Width    int
}

var Config Configuration

func GenerateImageList() ([]Image, error) {
	var imageList []Image
	fis, err := ioutil.ReadDir(Config.Directory)
	if err != nil {
		return imageList, err
	}
	errOpen := errors.New("")
	errOpen = nil
	messageOpenError := "The following files could not be opened:\n"
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		name := fi.Name()
		fullname := path.Join(Config.Directory, name)
		file, ferr := os.Open(fullname)
		defer file.Close()
		if ferr != nil {
			messageOpenError += fullname
			errOpen = errors.New(messageOpenError)
			continue
		}
		imageConfig, _, cerr := image.DecodeConfig(file)
		if cerr != nil {
			// The open file is not a supported image. Continue the process in the this case.
			continue
		}
		img := Image{name, imageConfig.Height, imageConfig.Width}
		imageList = append(imageList, img)
	}
	return imageList, errOpen
}

func SortImages(imageList []Image) {
	for _, img := range imageList {
		for j, cat := range Config.Categories {
			if cat.Match(img) {
				Config.Categories[j].Filenames = append(Config.Categories[j].Filenames, img.Filename)
				break
			}
		}
	}
}

func MoveImages() error {
	basedir := Config.Directory
	errMessage := "The following files could not be moved:"
	errMove := errors.New("")
	errMove = nil
	for _, cat := range Config.Categories {
		destdir := path.Join(basedir, cat.Name)
		for _, fname := range cat.Filenames {
			src := path.Join(basedir, fname)
			dst := path.Join(destdir, fname)
			err := os.Rename(src, dst)
			if err != nil {
				errMessage += "\n" + src
				errMove = errors.New(errMessage)
			}
			fmt.Printf("File %s moved to %s.\n", src, dst)
		}
	}
	return errMove
}

// Return true, if the image description img fullfills all the
// properties of the category cat, otherwise false.
func (cat *Category) Match(img Image) bool {
	if cat.Height > 0 {
		if cat.Height != img.Height {
			return false
		}
	}
	if cat.Width > 0 {
		if cat.Width != img.Width {
			return false
		}
	}
	if cat.MinHeight > 0 {
		if img.Height < cat.MinHeight {
			return false
		}
	}
	if cat.MinWidth > 0 {
		if img.Width < cat.MinWidth {
			return false
		}
	}
	if cat.MaxHeight > 0 {
		if img.Height > cat.MaxHeight {
			return false
		}
	}
	if cat.MaxWidth > 0 {
		if img.Width > cat.MaxWidth {
			return false
		}
	}
	if len(cat.Heights) > 0 {
		found := false
		for _, h := range cat.Heights {
			if h == img.Height {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	if len(cat.Widths) > 0 {
		found := false
		for _, w := range cat.Widths {
			if w == img.Width {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	nPixels := img.Height * img.Width
	if cat.MaxPixels > 0 {
		if nPixels > cat.MaxPixels {
			return false
		}
	}
	if cat.MinPixels > 0 {
		if nPixels < cat.MinPixels {
			return false
		}
	}
	return true
}
