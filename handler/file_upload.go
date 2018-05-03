package handler

import (
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
)

func SetUploadDir(uploadDir string) BoostedMiddleWare {
	return func(next boostedHandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			next(w, r, uploadDir)
		}
	}
}

func FileUpload(w http.ResponseWriter, r *http.Request, argv ... interface{}) {
	uploadDir := argv[0].(string)

	var buffer bytes.Buffer
	file, header, err := r.FormFile("file")
	if err != nil {
		writeError(w, err)
		return
	}
	defer file.Close()

	io.Copy(&buffer, file)

	err = ioutil.WriteFile(uploadDir+"/"+header.Filename, buffer.Bytes(), 0644)
	if err != nil {
		writeError(w, err)
		return
	}
	buffer.Reset()
}
