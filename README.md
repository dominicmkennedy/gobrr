# gobrr

gobrr is a golang package that will provide an in memory filesystem.
This filesystem aims to speedup temporary file I/O by storing everything in memory.
The gobrr filesystem will provide all necessary means of file interaction such as `os.File`, `file paths`, and `file descriptors`
gobrr is currently only adapted to work on POSIX based systems, and is still experimental please use at your own risk.

This package is largely inspired by [this](https://terinstock.com/post/2018/10/memfd_create-Temporary-in-memory-files-with-Go-and-Linux/) blog post by [Terin Stock](https://terinstock.com/) and I would like to thank him for his work
