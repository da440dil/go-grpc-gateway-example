//+build mage

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
)

func Build() error {
	gopath := os.Getenv("GOPATH")
	if len(gopath) == 0 {
		return errors.New("$GOPATH required")
	}

	protoc := "protoc"
	protopath := "--proto_path=proto"
	googleapis := fmt.Sprintf("--proto_path=%s",
		filepath.Join(gopath, "/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis"))
	swagger := fmt.Sprintf("--proto_path=%s",
		filepath.Join(gopath, "/src/github.com/grpc-ecosystem/grpc-gateway"))
	filename := "proto/hello.proto"

	var out string
	var err error

	// Generate gRPC stub
	out = "--go_out=plugins=grpc:proto"
	err = sh.Run(protoc, protopath /* include, */, googleapis, swagger, out, filename)
	if err != nil {
		return err
	}

	// Generate reverse-proxy
	out = "--grpc-gateway_out=logtostderr=true:proto"
	err = sh.Run(protoc, protopath, googleapis, swagger, out, filename)
	if err != nil {
		return err
	}

	// Generate swagger definitions
	out = "--swagger_out=logtostderr=true:proto"
	err = sh.Run(protoc, protopath, googleapis, swagger, out, filename)
	if err != nil {
		return err
	}

	return nil
}
