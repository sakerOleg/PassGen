package cmd

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "passgen",
	Short: "Generate passwords",
	Long: `Generate random passwords with customizable options.

	For example:
passgen -l 8 -d -s`,

	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		isDigits, _ := cmd.Flags().GetBool("digits")
		isSpecChars, _ := cmd.Flags().GetBool("specChars")
		isRepeatable, _ := cmd.Flags().GetInt("repeat")

		charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

		if isDigits {
			charset += "0123456789"
		}

		if isSpecChars {
			charset += "!@#$%^&*()_+[]{}|;:,.<>?-="
		}

		password := make([]byte, length)

		fmt.Print("Generated passwords:\n\n")

		for i := 0; i < isRepeatable; i++ {
			for i := range password {
				password[i] = charset[rand.Intn(len(charset))]
			}
			fmt.Printf("%s\n\n", password)
		}

		fmt.Print("\n")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("length", "l", 6, "Length of the generated password")
	rootCmd.Flags().IntP("repeat", "r", 1, "Number of the generated passwords")
	rootCmd.Flags().BoolP("digits", "d", false, "Include digits in the generated password")
	rootCmd.Flags().BoolP("specChars", "s", false, "Include special characters in the generated password")
}
