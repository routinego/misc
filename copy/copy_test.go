package copy

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const (
	content = "This is some sample content that will be written to the test file."

	srcFileName = "tempFrom.txt"
	dstFileName = "tempTo.txt"
)

var (
	tmp = os.TempDir()

	srcPath = filepath.Join(tmp, srcFileName)
	dstPath = filepath.Join(tmp, dstFileName)
)

func setup() error {
	file, err := os.OpenFile(srcPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 644)
	if err != nil {
		return fmt.Errorf("failed to create file: %s", err.Error())
	}
	defer file.Close()

	_, err = fmt.Fprint(file, content)
	if err != nil {
		return fmt.Errorf("failed writing content to file: %s", err.Error())
	}

	return nil
}

func TestContents(t *testing.T) {
	var err error

	if err = setup(); err != nil {
		t.Errorf("Setup failed: %s", err.Error())
	}
	defer teardown()

	err = Contents(srcPath, dstPath)
	if err != nil {
		t.Errorf("copy.Contents(%q, %q) failed with error: %s", srcPath, dstPath, err.Error())
		return
	}

	if _, err = os.Stat(dstPath); os.IsNotExist(err) {
		t.Errorf("Destination file was not created")
		return
	}

	dstContent, _ := ioutil.ReadFile(dstPath)

	if content != string(dstContent) {
		t.Errorf("Contents of source (%q) and destination (%q) file not the same.", content, string(dstContent))
	}
}

func teardown() {
	os.Remove(srcPath)
	os.Remove(dstPath)
}
