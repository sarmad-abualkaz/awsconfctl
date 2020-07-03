package cmd

import (
	"fmt"

	"github.com/spf13/cobra"	

	"awsconfctl/aws/secretsmanager"
)

var SMListCMD = &cobra.Command{
	Use: "sm-list-secrets",
	Short: "List Secrets Manger secrets.",
	Long: `Lists of secrets from Secrets Manger with names containing a string --contain <string>.`,
	Run: listSMSecrets,
}

func init(){
	rootCmd.AddCommand(SMListCMD)
	SMListCMD.Flags().StringVar(&contains, "contains", "", "Secret filter regex/string.")
	SMListCMD.Flags().Int64Var(&maxRes, "maxRes", 0, "Max respoonses from listing secrets.")
}

func listSMSecrets(SMListCMD *cobra.Command, args []string){
	profile,_ := rootCmd.Flags().GetString("profile")
	region,_ := rootCmd.Flags().GetString("region")
	stg,_ := SMListCMD.Flags().GetString("contains")
	maxRes,_ := SMListCMD.Flags().GetInt64("maxRes")

	// Check if string from --contains is blank
	if stg != "" {
		// Trigger 'ListSecrets' func from secretsmanager package if --contains is not blank
		secretsmanager.ListSecrets(stg, profile, region, maxRes)
	} else {
		// Error if --contains is blank
		err := fmt.Errorf("Error: contains cannot be blank.")
		fmt.Println(err.Error())		
	}
}
