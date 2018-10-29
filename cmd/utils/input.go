package utils

import (
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/jinzhu/gorm"
	"os"
)

func RequirePassword(db *gorm.DB) (string, uint) {
	fmt.Println("Enter Password: ")
	raw, err := gopass.GetPasswd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	user_id, out := CheckPassword(db, string(raw))
	if out {
		return string(raw), user_id
	}
	fmt.Println("Error: Incorrect password.")
	os.Exit(1)
	return "", user_id
}
