default: build

clean:
	rm -rf sign-server


deps:
	go get github.com/vpereira/signer
	go get github.com/gorilla/mux
	go get github.com/codegangsta/negroni
	go get code.google.com/p/go-uuid/uuid
	go get golang.org/x/crypto/openpgp/clearsign
	go get github.com/spf13/viper

test-deps:
	go get github.com/stretchr/testify/assert


test:
	go test

format:
	go fmt ./...

build:
	go build -o sign-server

export-keys:
	gpg --export > pubring.gpg
	gpg --export-secret-keys > secring.gpg
