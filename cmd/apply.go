package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"	
	"awsconfctl/yaml/unmarshal"
	"awsconfctl/aws/ssm"
	"awsconfctl/aws/secretsmanager"
)

var applyCmd = &cobra.Command{
	Use: "apply",
	Short: "Applies changes from a file.",
	Long: `Applies changes to SSM Params or AWS Secret from a file which either creates or updates configs.`,
	Run: applyFile,
}

func init(){
	rootCmd.AddCommand(applyCmd)
}

func applyFile(*cobra.Command, []string){
	file,_ := rootCmd.Flags().GetString("file")
	profile,_ := rootCmd.Flags().GetString("profile")
	region,_ := rootCmd.Flags().GetString("region")

	// Check if file passed is blank	
	if file != "" {
		
		// Check if file exists 
		if _, err := os.Stat(file); err == nil {

			// Enter execution if file exists (and err is nil)			
			fmt.Println("apply called:\n")
			
			// Retrieve the full 'configSetup' using ReadnUnmarshalYaml function:
			yamlRes, yamlResErr := unmarshal.ReadnUnmarshalYaml(file)
			
			// Checking for errors from ReadnUnmarshalYaml
			if yamlResErr != nil {
				fmt.Printf("Error parsing YAML in %s\n", file)
				fmt.Println(yamlResErr.Error())
			
			// Logic for apply:
			} else {
				// Looping through the list configSetup.configType.params:
				for i,v := range yamlRes.ConfigSetup.Params {
					// Check if 'key' exists per each parameter added:
					_, ok := v["key"]
					if ok {
						// Check if 'value' exists per each parameter added:
						_, okVal := v["value"]
						if okVal {
							// Check if the configType is 'systemManager' 
							if yamlRes.ConfigSetup.ConfigType == "systemManager" {
								// Check if 'type' exists per each parameter added:
								_, okType := v["type"]
								if 	okType {
									if v["type"] == "secureString" {
										_, okKeyId := v["KMSKey"]
										if okKeyId {
											// Trigger put parameter with a keyID and using a secureString type ssm
											ssm.PutParam(v["key"], v["value"], strings.Title(v["type"]), v["KMSKey"], profile, region)
										} else {
											// Send back problem on putting parameter without a keyID and using a secureString type ssm
											fmt.Printf("Parameter %v is a secureString but does not contain required 'KMSKey'\n", i)
										}
									} else if v["type"] == "string" {
										// Trigger put parameter without a keyID
										ssm.PutParam(v["key"], v["value"], strings.Title(v["type"]), "",profile, region)
									} else {
										// Send back problem on putting parameter not 'string' or 'secureString'										
										fmt.Printf("Parameter %v has issues, as this command can only take ssm parameter type 'string' or 'secureString'. The 'type: %s' is not recongnized or cannot be handled.\n", i, v["type"])
									}
								} else {
									// Send back problem on not providing 'type'
									fmt.Printf("Parameter %v does not contain required 'type'\n", i)
								}
							// Check if the configType is 'secretsManager' 								
							} else if yamlRes.ConfigSetup.ConfigType == "secretsManager" {
								// Trigger secerts manager creation
								_, okKeyId := v["KMSKey"]
								if okKeyId {
									// Trigger put parameter with a keyID and using a secureString type ssm
									secretsmanager.PutSecret(v["key"], v["value"], v["KMSKey"], profile, region)
								} else {
									// Send back problem on putting parameter without a keyID and using a secureString type ssm
									fmt.Printf("Parameter %v is a secret but does not contain required 'KMSKey'\n", i)
								}								
								// fmt.Println("Secret creation should be triggered here. Its not ready for this command.\n")
							} else {
								// Send back problem on putting parameter without either secureString or systemManager (don't expect to see this msg as 'awsconfctl/yaml/unmarshal' should catch it.)
								fmt.Printf("Something isn't correct, check the value of configType. It should be either 'SystemManager' or 'SecretsManager', currently its set to %s.\n", yamlRes.ConfigSetup.ConfigType)
							}
						} else {
							// Send back problem on putting parameter without providing a 'value'
							fmt.Printf("Parameter %v does not contain required 'value'.\n", i)
						}
					} else {
						// Send back problem on putting parameter without providing a 'key'
						fmt.Printf("Parameter %v does not contain required 'key'.\n", i)
					}
				}
			}
		} else {
			// Error if file does not exist locally
			fmt.Println(err.Error())
		}
	} else {
		// Send back error if flag --file or -f is blank
		err := fmt.Errorf("Error: file is blank.\n")
		fmt.Println(err.Error())
	}
}
