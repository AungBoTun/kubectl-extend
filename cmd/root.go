package cmd
import  
(
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/templates"
)
//k8s.io/kubectl/pkg/util/templates"

var rootCmd = &cobra.Command{
	Use:   "tlm",
	Short: "A kubeclt plugin to call API to display the return",
	Long:  `A kubeclt plugin to call API and display the return.`,
	Example: templates.Examples(`
	  #call the api and display the return
	   kubectl tlm
	`),
}
func Execute() {
	if err := rootCmd.Execute(); err !=nil {
		fmt.Fprintf(os.Stderr,"error: %v\n",err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(apiCmd)
}