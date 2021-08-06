# Go reverse shell

## Introduction
A simple reverse shell which can be compiled to use on Windows or Linux, depending on which file you use to compile.

## Setup
Just change the desired IP and port into `main.go`.

## Compile
### General
You can choose which architecture to use changing `GOARCH` value.
Also, you can select the desired platform with the parameter `GOOS`.

### Example for Windows
```sh
GOOS=windows GOARCH=386 go build -o binary/reverse.exe windows.go main.go
```

### Example for Linux

```sh
GOOS=linux GOARCH=amd64 go build -o binary/reverse linux.go main.go
```
