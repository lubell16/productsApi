package files

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupLocal(t *testing.T) (*Local, string, func()) {
	// create a temporary directory
	dir, err := os.MkdirTemp("", "files")
	if err != nil {
		t.Fatal(err)
	}

	l, err := NewLocal(dir, 1024*1000*5)
	if err != nil {
		t.Fatal(err)
	}

	return l, dir, func() {
		// cleanup function
		//os.RemoveAll(dir)
	}

}

func TestSavesContentsOfReader(t *testing.T) {
	savePath := "/1/test.png"
	fileContents := "Hello World"
	l, dir, cleanup := setupLocal(t)
	defer cleanup()

	err := l.Save(savePath, bytes.NewBuffer([]byte(fileContents)))
	assert.NoError(t, err)

	// check the file has been correctly written
	f, err := os.Open(filepath.Join(dir, savePath))
	assert.NoError(t, err)

	// check the contents of the file
	d, err := ioutil.ReadAll(f)
	assert.NoError(t, err)
	assert.Equal(t, fileContents, string(d))
}
