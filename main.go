package main

import "github.com/docker/docker/client"

func main() {
    _ = client.NewClientWithOpts
}