package main

import (
	"bytes"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/clearsign"
	"log"
	"os"
)

type Signer struct {
	Email    string
	Privring openpgp.EntityList
	Entity   *openpgp.Entity
	Path     string
}

func (s *Signer) GetKeyByEmail(keyring openpgp.EntityList, email string) *openpgp.Entity {
	for _, entity := range keyring {
		for _, ident := range entity.Identities {
			if ident.UserId.Email == email {
				return entity
			}
		}
	}
	return nil
}

func (s *Signer) SignIt(sha256 string, uuid string) {
	var buf bytes.Buffer
	w, err := clearsign.Encode(&buf, s.Entity.PrivateKey, nil)

	if err != nil {
		log.Fatal(err)
	}

	_, _ = w.Write([]byte(sha256))
	w.Close()
	ret := buf.Bytes()
	f, err := os.Create(s.Path + uuid)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(string(ret[:]))

}

// privring_path is where you exported your private ring
// with gpg --export-secret-keys > privring.gpg
func NewSigner(email string, privring_path string) Signer {

	signer := Signer{
		Email: email,
		Path:  "signatures/",
	}

	privRingKeyFile, err := os.Open(privring_path)

	if err != nil {
		log.Fatal(err)
	}

	privring, err := openpgp.ReadKeyRing(privRingKeyFile)

	if err != nil {
		log.Fatal(err)
	}

	signer.Entity = signer.GetKeyByEmail(privring, email)

	return signer
}
