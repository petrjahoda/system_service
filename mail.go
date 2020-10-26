package main

import (
	"github.com/petrjahoda/database"
	"gopkg.in/gomail.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func sendEmail(header string) {
	logInfo("MAIN", "Sending email about: "+header)
	timer := time.Now()
	companyName, host, port, username, password, recipient := readMailSettings()
	message := gomail.NewMessage()
	message.SetHeader("From", username)
	message.SetHeader("Subject", companyName+": "+header)
	message.SetBody("text/html", "Disc space will last for just 30 days, please consider bigger disc for data.")
	message.SetHeader("To", recipient)
	d := gomail.NewDialer(host, port, username, password)
	err := d.DialAndSend(message)
	if err != nil {
		logError("MAIN", "Email not sent: "+err.Error())
		return
	}
	logInfo("MAIN", "Email sent in "+time.Since(timer).String())
}

func readMailSettings() (string, string, int, string, string, string) {
	logInfo("MAIN", "Reading mail settings")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	if err != nil {
		logError("MAIN", "Problem opening database: "+err.Error())
		return "", "", 0, "", "", ""
	}

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
		logError("MAIN", "Problem parsing port for email, using default port 587 "+err.Error())
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
	logInfo("MAIN", "Mail settings read in "+time.Since(timer).String())
	return companyName, host, port, username, password, email
}
