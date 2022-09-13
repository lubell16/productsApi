package handlers

import (
	"io"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/lubell16/working/product-images/files"
)

// Files is a handler for reading and writing files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

// NewFiles creates a new File handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

// UploadREST implements the http.Handler interface
func (f *Files) UploadREST(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)
	// check the filepath is a valid name and file
	if id == "" || fn == "" {
		f.invalidUri(r.URL.String(), rw)
		return
	}

	f.saveFile(id, fn, rw, r.Body)
}

// UploadMultipart

func (f *Files) UploadMultipart(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected multipart data form", http.StatusBadRequest)
		return
	}

	id, ideErr := strconv.Atoi(r.FormValue("id"))
	f.log.Info("Process form from id", "id", id)

	if ideErr != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected id value ", http.StatusBadRequest)
		return
	}

	ff, mh, err := r.FormFile("file")
	if err != nil {
		f.log.Error("Bad request", "error", err)
		http.Error(rw, "Expected file ", http.StatusBadRequest)
		return
	}
	f.saveFile(r.FormValue("id"), mh.Filename, rw, ff)
}

func (f *Files) invalidUri(uri string, rw http.ResponseWriter) {
	f.log.Error("Invalid path", "path", uri)
	http.Error(rw, "Invalid file path should be in the format: /[id]/[filepath]", http.StatusBadRequest)
}

//saveFile saves the conteents of the request to a file
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r io.ReadCloser) {
	f.log.Info("Save file for product", "id", id, "path", path)
	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(rw, "unable to save file", http.StatusInternalServerError)
	}
}
