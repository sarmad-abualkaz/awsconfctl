package ssm

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Declare 'GetParamValue' function 
func GetParamValue(name string, profile string, region string)(string, error) {
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
		ssmClient := ssm.New(sess)

		// Calling 'GetParameter' from ssm client
		val, valerr := ssmClient.GetParameter(&ssm.GetParameterInput{
			Name: aws.String(name),
			WithDecryption: aws.Bool(true),
		})
		
		// Check if GetParameter returned an error from ssm client
		if valerr != nil {
			// Return empty string and the aws error
			return "", valerr
		} else {
			// Set a string variable for the &ssm.GetParameterOutput.'Parameter.Value'
			ssmVal := aws.StringValue(val.Parameter.Value)
			// Return 'Parameter.Value' and nil for error
			return ssmVal, nil
		}
	}
}
