package main

import (
	"fmt"
	"os"
	"os/exec"
)

type GrpcGenerator struct {
}

func (d *GrpcGenerator) Run(opt *Option, metaData *ServiceMetaData) (err error) {

	//protoc --go_out=plugins=grpc:. hello.proto
	outputParams := fmt.Sprintf("plugins=grpc:%s/generate/", opt.Output)
	cmd := exec.Command("protoc", "--go_out", outputParams, opt.Proto3Filename)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Printf("grpc generator failed, err:%v\n", err)
		return

	}
	return
}

func init() {
	gc := &GrpcGenerator{}

	Register("grpc generator", gc)
}