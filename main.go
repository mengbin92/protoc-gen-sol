package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

// var (
// 	newCmd = flag.String("new", "", "new a proto file")
// )

func main() {
	// flag.Parse()
	// if *newCmd != "" {
	// 	p := &proto.Proto{}
	// 	p.New(*newCmd)
	// 	os.Exit(0)
	// }

	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			generateFile(gen, f)
		}
		return nil
	})
}
