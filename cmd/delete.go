package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"	
	"awsconfctl/yaml/unmarshal"
	"awsconfctl/aws/ssm"
	"awsconfctl/aws/secretsmanager"
)

var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Deletes changes from a file.",
	Long: `Deletes SSM Params or AWS Secret from a file.`,
	Run: deleteFile,
}

func init(){
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&file, "file", "f", "", "File name.")
	deleteCmd.Flags().BoolVarP(&deleteForGood, "deleteForGood", "d", false, "Boolean for delete command against secrets - secret to be deleted without any recovery window.")
	deleteCmd.Flags().Int64Var(&recWindow, "recWindow", 7, "number of days that Secrets Manager waits before it can delete the secret.")

}

func deleteFile(deleteCmd *cobra.Command, args []string){
	file,_ := deleteCmd.Flags().GetString("file")
	profile,_ := rootCmd.Flags().GetString("profile")
	region,_ := rootCmd.Flags().GetString("region")
	recWindow,_ := deleteCmd.Flags().GetInt64("recWindow")
	deleteForGood,_ := deleteCmd.Flags().GetBool("deleteForGood")

	// Check if file passed is blank
	if file != "" {
		
		// Check if file exists 
		if _, err := os.Stat(file); err == nil {
			
			// Enter execution if file exists (and err is nil)
			fmt.Println("delete called:\n")
			
			// Retrieve the full 'configSetup' using ReadnUnmarshalYaml function:
			yamlRes, yamlResErr := unmarshal.ReadnUnmarshalYaml(file)
			
			// Checking for errors from ReadnUnmarshalYaml
			if yamlResErr != nil {
				fmt.Printf("Error parsing YAML in %s\n", file)
				fmt.Println(yamlResErr.Error())
			
			// Logic for delete:
			} else {
				// Looping through the list configSetup.configType.params:
				for i,v := range yamlRes.ConfigSetup.Params {
					// Check if 'key' exists per each parameter added:
					_, ok := v["key"]
					if ok {
						// Trigger a delete parameters from ssm package
						if yamlRes.ConfigSetup.ConfigType == "systemManager" {
							ssm.DeleteParam(v["key"], profile, region)
						// Trigger a delete parameters from secretsmanger package
						} else if yamlRes.ConfigSetup.ConfigType == "secretsManager" {
							secretsmanager.DeleteSecret(v["key"], recWindow, deleteForGood, profile, region)
						} else {
							// Send back problem on deleting parameter where ConfigType is not secretsManager or systemManager (don't expect to see this msg as 'awsconfctl/yaml/unmarshal' should catch it.)
							fmt.Printf("Something isn't correct, check the value of configType. It should be either 'SystemManager' or 'SecretsManager', currently its set to %s.\n", yamlRes.ConfigSetup.ConfigType)
						}
					} else {
						// Send back problem on deleting parameter without providing a 'key'
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
