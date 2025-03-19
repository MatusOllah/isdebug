# isdebug

[![Go Reference](https://pkg.go.dev/badge/github.com/MatusOllah/isdebug.svg)](https://pkg.go.dev/github.com/MatusOllah/isdebug) [![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/isdebug)](https://goreportcard.com/report/github.com/MatusOllah/isdebug) [![GitHub license](https://img.shields.io/github/license/MatusOllah/isdebug)](https://github.com/MatusOllah/isdebug/blob/main/LICENSE) [![Made in Slovakia](https://raw.githubusercontent.com/pedromxavier/flag-badges/refs/heads/main/badges/SK.svg)](https://www.youtube.com/watch?v=UqXJ0ktrmh0)

**isdebug** is a little cross-platform debugger detection package for Go.

## Features

* Detects whether a debugger is present at runtime
* Cross-platform support (Windows, macOS, Linux, BSD, Plan9, Android, iOS, etc...)
* Lightweight and easy to use

## Use Cases

* Anti-debugging and reverse engineering protection
* Game security and anti-cheat protection
* Secure applications
* Conditional debugging

## Installation

Run:

```sh
go get -u github.com/MatusOllah/isdebug
```

## Usage

```go
package main

import (
    "github.com/MatusOllah/isdebug"

    "fmt"
    "os"
)

func main() {
    for {
        if isdebug.IsDebug() {
            fmt.Println("Debugger detected!")
            os.Exit(0)
        }
    }
}
```
