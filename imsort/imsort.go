package imsort

type Configuration struct {
	Directory  string     `mapstructure:"directory"`
	Categories []Category `mapstructure:"dir"`
}

type Category struct {
	Name      string `mapstructure:"name"`
	Height    []int  `mapstructure:"height"`
	Width     []int  `mapstructure:"width"`
	MinPixels int    `mapstructure:"min_pixels"`
	MaxPixels int    `mapstructure:"max_pixels"`
	// HeightSet    bool
	// WidthSet     bool
	// MinPixelsSet bool
	// MaxPixelsSet bool
}

var Config Configuration
