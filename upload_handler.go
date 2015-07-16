package main

import (
	"encoding/json"
	"fmt"
	"github.com/vpereira/signer"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// TODO
// return on erros should be a json as well
func UploadHandler(res http.ResponseWriter, req *http.Request) {
	var (
		status int
		err    error
	)

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
			if outfile, err = os.Create(filepath.Join(Config_map_string["upload_dir"], hdr.Filename)); nil != err {
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

			uploaded_file := NewUploadedFile(hdr.Filename, "uploaded", written)

			signer := signer.NewSigner(Config_map_string["email"], "secring.gpg",
				Config_map_string["sign_dir"])

			go signer.SignIt(uploaded_file.Sha256, fmt.Sprintf("%s", uploaded_file.Sha256))

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
