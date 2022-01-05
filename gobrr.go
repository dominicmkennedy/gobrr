package gobrr

import (
	"os"
	"fmt"

	"golang.org/x/sys/unix"
)

func CreateEmptyMemfile() (*os.File, error) {
    fd, err := unix.MemfdCreate("", 0)
    if err != nil {
        return nil, fmt.Errorf("MemfdCreate: %s\n", err)
    }

    fp := fmt.Sprintf("/proc/self/fd/%d", fd)

    return os.NewFile(uintptr(fd), fp), nil
}

func CreateMemfileFromData(b []byte) (*os.File, error) {
    fd, err := unix.MemfdCreate("", 0)
    if err != nil {
        return nil, fmt.Errorf("MemfdCreate: %s\n", err)
    }

    if err := unix.Ftruncate(fd, int64(len(b))); err != nil {
        return nil, fmt.Errorf("Ftruncate: %s\n", err)
    }

    filemap, err := unix.Mmap(fd, 0, len(b), unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
    if err != nil {
        return nil, fmt.Errorf("Mmap: %s\n", err)
    }

    copy(filemap, b)

    if err := unix.Munmap(filemap); err != nil {
        return nil, fmt.Errorf("Munmap: %s\n", err)
    }

    fp := fmt.Sprintf("/proc/self/fd/%d", fd)

    return os.NewFile(uintptr(fd), fp), nil
}
