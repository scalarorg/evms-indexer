package btc

import (
	"fmt"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

// MerkleProof represents the proof path for a transaction
type MerkleProof struct {
	Position int               `json:"position"`
	Path     []*chainhash.Hash `json:"path"`
	Root     *chainhash.Hash   `json:"root"`
}

func (m *MerkleProof) ConcatProofPath() []byte {
	path := make([]byte, 0)
	for _, hash := range m.Path {
		path = append(path, hash.CloneBytes()...)
	}
	return path
}

// calculateMerkleProofPath calculates the Merkle proof path for a transaction at the given position
func calculateMerkleProofPath(txs []btcjson.TxRawResult, txPosition int) (*MerkleProof, error) {
	if txPosition < 0 || txPosition >= len(txs) {
		return nil, fmt.Errorf("invalid transaction position: %d, total transactions: %d", txPosition, len(txs))
	}

	if len(txs) == 0 {
		return nil, fmt.Errorf("no transactions provided")
	}

	// Convert transactions to hashes
	hashes := make([]*chainhash.Hash, len(txs))
	for i, tx := range txs {
		hash, err := chainhash.NewHashFromStr(tx.Txid)
		if err != nil {
			return nil, fmt.Errorf("failed to convert transaction hash to chainhash: %w", err)
		}
		hashes[i] = hash
	}

	// Build the Merkle tree and get the proof path
	proof, root := createMerkleBranchAndRoot(hashes, txPosition)

	return &MerkleProof{
		Position: txPosition,
		Path:     proof,
		Root:     root,
	}, nil
}

// createMerkleBranchAndRoot creates a Merkle branch and root for the given hashes and index
// This is a direct translation of the Rust function
func createMerkleBranchAndRoot(hashes []*chainhash.Hash, index int) ([]*chainhash.Hash, *chainhash.Hash) {
	// Create a copy of hashes to avoid modifying the original
	workingHashes := make([]*chainhash.Hash, len(hashes))
	copy(workingHashes, hashes)

	workingIndex := index
	merkle := make([]*chainhash.Hash, 0)

	// Continue until we have only one hash left (the root)
	for len(workingHashes) > 1 {
		// If odd number of hashes, duplicate the last one
		if len(workingHashes)%2 != 0 {
			last := workingHashes[len(workingHashes)-1]
			workingHashes = append(workingHashes, last)
		}

		// Calculate the sibling index
		if workingIndex%2 == 0 {
			workingIndex = workingIndex + 1
		} else {
			workingIndex = workingIndex - 1
		}

		// Add the sibling to the merkle branch
		merkle = append(merkle, workingHashes[workingIndex])

		// Move to the parent level
		workingIndex = workingIndex / 2

		// Create the next level of hashes by pairing and hashing
		nextLevel := make([]*chainhash.Hash, 0)
		for i := 0; i < len(workingHashes); i += 2 {
			left := workingHashes[i]
			right := workingHashes[i+1]
			combined := append(left[:], right[:]...)
			hashBytes := chainhash.DoubleHashB(combined)
			hash, _ := chainhash.NewHash(hashBytes)
			nextLevel = append(nextLevel, hash)
		}
		workingHashes = nextLevel
	}

	return merkle, workingHashes[0]
}

// VerifyMerkleProof verifies a Merkle proof for a transaction
func VerifyMerkleProof(txHash *chainhash.Hash, proof *MerkleProof) bool {
	if len(proof.Path) == 0 {
		// Single transaction, just verify the root
		return txHash.IsEqual(proof.Root)
	}

	// Start with the transaction hash
	currentHash := txHash

	// Follow the proof path
	for i, siblingHash := range proof.Path {
		// Determine if we're the left or right child
		isLeft := (proof.Position>>i)&1 == 0

		if isLeft {
			// We're on the left, sibling is on the right
			combined := append(currentHash[:], siblingHash[:]...)
			hashBytes := chainhash.DoubleHashB(combined)
			currentHash, _ = chainhash.NewHash(hashBytes)
		} else {
			// We're on the right, sibling is on the left
			combined := append(siblingHash[:], currentHash[:]...)
			hashBytes := chainhash.DoubleHashB(combined)
			currentHash, _ = chainhash.NewHash(hashBytes)
		}
	}

	// The final hash should match the root
	return currentHash.IsEqual(proof.Root)
}
