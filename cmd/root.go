package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "soc2",
	Short: "The soc2 CLI tool streamlines SOC2 compliance assessments for Ruby on Rails applications deployed on AWS.",
	Long: `Designed for developers and compliance teams, the soc2 CLI tool utilizes the Go programming language and Cobra framework to offer an automated solution for evaluating SOC2 compliance. By assessing key areas such as Security, Availability, Processing Integrity, Confidentiality, and Privacy, soc2 helps ensure that Ruby on Rails applications meet SOC2 standards. This tool simplifies the compliance workflow, providing insights and recommendations to address potential compliance gaps. For example:

soc2 assess`,
	// Updated Run function
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🌟 Welcome to the SOC2 Compliance Checker CLI Tool! 🌟")
		fmt.Println("Use 'soc2 --help' to see available commands and how to start assessing your application's SOC2 compliance.")
		cmd.Help() // This line will display the help text, including available commands.
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
