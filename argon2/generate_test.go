package argon2

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/crypto/argon2"
)

func TestGenerateRandomBytes(t *testing.T) {
	tests := []struct {
		name   string
		length uint32
	}{
		{"Generate 16 bytes", 16},
		{"Generate 32 bytes", 32},
		{"Generate 64 bytes", 64},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := generateRandomBytes(tc.length)
			if err != nil {
				t.Errorf("generateRandomBytes() error = %v", err)
				return
			}
			if len(got) != int(tc.length) {
				t.Errorf("generateRandomBytes() returned %d bytes, want %d", len(got), tc.length)
			}
		})
	}
}

func TestGenerateFromPlainText(t *testing.T) {
	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	password := "testpassword"

	encodedHash, err := generateFromPlainText(password, p)
	if err != nil {
		t.Fatalf("generateFromPlainText() error = %v", err)
	}

	// Check the format of the encoded hash
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		t.Errorf("Encoded hash has incorrect number of parts: got %d, want 6", len(parts))
	}

	if parts[1] != "argon2id" {
		t.Errorf("Incorrect algorithm identifier: got %s, want argon2id", parts[1])
	}

	var version int
	_, err = fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		t.Errorf("Error parsing version: %v", err)
	}
	if version != argon2.Version {
		t.Errorf("Incorrect version: got %d, want %d", version, argon2.Version)
	}

	var memory, iterations, parallelism int
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil {
		t.Errorf("Error parsing parameters: %v", err)
	}
	if uint32(memory) != p.memory || uint32(iterations) != p.iterations || uint8(parallelism) != p.parallelism {
		t.Errorf("Incorrect parameters: got m=%d,t=%d,p=%d, want m=%d,t=%d,p=%d",
			memory, iterations, parallelism, p.memory, p.iterations, p.parallelism)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		t.Errorf("Error decoding salt: %v", err)
	}
	if len(salt) != int(p.saltLength) {
		t.Errorf("Incorrect salt length: got %d, want %d", len(salt), p.saltLength)
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		t.Errorf("Error decoding hash: %v", err)
	}
	if len(hash) != int(p.keyLength) {
		t.Errorf("Incorrect hash length: got %d, want %d", len(hash), p.keyLength)
	}
}

func TestGenerate(t *testing.T) {
	password := "testpassword"
	encodedHash := Generate(password)

	// Verify that the generated hash can be successfully compared
	match, err := Compare(password, encodedHash)
	if err != nil {
		t.Fatalf("Error comparing password: %v", err)
	}
	if !match {
		t.Errorf("Generated hash does not match original password")
	}

	// Verify that a different password doesn't match
	match, err = Compare("wrongpassword", encodedHash)
	if err != nil {
		t.Fatalf("Error comparing password: %v", err)
	}
	if match {
		t.Errorf("Generated hash incorrectly matches wrong password")
	}
}
