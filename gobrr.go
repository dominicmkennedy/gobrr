package gobrr

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/sys/unix"
)

// CreateEmptyMemfile creates an empty file in memory, and returns an os.File pointer to that file.
// This file acts like any other file on disk, it may be read from and written to.
// Both a valid path and a valid file descriptor may be obtained for this file by using either File.Name(), or File.Fd() respectively.
// When the file is no longer needed it must be closed like any other file, however once it is closed its contents are lost
func CreateEmptyMemfile() (*os.File, error) {
	fd, err := unix.MemfdCreate("", 0)
	if err != nil {
		return nil, fmt.Errorf("MemfdCreate: %v", err)
	}

	fp := fmt.Sprintf("/proc/self/fd/%d", fd)

	return os.NewFile(uintptr(fd), fp), nil
}

// CreateMemfileFromData is equivalent to calling CreateEmptyMemfile then calling File.Write() on the result,
// however this function should theoretically outperform the aforementioned option due to its used of POSIX's mmap
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

// CopyFileToMemfile takes an os.File pointer and copies the data from that file into a memfile
func CopyFileToMemfile(oldFile *os.File) (*os.File, error) {
	oldFileInfo, err := oldFile.Stat()
	if err != nil {
		return nil, fmt.Errorf("Stat: %v", err)
	}

	size := oldFileInfo.Size()
	data := make([]byte, size)

	_, err = oldFile.Read(data)
	if err != io.EOF && err != nil {
		return nil, fmt.Errorf("error reading from: %s: %v", oldFile.Name(), err)
	}

	return CreateMemfileFromData(data)
}

// CopyFilePathToMemfile takes a file path and copies the data from the file at that location into a memfile
func CopyFilePathToMemfile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening: %s: %v", path, err)
	}

	return CopyFileToMemfile(file)
}

// CopyMemfileToBytes is just a speedy way to retrieve the data stored in a memfile
func CopyMemfileToBytes(memfile *os.File) ([]byte, error) {
	memfileInfo, err := memfile.Stat()
	if err != nil {
		return nil, fmt.Errorf("Stat: %v", err)
	}

	fd := memfile.Fd()
	size := memfileInfo.Size()
	data := make([]byte, size)

	filemap, err := unix.Mmap(int(fd), 0, int(size), unix.PROT_READ, unix.MAP_SHARED)
	if err != nil {
		return nil, fmt.Errorf("Mmap: %v", err)
	}

	copy(data, filemap)

	if err := unix.Munmap(filemap); err != nil {
		return nil, fmt.Errorf("Munmap: %v", err)
	}

	return data, nil
}
