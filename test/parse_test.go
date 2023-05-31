package test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"google.golang.org/protobuf/proto"
	descriptor "google.golang.org/protobuf/types/descriptorpb"
)

func TestReadProtoFile(t *testing.T) {
	file, err := os.Open("./file_descriptor.bin")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	protoFileBytes, _ := io.ReadAll(file)
	fmt.Printf(".proto:\n%s\n", string(protoFileBytes))

	protoFile := &descriptor.FileDescriptorSet{}
	if err := proto.Unmarshal(protoFileBytes, protoFile); err != nil {
		panic(err.Error())
	}

	for _, file := range protoFile.GetFile() {
		for _, message := range file.GetMessageType() {
			fmt.Printf("消息名称：%s\n", message.GetName())

			for _, field := range message.GetField() {
				fmt.Printf("字段名称：%s，类型：%s\n", field.GetName(), field.GetType().String())
			}
		}
	}

}
