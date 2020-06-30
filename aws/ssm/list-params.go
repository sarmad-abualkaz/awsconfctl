package ssm

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

// Declare ListParams
func ListParams(stg string, profile string, region string, maxRes int64) {
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
		ssmClient := ssm.New(sess)

		// Declare a describeInputs variable (to be used by client)
		var describeInputs *ssm.DescribeParametersInput

		// Check if max results is not zero
		if maxRes != 0 {
			// Set 'describeInputs' with a 'MaxResults'
			describeInputs = &ssm.DescribeParametersInput{
				ParameterFilters: []*ssm.ParameterStringFilter{
					{
						Key: aws.String("Name"),
						Option: aws.String("Contains"),
						Values: aws.StringSlice([]string{stg}),
					},
				},
				MaxResults: aws.Int64(maxRes),
			}
		// If max results is not zero	
		} else {
			// Set 'describeInputs' without a 'MaxResults'
			describeInputs = &ssm.DescribeParametersInput{
				ParameterFilters: []*ssm.ParameterStringFilter{
					{
						Key: aws.String("Name"),
						Option: aws.String("Contains"),
						Values: aws.StringSlice([]string{stg}),
					},
				},
			}
		}

		// C all the 'DescribeParameters' from ssm client
		paramList, paramListerr := ssmClient.DescribeParameters(describeInputs)
		
		// If 'DescribeParameters' returns an error
		if paramListerr != nil {
			if awsErr, ok := paramListerr.(awserr.Error); ok {
				// Get error details
				log.Println("Error:", awsErr.Code(), awsErr.Message())
				// Get original error
				if origErr := awsErr.OrigErr(); origErr != nil {
					// operate on original error.
				}
			} else {
				fmt.Println(err.Error())
			}
			return
		} else {
			// List ssm parameters if error is nil
			fmt.Printf("List of ssm parameters containing the string %s:\n", stg)
			// Loop through parameters and pring name
			for _,v := range paramList.Parameters {
				fmt.Println(aws.StringValue(v.Name))
			}
		}
	}
}
