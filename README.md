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
$ pixelize [OR] pixelize help [OR] pixelize man
```
This will print the manual.

Here is how it works :

/!\ The pixelizer needs to be used directly in the folder containing the image. This known limitation may change in the future. /!\

1. Print the manual :
```
$ pixelize [OR] pixelize help [OR] pixelize man
```
2. Pixelize an image :
```
$ pixelize [imageName] [X] [Y] [Z]
```
* imageName = name of the image, in the current directory, to pixelize.
* X = Width of each group of pixels to fuse (Ex : 5 will fuse every group of 5\*5 pixels from the image)
* Y = Number of results, with increasing X at each result (Ex : 5 3 will print 3 results while fusing groups of 5\*5 pixels for the first result, 6\*6 for the second, etc)
* Z = Change to the step increase between each result (Ex : 5 3 5 will print 3 results while fusing groups of 5\*5 pixels for the first result, 10\*10 for the second, etc)

3. Clear the folder of the result files :
```
$ pixelize clear
```
4. Redo the last pixelization :
```
$ pixelize redo
```
