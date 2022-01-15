# gobrr

This is an experimental package, and is unfit for production use. gobrr provides an in memory filesystem to any go program.
This filesystem aims to speedup temporary file I/O by storing everything in memory.
The gobrr filesystem will provide all necessary means of file interaction such as `os.File`, `file paths`, and `file descriptors`
gobrr is built as a wrapper around `memfd_create` and as such will only work on linux systems.

This package is largely inspired by [this](https://terinstock.com/post/2018/10/memfd_create-Temporary-in-memory-files-with-Go-and-Linux/) blog post by [Terin Stock](https://terinstock.com/) and I would like to thank him for his work
