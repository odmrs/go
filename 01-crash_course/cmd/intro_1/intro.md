# Introduction to GOLANG

Fun facts:
1 -> Statically and Strongly typed
2 -> GO is Compiled
3 -> Built in Concurrency
4 -> Very simplicity a lot because of garbage collector

## Variables

```go
var myVariable = "This is my Variables"
// or 
var myVariable string
```

## Packages are files!!
Everything inside your project is a module:
module(
    package 1:
    folder 1:
        file_1.go
        file_1.go
        file_1.go
)

## Start a project

A best pratices is initalize a module with the name of your repository like this repo:

```shell
go mod init github.com/odmrs/go
```

## First file

```go
package main // Told to compiler that is the main file, and the compiler will execute this file first

import "fmt" // Here the compiler will import the "fmt" module

func main() {
    fmt.Println("Hello World!") // fmt will call the funciton fmt
}
```
