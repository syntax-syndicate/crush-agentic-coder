package cmd

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/charmbracelet/x/term"
	"github.com/spf13/cobra"
)

func init() {
	runCmd.Flags().BoolP("quiet", "q", false, "Hide spinner")
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run [prompt...]",
	Short: "Run a single non-interactive prompt",
	Long: `Run a single prompt in non-interactive mode and exit.
The prompt can be provided as arguments or piped from stdin.`,
	Example: `
# Run a simple prompt
crush run Explain the use of context in Go

# Pipe input from stdin
echo "What is this code doing?" | crush run

# Run with quiet mode (no spinner)
crush run -q "Generate a README for this project"
  `,
	RunE: func(cmd *cobra.Command, args []string) error {
		quiet, _ := cmd.Flags().GetBool("quiet")

		app, err := setupApp(cmd)
		if err != nil {
			return err
		}
		defer app.Shutdown()

		prompt := strings.Join(args, " ")

		prompt, err = maybePrependStdin(prompt)
		if err != nil {
			slog.Error("Failed to read from stdin", "error", err)
			return err
		}

		if prompt == "" {
			return fmt.Errorf("no prompt provided")
		}

		// Run non-interactive flow using the App method
		return app.RunNonInteractive(cmd.Context(), prompt, quiet)
	},
}

func maybePrependStdin(prompt string) (string, error) {
	if term.IsTerminal(os.Stdin.Fd()) {
		return prompt, nil
	}
	fi, err := os.Stdin.Stat()
	if err != nil {
		return prompt, err
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		return prompt, nil
	}
	bts, err := io.ReadAll(os.Stdin)
	if err != nil {
		return prompt, err
	}
	return string(bts) + "\n\n" + prompt, nil
}