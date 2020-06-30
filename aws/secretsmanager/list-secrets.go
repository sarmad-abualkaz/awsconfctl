package secretsmanager

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

// Declare 'ListSecrets' function
func ListSecrets(stg string, profile string, region string, maxRes int64){
	// Set AWS session
	sess, err := session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		Profile: profile,
	
		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region: aws.String(region),
		},
	})
	
	// Check for errors from AWS session setup
	if err != nil {
		fmt.Printf("Error failed to load aws credentials.\n")
		fmt.Println(err.(awserr.Error))
		return

	// Successful AWS session setup	
	} else {
		
		// Create a new instance of the service's client with a Session.
		smClient := secretsmanager.New(sess)

		// Declare a listInputs variable (to be used by client)
		var listInputs *secretsmanager.ListSecretsInput

		// If max responses is not set to zero
		if maxRes != 0 {
			// listInputs with MaxResults
			listInputs = &secretsmanager.ListSecretsInput{
				MaxResults: aws.Int64(maxRes),
			}
		} else {
			// if max responses is set to zero - have a listInputs without MaxResults
			listInputs = &secretsmanager.ListSecretsInput{}
		}
		
		// Call the 'ListSecrets' from the secretsmanger client
		secretListOut, secretListErr := smClient.ListSecrets(listInputs)

		// If 'ListSecrets' returns an error
		if secretListErr != nil {
			if awsErr, ok := secretListErr.(awserr.Error); ok {
				// Get error details
				log.Println("Error:", awsErr.Code(), awsErr.Message())
				// Get original error
				if origErr := awsErr.OrigErr(); origErr != nil {
					// operate on original error.
				}
			} else {
				fmt.Println(err.Error())
			}
			// Exit if error
			return
		} else {
			// Success if call to 'ListSecrets' from the secretsmanger client does not return an error
			fmt.Printf("List of ssm secrets containing the string %s:\n", stg)
			// Loop through each element in secretListOut.SecretList and print name
			for _,v := range secretListOut.SecretList {
				if strings.Contains(aws.StringValue(v.Name), stg){
					fmt.Println(aws.StringValue(v.Name))
				}
			}
		}
	}		
}
