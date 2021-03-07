# Pixelizer
Pixelizer is an image pixelizer written in go.
It allows very simple and intuitive CLI pixelization.

# Installation
To install Pixelizer, you need to : 
* Install Go [here](https://golang.org/)
* Run this command :
```
$ go get -u github.com/eleby/pixelizer/...
```
* Optional - Install the package to use a cmd shortcut :
First navigate to your go sources repository, then : 
```
$ cd github.com/eleby/pixelizer/main
```
Then build the package (No install, since go install doesn't allow renaming) : 
```
$ go build -o pixelize
```
And move it to your $GOBIN or $GOPATH/bin : 
```
$ sudo mv pixelize [your $GOBIN or $GOPATH/bin]
```

# Usage
Pixelizer is meant to be simple to use, error-free and intuitive.
To be reminded of the usage, you can sure read this readme, but you can also open your terminal and use :
```
$ pixelize / $ pixelize man / $ pixelize help
```
This will print the manual.

Here is how it works :
1. Pixelize an image (Basic syntax) :
```
$ pixelize [imageName] [X] [Y] [Z] [...]
```
* imageName = path and name of the image to pixelize.
* X = Width of each group of pixels to fuse (Ex : 5 will fuse every group of 5\*5 pixels from the image)
* Y = Number of results, with increasing X at each result (Ex : 5 3 will print 3 results while fusing groups of 5\*5 pixels for the first result, 6\*6 for the second, etc)
* Z = Change to the step increase between each result (Ex : 5 3 5 will print 3 results while fusing groups of 5\*5 pixels for the first result, 10\*10 for the second, etc)
* ... = Secondary arguments

If not set, X, Y and Z will default to 1, meaning that "pixelize img.png" will output one image with a pixelization level of 1.

2. Secondary arguments to add to modify the basic syntax's behavior :
   
* Log the essential of the program activity in terminal :
```
$ pixelize [basic syntax] [log]
```
* Log more of the program activity in terminal :
```
$ pixelize [basic syntax] [verb]
```
* Print every result image in terminal using ANSI codes and register these in the program save directory (HomeDirectory/pixelizer) :
```
$ pixelize [basic syntax] [print]
```
* Reduce the size of the image so each pixelized group of colors will display as a pixel.
```
$ pixelize [basic syntax] [pixel]
```

The pixel argument cannot be used while outputting gifs. If both are specified, only the pixel one will be taken into account.
* Output a gif containing all of the result images :
```
$ pixelize [basic syntax] [gif] [reverse] [full]
```

You can add the optional parameter "reverse" or "full" to change the gif animation.

The "reverse" parameter will reverse the animation.

The "full" parameter will add the reversed animation to the normal one and make a loop of them.

4. Clear the folder of the result files :
```
$ pixelize clear
```
5. Redo the last pixelization :
```
$ pixelize redo
```
The "clear" and "redo" parameters can be used with the "log" or "verb" secondary arguments.


