package utils

type ByteReader interface {
	ReadByte() (byte, error)
}
