package unmarshal

import (
	"errors"
	"strings"
	
	"io/ioutil"
	"gopkg.in/yaml.v3"
)

// Main struct for YamlConfig.ConfigSetup
type YamlConfig struct {
	ConfigSetup struct {
		ConfigType string `yaml:"configType"`
		Params []map[string]string `yaml:"params"`
	} `yaml:"configSetup"`
}

// Reads and Unmarshals Yaml File
func ReadnUnmarshalYaml(file string)(*YamlConfig, error){
	
	// Read file content
	yamlFile, err := ioutil.ReadFile(file)

	// Error if reading content fails 
	if err != nil {
		return &YamlConfig{}, err
	}

	var yamlConfig *YamlConfig

	// Unmarshal Yaml file from content
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	
	// Error if unmarshal Yaml fails 
	if err != nil {
		return &YamlConfig{}, err
	}
	
	// Check if 'ConfigSetup.ConfigType' is blank
	if yamlConfig.ConfigSetup.ConfigType == ""{
	    // Error if 'ConfigSetup.ConfigType' is blank
		return &YamlConfig{}, errors.New("Error: ConfigType is empty.\n")
	// Check if 'ConfigSetup.ConfigType' is either `systemManager` or `secretsManager`
	} else if strings.TrimSpace(yamlConfig.ConfigSetup.ConfigType) == "systemManager" || strings.TrimSpace(yamlConfig.ConfigSetup.ConfigType) == "secretsManager"{
		// Ensure if list of 'ConfigSetup.ConfigType.Params' has elements
		if len(yamlConfig.ConfigSetup.Params) == 0{
			// If 'ConfigSetup.ConfigType.Params' is empty
			return &YamlConfig{}, errors.New("Error: Params is empty.\n")
		} else {
			// Return yamlConfig back if all checks are correct
			return yamlConfig, nil
		}
	} else {
		// Return error if 'ConfigSetup.ConfigType'  is NOT either `systemManager` or `secretsManager`
		return &YamlConfig{}, errors.New("Error: ConfigType must be either 'systemManager' or 'secretsManager'.\n")
	}
}
