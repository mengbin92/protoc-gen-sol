package main

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

// generateFile generates a _http.pb.go file containing kratos errors definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + ".pb.sol"
	if file.GoImportPath == "" {
		file.GoImportPath = protogen.GoImportPath(fmt.Sprintf(`option go_package = "./;%s";`, file.GeneratedFilenamePrefix))
	}
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	g.P()

	return g
}
