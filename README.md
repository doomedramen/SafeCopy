# SafeCopy

## Overview

Safecopy is a command line tool for integrity checking a file when it is copied.
It gets the md5 sum of a file before copying it and then comparing that to the md5 sum of the copied file.
* md5sum file
* copy file
* check md5 of copied file
* delete original


## Installation

To install SafeCopy run:
```
$ go get github.com/wookoouk/SafeCopy.git
$ export PATH=$PATH:$GOPATH/bin
```

It is reccomended that you alias the `SafeCopy` command to something easier to type like `sc`
```
$ alias sc=SafeCopy
```

## Usage

```
$ sc ~/orginal_file.txt ~/copied_file.txt
```
