package main

type ByteReader interface {
	ReadByte() (c byte, err error)
}
type ByteWriter interface {
	WriteByte(c byte) error
}

type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}
type RuneScanner interface {
	RuneReader
	UnreadRune() error
}
