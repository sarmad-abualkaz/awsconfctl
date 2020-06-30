package ssm

import (
	"fmt"
	"log"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/aws/awserr"

)

// Declare 'PutParam' func
func PutParam(name string , value string, ssmType string, kmsKey string , profile string, region string ){
	
	// Declare PutParameterInput struct
	var putInputs *ssm.PutParameterInput

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

		// Decalre parameter exist (isExist) variable
		isExist, isExisterr := GetParamValue(name, profile, region)

		// Check if parameter already exists
		if isExist != "" {
			// Check if parameter value will change when calling this function
			if isExist == value {
				// Print message that parameter will not change
				fmt.Printf("%s account(%s): '%s' unchanged.\n", profile, region, name)
			} else {
				// create a putInputs logic
				if kmsKey == "" {
					// Create 'putInputs' struct with 'Overwrite' == true and no KeyId
					putInputs = &ssm.PutParameterInput{
					Name: aws.String(name),
					Overwrite: aws.Bool(true),
					Type: aws.String(ssmType),
					Value: aws.String(value),
					}

				} else {
					// Create 'putInputs' struct with 'Overwrite' == true and a KeyId
					putInputs = &ssm.PutParameterInput{
						Name: aws.String(name),
						Overwrite: aws.Bool(true),
						Type: aws.String(ssmType),
						Value: aws.String(value),
						KeyId: aws.String(kmsKey),
					}				
				}
				
				// Call ssm.New(sess).PutParameter to overwrite a parameter
				_, err := ssmClient.PutParameter(putInputs)

				// Check if the above request failed with err
				if err != nil {
					// Print failed message if the PutParameter request failed
					fmt.Printf("Failed to update ssm parameters - %s account(%s): '%s'.\n", profile, region, name)
					// Print error logic 
					if awsErr, ok := err.(awserr.Error); ok {
						// Get error details
						log.Println("Error:", awsErr.Code(), awsErr.Message())
						// Get original error		
					} else {
						// Print error otherwise
						fmt.Println(err.Error())
					}
				} else {
					// Print success message if err returned nil from PutParameter call
					fmt.Printf("%s account(%s): '%s' configured.\n", profile, region, name)
				}
			} 
		// Logic for a new parameter
		} else if isExisterr != nil {
			if kmsKey == "" { 
				// Create 'putInputs' struct without 'Overwrite' (defaulting to false) and no KeyID
				putInputs = &ssm.PutParameterInput{
					Name: aws.String(name),
					Type: aws.String(ssmType),
					Value: aws.String(value),
				}
			} else {
				// Create 'putInputs' struct without 'Overwrite' (defaulting to false) and a KeyID
				putInputs = &ssm.PutParameterInput{
					Name: aws.String(name),
					Type: aws.String(ssmType),
					Value: aws.String(value),
					KeyId: aws.String(kmsKey),
				}							
			}
		
			// Call ssm.New(sess).PutParameter for new parameter
			_, errNew := ssmClient.PutParameter(putInputs)

			if errNew != nil {
				// Print failed message if the PutParameter request failed
				fmt.Printf("Failed to update ssm parameters - %s account(%s): '%s'.\n", profile, region, name)
				// Print error logic 
				if awsErr, ok := errNew.(awserr.Error); ok {
					// Get error details
					log.Println("Error:", awsErr.Code(), awsErr.Message())
					// Get original error		
				} else {
					// Print error otherwise
					fmt.Println(errNew.Error())
				}
			} else {
				// Print success message if err returned nil from PutParameter call
				fmt.Printf("%s account(%s): '%s' created.\n", profile, region, name)
			}		
		} else {
			fmt.Printf("Something went wrong with putting parameter %s account(%s): '%s'.\n", profile, region, name)
		}		

	}
}
