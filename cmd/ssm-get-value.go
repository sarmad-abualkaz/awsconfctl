package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/aws/aws-sdk-go/aws/awserr"

	"awsconfctl/aws/ssm"
)

var SSMValueCMD = &cobra.Command{
	Use: "ssm-get-value",
	Short: "Gets value of an SSM Parameter.",
	Long: `Retreives the value of an SSM Parameter.`,
	Run: getSSMValue,
}

func init(){
	rootCmd.AddCommand(SSMValueCMD)
	SSMValueCMD.Flags().StringVar(&name, "name", "", "SSM Parameter name.")
}

func getSSMValue(SSMValueCMD *cobra.Command, args []string){
	paramName,_ := SSMValueCMD.Flags().GetString("name")
	profile,_ := rootCmd.Flags().GetString("profile")
	region,_ := rootCmd.Flags().GetString("region")

	// Trigger 'GetParamValue' func from ssm package
	ssmVal, valerr := ssm.GetParamValue(paramName, profile, region)
	
	// Check if 'GetParamValue' returned an error
	if valerr != nil {
		fmt.Printf("Failed to retreive list of ssm parameters.\n")
		if awsErr, ok := valerr.(awserr.Error); ok {
			// Get error details
			log.Println("Error:", awsErr.Code(), awsErr.Message())
			// Get original error		
		} else {
			fmt.Println(valerr.Error())
		}
	} else {
		// Display success message otherwise
		fmt.Printf("Value for ssm parameter %s:\n", name)
		fmt.Println(ssmVal)
	}

}
