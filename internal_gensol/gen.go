package internal_gensol

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func GenerateFile(gen *protogen.Plugin, file *protogen.File) {
	fmt.Printf("filename: %s\n", file.GeneratedFilenamePrefix+".pb.go")
}
