package cmd

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/khainguyen95/pmt/cmd/utils"
	"github.com/spf13/cobra"
	"os"
)

func DeleteInfo(db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete info of application.",
		Long:  `Delete info of application.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runDeleteInfo(db, args[0])
		},
	}
	return cmd
}
func runDeleteInfo(db *gorm.DB, application string) {
	_, user_id := u.RequirePassword(db)
	if application == "" {
		fmt.Println("Error: Name application must not be empty.")
		os.Exit(1)
	}
	if u.CheckAppExist(db, application) {
		fmt.Println("Error: Application not exists.")
		os.Exit(1)
	}
	u.DeleteAppInfo(db, application, user_id)
	fmt.Println("The record has been deleted!")
}
