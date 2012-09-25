package oci

import (
	"testing"
)

func TestEnvCreation(t *testing.T) {
	var h, e = OCIEnvCreate(OCI_DEFAULT)
	if e != nil {
		t.Error("OCIEnvCreate returned error")
	}

	if OCIHandleFree(h) != nil {
		t.Error("OCIHandleFree returned error")
	}
}
