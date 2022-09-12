package files

import "io"

//Storage defines the behavior for file operations
//IImplementations may be of the time local disk, clout storage, etc
type Storage interface {
	Save(path string, file io.Reader) error
}
