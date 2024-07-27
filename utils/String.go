package utils

import "fmt"

type Stringify interface {
	String() string
}

func BoolToString(value bool) string {
	if value {
		return "true"
	}
	return "false"
}

func ByteToString(value byte) string {
	return fmt.Sprint(value)
}

func ShortToString(value int16) string {
	return fmt.Sprint(value)
}

func UShortToString(value uint16) string {
	return fmt.Sprint(value)
}

func IntToString(value int32) string {
	return fmt.Sprint(value)
}

func UIntToString(value uint32) string {
	return fmt.Sprint(value)
}

func ULongToString(value uint64) string {
	return fmt.Sprint(value)
}


func UInt32ArrayToString(value []uint32) string {
	return fmt.Sprint(value)
}
