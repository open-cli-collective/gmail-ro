package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/piekstra/gmail-ro/internal/gmail"
	"github.com/spf13/cobra"
)

var threadJSONOutput bool

func init() {
	rootCmd.AddCommand(threadCmd)
	threadCmd.Flags().BoolVarP(&threadJSONOutput, "json", "j", false, "Output result as JSON")
}

var threadCmd = &cobra.Command{
	Use:   "thread <thread-id>",
	Short: "Read a full conversation thread",
	Long: `Read all messages in a Gmail conversation thread.

The thread ID is the same as the message ID for any message in the thread.
Use the search command to find message/thread IDs.

Examples:
  gmail-ro thread 18abc123def456
  gmail-ro thread 18abc123def456 --json`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client, err := gmail.NewClient(ctx)
		if err != nil {
			return err
		}

		messages, err := client.GetThread(args[0])
		if err != nil {
			return err
		}

		if len(messages) == 0 {
			fmt.Println("No messages found in thread.")
			return nil
		}

		if threadJSONOutput {
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			return enc.Encode(messages)
		}

		fmt.Printf("Thread contains %d message(s)\n\n", len(messages))
		for i, msg := range messages {
			fmt.Printf("=== Message %d of %d ===\n", i+1, len(messages))
			fmt.Printf("ID: %s\n", msg.ID)
			fmt.Printf("From: %s\n", msg.From)
			fmt.Printf("To: %s\n", msg.To)
			fmt.Printf("Subject: %s\n", msg.Subject)
			fmt.Printf("Date: %s\n", msg.Date)
			fmt.Println("\n--- Body ---\n")
			fmt.Println(msg.Body)
			fmt.Println()
		}

		return nil
	},
}
