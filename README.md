# Wallsort

Wallsort is a command line utility to sort images (mostly intended to sort wallpapers)
according to user defined rules.

If you have a directory which contains all your wallpapers, you might want to
sort them into subdirectories according to their sizes. Wallsort will do this
for you. It currently supports JPEG and PNG images.


## Installation

Run this to install wallsort:
```
go get -u github.com/goggle/wallsort
```
You should now have a executable file called `wallsort` in your `$GOPATH/bin` directory.


## Configuration

Wallsort uses a [TOML] (https://github.com/toml-lang/toml) configuration file called
`config.toml`. It can be stored in the directory `$HOME/.config/wallsort`,
`$HOME/.wallsort` or the directory, where the executable is stored.

When wallsort is executed for the first time, it can automatically create a
default configuration file. It will look like this:
```
directory = "/home/user/Pictures/wallpapers"

[[category]]
name = "1080p"
width = 1920
min_height = 1080
max_height = 1200

[[category]]
name = "1440p"
width = 2560
min_height = 1440
max_height = 1600

[[category]]
name = "QHDplus"
height = 1800
width = 3200

[[category]]
name = "UHD"
height = 2160
width = 3840

[[category]]
name = "lowres"
max_height = 1080
max_width = 1920
max_pixels = 2073599

[[category]]
name = "highres"
min_height = 2160
min_width = 3840
min_pixels = 8294401

[[category]]
name = "others"
```
The configuration file needs to have an entry ```directory = "/path/to/wallpapers"```.
This sets the base image directory.
You can now set rules to sort your images by defining a ```[[category]]``` section.
Valid fields in a ```[[category]]``` section are:
* ```name```: Defines a name for this category (e.g. ```name = "1920x1080"```).
* ```height```: Defines the image height in pixels (e.g. ```height = 1080```).
* ```width```: Defines the image width in pixels (e.g. ```width = 1920```).
