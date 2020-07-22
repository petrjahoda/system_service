package main

import (
	"database/sql"
	"github.com/kardianos/service"
	"github.com/petrjahoda/zapsi_database"
	"github.com/ricochet2200/go-disk-usage/du"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

const version = "2020.3.1.22"
const programName = "System Service"
const programDescription = "Creates database and checks system data"
const config = "user=postgres password=Zps05..... dbname=zapsi3 host=zapsidatabase port=5432 sslmode=disable"
const downloadInSeconds = 86400

var serviceRunning = false
var processRunning = false

type Result struct {
	Size float32
}

type program struct{}

func (p *program) Start(s service.Service) error {
	LogInfo("MAIN", "Starting "+programName+" on "+s.Platform())
	serviceRunning = true
	go p.run()
	return nil
}

func (p *program) run() {
	LogInfo("MAIN", "Program version "+version+" started")
	for {
		start := time.Now()
		LogInfo("MAIN", "Program running")
		CompleteDatabaseCheck()
		if time.Since(start) < (downloadInSeconds*time.Second) && serviceRunning {
			processRunning = true
			sleeptime := downloadInSeconds*time.Second - time.Since(start)
			databaseSizeMegaBytes := GetDatabaseSize()
			databaseGrowthInMegaBytes := GetGrowthOfDatabase(databaseSizeMegaBytes)
			if databaseGrowthInMegaBytes != 0 {
				discSpaceMegaBytes := GetDiscSpace()
				estimatedDiscSpaceDays := CalculateDiscSpace(databaseGrowthInMegaBytes, discSpaceMegaBytes)
				WriteNewSystemRecord(databaseSizeMegaBytes, databaseGrowthInMegaBytes, discSpaceMegaBytes, estimatedDiscSpaceDays)
				SendEmail(estimatedDiscSpaceDays)
			} else {
				LogInfo("MAIN", "No change in database size")
			}
			LogInfo("MAIN", "Sleeping for "+sleeptime.String())
			time.Sleep(sleeptime)
			processRunning = false
		}
	}
}

func (p *program) Stop(s service.Service) error {
	serviceRunning = false
	if processRunning {
		LogInfo("MAIN", "Stopping...")
		time.Sleep(1 * time.Second)
	}
	LogInfo("MAIN", "Stopped on platform "+s.Platform())
	return nil
}

func main() {
	serviceConfig := &service.Config{
		Name:        programName,
		DisplayName: programName,
		Description: programDescription,
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		LogError("MAIN", err.Error())
	}
	err = s.Run()
	if err != nil {
		LogError("MAIN", "Problem starting "+serviceConfig.Name)
	}
}

func SendEmail(estimatedDiscSpaceDays float32) {
	if estimatedDiscSpaceDays < 30 {
		LogInfo("MAIN", "Sending email about low disc space")
		timer := time.Now()
		companyName := GetCompanyName()
		recipient := GetRecipient()
		err, host, port, username, password, _ := UpdateMailSettings()
		if err != nil {
			return
		}
		m := gomail.NewMessage()
		m.SetHeader("From", username)
		m.SetHeader("Subject", companyName+": low disc space")
		m.SetBody("text/html", "Disc space is already free for only 30 days.")
		m.SetHeader("To", recipient)

		d := gomail.NewDialer(host, port, username, password)
		if emailSentError := d.DialAndSend(m); emailSentError != nil {
			LogError("MAIN", "Email not sent: "+emailSentError.Error()+", elapsed: "+time.Since(timer).String())
		} else {
			LogInfo("MAIN", "Email sent, elapsed: "+time.Since(timer).String())
		}
	}
}

func GetRecipient() string {
	LogInfo("MAIN", "Downloading recipient email name")
	timer := time.Now()
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=Zps05..... dbname=postgres host=localhost port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening  database: "+err.Error())
		return ""
	}
	var company zapsi_database.Setting
	db.Where("name=?", "email").Find(&company)
	LogInfo("MAIN", "Recipient email ["+company.Value+"] downloaded from database, elapsed: "+time.Since(timer).String())
	return company.Value
}

func GetCompanyName() string {
	LogInfo("MAIN", "Downloading company name")
	timer := time.Now()
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=Zps05..... dbname=postgres host=zapsidatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening  database: "+err.Error())
		return ""
	}
	var company zapsi_database.Setting
	db.Where("name=?", "company").Find(&company)
	LogInfo("MAIN", "Company name ["+company.Value+"] downloaded from database, elapsed: "+time.Since(timer).String())
	return company.Value
}

func UpdateMailSettings() (error, string, int, string, string, string) {
	LogInfo("MAIN", "Updating mail settings")
	timer := time.Now()
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=Zps05..... dbname=postgres host=localhost port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening  database: "+err.Error())
		return nil, "", 0, "", "", ""
	}
	var settingsHost zapsi_database.Setting
	db.Where("name=?", "host").Find(&settingsHost)
	host := settingsHost.Value
	var settingsPort zapsi_database.Setting
	db.Where("name=?", "port").Find(&settingsPort)
	port, err := strconv.Atoi(settingsPort.Value)
	if err != nil {
		LogError("MAIN", "Problem parsing port for email, using default port 587 "+err.Error())
		port = 587
	}
	var settingsUsername zapsi_database.Setting
	db.Where("name=?", "username").Find(&settingsUsername)
	username := settingsUsername.Value
	var settingsPassword zapsi_database.Setting
	db.Where("name=?", "password").Find(&settingsPassword)
	password := settingsPassword.Value
	var settingsEmail zapsi_database.Setting
	db.Where("name=?", "email").Find(&settingsEmail)
	email := settingsEmail.Value
	LogInfo("MAIN", "Mail settings updated, elapsed: "+time.Since(timer).String())
	return err, host, port, username, password, email
}

func WriteNewSystemRecord(databaseSizeMegaBytes float32, databaseGrowthInMegaBytes float32, discSpaceMegaBytes float32, estimatedDiscSpaceDays float32) {
	LogInfo("MAIN", "Writing new record to database")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening  database: "+err.Error())
		return
	}
	var newSystemRecord zapsi_database.SystemRecord
	newSystemRecord.DatabaseSizeInMegaBytes = databaseSizeMegaBytes
	newSystemRecord.DatabaseGrowthInMegaBytes = databaseGrowthInMegaBytes
	newSystemRecord.DiscFreeSizeInMegaBytes = discSpaceMegaBytes
	newSystemRecord.EstimatedDiscFreeSizeInDays = estimatedDiscSpaceDays
	db.Save(&newSystemRecord)
	LogInfo("MAIN", "New record saved, elapsed: "+time.Since(timer).String())

}

func CalculateDiscSpace(databaseGrowthInMegaBytes float32, discSpaceMegaBytes float32) float32 {
	return discSpaceMegaBytes / databaseGrowthInMegaBytes
}

func GetDiscSpace() float32 {
	LogInfo("MAIN", "Getting free disk space")
	timer := time.Now()
	usage := du.NewDiskUsage("/")
	LogInfo("MAIN", "Disc free space calculated, elapsed: "+time.Since(timer).String())
	return float32(usage.Available() / (1024 * 1024))
}

func GetGrowthOfDatabase(actualDatabaseSize float32) float32 {
	LogInfo("MAIN", "Getting database growth")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening  database: "+err.Error())
		return 0
	}
	var systemRecord zapsi_database.SystemRecord
	db.Last(&systemRecord)
	growth := actualDatabaseSize - systemRecord.DatabaseSizeInMegaBytes
	LogInfo("MAIN", "Database growth calculated, elapsed: "+time.Since(timer).String())
	return growth
}

func GetDatabaseSize() float32 {
	LogInfo("MAIN", "Getting database size")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening  database: "+err.Error())
		return 0
	}
	var result Result
	db.Raw("SELECT pg_database_size('zapsi3')/1000000 as Size;").Scan(&result)
	LogInfo("MAIN", "Database size calculated, elapsed: "+time.Since(timer).String())
	return result.Size
}

func CompleteDatabaseCheck() {
	firstRunCheckComplete := false
	for firstRunCheckComplete == false {
		processRunning = true
		databaseOk := CheckDatabase()
		tablesOk, err := CheckTables()
		if err != nil {
			LogInfo("MAIN", "Problem creating tables: "+err.Error())
			time.Sleep(1 * time.Second)
		}
		if databaseOk && tablesOk {
			WriteProgramVersionIntoSettings()
			firstRunCheckComplete = true
		}
		processRunning = false
	}
}

func CheckTables() (bool, error) {
	LogInfo("MAIN", "Checking tables")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening database: "+err.Error()+", elapsed: "+time.Since(timer).String())
		return false, err
	}

	if !db.Migrator().HasTable(&zapsi_database.DeviceType{}) {
		LogInfo("MAIN", "DeviceTypeId table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.DeviceType{})
		if err != nil {
			LogError("MAIN", "Cannot create devicetype table")
		}
		zapsi := zapsi_database.DeviceType{Name: "Zapsi"}
		db.Create(&zapsi)
		zapsiTouch := zapsi_database.DeviceType{Name: "Zapsi Touch"}
		db.Create(&zapsiTouch)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.DeviceType{})
		if err != nil {
			LogError("MAIN", "Cannot update devicetype table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.Device{}) {
		LogInfo("MAIN", "Device table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Device{})
		if err != nil {
			LogError("MAIN", "Cannot create device table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Device{})
		if err != nil {
			LogError("MAIN", "Cannot update device table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.Setting{}) {
		LogInfo("MAIN", "Setting table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Setting{})
		if err != nil {
			LogError("MAIN", "Cannot create setting table")
		}
		company := zapsi_database.Setting{Name: "company", Value: "zapsi"}
		db.Create(&company)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Setting{})
		if err != nil {
			LogError("MAIN", "Cannot update setting table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.DevicePortType{}) {
		LogInfo("MAIN", "DevicePortType table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.DevicePortType{})
		if err != nil {
			LogError("MAIN", "Cannot create deviceporttype table")
		}
		digital := zapsi_database.DevicePortType{Name: "Digital"}
		db.Create(&digital)
		analog := zapsi_database.DevicePortType{Name: "Analog"}
		db.Create(&analog)
		serial := zapsi_database.DevicePortType{Name: "Serial"}
		db.Create(&serial)
		special := zapsi_database.DevicePortType{Name: "Special"}
		db.Create(&special)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.DevicePortType{})
		if err != nil {
			LogError("MAIN", "Cannot update deviceporttype table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.DevicePort{}) {
		LogInfo("MAIN", "DevicePort table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.DevicePort{})
		if err != nil {
			LogError("MAIN", "Cannot create deviceport table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.DevicePort{})
		if err != nil {
			LogError("MAIN", "Cannot update deviceport table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.DevicePortAnalogRecord{}) {
		LogInfo("MAIN", "DevicePortAnalogRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.DevicePortAnalogRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create deviceportanalogrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.DevicePortAnalogRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update deviceportanalogrecord table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.DevicePortDigitalRecord{}) {
		LogInfo("MAIN", "DevicePortDigitalRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.DevicePortDigitalRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create deviceportdigitalrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.DevicePortDigitalRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update deviceportdigitalrecord table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.DevicePortSerialRecord{}) {
		LogInfo("MAIN", "DevicePortSerialRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.DevicePortSerialRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create deviceportserialrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.DevicePortSerialRecord{})
		if err != nil {
			LogError("MAIN", "Cannot UPDATE deviceportserialrecord table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.PackageType{}) {
		LogInfo("MAIN", "PackageType table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.PackageType{})
		if err != nil {
			LogError("MAIN", "Cannot create packagetype table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.PackageType{})
		if err != nil {
			LogError("MAIN", "Cannot UPDATE packagetype table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.Product{}) {
		LogInfo("MAIN", "Product table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Product{})
		if err != nil {
			LogError("MAIN", "Cannot create product table")
		}
		product := zapsi_database.Product{Name: "Product"}
		db.Create(&product)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Product{})
		if err != nil {
			LogError("MAIN", "Cannot update product table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.WorkplaceSection{}) {
		LogInfo("MAIN", "WorkplaceSection table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.WorkplaceSection{})
		if err != nil {
			LogError("MAIN", "Cannot create workplacesection table")
		}
		machines := zapsi_database.WorkplaceSection{Name: "Machines"}
		db.Create(&machines)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.WorkplaceSection{})
		if err != nil {
			LogError("MAIN", "Cannot update workplacesection table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.Workplace{}) {
		LogInfo("MAIN", "Workplace table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Workplace{})
		if err != nil {
			LogError("MAIN", "Cannot create workplace table")
		}

	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Workplace{})
		if err != nil {
			LogError("MAIN", "Cannot update workplace table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.Order{}) {
		LogInfo("MAIN", "Order table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Order{})
		if err != nil {
			LogError("MAIN", "Cannot create order table")
		}
		order := zapsi_database.Order{Name: "Internal", DateTimeRequest: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}}
		db.Create(&order)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Order{})
		if err != nil {
			LogError("MAIN", "Cannot update order table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.Package{}) {
		LogInfo("MAIN", "Package table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Package{})
		if err != nil {
			LogError("MAIN", "Cannot create package table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Package{})
		if err != nil {
			LogError("MAIN", "Cannot update package table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.FaultType{}) {
		LogInfo("MAIN", "FaultType table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.FaultType{})
		if err != nil {
			LogError("MAIN", "Cannot create faulttype table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.FaultType{})
		if err != nil {
			LogError("MAIN", "Cannot update faulttype table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.Fault{}) {
		LogInfo("MAIN", "Fault table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Fault{})
		if err != nil {
			LogError("MAIN", "Cannot create fault table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Fault{})
		if err != nil {
			LogError("MAIN", "Cannot update fault table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.BreakdownType{}) {
		LogInfo("MAIN", "BreakdownType table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.BreakdownType{})
		if err != nil {
			LogError("MAIN", "Cannot create breakdowntype table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.BreakdownType{})
		if err != nil {
			LogError("MAIN", "Cannot create update table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.Breakdown{}) {
		LogInfo("MAIN", "Breakdown table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Breakdown{})
		if err != nil {
			LogError("MAIN", "Cannot create breakdown table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Breakdown{})
		if err != nil {
			LogError("MAIN", "Cannot update breakdown table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.DowntimeType{}) {
		LogInfo("MAIN", "DowntimeType table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.DowntimeType{})
		if err != nil {
			LogError("MAIN", "Cannot create downtimetype table")
		}
		system := zapsi_database.DowntimeType{Name: "System"}
		db.Create(&system)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.DowntimeType{})
		if err != nil {
			LogError("MAIN", "Cannot update downtimetype table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.Downtime{}) {
		LogInfo("MAIN", "Downtime table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Downtime{})
		if err != nil {
			LogError("MAIN", "Cannot create downtime table")
		}
		system := zapsi_database.DowntimeType{}
		db.Where("Name = ?", "System").Find(&system)
		noReasonDowntime := zapsi_database.Downtime{
			Name:           "No reason Downtime",
			DowntimeTypeID: 1,
		}
		db.Create(&noReasonDowntime)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Downtime{})
		if err != nil {
			LogError("MAIN", "Cannot update downtime table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.UserType{}) {
		LogInfo("MAIN", "UserType table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.UserType{})
		if err != nil {
			LogError("MAIN", "Cannot create usertype table")
		}
		operator := zapsi_database.UserType{Name: "Operator"}
		db.Create(&operator)
		zapsi := zapsi_database.UserType{Name: "Zapsi"}
		db.Create(&zapsi)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.UserType{})
		if err != nil {
			LogError("MAIN", "Cannot update usertype table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.UserRole{}) {
		LogInfo("MAIN", "UserRole table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.UserRole{})
		if err != nil {
			LogError("MAIN", "Cannot create userrole table")
		}
		admin := zapsi_database.UserRole{Name: "Administrator"}
		db.Create(&admin)
		powerUser := zapsi_database.UserRole{Name: "PowerUser"}
		db.Create(&powerUser)
		user := zapsi_database.UserRole{Name: "User"}
		db.Create(&user)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.UserRole{})
		if err != nil {
			LogError("MAIN", "Cannot update userrole table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.User{}) {
		LogInfo("MAIN", "User table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.User{})
		if err != nil {
			LogError("MAIN", "Cannot create user table")
		}
		userRole := zapsi_database.UserRole{}
		db.Where("Name = ?", "Administrator").Find(&userRole)
		userType := zapsi_database.UserType{}
		db.Where("Name = ?", "Zapsi").Find(&userType)
		password := hashAndSalt([]byte("54321"))
		zapsiUser := zapsi_database.User{
			FirstName:  "Zapsi",
			SecondName: "Zapsi",
			Password:   password,
			UserRoleID: 1,
			UserTypeID: 1,
		}
		db.Create(&zapsiUser)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.User{})
		if err != nil {
			LogError("MAIN", "Cannot update user table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.Workshift{}) {
		LogInfo("MAIN", "Workshift table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Workshift{})
		if err != nil {
			LogError("MAIN", "Cannot create workshift table")
		}
		firstShiftStart := time.Date(2000, 1, 1, 6, 0, 0, 0, time.Local)
		firstShiftEnd := time.Date(2000, 1, 1, 14, 0, 0, 0, time.Local)
		firstShift := zapsi_database.Workshift{Name: "First Shift", WorkshiftStart: firstShiftStart, WorkshiftEnd: firstShiftEnd}
		db.Create(&firstShift)
		secondShiftStart := time.Date(2000, 1, 1, 14, 0, 0, 0, time.Local)
		secondShiftEnd := time.Date(2000, 1, 1, 22, 0, 0, 0, time.Local)
		secondShift := zapsi_database.Workshift{Name: "Second Shift", WorkshiftStart: secondShiftStart, WorkshiftEnd: secondShiftEnd}
		db.Create(&secondShift)
		thirdShiftStart := time.Date(2000, 1, 1, 22, 0, 0, 0, time.Local)
		thirdShiftEnd := time.Date(2000, 1, 2, 6, 0, 0, 0, time.Local)
		thirdShift := zapsi_database.Workshift{Name: "Third Shift", WorkshiftStart: thirdShiftStart, WorkshiftEnd: thirdShiftEnd}
		db.Create(&thirdShift)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Workshift{})
		if err != nil {
			LogError("MAIN", "Cannot update workshift table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.WorkplaceMode{}) {
		LogInfo("MAIN", "Workplacemode table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.WorkplaceMode{})
		if err != nil {
			LogError("MAIN", "Cannot create workplacemode table")
		}
		mode := zapsi_database.WorkplaceMode{Name: "Production", DowntimeDuration: time.Second * 300, PoweroffDuration: time.Second * 300}
		db.Create(&mode)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.WorkplaceMode{})
		if err != nil {
			LogError("MAIN", "Cannot update workplacemode table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.Part{}) {
		LogInfo("MAIN", "Part table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Part{})
		if err != nil {
			LogError("MAIN", "Cannot create part table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Part{})
		if err != nil {
			LogError("MAIN", "Cannot update part table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.State{}) {
		LogInfo("MAIN", "State table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.State{})
		if err != nil {
			LogError("MAIN", "Cannot create state table")
		}
		production := zapsi_database.State{Name: "Production", Color: "#89AB0F"}
		db.Create(&production)
		downtime := zapsi_database.State{Name: "Downtime", Color: "#E6AD3C"}
		db.Create(&downtime)
		poweroff := zapsi_database.State{Name: "Poweroff", Color: "#DE6B59"}
		db.Create(&poweroff)
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.State{})
		if err != nil {
			LogError("MAIN", "Cannot update state table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.WorkplaceWorkshift{}) {
		LogInfo("MAIN", "WorkplaceWorkshift table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.WorkplaceWorkshift{})
		if err != nil {
			LogError("MAIN", "Cannot create workplaceworkshift table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Workshift{})
		if err != nil {
			LogError("MAIN", "Cannot update workplaceworkshift table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.WorkplacePort{}) {
		LogInfo("MAIN", "WorkplacePort table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.WorkplacePort{})
		if err != nil {
			LogError("MAIN", "Cannot create workplaceport table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.WorkplacePort{})
		if err != nil {
			LogError("MAIN", "Cannot update workplaceport table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.StateRecord{}) {
		LogInfo("MAIN", "StateRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.StateRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create staterecord table")
		}

	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.StateRecord{})
		if err != nil {
			LogError("MAIN", "update create staterecord table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.Operation{}) {
		LogInfo("MAIN", "Operation table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Operation{})
		if err != nil {
			LogError("MAIN", "Cannot create operation table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Operation{})
		if err != nil {
			LogError("MAIN", "Cannot update operation table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.OrderRecord{}) {
		LogInfo("MAIN", "OrderRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.OrderRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create orderrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.OrderRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update orderrecord table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.PartRecord{}) {
		LogInfo("MAIN", "PartRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.PartRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create partrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.PartRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update partrecord table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.PackageRecord{}) {
		LogInfo("MAIN", "PackageRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.PackageRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create packagerecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.PackageRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update packagerecord table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.FaultRecord{}) {
		LogInfo("MAIN", "FaultRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.FaultRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create faultrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.FaultRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update faultrecord table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.BreakdownRecord{}) {
		LogInfo("MAIN", "BreakdownRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.BreakdownRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create breakdownrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.BreakdownRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update breakdownrecord table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.DownTimeRecord{}) {
		LogInfo("MAIN", "DownTimeRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.DownTimeRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create downtimerecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.DownTimeRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update downtimerecord table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.UserRecord{}) {
		LogInfo("MAIN", "UserRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.UserRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create userrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.UserRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update userrecord table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.Alarm{}) {
		LogInfo("MAIN", "Alarm table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.Alarm{})
		if err != nil {
			LogError("MAIN", "Cannot create alarm table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.Alarm{})
		if err != nil {
			LogError("MAIN", "Cannot update alarm table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.AlarmRecord{}) {
		LogInfo("MAIN", "AlarmRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.AlarmRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create alarmrecord table")
		}

	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.AlarmRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update alarmrecord table")
		}
	}

	if !db.Migrator().HasTable(&zapsi_database.SystemRecord{}) {
		LogInfo("MAIN", "SystemRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.SystemRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create systemrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.SystemRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update systemrecord table")
		}
	}
	if !db.Migrator().HasTable(&zapsi_database.DeviceWorkplaceRecord{}) {
		LogInfo("MAIN", "DeviceWorkplaceRecord table not exists, creating")
		err := db.Migrator().CreateTable(&zapsi_database.DeviceWorkplaceRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create deviceworkplacerecord table")
		}

	} else {
		err := db.Migrator().AutoMigrate(&zapsi_database.DeviceWorkplaceRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update deviceworkplacerecord table")
		}
	}
	LogInfo("MAIN", "Tables created, elapsed: "+time.Since(timer).String())
	return true, nil

}

func CheckDatabase() bool {
	timer := time.Now()
	LogInfo("MAIN", "Checking database")
	_, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Database zapsi3 does not exist")
		db, err := gorm.Open(postgres.Open("user=postgres password=Zps05..... dbname=postgres host=localhost port=5432 sslmode=disable"), &gorm.Config{})
		if err != nil {
			LogError("MAIN", "Problem opening database: "+err.Error()+", elapsed: "+time.Since(timer).String())
			return false
		}
		db = db.Exec("CREATE DATABASE zapsi3;")
		if db.Error != nil {
			LogError("MAIN", "Cannot create database zapsi3")
		}
		LogInfo("MAIN", "Database zapsi3 created, elapsed: "+time.Since(timer).String())
		return true

	}
	LogInfo("MAIN", "Database zapsi3 exists, elapsed: "+time.Since(timer).String())
	return true

}

func WriteProgramVersionIntoSettings() {
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening  database: "+err.Error())
		return
	}
	var existingSettings zapsi_database.Setting
	db.Where("name=?", programName).Find(&existingSettings)
	existingSettings.Name = programName
	existingSettings.Value = version
	db.Save(&existingSettings)
	LogInfo("MAIN", "Updated version in database for "+programName)
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
