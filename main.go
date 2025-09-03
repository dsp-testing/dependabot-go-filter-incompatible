package main

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

func main() {
	// Just importing to satisfy Go module requirements
	_ = zap.NewNop()
	_ = bcrypt.DefaultCost
	_ = grpc.WithInsecure()
}
