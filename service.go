package main

import (
	"database/sql"
	"github.com/petrjahoda/database"
	"github.com/ricochet2200/go-disk-usage/du"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func SaveNewSystemRecord(databaseSizeMegaBytes float32, databaseGrowthInMegaBytes float32, discSpaceMegaBytes float32, estimatedDiscSpaceDays float32) {
	LogInfo("MAIN", "Saving new system record")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening database: "+err.Error())
		return
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	var newSystemRecord database.SystemRecord
	newSystemRecord.DatabaseSizeInMegaBytes = databaseSizeMegaBytes
	newSystemRecord.DatabaseGrowthInMegaBytes = databaseGrowthInMegaBytes
	newSystemRecord.DiscFreeSizeInMegaBytes = discSpaceMegaBytes
	newSystemRecord.EstimatedDiscFreeSizeInDays = estimatedDiscSpaceDays
	db.Save(&newSystemRecord)
	LogInfo("MAIN", "New system record saved in: "+time.Since(timer).String())

}

func CalculateEstimatedDiscSpaceInDays(databaseGrowthInMegaBytes float32, discSpaceMegaBytes float32) float32 {
	return discSpaceMegaBytes / databaseGrowthInMegaBytes
}

func CalculateFreeDiscSpace() float32 {
	LogInfo("MAIN", "Getting free disk space")
	timer := time.Now()
	usage := du.NewDiskUsage("/")
	LogInfo("MAIN", "Disc free space calculated in "+time.Since(timer).String())
	return float32(usage.Available() / (1024 * 1024))
}

func CalculateGrowthOfDatabase(actualDatabaseSize float32) float32 {
	LogInfo("MAIN", "Calculating database growth")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening database: "+err.Error())
		return 0
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	var systemRecord database.SystemRecord
	db.Last(&systemRecord)
	growth := actualDatabaseSize - systemRecord.DatabaseSizeInMegaBytes
	LogInfo("MAIN", "Database growth calculated in "+time.Since(timer).String())
	return growth
}

func ReadDatabaseSize() float32 {
	LogInfo("MAIN", "Reading database size")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening database: "+err.Error())
		return 0
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	var size float32
	db.Raw("SELECT pg_database_size('version3')/1000000 as size;").Scan(&size)
	LogInfo("MAIN", "Database size read in "+time.Since(timer).String())
	return size
}

func CheckDatabase() {
	LogInfo("MAIN", "Complete database check started")
	start := time.Now()
	firstRunCheckComplete := false
	for firstRunCheckComplete == false {
		processRunning = true
		databaseOk := CheckDatabaseOnly()
		if databaseOk {
			tablesOk := CheckTablesOnly()
			if tablesOk {
				WriteProgramVersionIntoSettings()
				firstRunCheckComplete = true
			} else {
				time.Sleep(10 * time.Second)
			}
		} else {
			time.Sleep(10 * time.Second)
		}
		processRunning = false
	}
	LogInfo("MAIN", "Database completely checked in "+time.Since(start).String())
}

func CheckTablesOnly() bool {
	LogInfo("MAIN", "Checking tables")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening database: "+err.Error())
		return false
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	if !db.Migrator().HasTable(&database.DeviceType{}) {
		LogInfo("MAIN", "DeviceTypeId table not exists, creating")
		err := db.Migrator().CreateTable(&database.DeviceType{})
		if err != nil {
			LogError("MAIN", "Cannot create devicetype table")
		}
		zapsi := database.DeviceType{Name: "Zapsi"}
		db.Create(&zapsi)
		zapsiTouch := database.DeviceType{Name: "Zapsi Touch"}
		db.Create(&zapsiTouch)
	} else {
		err := db.Migrator().AutoMigrate(&database.DeviceType{})
		if err != nil {
			LogError("MAIN", "Cannot update devicetype table")
		}
	}
	if !db.Migrator().HasTable(&database.Device{}) {
		LogInfo("MAIN", "Device table not exists, creating")
		err := db.Migrator().CreateTable(&database.Device{})
		if err != nil {
			LogError("MAIN", "Cannot create device table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.Device{})
		if err != nil {
			LogError("MAIN", "Cannot update device table")
		}
	}
	if !db.Migrator().HasTable(&database.Setting{}) {
		LogInfo("MAIN", "Setting table not exists, creating")
		err := db.Migrator().CreateTable(&database.Setting{})
		if err != nil {
			LogError("MAIN", "Cannot create setting table")
		}
		company := database.Setting{Name: "company", Value: "Company"}
		db.Create(&company)
		timezone := database.Setting{Name: "timezone", Value: "Europe/Prague"}
		db.Create(&timezone)
		host := database.Setting{Name: "host", Value: ""}
		db.Create(&host)
		port := database.Setting{Name: "port", Value: ""}
		db.Create(&port)
		username := database.Setting{Name: "username", Value: ""}
		db.Create(&username)
		password := database.Setting{Name: "password", Value: ""}
		db.Create(&password)
		email := database.Setting{Name: "email", Value: ""}
		db.Create(&email)
	} else {
		err := db.Migrator().AutoMigrate(&database.Setting{})
		if err != nil {
			LogError("MAIN", "Cannot update setting table")
		}
	}
	if !db.Migrator().HasTable(&database.DevicePortType{}) {
		LogInfo("MAIN", "DevicePortType table not exists, creating")
		err := db.Migrator().CreateTable(&database.DevicePortType{})
		if err != nil {
			LogError("MAIN", "Cannot create deviceporttype table")
		}
		digital := database.DevicePortType{Name: "Digital"}
		db.Create(&digital)
		analog := database.DevicePortType{Name: "Analog"}
		db.Create(&analog)
		serial := database.DevicePortType{Name: "Serial"}
		db.Create(&serial)
		special := database.DevicePortType{Name: "Special"}
		db.Create(&special)
	} else {
		err := db.Migrator().AutoMigrate(&database.DevicePortType{})
		if err != nil {
			LogError("MAIN", "Cannot update deviceporttype table")
		}
	}
	if !db.Migrator().HasTable(&database.DevicePort{}) {
		LogInfo("MAIN", "DevicePort table not exists, creating")
		err := db.Migrator().CreateTable(&database.DevicePort{})
		if err != nil {
			LogError("MAIN", "Cannot create deviceport table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.DevicePort{})
		if err != nil {
			LogError("MAIN", "Cannot update deviceport table")
		}
	}
	if !db.Migrator().HasTable(&database.DevicePortAnalogRecord{}) {
		LogInfo("MAIN", "DevicePortAnalogRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.DevicePortAnalogRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create deviceportanalogrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.DevicePortAnalogRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update deviceportanalogrecord table")
		}
	}
	if !db.Migrator().HasTable(&database.DevicePortDigitalRecord{}) {
		LogInfo("MAIN", "DevicePortDigitalRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.DevicePortDigitalRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create deviceportdigitalrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.DevicePortDigitalRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update deviceportdigitalrecord table")
		}
	}
	if !db.Migrator().HasTable(&database.DevicePortSerialRecord{}) {
		LogInfo("MAIN", "DevicePortSerialRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.DevicePortSerialRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create deviceportserialrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.DevicePortSerialRecord{})
		if err != nil {
			LogError("MAIN", "Cannot UPDATE deviceportserialrecord table")
		}
	}
	if !db.Migrator().HasTable(&database.PackageType{}) {
		LogInfo("MAIN", "PackageType table not exists, creating")
		err := db.Migrator().CreateTable(&database.PackageType{})
		if err != nil {
			LogError("MAIN", "Cannot create packagetype table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.PackageType{})
		if err != nil {
			LogError("MAIN", "Cannot UPDATE packagetype table")
		}
	}
	if !db.Migrator().HasTable(&database.Product{}) {
		LogInfo("MAIN", "Product table not exists, creating")
		err := db.Migrator().CreateTable(&database.Product{})
		if err != nil {
			LogError("MAIN", "Cannot create product table")
		}
		product := database.Product{Name: "Product"}
		db.Create(&product)
	} else {
		err := db.Migrator().AutoMigrate(&database.Product{})
		if err != nil {
			LogError("MAIN", "Cannot update product table")
		}
	}
	if !db.Migrator().HasTable(&database.WorkplaceSection{}) {
		LogInfo("MAIN", "WorkplaceSection table not exists, creating")
		err := db.Migrator().CreateTable(&database.WorkplaceSection{})
		if err != nil {
			LogError("MAIN", "Cannot create workplacesection table")
		}
		machines := database.WorkplaceSection{Name: "Machines"}
		db.Create(&machines)
	} else {
		err := db.Migrator().AutoMigrate(&database.WorkplaceSection{})
		if err != nil {
			LogError("MAIN", "Cannot update workplacesection table")
		}
	}
	if !db.Migrator().HasTable(&database.WorkplaceMode{}) {
		LogInfo("MAIN", "Workplacemode table not exists, creating")
		err := db.Migrator().CreateTable(&database.WorkplaceMode{})
		if err != nil {
			LogError("MAIN", "Cannot create workplacemode table")
		}
		mode := database.WorkplaceMode{Name: "Production", DowntimeDuration: time.Second * 300, PoweroffDuration: time.Second * 300}
		db.Create(&mode)
	} else {
		err := db.Migrator().AutoMigrate(&database.WorkplaceMode{})
		if err != nil {
			LogError("MAIN", "Cannot update workplacemode table")
		}
	}
	if !db.Migrator().HasTable(&database.Workplace{}) {
		LogInfo("MAIN", "Workplace table not exists, creating")
		err := db.Migrator().CreateTable(&database.Workplace{})
		if err != nil {
			LogError("MAIN", "Cannot create workplace table")
		}

	} else {
		err := db.Migrator().AutoMigrate(&database.Workplace{})
		if err != nil {
			LogError("MAIN", "Cannot update workplace table")
		}
	}

	if !db.Migrator().HasTable(&database.Order{}) {
		LogInfo("MAIN", "Order table not exists, creating")
		err := db.Migrator().CreateTable(&database.Order{})
		if err != nil {
			LogError("MAIN", "Cannot create order table")
		}
		order := database.Order{Name: "Internal", DateTimeRequest: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}}
		db.Create(&order)
	} else {
		err := db.Migrator().AutoMigrate(&database.Order{})
		if err != nil {
			LogError("MAIN", "Cannot update order table")
		}
	}
	if !db.Migrator().HasTable(&database.Package{}) {
		LogInfo("MAIN", "Package table not exists, creating")
		err := db.Migrator().CreateTable(&database.Package{})
		if err != nil {
			LogError("MAIN", "Cannot create package table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.Package{})
		if err != nil {
			LogError("MAIN", "Cannot update package table")
		}
	}

	if !db.Migrator().HasTable(&database.FaultType{}) {
		LogInfo("MAIN", "FaultType table not exists, creating")
		err := db.Migrator().CreateTable(&database.FaultType{})
		if err != nil {
			LogError("MAIN", "Cannot create faulttype table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.FaultType{})
		if err != nil {
			LogError("MAIN", "Cannot update faulttype table")
		}
	}
	if !db.Migrator().HasTable(&database.Fault{}) {
		LogInfo("MAIN", "Fault table not exists, creating")
		err := db.Migrator().CreateTable(&database.Fault{})
		if err != nil {
			LogError("MAIN", "Cannot create fault table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.Fault{})
		if err != nil {
			LogError("MAIN", "Cannot update fault table")
		}
	}

	if !db.Migrator().HasTable(&database.BreakdownType{}) {
		LogInfo("MAIN", "BreakdownType table not exists, creating")
		err := db.Migrator().CreateTable(&database.BreakdownType{})
		if err != nil {
			LogError("MAIN", "Cannot create breakdowntype table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.BreakdownType{})
		if err != nil {
			LogError("MAIN", "Cannot create update table")
		}
	}
	if !db.Migrator().HasTable(&database.Breakdown{}) {
		LogInfo("MAIN", "Breakdown table not exists, creating")
		err := db.Migrator().CreateTable(&database.Breakdown{})
		if err != nil {
			LogError("MAIN", "Cannot create breakdown table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.Breakdown{})
		if err != nil {
			LogError("MAIN", "Cannot update breakdown table")
		}
	}

	if !db.Migrator().HasTable(&database.DowntimeType{}) {
		LogInfo("MAIN", "DowntimeType table not exists, creating")
		err := db.Migrator().CreateTable(&database.DowntimeType{})
		if err != nil {
			LogError("MAIN", "Cannot create downtimetype table")
		}
		system := database.DowntimeType{Name: "System"}
		db.Create(&system)
	} else {
		err := db.Migrator().AutoMigrate(&database.DowntimeType{})
		if err != nil {
			LogError("MAIN", "Cannot update downtimetype table")
		}
	}
	if !db.Migrator().HasTable(&database.Downtime{}) {
		LogInfo("MAIN", "Downtime table not exists, creating")
		err := db.Migrator().CreateTable(&database.Downtime{})
		if err != nil {
			LogError("MAIN", "Cannot create downtime table")
		}
		system := database.DowntimeType{}
		db.Where("Name = ?", "System").Find(&system)
		noReasonDowntime := database.Downtime{
			Name:           "No reason Downtime",
			DowntimeTypeID: 1,
		}
		db.Create(&noReasonDowntime)
	} else {
		err := db.Migrator().AutoMigrate(&database.Downtime{})
		if err != nil {
			LogError("MAIN", "Cannot update downtime table")
		}
	}

	if !db.Migrator().HasTable(&database.UserType{}) {
		LogInfo("MAIN", "UserType table not exists, creating")
		err := db.Migrator().CreateTable(&database.UserType{})
		if err != nil {
			LogError("MAIN", "Cannot create usertype table")
		}
		operator := database.UserType{Name: "Operator"}
		db.Create(&operator)
		admin := database.UserType{Name: "Admin"}
		db.Create(&admin)
	} else {
		err := db.Migrator().AutoMigrate(&database.UserType{})
		if err != nil {
			LogError("MAIN", "Cannot update usertype table")
		}
	}

	if !db.Migrator().HasTable(&database.UserRole{}) {
		LogInfo("MAIN", "UserRole table not exists, creating")
		err := db.Migrator().CreateTable(&database.UserRole{})
		if err != nil {
			LogError("MAIN", "Cannot create userrole table")
		}
		admin := database.UserRole{Name: "Administrator"}
		db.Create(&admin)
		powerUser := database.UserRole{Name: "PowerUser"}
		db.Create(&powerUser)
		user := database.UserRole{Name: "User"}
		db.Create(&user)
	} else {
		err := db.Migrator().AutoMigrate(&database.UserRole{})
		if err != nil {
			LogError("MAIN", "Cannot update userrole table")
		}
	}
	if !db.Migrator().HasTable(&database.User{}) {
		LogInfo("MAIN", "User table not exists, creating")
		err := db.Migrator().CreateTable(&database.User{})
		if err != nil {
			LogError("MAIN", "Cannot create user table")
		}
		userRole := database.UserRole{}
		db.Where("Name = ?", "Administrator").Find(&userRole)
		userType := database.UserType{}
		db.Where("Name = ?", "Admin").Find(&userType)
		password := hashAndSalt([]byte("54321"))
		adminUser := database.User{
			FirstName:  "Admin",
			SecondName: "Admin",
			Password:   password,
			UserRoleID: 1,
			UserTypeID: 1,
		}
		db.Create(&adminUser)
	} else {
		err := db.Migrator().AutoMigrate(&database.User{})
		if err != nil {
			LogError("MAIN", "Cannot update user table")
		}
	}

	if !db.Migrator().HasTable(&database.Workshift{}) {
		LogInfo("MAIN", "Workshift table not exists, creating")
		err := db.Migrator().CreateTable(&database.Workshift{})
		if err != nil {
			LogError("MAIN", "Cannot create workshift table")
		}
		firstShiftStart := time.Date(2000, 1, 1, 6, 0, 0, 0, time.Local)
		firstShiftEnd := time.Date(2000, 1, 1, 14, 0, 0, 0, time.Local)
		firstShift := database.Workshift{Name: "First Shift", WorkshiftStart: firstShiftStart, WorkshiftEnd: firstShiftEnd}
		db.Create(&firstShift)
		secondShiftStart := time.Date(2000, 1, 1, 14, 0, 0, 0, time.Local)
		secondShiftEnd := time.Date(2000, 1, 1, 22, 0, 0, 0, time.Local)
		secondShift := database.Workshift{Name: "Second Shift", WorkshiftStart: secondShiftStart, WorkshiftEnd: secondShiftEnd}
		db.Create(&secondShift)
		thirdShiftStart := time.Date(2000, 1, 1, 22, 0, 0, 0, time.Local)
		thirdShiftEnd := time.Date(2000, 1, 2, 6, 0, 0, 0, time.Local)
		thirdShift := database.Workshift{Name: "Third Shift", WorkshiftStart: thirdShiftStart, WorkshiftEnd: thirdShiftEnd}
		db.Create(&thirdShift)
	} else {
		err := db.Migrator().AutoMigrate(&database.Workshift{})
		if err != nil {
			LogError("MAIN", "Cannot update workshift table")
		}
	}

	if !db.Migrator().HasTable(&database.Part{}) {
		LogInfo("MAIN", "Part table not exists, creating")
		err := db.Migrator().CreateTable(&database.Part{})
		if err != nil {
			LogError("MAIN", "Cannot create part table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.Part{})
		if err != nil {
			LogError("MAIN", "Cannot update part table")
		}
	}

	if !db.Migrator().HasTable(&database.State{}) {
		LogInfo("MAIN", "State table not exists, creating")
		err := db.Migrator().CreateTable(&database.State{})
		if err != nil {
			LogError("MAIN", "Cannot create state table")
		}
		production := database.State{Name: "Production", Color: "#89AB0F"}
		db.Create(&production)
		downtime := database.State{Name: "Downtime", Color: "#E6AD3C"}
		db.Create(&downtime)
		poweroff := database.State{Name: "Poweroff", Color: "#DE6B59"}
		db.Create(&poweroff)
	} else {
		err := db.Migrator().AutoMigrate(&database.State{})
		if err != nil {
			LogError("MAIN", "Cannot update state table")
		}
	}

	if !db.Migrator().HasTable(&database.WorkplaceWorkshift{}) {
		LogInfo("MAIN", "WorkplaceWorkshift table not exists, creating")
		err := db.Migrator().CreateTable(&database.WorkplaceWorkshift{})
		if err != nil {
			LogError("MAIN", "Cannot create workplaceworkshift table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.Workshift{})
		if err != nil {
			LogError("MAIN", "Cannot update workplaceworkshift table")
		}
	}

	if !db.Migrator().HasTable(&database.WorkplacePort{}) {
		LogInfo("MAIN", "WorkplacePort table not exists, creating")
		err := db.Migrator().CreateTable(&database.WorkplacePort{})
		if err != nil {
			LogError("MAIN", "Cannot create workplaceport table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.WorkplacePort{})
		if err != nil {
			LogError("MAIN", "Cannot update workplaceport table")
		}
	}
	if !db.Migrator().HasTable(&database.StateRecord{}) {
		LogInfo("MAIN", "StateRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.StateRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create staterecord table")
		}

	} else {
		err := db.Migrator().AutoMigrate(&database.StateRecord{})
		if err != nil {
			LogError("MAIN", "update create staterecord table")
		}
	}

	if !db.Migrator().HasTable(&database.Operation{}) {
		LogInfo("MAIN", "Operation table not exists, creating")
		err := db.Migrator().CreateTable(&database.Operation{})
		if err != nil {
			LogError("MAIN", "Cannot create operation table")
		}
		operation := database.Operation{
			Name:    "Operation",
			OrderID: 1,
			Barcode: 0,
			Note:    "",
		}
		db.Create(&operation)
	} else {
		err := db.Migrator().AutoMigrate(&database.Operation{})
		if err != nil {
			LogError("MAIN", "Cannot update operation table")
		}
	}

	if !db.Migrator().HasTable(&database.OrderRecord{}) {
		LogInfo("MAIN", "OrderRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.OrderRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create orderrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.OrderRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update orderrecord table")
		}
	}

	if !db.Migrator().HasTable(&database.PartRecord{}) {
		LogInfo("MAIN", "PartRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.PartRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create partrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.PartRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update partrecord table")
		}
	}

	if !db.Migrator().HasTable(&database.PackageRecord{}) {
		LogInfo("MAIN", "PackageRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.PackageRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create packagerecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.PackageRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update packagerecord table")
		}
	}

	if !db.Migrator().HasTable(&database.FaultRecord{}) {
		LogInfo("MAIN", "FaultRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.FaultRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create faultrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.FaultRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update faultrecord table")
		}
	}

	if !db.Migrator().HasTable(&database.BreakdownRecord{}) {
		LogInfo("MAIN", "BreakdownRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.BreakdownRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create breakdownrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.BreakdownRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update breakdownrecord table")
		}
	}

	if !db.Migrator().HasTable(&database.DownTimeRecord{}) {
		LogInfo("MAIN", "DownTimeRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.DownTimeRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create downtimerecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.DownTimeRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update downtimerecord table")
		}
	}

	if !db.Migrator().HasTable(&database.UserRecord{}) {
		LogInfo("MAIN", "UserRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.UserRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create userrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.UserRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update userrecord table")
		}
	}

	if !db.Migrator().HasTable(&database.Alarm{}) {
		LogInfo("MAIN", "Alarm table not exists, creating")
		err := db.Migrator().CreateTable(&database.Alarm{})
		if err != nil {
			LogError("MAIN", "Cannot create alarm table")
		}

	} else {
		err := db.Migrator().AutoMigrate(&database.Alarm{})
		if err != nil {
			LogError("MAIN", "Cannot update alarm table")
		}
	}
	if !db.Migrator().HasTable(&database.AlarmRecord{}) {
		LogInfo("MAIN", "AlarmRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.AlarmRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create alarmrecord table")
		}

	} else {
		err := db.Migrator().AutoMigrate(&database.AlarmRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update alarmrecord table")
		}
	}

	if !db.Migrator().HasTable(&database.SystemRecord{}) {
		LogInfo("MAIN", "SystemRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.SystemRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create systemrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.SystemRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update systemrecord table")
		}
	}
	if !db.Migrator().HasTable(&database.DeviceWorkplaceRecord{}) {
		LogInfo("MAIN", "DeviceWorkplaceRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.DeviceWorkplaceRecord{})
		if err != nil {
			LogError("MAIN", "Cannot create deviceworkplacerecord table")
		}

	} else {
		err := db.Migrator().AutoMigrate(&database.DeviceWorkplaceRecord{})
		if err != nil {
			LogError("MAIN", "Cannot update deviceworkplacerecord table")
		}
	}
	LogInfo("MAIN", "Tables checked in "+time.Since(timer).String())
	return true

}

func CheckDatabaseOnly() bool {
	LogInfo("MAIN", "Checking database")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogWarning("MAIN", "Database does not exist")
		db, err := gorm.Open(postgres.Open(postgresConfig), &gorm.Config{})
		if err != nil {
			LogError("MAIN", "Problem opening database: "+err.Error())
			return false
		}
		sqlDB, err := db.DB()
		defer sqlDB.Close()
		db = db.Exec("CREATE DATABASE version3;")
		if db.Error != nil {
			LogError("MAIN", "Cannot create database version3: "+err.Error())
		}
		LogInfo("MAIN", "Database created")
		return true

	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	LogInfo("MAIN", "Database checked in "+time.Since(timer).String())
	return true

}

func WriteProgramVersionIntoSettings() {
	LogInfo("MAIN", "Writing program version into settings")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		LogError("MAIN", "Problem opening database: "+err.Error())
		return
	}
	sqlDB, err := db.DB()
	defer sqlDB.Close()
	var existingSettings database.Setting
	db.Where("name=?", serviceName).Find(&existingSettings)
	existingSettings.Name = serviceName
	existingSettings.Value = version
	db.Save(&existingSettings)
	LogInfo("MAIN", "Program version written into settings in "+time.Since(timer).String())
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
