package proto

import (
	"fmt"

	"github.com/spf13/cobra"
)

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "New a proto file",
	Long:  "New a proto file. Usage: protoc-gen-solidity new example.proto",
	Run:   runNew,
}

func runNew(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Need a proto file path")
		return
	}

	name := args[1]
	p := &Proto{}
	if err := p.New(name); err != nil {
		fmt.Printf("new proto template error: %s\n", err.Error())
		return
	}
}
