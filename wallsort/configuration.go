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

func SetBaseDirectory(dir string) {
	Config.Directory = dir
}
