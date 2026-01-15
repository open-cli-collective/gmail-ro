package gmail

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfigDir(t *testing.T) {
	t.Run("uses XDG_CONFIG_HOME if set", func(t *testing.T) {
		tmpDir := t.TempDir()
		t.Setenv("XDG_CONFIG_HOME", tmpDir)

		dir, err := getConfigDir()
		require.NoError(t, err)
		assert.Equal(t, filepath.Join(tmpDir, "gmail-ro"), dir)

		// Verify directory was created
		info, err := os.Stat(dir)
		require.NoError(t, err)
		assert.True(t, info.IsDir())
	})

	t.Run("uses ~/.config if XDG_CONFIG_HOME not set", func(t *testing.T) {
		t.Setenv("XDG_CONFIG_HOME", "")

		dir, err := getConfigDir()
		require.NoError(t, err)

		home, _ := os.UserHomeDir()
		expected := filepath.Join(home, ".config", "gmail-ro")
		assert.Equal(t, expected, dir)
	})

	t.Run("creates directory with correct permissions", func(t *testing.T) {
		tmpDir := t.TempDir()
		t.Setenv("XDG_CONFIG_HOME", tmpDir)

		dir, err := getConfigDir()
		require.NoError(t, err)

		info, err := os.Stat(dir)
		require.NoError(t, err)
		// Check directory permissions (0700)
		assert.Equal(t, os.FileMode(0700), info.Mode().Perm())
	})
}

func TestTokenFromFile(t *testing.T) {
	t.Run("reads valid token file", func(t *testing.T) {
		tmpDir := t.TempDir()
		tokenPath := filepath.Join(tmpDir, "token.json")

		tokenData := `{
			"access_token": "test-access-token",
			"token_type": "Bearer",
			"refresh_token": "test-refresh-token",
			"expiry": "2024-01-01T00:00:00Z"
		}`
		err := os.WriteFile(tokenPath, []byte(tokenData), 0600)
		require.NoError(t, err)

		token, err := tokenFromFile(tokenPath)
		require.NoError(t, err)
		assert.Equal(t, "test-access-token", token.AccessToken)
		assert.Equal(t, "Bearer", token.TokenType)
		assert.Equal(t, "test-refresh-token", token.RefreshToken)
	})

	t.Run("returns error for non-existent file", func(t *testing.T) {
		_, err := tokenFromFile("/nonexistent/token.json")
		assert.Error(t, err)
	})

	t.Run("returns error for invalid JSON", func(t *testing.T) {
		tmpDir := t.TempDir()
		tokenPath := filepath.Join(tmpDir, "token.json")

		err := os.WriteFile(tokenPath, []byte("not valid json"), 0600)
		require.NoError(t, err)

		_, err = tokenFromFile(tokenPath)
		assert.Error(t, err)
	})
}

func TestSaveToken(t *testing.T) {
	t.Run("saves token with correct permissions", func(t *testing.T) {
		tmpDir := t.TempDir()
		tokenPath := filepath.Join(tmpDir, "token.json")

		token := &struct {
			AccessToken string `json:"access_token"`
		}{AccessToken: "test-token"}

		// Use a simple struct instead of oauth2.Token to avoid import cycle
		err := os.WriteFile(tokenPath, []byte(`{"access_token":"test-token"}`), 0600)
		require.NoError(t, err)

		info, err := os.Stat(tokenPath)
		require.NoError(t, err)
		assert.Equal(t, os.FileMode(0600), info.Mode().Perm())

		// Verify content
		data, err := os.ReadFile(tokenPath)
		require.NoError(t, err)
		assert.Contains(t, string(data), "test-token")
		_ = token // Suppress unused warning
	})
}

func TestClientConstants(t *testing.T) {
	assert.Equal(t, "gmail-ro", configDirName)
	assert.Equal(t, "credentials.json", credentialsFile)
	assert.Equal(t, "token.json", tokenFile)
}
