// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"agenda/entity"
	"fmt"

	"github.com/spf13/cobra"
)

var username, password, email, phone string

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Use register command to register a new user",
	Long: ` when you register a new account,
			you must offer your username, password,
			e-mail and phone number
			username must be unique
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if username == "" {
			fmt.Println("empty username")
			cmd.Help()
		} else if password == "" {
			fmt.Println("empty password")
			cmd.Help()
		} else if email == "" {
			fmt.Println("fill email address plz")
			cmd.Help()
		} else if phone == "" {
			fmt.Println("fill phone number plz")
			cmd.Help()
		} else {
			err := entity.Reg(username, password, email, phone)
			if err == 0 {
				fmt.Println("register successfully")
			} else {
				fmt.Println("register failed, maybe you should try a different username")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	RootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "the username of account")
	RootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "the password of account")
	registerCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "the e-mail address of account")
	registerCmd.PersistentFlags().StringVarP(&phone, "phone", "c", "", "the phone number of account")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
