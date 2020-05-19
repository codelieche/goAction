package demo

import "testing"

func TestDeleteBlobFromAzure(t *testing.T) {
	containerName := "test"
	blobName := "golang.html"

	if err := DeleteBlobFromAzure(containerName, blobName); err != nil {
		t.Error(err)
	} else {
		t.Log("Delete Success")
	}
}
