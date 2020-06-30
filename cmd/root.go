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
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "File name.")
	rootCmd.PersistentFlags().StringVar(&name, "name", "", "Config Parameter name.")
	rootCmd.PersistentFlags().StringVar(&contains, "contains", "", "Parameter filter regex/string.")
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "dev", "AWS profile name.")	
	rootCmd.PersistentFlags().StringVar(&region, "region", region, "AWS region.")
	rootCmd.PersistentFlags().Int64Var(&maxRes, "maxRes", 0, "Max respoonses from listing configs.")
	rootCmd.PersistentFlags().BoolVarP(&deleteForGood, "deleteForGood", "d", false, "Boolean for delete command against secrets - secret to be deleted without any recovery window.")
	rootCmd.PersistentFlags().Int64Var(&recWindow, "recWindow", 7, "number of days that Secrets Manager waits before it can delete the secret.")
}

func Execute(){
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)		
	}
}
