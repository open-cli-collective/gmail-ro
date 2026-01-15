package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCommand(t *testing.T) {
	t.Run("has correct use", func(t *testing.T) {
		assert.Equal(t, "read <message-id>", readCmd.Use)
	})

	t.Run("requires exactly one argument", func(t *testing.T) {
		err := readCmd.Args(readCmd, []string{})
		assert.Error(t, err)

		err = readCmd.Args(readCmd, []string{"msg123"})
		assert.NoError(t, err)

		err = readCmd.Args(readCmd, []string{"msg1", "msg2"})
		assert.Error(t, err)
	})

	t.Run("has json flag", func(t *testing.T) {
		flag := readCmd.Flags().Lookup("json")
		assert.NotNil(t, flag)
		assert.Equal(t, "j", flag.Shorthand)
		assert.Equal(t, "false", flag.DefValue)
	})

	t.Run("has short description", func(t *testing.T) {
		assert.NotEmpty(t, readCmd.Short)
		assert.Contains(t, readCmd.Short, "message")
	})

	t.Run("long description mentions message ID source", func(t *testing.T) {
		assert.Contains(t, readCmd.Long, "search")
	})
}
