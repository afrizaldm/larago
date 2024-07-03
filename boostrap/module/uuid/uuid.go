package v7

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UUID struct {
}

func NewUUID() *UUID {
	uuid := &UUID{}
	return uuid
}

func Generate4() string {

	uuid := make([]byte, 16)
	rand.Read(uuid)
	// Setting the version to 4
	uuid[6] = (uuid[6] & 0x0f) | (4 << 4)
	// Setting the variant to RFC 4122
	uuid[8] = (uuid[8] & 0xbf) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", binary.BigEndian.Uint32(uuid[:4]), binary.BigEndian.Uint16(uuid[4:6]), binary.BigEndian.Uint16(uuid[6:8]), binary.BigEndian.Uint16(uuid[8:10]), uuid[10:])
}

func (u *UUID) Generate5() string {

	namespace := uuid.Must(uuid.NewUUID())
	name := "example"
	uuid := uuid.NewSHA1(namespace, []byte(name))
	return uuid.String()
}

func (u *UUID) Generate6() string {

	now := time.Now().UnixMilli()
	uuid := make([]byte, 16)
	binary.BigEndian.PutUint64(uuid[:6], uint64(now))
	rand.Read(uuid[6:])
	// Setting the version to 6
	uuid[6] = (uuid[6] & 0x0f) | (6 << 4)
	// Setting the variant to RFC 4122
	uuid[8] = (uuid[8] & 0xbf) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", binary.BigEndian.Uint32(uuid[:4]), binary.BigEndian.Uint16(uuid[4:6]), binary.BigEndian.Uint16(uuid[6:8]), binary.BigEndian.Uint16(uuid[8:10]), uuid[10:])
}

func (u *UUID) Generate7() string {

	now := time.Now().UnixMilli()
	uuid := make([]byte, 16)
	binary.BigEndian.PutUint64(uuid[:8], uint64(now))
	rand.Read(uuid[8:])
	// Setting the version to 7
	uuid[6] = (uuid[6] & 0x0f) | (7 << 4)
	// Setting the variant to RFC 4122
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", binary.BigEndian.Uint32(uuid[:4]), binary.BigEndian.Uint16(uuid[4:6]), binary.BigEndian.Uint16(uuid[6:8]), binary.BigEndian.Uint16(uuid[8:10]), uuid[10:])
}

func (u *UUID) Generate8() string {

	uuid := make([]byte, 16)
	rand.Read(uuid)
	// Setting the version to 8 (custom)
	uuid[6] = (uuid[6] & 0x0f) | (8 << 4)
	// Setting the variant to RFC 4122
	uuid[8] = (uuid[8] & 0xbf) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", binary.BigEndian.Uint32(uuid[:4]), binary.BigEndian.Uint16(uuid[4:6]), binary.BigEndian.Uint16(uuid[6:8]), binary.BigEndian.Uint16(uuid[8:10]), uuid[10:])
}
