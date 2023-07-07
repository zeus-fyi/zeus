package merklize

import (
	"testing"

	"github.com/cbergoon/merkletree"
	"github.com/stretchr/testify/suite"
)

type MerkleTreeTestSuite struct {
	suite.Suite
}

func (s *MerkleTreeTestSuite) TestMerklizeList() {
	item1 := MerkleTreeContent{Value: "item1"}
	item2 := MerkleTreeContent{Value: "item2"}
	var list []merkletree.Content
	list = append(list, item1)
	list = append(list, item2)

	t := NewMerkleTree(list)

	verified, err := t.VerifyTree()
	s.Assert().Nil(err)
	s.Assert().True(verified)
}

func TestMerkleTreeTestSuite(t *testing.T) {
	suite.Run(t, new(MerkleTreeTestSuite))
}
