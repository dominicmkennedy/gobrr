package gobrr

import (
	"bytes"
	"io"
	"testing"
)

func TestCreateEmptyMemfile(t *testing.T) {
	file, err := CreateEmptyMemfile()
	if err != nil {
		t.Errorf("file creation failed: %s\n", err)
	}

	stat, err := file.Stat()
	if err != nil {
		t.Errorf("could not retrive file stats: %s\n", err)
	}

	if stat.Size() != 0 {
		t.Errorf("file size is not zero\n")
	}
}

func TestCreateMemfileFromData(t *testing.T) {
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	buf := make([]byte, len(data))

	file, err := CreateMemfileFromData(data)
	if err != nil {
		t.Errorf("file creation failed: %s\n", err)
	}

	if _, err := file.Read(buf); err != io.EOF && err != nil {
		t.Errorf("error reading file: %v\n", err)
	}

	if !bytes.Equal(data, buf) {
		t.Errorf("file did not contain the proper contents\nexpected: %v\nreceived: %v\n", data, buf)
	}
}

func TestCopyFileToMemfile(t *testing.T) {
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	buf := make([]byte, len(data))

	oldFile, err := CreateMemfileFromData(data)
	if err != nil {
		t.Errorf("file creation failed: %s\n", err)
	}

	file, err := CopyFileToMemfile(oldFile)
	if err != nil {
		t.Errorf("error copying to memfile: %v\n", err)
	}

	if _, err := file.Read(buf); err != io.EOF && err != nil {
		t.Errorf("error reading file: %v\n", err)
	}

	if !bytes.Equal(data, buf) {
		t.Errorf("file did not contain the proper contents\nexpected: %v\nreceived: %v\n", data, buf)
	}
}

//TODO
func TestCopyFilePathToMemfile(t *testing.T) {
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	buf := make([]byte, len(data))

	oldFile, err := CreateMemfileFromData(data)
	if err != nil {
		t.Errorf("file creation failed: %s\n", err)
	}

	file, err := CopyFilePathToMemfile(oldFile.Name())
	if err != nil {
		t.Errorf("error copying to memfile: %v\n", err)
	}

	if _, err := file.Read(buf); err != io.EOF && err != nil {
		t.Errorf("error reading file: %v\n", err)
	}

	if !bytes.Equal(data, buf) {
		t.Errorf("file did not contain the proper contents\nexpected: %v\nreceived: %v\n", data, buf)
	}
}

func TestCopyMemfileToBytes(t *testing.T) {
	data := []byte("abcdefghijklmnopqrstuvwxyz")

	file, err := CreateMemfileFromData(data)
	if err != nil {
		t.Errorf("file creation failed: %s\n", err)
	}

	buf, err := CopyMemfileToBytes(file)
	if err != nil {
		t.Errorf("error copying mem file to bytes: %v\n", err)
	}

	if !bytes.Equal(data, buf) {
		t.Errorf("file did not contain the proper contents\nexpected: %v\nreceived: %v\n", data, buf)
	}
}
