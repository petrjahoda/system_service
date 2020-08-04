package main

import (
	"github.com/petrjahoda/database"
	"gopkg.in/gomail.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func SendEmail(header string) {
	LogInfo("MAIN", "Sending email about: "+header)
	timer := time.Now()
	companyName, host, port, username, password, recipient := ReadMailSettings()
	message := gomail.NewMessage()
	message.SetHeader("From", username)
	message.SetHeader("Subject", companyName+": "+header)
	message.SetBody("text/html", "Disc space will last for just 30 days, please consider bigger disc for data.")
	message.SetHeader("To", recipient)
	d := gomail.NewDialer(host, port, username, password)
	err := d.DialAndSend(message)
	if err != nil {
		LogError("MAIN", "Email not sent: "+err.Error())
		return
	}
	LogInfo("MAIN", "Email sent in "+time.Since(timer).String())
}

func ReadMailSettings() (string, string, int, string, string, string) {
	LogInfo("MAIN", "Reading mail settings")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening database: "+err.Error())
		return "", "", 0, "", "", ""
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	var settingsHost database.Setting
	db.Where("name=?", "host").Find(&settingsHost)
	var company database.Setting
	db.Where("name=?", "company").Find(&company)
	companyName := company.Value
	host := settingsHost.Value
	var settingsPort database.Setting
	db.Where("name=?", "port").Find(&settingsPort)
	port, err := strconv.Atoi(settingsPort.Value)
	if err != nil {
		LogError("MAIN", "Problem parsing port for email, using default port 587 "+err.Error())
		port = 587
	}
	var settingsUsername database.Setting
	db.Where("name=?", "username").Find(&settingsUsername)
	username := settingsUsername.Value
	var settingsPassword database.Setting
	db.Where("name=?", "password").Find(&settingsPassword)
	password := settingsPassword.Value
	var settingsEmail database.Setting
	db.Where("name=?", "email").Find(&settingsEmail)
	email := settingsEmail.Value
	LogInfo("MAIN", "Mail settings read in "+time.Since(timer).String())
	return companyName, host, port, username, password, email
}
