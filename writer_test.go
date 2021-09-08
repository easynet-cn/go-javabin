package javabin

import (
	"os"
	"testing"

	"github.com/google/uuid"
)

func TestUUID(t *testing.T) {
	uuid, err := uuid.Parse("c48d4bee-e998-41b0-b84d-1f29814726b7")

	if err != nil {
		t.Error(err)
	}

	f, err1 := os.Create("./uuid-go.bin")

	if err1 != nil {
		t.Error(err1)
	}

	defer f.Close()

	data := UUIDBytes(uuid, -4856846361193249489)

	f.Write(data)
}

func TestString(t *testing.T) {
	str := "this is a test.这是一个测试。"

	f, err := os.Create("./string-go.bin")

	if err != nil {
		t.Error(err)
	}

	defer f.Close()

	data := StringBytes(str)

	f.Write(data)
}
