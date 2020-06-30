package secretsmanager

import (
	"fmt"
	"log"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

// Decalre 'DeleteSecret' function
func DeleteSecret(name string, recWindow int64, force bool, profile string, region string){
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
		return
	
	// Successful AWS session setup	
	} else {
		// Setting ssm client
		smClient := secretsmanager.New(sess)

		// Check if parameter already exists
		isExist, isExisterr := GetSecretValue(name, profile, region)

		// Check if parameter already exists	
		if isExist != "" {
			// Check if both force and recovery window are set/required
			if force == true && recWindow > 0 {
				// if true send back error since those two conditions mutually exclusive
				fmt.Printf("Error: cannot have both --recWindow set to more than zero and --deleteForGood/-d set to 'true'\n")
			// Check if only force is set (and recovery window is set to zero) 	
			} else if force == true && recWindow == 0 {
				// Trigger 'DeleteSecret' from from secretsmanager Client (with ForceDeleteWithoutRecovery set to 'true'):
				_, deleteErr := smClient.DeleteSecret(&secretsmanager.DeleteSecretInput{
					ForceDeleteWithoutRecovery: aws.Bool(force),
					SecretId: aws.String(name),
				})
				// If check if 'DeleteSecret' errors out from AWS client
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
					// If error is nil from AWS client send back send success message
					fmt.Printf("%s account(%s): '%s' deleted.\n", profile, region, name)
				}
			// Check if only recover window is set (greater than zero) and force is set to false
			} else if force == false && recWindow > 0 {
				// Trigger 'DeleteSecret' from from secretsmanager Client (with RecoveryWindowInDays setup):
				_, deleteErr := smClient.DeleteSecret(&secretsmanager.DeleteSecretInput{
					SecretId: aws.String(name),
					RecoveryWindowInDays: aws.Int64(recWindow),
				})
				// If check if 'DeleteSecret' errors out from AWS client
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
					// If error is nil from AWS client send back send success message
					fmt.Printf("%s account(%s): '%s' deleted.\n", profile, region, name)
				}
			} else {
				// Error in case entering a 3rd condition
				fmt.Errorf("Error: could not delete %s account(%s): '%s'. Check configurations.", profile, region, name)
				return			
			}
			// Send back error if AWS Secret does not exist
		} else if isExisterr != nil {
			fmt.Printf("%s account(%s): '%s' does not exist.\n", profile, region, name)	
		}		
		
	}
}
