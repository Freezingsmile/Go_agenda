package cmd

import (
	"fmt"
	"agenda/entity"
	"github.com/spf13/cobra"
)

// quCmd represents the qu command
var quCmd = &cobra.Command{
	Use:   "qu",
	Short: "to find user infomation ",
	Long: `All user's information.`,
	Run: func(cmd *cobra.Command, args []string) {
		                                                                                                                                                                                                            
		nuser := entity.GetAllUsers();
		for _,i:= range nuser {
				fmt.Println("----------------")
					fmt.Println("Username: ", i.Name)
					fmt.Println("Telephone number: ", i.Tel)
					fmt.Println("Email: ", i.Email)
					fmt.Println("----------------")
				}
	},
}

func init() {
	RootCmd.AddCommand(quCmd)
	quCmd.Flags().StringP("userinfo", "q", "", "user info display")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
