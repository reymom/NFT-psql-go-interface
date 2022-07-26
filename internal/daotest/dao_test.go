//go:build integration_tests

package extensiontest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testEnv *testEnvironment

const connectionString = "postgresql://nft_store:nft_store_pw@localhost:5432/nft_store"

func TestSetup(t *testing.T) {
	t.Run("Setup NFT Store Connection", extensionTest)
	t.Run("Users", userTests)
	t.Run("NFTs", nftTests)
}

func extensionTest(t *testing.T) {
	a := assert.New(t)
	var e error
	testEnv, e = newTestEnvironment(connectionString, connectionString, t)
	a.NoError(e)

	t.Cleanup(testEnv.cleanup)
}
