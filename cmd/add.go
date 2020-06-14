/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/adlio/trello"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a card",
	Long:  `Add a card on the specified list of the board.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("add called")
		//fmt.Println(config)
		client := trello.NewClient(config.Key, config.Token)

		boardName, err := cmd.PersistentFlags().GetString("board")
		if err != nil {
			panic(err)
		}
		if boardName == "" {
			boardName = config.Board
			if boardName == "" {
				fmt.Println("board not provided")
				os.Exit(1)
			}
		}
		//fmt.Println(boardName)

		listName, err := cmd.PersistentFlags().GetString("list")
		if err != nil {
			panic(err)
		}
		if listName == "" {
			fmt.Println("list not provided")
			os.Exit(1)
		}
		//fmt.Println(listName)

		cardName := cmd.Flags().Arg(0)
		description := cmd.Flags().Arg(1)
		if cardName == "" {
			fmt.Println("card name not provided")
			os.Exit(1)
		}

		member, err := client.GetMember(config.MemberID, trello.Defaults())
		if err != nil {
			panic(err)
		}
		boards, err := member.GetBoards(trello.Defaults())
		if err != nil {
			panic(err)
		}
		for _, b := range boards {
			if b.Name == boardName {
				lists, _ := b.GetLists(trello.Defaults())
				for _, l := range lists {
					if l.Name == listName {
						c := &trello.Card{Name: cardName, Desc: description}
						err := l.AddCard(c, trello.Defaults())
						if err == nil {
							fmt.Println(c.ShortURL + " created")
						} else {
							panic(err)
						}
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.PersistentFlags().StringP("board", "b", "", "board name (required if not defined in config file)")
	addCmd.PersistentFlags().StringP("list", "l", "", "list name (required)")
}
