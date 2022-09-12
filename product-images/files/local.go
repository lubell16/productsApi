package files

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

// Local is an implementation of the Storage interfadc with works with
// the local disk on the current machine
type Local struct {
	maxFileSize int // maximum number of bytes for files
	basePath    string
}

// NewLocal creates a new Local filesttem with the given base path
// basePath ios the base directory to save files to
// maxSize is the max number of bytes that a file can be
func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	return &Local{basePath: p}, nil
}

// Save the contents of the Wreiter to the given path
// path is a relative path, basePath will be appended
func (l *Local) Save(path string, contents io.Reader) error {
	// get the full path for the file
	fp := l.fullPath(path)

	// get the directory and make sure it exists
	d := filepath.Dir(fp)
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return xerrors.Errorf("Unable to create directory: %w", err)

	}
	// if the file exists delete it
	_, err = os.Stat(fp)
	if err == nil {
		err = os.Remove(fp)
		if err != nil {
			return xerrors.Errorf("Unable to delete file: %w", err)
		}
		// if this is anything other than a not exists error
	} else if !os.IsNotExist(err) {
		return xerrors.Errorf("Unable to get file info: %w", err)
	}
	// create a new file at the path

	f, err := os.Create(fp)
	if err != nil {
		return xerrors.Errorf("Unable to create file: %w", err)
	}
	defer f.Close()

	// write the contents to the new file
	// ensure thyat we are not writing greater than max bytes
	_, err = io.Copy(f, contents)
	if err != nil {
		return xerrors.Errorf("Unable to write to file: %w", err)
	}
	return nil
}

func (l *Local) fullPath(path string) string {
	// append the given path to the base path
	return filepath.Join(l.basePath, path)
}
