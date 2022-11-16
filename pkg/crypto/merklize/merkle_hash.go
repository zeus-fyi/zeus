package merklize

import (
	"crypto/sha256"

	"github.com/cbergoon/merkletree"
)

type MerkleTree struct {
	*merkletree.MerkleTree
}

type MerkleTreeContent struct {
	Value string
}

// CalculateHash hashes the values of a TestContent
func (t MerkleTreeContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.Value)); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (t MerkleTreeContent) Equals(other merkletree.Content) (bool, error) {
	return t.Value == other.(MerkleTreeContent).Value, nil
}

func NewMerkleTree(list []merkletree.Content) *merkletree.MerkleTree {
	t, err := merkletree.NewTree(list)
	if err != nil {
		panic(err)
	}
	return t
}
