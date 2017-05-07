package copy

import (
	"fmt"
	"io"
	"os"
)

// Contents copies the contents of the file on path 'from' to file on path 'to'
// creating the destination file, if necessary.
func Contents(from, to string) error {
	src, err := os.OpenFile(from, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed opening source file: %s", err.Error())

	}
	defer src.Close()

	dst, err := os.OpenFile(to, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed creating file: %s", err.Error())
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("failed copying file: %s", err.Error())

	}

	return nil
}
