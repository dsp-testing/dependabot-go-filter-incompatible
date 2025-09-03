package main

import (
    "github.com/docker/docker/client"
)

func main() {
    // Just importing to satisfy Go module requirements
    _ = client.DefaultDockerHost
}
