package demo

import (
	"fmt"
	"testing"
)

func TestAzureBlobBase(t *testing.T) {
	containerName := "test"
	if err := AzureBlobBase(containerName); err != nil {
		t.Error(err)
	} else {
		fmt.Println("执行AzureBlobBase成功")
	}
}
