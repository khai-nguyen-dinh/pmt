package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/jinzhu/gorm"
	"github.com/khainguyen95/pmt/cmd/utils"
	"github.com/spf13/cobra"
	"strings"
)

func CreateUser(db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create new User.",
		Long:  `Create news User.`,
		Run: func(cmd *cobra.Command, args []string) {
			create(db)
			fmt.Println("Create user success!.Use '$pmt config list' to view list user.")
		},
	}
	return cmd
}

func create(db *gorm.DB) {
	user := &utils.User{Username: initUsername(db), Password: initPassword()}
	utils.CreateUser(db, user)
}

func initUsername(db *gorm.DB) string {
	for {
		fmt.Println("Enter your Username: ")
		var raw string
		fmt.Scanln(&raw)
		str, err := checkValidateUsername(db, raw)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return str
	}
}
func initPassword() string {
	for {
		fmt.Println("Enter your Password: ")
		raw, err := gopass.GetPasswd()
		str, err := checkValidatePassword(string(raw))
		if err != nil {
			fmt.Println(err)
			continue
		}
		// encode password to md5
		data := []byte(str)
		output := md5.Sum(data)
		return hex.EncodeToString(output[:])
	}
}

func checkValidatePassword(raw string) (string, error) {
	length := len(raw)
	if length < 8 || length > 11 {
		return "", errors.New("Error: Password must be at least 8 characters, no more than 15 characters!")
	}
	return raw, nil
}
func checkValidateUsername(db *gorm.DB, raw string) (string, error) {
	if raw == "" {
		return "", errors.New("Error: Username must NOT empty!")
	}
	if strings.Contains(raw, " ") {
		return "", errors.New("Error: Username cannot have only blank spaces!")
	}
	if ! utils.CheckUserExist(db, raw) {
		return "", errors.New("Error: Username already exists!")
	}
	return raw, nil
}
