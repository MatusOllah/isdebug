# isdebug

**isdebug** is a cross-platform debugger detection package for Go.

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
