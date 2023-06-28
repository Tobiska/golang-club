package main

type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}
