//+build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build() error {
	return sh.Run("protoc", "--go_out=plugins=grpc:.", "proto/hello.proto")
}
