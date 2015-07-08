package main

import (
	"code.google.com/p/go-uuid/uuid"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	//"golang.org/x/crypto/openpgp/clearsign"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	// "strconv"
)

type UploadedFile struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Path   string    `json:"path"`
	Bytes  int64     `json:"bytes"`
	Sha256 string    `json:"sha256"`
	Url    string    `json:"url"`
}

func (u *UploadedFile) GenPath(path string) string {
	// TODO
	// defensive programming
	return "/file/" + path + u.Name
}

func (u *UploadedFile) GetSha256(path string) string {
	// TODO
	// implement it :)
	hasher := sha256.New()
	s, err := ioutil.ReadFile(path + u.Name)
	hasher.Write(s)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

func (u *UploadedFile) GenURL() string {
	return fmt.Sprintf("http://localhost:8080/file/%s", u.Id)
}
func NewUploadedFile(fileName string, path string, bytes int64) UploadedFile {

	u := UploadedFile{
		Name:  fileName,
		Bytes: bytes,
	}
	u.Id = uuid.NewUUID()
	u.Path = u.GenPath(path)
	u.Sha256 = u.GetSha256(path)
	u.Url = u.GenURL()
	return u
}

func UploadHandler(res http.ResponseWriter, req *http.Request) {

	var (
		status int
		err    error
	)

	// add defer for open files

	defer func() {
		if nil != err {
			http.Error(res, err.Error(), status)
		}
	}()

	// parse request
	const _24K = (1 << 20) * 24
	if err = req.ParseMultipartForm(_24K); nil != err {
		status = http.StatusInternalServerError
		return
	}

	for _, fheaders := range req.MultipartForm.File {
		for _, hdr := range fheaders {

			// open uploaded
			var infile multipart.File
			if infile, err = hdr.Open(); nil != err {
				status = http.StatusInternalServerError
				return
			}

			// open destination
			var outfile *os.File
			if outfile, err = os.Create("./uploaded/" + hdr.Filename); nil != err {
				status = http.StatusInternalServerError
				return
			}

      defer outfile.Close()

			// 32K buffer copy
			var written int64
			if written, err = io.Copy(outfile, infile); nil != err {
				status = http.StatusInternalServerError
				return
			}

			uploaded_file := NewUploadedFile(hdr.Filename, "uploaded/", written)

      signer := NewSigner("tarball-signer@example.org","secring.gpg")

      go signer.SignIt(fmt.Sprintf("uploaded/%s",hdr.Filename), fmt.Sprintf("%s",uploaded_file.Id))

  		js, err := json.Marshal(uploaded_file)

			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.Write(js)

		}
	}

}
