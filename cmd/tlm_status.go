
package cmd
import (
	"fmt"
	 "kubectl-extend/pkg/api"
	 "github.com/spf13/cobra"
)
var (
	tenant string
	configPath string
)
var apiCmd = &cobra.Command{
	Use: "status",
	Short: "call the api and display the return",
	Run: func(cmd *cobra.Command, args []string){
	   runPlugin(tenant,configPath)
    },

}
func init() {
	apiCmd.Flags().StringVarP(&tenant,"tenant","t","","tenant ID (required)")
	apiCmd.Flags().StringVarP(&configPath,"config","c","","config.yaml (required)")
	apiCmd.MarkFlagRequired("tenant")
	apiCmd.MarkFlagRequired("config")
}
func runPlugin(tenant string, configPath string) {
	config,err := api.ReadConfig(configPath)
	if err !=nil {
		fmt.Fprintf(os.Stdeer, "error %v \n",err)
		os.Exit(1)
	}
	
	
	apiResonse,err :=api.CallAPI()
	if err !=nil {
		fmt.Fprintf(os.Stdeer,"error: %v\n",err)
		os.Exit(1)
	}
	fmt.Printf("API Response: %s\n",apiResponse)
}