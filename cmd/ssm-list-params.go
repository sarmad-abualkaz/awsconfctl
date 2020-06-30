package cmd

import (
	"fmt"

	"github.com/spf13/cobra"	

	"awsconfctl/aws/ssm"
)

var SSMListCMD = &cobra.Command{
	Use: "ssm-list-params",
	Short: "List SSM Parameters.",
	Long: `Lists of SSM Parameters names containing a string --contain <string>.`,
	Run: listSSMParams,
}

func init(){
	rootCmd.AddCommand(SSMListCMD)
}

func listSSMParams(*cobra.Command, []string){
	stg,_ := rootCmd.Flags().GetString("contains")
	profile,_ := rootCmd.Flags().GetString("profile")
	region,_ := rootCmd.Flags().GetString("region")
	maxRes,_ := rootCmd.Flags().GetInt64("maxRes")

	// Check if string from --contains is blank
	if stg != "" {
		// Trigger 'ListParams' func from ssm package
		ssm.ListParams(stg, profile, region, maxRes)
	} else {
		// Error if --contains is blank
		err := fmt.Errorf("Error: contains cannot be blank.")
		fmt.Println(err.Error())		
	}
}
