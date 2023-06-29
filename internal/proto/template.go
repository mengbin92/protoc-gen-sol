package proto

import (
	"os"

	"github.com/pkg/errors"
)

const protoTemplate = `
syntax = "proto3";

package example;

option go_package = "./;example";

message Person {
  string name = 1;
  int32 real_age = 2;
  repeated string email = 3;
}
`

func (p *Proto) New(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return errors.Wrapf(err, "open file: %s error", path)
	}
	defer f.Close()

	f.WriteString(protoTemplate)

	return nil
}
