package publicid

import (
	"fmt"

	skip32 "stretches-common-api/publicid/skip32"

	uuid "github.com/satori/go.uuid"
)

func GeneratePublicIdFromId(id int) string {
	return uuid.NewV1().String()
}
func GeneratePublicId(id int) string {
	return uuid.NewV1().String()
}
func Obfuscate32bit(id uint32) uint32 {
	obfu, err := skip32.New([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA})
	if err != nil {
		fmt.Printf("Key is not 10 bytes length")
	}
	return obfu.Obfus(id)
}
func Unobfuscate32bit(pid uint32) uint32 {
	obfu, err := skip32.New([]byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA})
	if err != nil {
		fmt.Printf("Key is not 10 bytes length")
	}
	return obfu.Unobfus(pid)
}
