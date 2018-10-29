package cmd

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/khainguyen95/pmt/cmd/utils"
	"github.com/spf13/cobra"
)

func ChangeUser(db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "change [USERNAME]",
		Short: "Change to different User.",
		Long:  `Change to different User.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runChangeUser(db, args[0])
		},
	}
	return cmd
}
func runChangeUser(db *gorm.DB, username string) {
	if utils.CheckUserExist(db, username) {
		fmt.Println("Error: Username not exists!")
		return
	}
	utils.ChangeUserActive(db, username)
	fmt.Println("Change User success!")
}
