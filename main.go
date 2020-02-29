package main

import (
	"github.com/jinzhu/gorm"
	"github.com/kardianos/service"
	"github.com/petrjahoda/zapsi_database"
	"github.com/ricochet2200/go-disk-usage/du"
	"gopkg.in/gomail.v2"
	"strconv"
	"strings"
	"time"
)

const version = "2020.1.2.28"
const programName = "System Service"
const programDesription = "Creates database and checks system data"
const deleteLogsAfter = 240 * time.Hour
const downloadInSeconds = 86400

type Result struct {
	Size float32
}

type program struct{}

func (p *program) Start(s service.Service) error {
	LogInfo("MAIN", "Starting "+programName+" on "+s.Platform())
	go p.run()
	return nil
}

func (p *program) run() {
	LogDirectoryFileCheck("MAIN")
	LogInfo("MAIN", "Program version "+version+" started")
	CreateConfigIfNotExists()
	LoadSettingsFromConfigFile()
	LogDebug("MAIN", "Using ["+DatabaseType+"] on "+DatabaseIpAddress+":"+DatabasePort+" with database "+DatabaseName)
	WriteProgramVersionIntoSettings()
	for {
		start := time.Now()
		LogInfo("MAIN", "Program running")
		DeleteOldLogFiles()
		CompleteDatabaseCheck()
		if time.Since(start) < (downloadInSeconds * time.Second) {
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
		}
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
	connectionString, dialect := zapsi_database.CheckDatabaseType(DatabaseType, DatabaseIpAddress, DatabasePort, DatabaseLogin, DatabaseName, DatabasePassword)
	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		LogError("MAIN", "Problem opening "+DatabaseName+" database: "+err.Error())
		return ""
	}
	var company zapsi_database.Setting
	db.Where("name=?", "email").Find(&company)
	defer db.Close()
	LogInfo("MAIN", "Recipient email ["+company.Value+"] downloaded from database, elapsed: "+time.Since(timer).String())
	return company.Value
}

func GetCompanyName() string {
	LogInfo("MAIN", "Downloading company name")
	timer := time.Now()
	connectionString, dialect := zapsi_database.CheckDatabaseType(DatabaseType, DatabaseIpAddress, DatabasePort, DatabaseLogin, DatabaseName, DatabasePassword)
	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		LogError("MAIN", "Problem opening "+DatabaseName+" database: "+err.Error())
		return ""
	}
	var company zapsi_database.Setting
	db.Where("name=?", "company").Find(&company)
	defer db.Close()
	LogInfo("MAIN", "Company name ["+company.Value+"] downloaded from database, elapsed: "+time.Since(timer).String())
	return company.Value
}

func UpdateMailSettings() (error, string, int, string, string, string) {
	LogInfo("MAIN", "Updating mail settings")
	timer := time.Now()
	connectionString, dialect := zapsi_database.CheckDatabaseType(DatabaseType, DatabaseIpAddress, DatabasePort, DatabaseLogin, DatabaseName, DatabasePassword)
	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		LogError("MAIN", "Problem opening "+DatabaseName+" database: "+err.Error())
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
	defer db.Close()
	LogInfo("MAIN", "Mail settings updated, elapsed: "+time.Since(timer).String())
	return err, host, port, username, password, email
}

func WriteNewSystemRecord(databaseSizeMegaBytes float32, databaseGrowthInMegaBytes float32, discSpaceMegaBytes float32, estimatedDiscSpaceDays float32) {
	LogInfo("MAIN", "Writing new record to database")
	timer := time.Now()
	connectionString, dialect := zapsi_database.CheckDatabaseType(DatabaseType, DatabaseIpAddress, DatabasePort, DatabaseLogin, DatabaseName, DatabasePassword)
	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		LogError("MAIN", "Problem opening "+DatabaseName+" database: "+err.Error())
		return
	}
	defer db.Close()
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
	platform := service.Platform()
	if strings.Contains(platform, "windows") {
		usage := du.NewDiskUsage("C:\\")
		LogInfo("MAIN", "Disc free space calculated, elapsed: "+time.Since(timer).String())
		return float32(usage.Available() / (1024 * 1024))
	}
	usage := du.NewDiskUsage("/")
	LogInfo("MAIN", "Disc free space calculated, elapsed: "+time.Since(timer).String())
	return float32(usage.Available() / (1024 * 1024))
}

func GetGrowthOfDatabase(actualDatabaseSize float32) float32 {
	LogInfo("MAIN", "Getting database growth")
	timer := time.Now()
	connectionString, dialect := zapsi_database.CheckDatabaseType(DatabaseType, DatabaseIpAddress, DatabasePort, DatabaseLogin, DatabaseName, DatabasePassword)
	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		LogError("MAIN", "Problem opening "+DatabaseName+" database: "+err.Error())
		return 0
	}
	defer db.Close()
	var systemRecord zapsi_database.SystemRecord
	db.Last(&systemRecord)
	growth := actualDatabaseSize - systemRecord.DatabaseSizeInMegaBytes
	LogInfo("MAIN", "Database growth calculated, elapsed: "+time.Since(timer).String())
	return growth
}

func GetDatabaseSize() float32 {
	LogInfo("MAIN", "Getting database size")
	timer := time.Now()
	connectionString, dialect := zapsi_database.CheckDatabaseType(DatabaseType, DatabaseIpAddress, DatabasePort, DatabaseLogin, DatabaseName, DatabasePassword)
	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		LogError("MAIN", "Problem opening "+DatabaseName+" database: "+err.Error())
		return 0
	}
	defer db.Close()
	var result Result
	switch DatabaseType {
	case "postgres":
		db.Raw("SELECT pg_database_size('" + DatabaseName + "')/1000000 as Size;").Scan(&result)
	case "mariadb":
		db.Raw("SELECT SUM(ROUND(((DATA_LENGTH + INDEX_LENGTH) / 1024 / 1024), 2)) AS \"size\" FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = \"" + DatabaseName + "\"").Scan(&result)
	case "mssql":
		db.Raw("SELECT SUM((size*8)/1024) size FROM sys.master_files WHERE DB_NAME(database_id) = '" + DatabaseName + "'").Scan(&result)
	}

	LogInfo("MAIN", "Database size calculated, elapsed: "+time.Since(timer).String())
	return result.Size
}
func (p *program) Stop(s service.Service) error {
	LogInfo("MAIN", "Stopped on platform "+s.Platform())
	return nil
}

func main() {
	serviceConfig := &service.Config{
		Name:        programName,
		DisplayName: programName,
		Description: programDesription,
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

func CompleteDatabaseCheck() {
	firstRunCheckComplete := false
	for firstRunCheckComplete == false {
		databaseOk := zapsi_database.CheckDatabase(DatabaseType, DatabaseIpAddress, DatabasePort, DatabaseLogin, DatabaseName, DatabasePassword)
		tablesOk, err := zapsi_database.CheckTables(DatabaseType, DatabaseIpAddress, DatabasePort, DatabaseLogin, DatabaseName, DatabasePassword)
		if err != nil {
			LogInfo("MAIN", "Problem creating tables: "+err.Error())
		}
		if databaseOk && tablesOk {
			WriteProgramVersionIntoSettings()
			firstRunCheckComplete = true
		}
	}
}

func WriteProgramVersionIntoSettings() {
	connectionString, dialect := zapsi_database.CheckDatabaseType(DatabaseType, DatabaseIpAddress, DatabasePort, DatabaseLogin, DatabaseName, DatabasePassword)
	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		LogError("MAIN", "Problem opening "+DatabaseName+" database: "+err.Error())
		return
	}
	defer db.Close()
	var settings zapsi_database.Setting
	db.Where("name=?", programName).Find(&settings)
	settings.Name = programName
	settings.Value = version
	db.Save(&settings)
	LogDebug("MAIN", "Updated version in database for "+programName)
}
