// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package base58

import (
	"crypto/sha256"
	"errors"
	"strings"
)

// ErrChecksum indicates that the checksum of a check-encoded string does not verify against
// the checksum.
var ErrChecksum = errors.New("checksum error")

// ErrInvalidFormat indicates that the check-encoded string has an invalid format.
var ErrInvalidFormat = errors.New("invalid format: VC and/or checksum bytes missing")

// checksum: first four bytes of sha256^2
func checksum(input []byte) (cksum [4]byte) {
	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])
	copy(cksum[:], h2[:4])
	return
}

// CheckEncode prepends a version byte and appends a four byte checksum.
func CheckEncode(input []byte) string {
	b := make([]byte, 0, 1+len(input)+4)
	b = append(b, 0x01)
	b = append(b, input[:]...)
	cksum := checksum(b)
	b = append(b, cksum[:]...)
	addr := "ABEY" + Encode(b)
	return addr
}

// CheckDecode decodes a string that was encoded with CheckEncode and verifies the checksum.
func CheckDecode(input string) (result []byte, err error) {
	if 0 != strings.Compare("ABEY",input[0:4]) {
		return nil,ErrInvalidFormat
	}
	decoded := Decode(input[4:])
	if len(decoded) < 5 {
		return nil, ErrInvalidFormat
	}
	
	var cksum [4]byte
	copy(cksum[:], decoded[len(decoded)-4:])
	if checksum(decoded[:len(decoded)-4]) != cksum {
		return nil, ErrChecksum
	}
	payload := decoded[:len(decoded)-4]
	result = append(result, payload...)
	return
}
