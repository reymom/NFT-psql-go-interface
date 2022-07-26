//go:build integration_tests

package extensiontest

func userTests(t *testing.T) {
	t.Run("Manage Users", testManageUsers)
}

func testManageUsers(t *testing.T) {
	a := assert.New(t)
	testEnv.cleanup()

	//todo
}
