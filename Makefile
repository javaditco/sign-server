default: build

clean:
	rm -rf sign-server


deps:
	go get github.com/gorilla/mux
	go get github.com/codegangsta/negroni
	go get code.google.com/p/go-uuid/uuid
	go get golang.org/x/crypto/openpgp/clearsign


format:
	go fmt ./...

build:
	go build -o sign-server

export-keys:
	gpg --export > pubring.gpg
	gpg --export-secret-keys > secring.gpg
