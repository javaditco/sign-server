package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/vpereira/signer"
	"testing"
)

func TestStub(t *testing.T) {
	assert.True(t, true, "This is good. Canary test passing")
}

func TestUploadedFile(t *testing.T) {
	tu := UploadedFile{Name: "Foo", Bytes: 31337}
	assert.NotNil(t, tu, "We are expecting a UploadedFile object")
	assert.Equal(t, "/foo/Foo", tu.GenPath("/foo/"))
}

func TestNewUploadedFile(t *testing.T) {
	tu := NewUploadedFile("README.md", "./", 31337)
	assert.NotNil(t, tu, "We are expecting a UploadedFile object")
}

func TestNewSigner(t *testing.T) {
	s := signer.NewSigner("tarball-signer@example.org", "secring.gpg", "./signatures")
	assert.NotNil(t, s, "We are expecting a UploadedFile object")
}

func TestSignIt(t *testing.T) {
	sha256 := "f6f24a11d7cbbbc6d9440aca2eba0f6498755ca90adea14c5e233bf4c04bd928"
	s := signer.NewSigner("tarball-signer@example.org", "secring.gpg", "./signatures")
	s.SignIt(sha256, sha256)
}
