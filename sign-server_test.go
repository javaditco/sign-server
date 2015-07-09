package main

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStub(t *testing.T) {
	assert.True(t, true, "This is good. Canary test passing")
}

func TestUploadedFile(t *testing.T) {
	tu := UploadedFile{Id: uuid.NewUUID(), Name: "Foo", Bytes: 31337}
	assert.NotNil(t, tu, "We are expecting a UploadedFile object")
	assert.Equal(t, "/foo/Foo", tu.GenPath("/foo/"))
}

func TestNewUploadedFile(t *testing.T) {
	tu := NewUploadedFile("README.md", "./", 31337)
	assert.NotNil(t, tu, "We are expecting a UploadedFile object")
}
