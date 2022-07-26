//go:build integration_tests

package extensiontest

func nftTests(t *testing.T) {
	t.Run("Create NFTs", testCreateNfts)
	t.Run("Remove NFTs", testRemoveNfts)
	t.Run("List NFTs", testListNfts)
}

func testCreateNfts(t *testing.T) {
	a := assert.New(t)
	testEnv.cleanup()

	//todo
}

func testRemoveNfts(t *testing.T) {
	a := assert.New(t)
	testEnv.cleanup()

	//todo
}

func testListNfts(t *testing.T) {
	a := assert.New(t)
	testEnv.cleanup()

	//todo
}
