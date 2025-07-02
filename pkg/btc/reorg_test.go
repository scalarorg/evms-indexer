package btc

import (
	"context"
	"testing"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/scalarorg/evms-indexer/pkg/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDBAdapter is a mock implementation of DBAdapter for testing
type MockDBAdapter struct {
	mock.Mock
}

func (m *MockDBAdapter) GetBlockHashByHeight(ctx context.Context, height int64) (*chainhash.Hash, error) {
	args := m.Called(ctx, height)
	return args.Get(0).(*chainhash.Hash), args.Error(1)
}

func (m *MockDBAdapter) GetBlockHeaderByHeight(ctx context.Context, height int64) (*db.BlockHeaderLite, error) {
	args := m.Called(ctx, height)
	return args.Get(0).(*db.BlockHeaderLite), args.Error(1)
}

func (m *MockDBAdapter) DeleteBlockAndTxsFromHeight(ctx context.Context, height int64) error {
	args := m.Called(ctx, height)
	return args.Error(0)
}

func TestDetectAndHandleReorg_NoReorg(t *testing.T) {
	mockDB := &MockDBAdapter{}
	handler := &ReorgHandler{DB: mockDB}

	// Create test data
	prevHash, _ := chainhash.NewHashFromStr("0000000000000000000000000000000000000000000000000000000000000001")
	currentHash, _ := chainhash.NewHashFromStr("0000000000000000000000000000000000000000000000000000000000000002")
	newHash, _ := chainhash.NewHashFromStr("0000000000000000000000000000000000000000000000000000000000000003")

	newBlockHeader := &db.BlockHeaderLite{
		Height:   2,
		Hash:     newHash,
		PrevHash: currentHash,
	}

	tipHeader := &db.BlockHeaderLite{
		Height:   1,
		Hash:     currentHash,
		PrevHash: prevHash,
	}

	// Mock the database call
	mockDB.On("GetBlockHeaderByHeight", mock.Anything, int64(1)).Return(tipHeader, nil)

	// Test
	err := handler.DetectAndHandleReorg(context.Background(), newBlockHeader)

	// Assert
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestDetectAndHandleReorg_GenesisBlock(t *testing.T) {
	mockDB := &MockDBAdapter{}
	handler := &ReorgHandler{DB: mockDB}

	// Create test data for genesis block
	genesisHash, _ := chainhash.NewHashFromStr("0000000000000000000000000000000000000000000000000000000000000001")
	nullHash, _ := chainhash.NewHashFromStr("0000000000000000000000000000000000000000000000000000000000000000")

	newBlockHeader := &db.BlockHeaderLite{
		Height:   0,
		Hash:     genesisHash,
		PrevHash: nullHash,
	}

	// Test
	err := handler.DetectAndHandleReorg(context.Background(), newBlockHeader)

	// Assert
	assert.NoError(t, err)
	// No database calls should be made for genesis block
	mockDB.AssertNotCalled(t, "GetBlockHeaderByHeight")
}
