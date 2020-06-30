package secretsmanager

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// Decalre 'GetSecretValue' function
func GetSecretValue(name string, profile string, region string)(string, error) {
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
		return "", err
	
	// Successful AWS session setup
	} else {
	
		// Create a new instance of the service's client with a Session.
		smClient := secretsmanager.New(sess)

		// Calling 'GetSecretValue' from secretsmanager client
		val, valerr := smClient.GetSecretValue(&secretsmanager.GetSecretValueInput{
			SecretId: aws.String(name),
		})
		
		// Check if GetSecretValue returned an error
		if valerr != nil {
			// Return blank string and error
			return "", valerr
		} else {
			// Set a string variable for the &secretsmanager.GetSecretValueOutput.'SecretString'
			secretVal := aws.StringValue(val.SecretString)
			// Return 'SecretString' and nil for error
			return secretVal, nil
		}
	}
}
