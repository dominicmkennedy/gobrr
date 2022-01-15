package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/sys/unix"
)

func CreateEmptyMemfile() (*os.File, error) {
	fd, err := unix.MemfdCreate("", 0)
	if err != nil {
		return nil, fmt.Errorf("MemfdCreate: %v", err)
	}

	fp := fmt.Sprintf("/proc/self/fd/%d", fd)

	return os.NewFile(uintptr(fd), fp), nil
}

func CreateMemfileFromData(b []byte) (*os.File, error) {
	fd, err := unix.MemfdCreate("", 0)
	if err != nil {
		return nil, fmt.Errorf("MemfdCreate: %v", err)
	}

	if err := unix.Ftruncate(fd, int64(len(b))); err != nil {
		return nil, fmt.Errorf("Ftruncate: %v", err)
	}

	filemap, err := unix.Mmap(fd, 0, len(b), unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
	if err != nil {
		return nil, fmt.Errorf("Mmap: %v", err)
	}

	copy(filemap, b)

	if err := unix.Munmap(filemap); err != nil {
		return nil, fmt.Errorf("Munmap: %v", err)
	}

	fp := fmt.Sprintf("/proc/self/fd/%d", fd)

	return os.NewFile(uintptr(fd), fp), nil
}

func CopyFileToMemfile(oldFile *os.File) (*os.File, error) {
	oldFileInfo, err := oldFile.Stat()
	if err != nil {
		return nil, fmt.Errorf("Stat: %v", err)
	}

	size := oldFileInfo.Size()
	data := make([]byte, size)

	bytesRead, err := oldFile.Read(data)
	if err != io.EOF && bytesRead != 0 {
		return nil, fmt.Errorf("error reading from: %s: %v", oldFile.Name(), err)
	}

	return CreateMemfileFromData(data)
}

func CopyFilePathToMemfile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
        return nil, fmt.Errorf("error opening: %s: %v", path, err)
	}

	return CopyFileToMemfile(file)
}
