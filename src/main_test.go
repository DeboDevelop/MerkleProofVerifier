package main

import (
	"testing"

	"github.com/DeboDevelop/MerkleProofVerifier/hasher"
	"github.com/DeboDevelop/MerkleProofVerifier/tree"
	"github.com/DeboDevelop/MerkleProofVerifier/verifier"
)

func TestSingle(t *testing.T) {
	dataList := []string{"etc", "pi", "chi", "pki", "ro", "gdb", "libnl", "gss", "ldap", "opt", "bare"}
	m := tree.NewMerkleTree(dataList, hasher.SHA256Hasher)
	keyPath := "etc/pi/ro/opt"
	w, err := m.GenWitnessSingleLeaf(keyPath)
	if err != nil {
		t.Error(err)
	}
	c := m.GetCommitment()
	if !verifier.VerifySingleLeaf(c, w, keyPath, hasher.SHA256Hasher) {
		t.Error("Single Leaf Verification vailed! The commitment and derived commitment didn't match.")
	}
}
