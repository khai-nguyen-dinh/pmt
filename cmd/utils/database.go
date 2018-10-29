package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/user"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(16);not null"`
	Active   int    `gorm:"type:tinyint(1);not null;default:0"`
	Data     []Data
}

type Data struct {
	gorm.Model
	Application string `gorm:"type:varchar(255);not null"`
	Password    string `gorm:"type:text"`
	Raw         string `gorm:"type:text"`
	UserId      uint   `sql:"type:int REFERENCES users(id)"`
}

func InitDB() *gorm.DB {
	usr, err := user.Current()
	if err != nil {
		fmt.Print(err)
	}
	if _, err := os.Stat(usr.HomeDir + "/.pmt"); os.IsNotExist(err) {
		os.MkdirAll(usr.HomeDir+"/.pmt", os.ModePerm);
	}
	db, err := gorm.Open("sqlite3", usr.HomeDir+"/.pmt/secrete.data")
	if err != nil {
		fmt.Print(err)
	}

	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	if !db.HasTable(&Data{}) {
		db.Exec("PRAGMA foreign_keys = ON;")
		db.CreateTable(&Data{})
	}
	return db
}

func CreateUser(db *gorm.DB, user *User) {
	// set default user after create success.
	users := ListAllUser(db)
	if len(users) == 0 {
		user.Active = 1
	}
	db.Create(user)
}
func CreateInfo(db *gorm.DB, data Data) {
	db.Create(&data)
}

func ListInfo(db *gorm.DB, user_id uint) [] Data {
	var data [] Data
	db.Where("user_id = ?", user_id).Find(&data)
	return data
}

func ListAllUser(db *gorm.DB) ([] User) {
	var users [] User
	db.Find(&users)
	return users
}
func ChangeUserActive(db *gorm.DB, username string) {
	var users [] User
	// Unactive all user.
	db.Where("active = ?", 1).Find(&users)
	for i := 0; i < len(users); i += 1 {
		db.Model(&users[i]).Update("active", 0)
		db.Save(&users[i])
	}
	var user User
	// Update new active user.
	db.Where("username = ?", username).First(&user)
	db.Model(&user).Update("active", 1)
	db.Save(&user)
}
func CheckUserExist(db *gorm.DB, username string) bool {
	var user User
	return db.Where("username = ?", username).Find(&user).RecordNotFound()
}

func CheckAppExist(db *gorm.DB, app string) bool {
	var data Data
	return db.Where("application = ?", app).Find(&data).RecordNotFound()
}

func CheckPassword(db *gorm.DB, password string) (uint, bool) {
	var user User
	data := []byte(password)
	output := md5.Sum(data)
	db.Where("active = ?", 1).First(&user)
	if user.Password == hex.EncodeToString(output[:]) {
		return user.ID, true
	}
	return user.ID, false
}

func GetUserActive(db *gorm.DB) User {
	var user User
	db.Where("active = ?", 1).First(&user)
	return user
}
func GetAppInfo(db *gorm.DB, application string, user_id uint) Data {
	var data Data
	db.Where("application = ? AND user_id = ?", application, user_id).First(&data)
	return data
}
func DeleteAppInfo(db *gorm.DB, application string, user_id uint) {
	var data Data
	db.Where("application = ? AND user_id = ?", application, user_id).First(&data)
	db.Delete(&data)
}
