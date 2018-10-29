package cmd

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/khainguyen95/pmt/cmd/utils"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func DetailInfo(db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show info of application.",
		Long:  `Show all info of application.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runDetailInfo(db, args[0])
		},
	}
	return cmd
}
func runDetailInfo(db *gorm.DB, application string) {
	password, user_id := u.RequirePassword(db)
	if application == "" {
		fmt.Println("Error: Name application must not be empty.")
		os.Exit(1)
	}
	if u.CheckAppExist(db, application) {
		fmt.Println("Error: Application not exists.")
		os.Exit(1)
	}
	data := u.GetAppInfo(db, application, user_id)
	fmt.Printf("Application: %s\n", data.Application)
	secret, err := u.DecryptString(data.Password, password)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Password: %s\n", secret)
	if data.Raw != "" {
		other, err := u.DecryptString(data.Raw, password)
		if err != nil {
			fmt.Println(err)
		}
		s := strings.Split(other, "|")
		for _, v := range s {
			tmp := strings.Split(v, "|")
			fmt.Println("%s: %s\n", tmp[0], tmp[1])
		}
	}
}
