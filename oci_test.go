package oci

import (
	"testing"
)

func TestEnvCreation(t *testing.T) {
	var env, ctx *OCIHandle
	var e error

	env, e = OCIEnvCreate(OCI_DEFAULT)
	if e != nil {
		t.Error("OCIEnvCreate returned error")
	}

	ctx, e = OCILogon(env, "goci", "goci", "GOCI")
	if e != nil {
		t.Error("OCILogon failed")
	}

	if OCIHandleFree(ctx) != nil {
		t.Error("OCIHandleFree returned error")
	}

	if OCIHandleFree(env) != nil {
		t.Error("OCIHandleFree returned error")
	}
}
