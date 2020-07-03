package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"awsconfctl/config"
)

// Decalre flags for root command 
var profile string
var region string
var name string
var contains string
var maxRes int64
var file string
var deleteForGood bool
var recWindow int64

// Declare root command
var rootCmd =&cobra.Command {
	Use: "awsconfctl",
	Short: "Read, create or update configs in AWS SSM Parameter or AWS Secrets.",
	Long: `This creates, updates or delete configurations paremeters in either AWS SSM Parameters or AWS Secrets.`,
}

// Init function
func init(){
	region := config.ConfigEnv.AWS.AWS_REGION
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "dev", "AWS profile name.")	
	rootCmd.PersistentFlags().StringVar(&region, "region", region, "AWS region.")
}

func Execute(){
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)		
	}
}
