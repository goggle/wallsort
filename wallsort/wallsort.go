package wallsort

import (
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path"
)

type Configuration struct {
	Directory  string     `mapstructure:"directory"`
	Categories []Category `mapstructure:"dir"`
}

type Category struct {
	Name         string `mapstructure:"name"`
	Height       []int  `mapstructure:"height"`
	Width        []int  `mapstructure:"width"`
	MinPixels    int    `mapstructure:"min_pixels"`
	MaxPixels    int    `mapstructure:"max_pixels"`
	SingleHeight int    `mapstructure:"height"`
	SingleWidth  int    `mapstructure:"width"`

	// The filenames in the base directory, which belong into this category.
	Filenames []string
}

type Image struct {
	Filename string
	Height   int
	Width    int
}

var Config Configuration

func adjustConfiguration(config *Configuration) {
	for i, cat := range config.Categories {
		if cat.SingleHeight != 0 {
			sl := make([]int, 2)
			sl[0] = cat.SingleHeight
			sl[1] = cat.SingleHeight
			config.Categories[i].Height = sl
		}
		if cat.SingleWidth != 0 {
			sl := make([]int, 2)
			sl[0] = cat.SingleWidth
			sl[1] = cat.SingleWidth
			config.Categories[i].Width = sl
		}
		if len(cat.Height) == 1 {
			config.Categories[i].Height = append(config.Categories[i].Height, cat.Height[0])
		}
		if len(cat.Width) == 1 {
			config.Categories[i].Width = append(config.Categories[i].Width, cat.Width[0])
		}
	}
}

func GenerateImageList() ([]Image, error) {
	var imageList []Image
	fis, err := ioutil.ReadDir(Config.Directory)
	if err != nil {
		return imageList, err
	}
	errOpen := errors.New("")
	errOpen = nil
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		name := fi.Name()
		fullname := path.Join(Config.Directory, name)
		file, ferr := os.Open(fullname)
		defer file.Close()
		if ferr != nil {
			errOpen = errors.New("Could not open every file in the base directory.")
			continue
		}
		imageConfig, _, cerr := image.DecodeConfig(file)
		if cerr != nil {
			continue
		}
		img := Image{name, imageConfig.Height, imageConfig.Width}
		imageList = append(imageList, img)
	}
	return imageList, errOpen
}

func SortImages(imageList []Image) {
	for _, img := range imageList {
		for _, cat := range Config.Categories {
			fmt.Println(img, cat)

		}
	}
}

func (cat *Category) Match(img Image) bool {
	return true

}
