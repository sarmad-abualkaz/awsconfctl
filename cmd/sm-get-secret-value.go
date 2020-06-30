package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/aws/aws-sdk-go/aws/awserr"

	"awsconfctl/aws/secretsmanager"
)

var SMValueCMD = &cobra.Command{
	Use: "sm-get-secret-value",
	Short: "Gets value of an Secrets Manager value.",
	Long: `Retreives the value of an Secrets Manager secret.`,
	Run: getSMValue,
}

func init(){
	rootCmd.AddCommand(SMValueCMD)
}

func getSMValue(*cobra.Command, []string){
	secretName,_ := rootCmd.Flags().GetString("name")
	profile,_ := rootCmd.Flags().GetString("profile")
	region,_ := rootCmd.Flags().GetString("region")
	
	// Trigger 'GetSecretValue' func from secretsmanager package
	secretVal, valerr := secretsmanager.GetSecretValue(secretName, profile, region)
	
	// Check if 'GetSecretValue' returned error
	if valerr != nil {
		fmt.Printf("Failed to retreive list of secret from Secrets Manager.\n")
		if awsErr, ok := valerr.(awserr.Error); ok {
			// Get error details
			log.Println("Error:", awsErr.Code(), awsErr.Message())
			// Get original error		
		} else {
			fmt.Println(valerr.Error())
		}
	} else {
		// Display success message otherwise
		fmt.Printf("Value for secret %s:\n", name)
		fmt.Println(secretVal)
	}

}
