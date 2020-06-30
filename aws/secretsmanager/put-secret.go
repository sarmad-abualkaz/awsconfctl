package secretsmanager

import (
	"fmt"
	"log"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

func PutSecret(name string , value string, kmsKey string , profile string, region string ){

	// Set AWS session
	sess, sessErr := session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		Profile: profile,
		
		// Provide SDK Config options, such as Region.
		Config: aws.Config{
			Region: aws.String(region),
		},
	})

	if sessErr != nil {
		fmt.Printf("Error failed to load aws credentials.\n")
		fmt.Println(sessErr.(awserr.Error))
	}

	// Setting secretsmanager client
	smClient := secretsmanager.New(sess)

	// Decalre secret exist (isExist) variable
	isExist, isExisterr := GetSecretValue(name, profile, region)

	// Check if secret already exists
	if isExist != "" {
		// Check if secret value will change when calling this function
		if isExist == value {
			// Print message that secret will not change
			fmt.Printf("%s account(%s): '%s' unchanged.\n", profile, region, name)
		} else {
			// Create 'updateInputs' to overwrite current value or configurations
			_, updateErr := smClient.UpdateSecret(&secretsmanager.UpdateSecretInput{
				KmsKeyId: aws.String(kmsKey),
				SecretId: aws.String(name),
				SecretString: aws.String(value),
			})
			
			if updateErr != nil {
				// Print failed message if the PutSecret request failed
				fmt.Printf("Failed to update secret - %s account(%s): '%s'.\n", profile, region, name)
				// Print error logic 
				if awsErr, ok := updateErr.(awserr.Error); ok {
					// Get error details
					log.Println("Error:", awsErr.Code(), awsErr.Message())
					// Get original error		
				} else {
					// Print error otherwise
					fmt.Println(updateErr.Error())
				}
			} else {
				// Print success message if err returned nil from PutParameter call
				fmt.Printf("%s account(%s): '%s' configured.\n", profile, region, name)
			}
		}
	} else if isExisterr != nil {
		
		_, createErr := smClient.CreateSecret(&secretsmanager.CreateSecretInput{
			KmsKeyId: aws.String(kmsKey),
			Name: aws.String(name),
			SecretString: aws.String(value),
		})

		if createErr != nil {
			// Print failed message if the PutParameter request failed
			fmt.Printf("Failed to update ssm parameters - %s account(%s): '%s'.\n", profile, region, name)
			// Print error logic 
			if awsErr, ok := createErr.(awserr.Error); ok {
				// Get error details
				log.Println("Error:", awsErr.Code(), awsErr.Message())
				// Get original error		
			} else {
				// Print error otherwise
				fmt.Println(createErr.Error())
			}
		} else {
			// Print success message if err returned nil from PutParameter call
			fmt.Printf("%s account(%s): '%s' created.\n", profile, region, name)
		}					
	} else {
		fmt.Printf("Something went wrong with putting parameter %s account(%s): '%s'.\n", profile, region, name)
	}
}
