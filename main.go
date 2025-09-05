package main

import (
    "fmt"
)

func main() {
   version := utils.GetVersion()
   fmt.Printf("Using private-go-utils version: %s\n", version)
}
