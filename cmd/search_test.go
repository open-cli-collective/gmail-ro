package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchCommand(t *testing.T) {
	t.Run("has correct use", func(t *testing.T) {
		assert.Equal(t, "search <query>", searchCmd.Use)
	})

	t.Run("requires exactly one argument", func(t *testing.T) {
		err := searchCmd.Args(searchCmd, []string{})
		assert.Error(t, err)

		err = searchCmd.Args(searchCmd, []string{"query"})
		assert.NoError(t, err)

		err = searchCmd.Args(searchCmd, []string{"query1", "query2"})
		assert.Error(t, err)
	})

	t.Run("has max flag", func(t *testing.T) {
		flag := searchCmd.Flags().Lookup("max")
		assert.NotNil(t, flag)
		assert.Equal(t, "m", flag.Shorthand)
		assert.Equal(t, "10", flag.DefValue)
	})

	t.Run("has json flag", func(t *testing.T) {
		flag := searchCmd.Flags().Lookup("json")
		assert.NotNil(t, flag)
		assert.Equal(t, "j", flag.Shorthand)
		assert.Equal(t, "false", flag.DefValue)
	})

	t.Run("has examples in long description", func(t *testing.T) {
		assert.Contains(t, searchCmd.Long, "from:")
		assert.Contains(t, searchCmd.Long, "subject:")
		assert.Contains(t, searchCmd.Long, "is:unread")
	})
}
