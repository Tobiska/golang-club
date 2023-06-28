package main

type Closer interface {
	Close() error
}
