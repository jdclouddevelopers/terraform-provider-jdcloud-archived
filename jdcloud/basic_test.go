package jdcloud

import (
	"log"
	"os"
	"testing"
)

// This step is necessary since we need to pass the
// Secret key and public key to begin our testing
func testAccPreCheck(t *testing.T) {
	if accessKey := os.Getenv("access_key"); accessKey == "" {
		t.Fatalf("parameter : access_key must be set to complete testing")
	}
	if secretKey := os.Getenv("secret_key"); secretKey == "" {
		t.Fatalf("parameter : secret_key must be set to complete testing")
	}
	if regionID := os.Getenv("region"); regionID == "" {
		log.Println("region was not set, now they were set to cn-north-1")
		os.Setenv("region", "cn-north-1")
	}
}
