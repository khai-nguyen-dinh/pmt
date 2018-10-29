package cmd

import (
	"fmt"
	"github.com/jinzhu/gorm"
	u "github.com/khainguyen95/pmt/cmd/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func CreateInfo(db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create new info of application.",
		Long:  `Create news info of application.`,
		Run: func(cmd *cobra.Command, args []string) {
			runCreteaInfo(db)
		},
	}
	cmd.Flags().StringP("application", "a", "", "Define your name application.")
	cmd.Flags().StringP("password", "p", "", "Password of application.")
	cmd.Flags().StringP("other", "o", "", "Add other secrets of the application (format: <field1>:secret1|<field2>:secret2)")
	viper.BindPFlag("application", cmd.Flags().Lookup("application"))
	viper.BindPFlag("password", cmd.Flags().Lookup("password"))
	viper.BindPFlag("other", cmd.Flags().Lookup("other"))

	return cmd
}

func runCreteaInfo(db *gorm.DB) {
	password, user_id := u.RequirePassword(db)
	data := getInputData(db, password)
	data.UserId = user_id
	u.CreateInfo(db, data)
	fmt.Println("Create info success! Run '$pmt list' to view list application.")
}

func getInputData(db *gorm.DB, password string) u.Data {
	var data u.Data
	var err error
	if viper.GetString("application") == "" {
		fmt.Println("Error: Name application must not be empty.")
		os.Exit(1)
	}
	if !u.CheckAppExist(db, viper.GetString("application")) {
		fmt.Println("Error: Application already exists.")
		os.Exit(1)
	}
	data.Application = viper.GetString("application")
	data.Password, err = u.EncryptString(viper.GetString("password"), password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if raw := viper.GetString("other"); raw != "" {
		data.Raw, err = u.EncryptString(viper.GetString("other"), password)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return data
}
