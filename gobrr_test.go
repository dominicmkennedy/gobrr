package gobrr

import (
	"bytes"
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

    file, err := CreateMemfileFromData(data)
    if err != nil {
        t.Errorf("file creation failed: %s\n", err)
    }

    buf := make([]byte, len(data))

    file.Read(buf)

    if !bytes.Equal(data, buf) {
        t.Errorf("file did not contain the proper contents")
    }
}
