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
	SSMListCMD.Flags().StringVar(&contains, "contains", "", "SSM Parameter filter regex/string.")
	SSMListCMD.Flags().Int64Var(&maxRes, "maxRes", 0, "Max respoonses from listing SSM Parameters.")
}

func listSSMParams(SSMListCMD *cobra.Command, args []string){
	profile,_ := rootCmd.Flags().GetString("profile")
	region,_ := rootCmd.Flags().GetString("region")
	stg,_ := SSMListCMD.Flags().GetString("contains")
	maxRes,_ := SSMListCMD.Flags().GetInt64("maxRes")

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
