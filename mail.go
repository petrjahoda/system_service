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
	companyName := GetCompanyName()
	recipient := GetRecipient()
	err, host, port, username, password, _ := UpdateMailSettings()
	if err != nil {
		return
	}
	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("Subject", companyName+": "+header)
	m.SetBody("text/html", "Disc space will last for just 30 days, please consider bigger disc for data.")
	m.SetHeader("To", recipient)
	d := gomail.NewDialer(host, port, username, password)
	if emailSentError := d.DialAndSend(m); emailSentError != nil {
		LogError("MAIN", "Email not sent: "+emailSentError.Error())
	} else {
		LogInfo("MAIN", "Email sent int "+time.Since(timer).String())
	}
}

func GetCompanyName() string {
	LogInfo("MAIN", "Downloading company name")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening  database: "+err.Error())
		return ""
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	var company database.Setting
	db.Where("name=?", "company").Find(&company)
	LogInfo("MAIN", "Company name ["+company.Value+"] downloaded from database, elapsed: "+time.Since(timer).String())
	return company.Value
}

func GetRecipient() string {
	LogInfo("MAIN", "Downloading recipient email name")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening  database: "+err.Error())
		return ""
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	var company database.Setting
	db.Where("name=?", "email").Find(&company)
	LogInfo("MAIN", "Recipient email ["+company.Value+"] downloaded from database, elapsed: "+time.Since(timer).String())
	return company.Value
}

func UpdateMailSettings() (error, string, int, string, string, string) {
	LogInfo("MAIN", "Updating mail settings")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening  database: "+err.Error())
		return nil, "", 0, "", "", ""
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	var settingsHost database.Setting
	db.Where("name=?", "host").Find(&settingsHost)
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
	LogInfo("MAIN", "Mail settings updated, elapsed: "+time.Since(timer).String())
	return err, host, port, username, password, email
}
