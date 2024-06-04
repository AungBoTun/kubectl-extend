package api

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"encoding/json"
)
type Config struct{
	APIURL string `yaml:"apiURL"`
}
type APIResponse struct {
	Conditions []struct {
		LastTransitionTime time.Time `json:"lastTransitionTime"`
		ObservedGeneration int       `json:"observedGeneration"`
		Message            string    `json:"message"`
		Type               string    `json:"type"`
		Status             string    `json:"status"`
		Reason             string    `json:"reason"`
	} `json:"conditions"`
	ComponentsStatus []struct {
		LastTransitionTime time.Time `json:"lastTransitionTime"`
		Ready              bool      `json:"ready"`
		State              string    `json:"state"`
		ResourceKind       string    `json:"resourceKind"`
		ResourceName       string    `json:"resourceName"`
		ResourceNamespace  string    `json:"resourceNamespace"`
	} `json:"componentsStatus"`
	ResourceInCluster bool `json:"resourceInCluster"`
}
func ReadConfig(configPath string)(Config,error) {
	var config Config
	yamlFile,err := os.Open(configPath)

	if err !=nil {
		return config, fmt.Errorf("error opening YAML file: %w",err)
	}
	defer yamlFile.Close()
	byteValue, err := ioutil.ReadAll(yamlFile)
	if err !=nil {
		return config, fmt.Errorf("error opening YAML file: %w",err)
	}
	err = yaml.Unmarshal(byteValue,&config)
	if err !=nil {
		return config, fmt.Errorf("error opening YAML file: %w",err)
	}
	return config,nil
}
func CallAPI(tenant string,apiURL string) (APIResponse,error) {
	var apiResponse APIResponse
	if tenant == "" {
		return apiResponse, fmt.Errorf("tenenat Id is required")
	}
	if apiURL == "" {
		return apiResponse, fmt.Errorf("API URL is required")
	}
	formattedURL :=fmt.Sprintf(apiURL,tenant)

   resp, err := http.Get(formattedURL)
   if err !=nil {
	return apiResponse, fmt.Errorf("error opening API : %w",err)
   }
   defer resp.Body.Close()
   if resp.StatusCode != http.StatusOK {
	return apiResponse, fmt.Errorf("non-200 status code : %w",resp.StatusCode)
   }
   body, err := ioutil.ReadAll(resp.Body)
   if err !=nil {
	return apiResponse, fmt.Errorf("error reading body : %w",err)
   }
   
   err = json.Unmarshal(body, &apiResponse);
   if err !=nil {
	return apiResponse, fmt.Errorf("error reading Json response : %w",err)
   }
   return apiResponse, nil 
 }