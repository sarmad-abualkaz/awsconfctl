package ssm

import (
	"fmt"
	"log"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

// Declare 'DeleteParam' function 
func DeleteParam(name string, profile string, region string){
	
	// Set AWS session
	sess, sessErr := session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		Profile: profile,
	
		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region: aws.String(region),
		},
	})

	// Check for errors from AWS session setup
	if sessErr != nil {
		fmt.Printf("Error failed to load aws credentials.\n")
		fmt.Println(sessErr.(awserr.Error))
	
	// Successful AWS session setup	
	} else {
	
		// Setting ssm client
		ssmClient := ssm.New(sess)

		// Check if parameter already exists
		isExist, isExisterr := GetParamValue(name, profile, region)

		// Check if parameter already exists	
		if isExist != "" {
			_, deleteErr := ssmClient.DeleteParameter(&ssm.DeleteParameterInput{
				Name: aws.String(name),
			})
			// If delete errors: 
			if deleteErr != nil {
				// Print failed message if the DeleteSecret request failed
				fmt.Printf("Failed to delete secert - %s account(%s): '%s'.\n", profile, region, name)
				// Print error logic 
				if awsErr, ok := deleteErr.(awserr.Error); ok {
					// Get error details
					log.Println("Error:", awsErr.Code(), awsErr.Message())
					// Get original error		
				} else {
					// Print error otherwise
					fmt.Println(deleteErr.Error())
				}				
			} else {
				// Send success message
				fmt.Printf("%s account(%s): '%s' deleted.\n", profile, region, name)
			}		
			// If parameter does not exist	
		} else if isExisterr != nil {
			fmt.Printf("%s account(%s): '%s' does not exist.\n", profile, region, name)	
		}
	}
}
