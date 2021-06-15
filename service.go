package main

import (
	"database/sql"
	"github.com/petrjahoda/database"
	"github.com/ricochet2200/go-disk-usage/du"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func createNewSystemRecord(databaseSizeMegaBytes float32, databaseGrowthInMegaBytes float32, discSpaceMegaBytes float32, estimatedDiscSpaceDays float32) {
	logInfo("MAIN", "Saving new system record")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	if err != nil {
		logError("MAIN", "Problem opening database: "+err.Error())
		return
	}
	var newSystemRecord database.SystemRecord
	newSystemRecord.DatabaseSizeInMegaBytes = databaseSizeMegaBytes
	newSystemRecord.DatabaseGrowthInMegaBytes = databaseGrowthInMegaBytes
	newSystemRecord.DiscFreeSizeInMegaBytes = discSpaceMegaBytes
	newSystemRecord.EstimatedDiscFreeSizeInDays = estimatedDiscSpaceDays
	db.Save(&newSystemRecord)
	logInfo("MAIN", "New system record saved in: "+time.Since(timer).String())

}

func calculateEstimatedDiscSpaceInDays(databaseGrowthInMegaBytes float32, discSpaceMegaBytes float32, hours float64) float32 {
	return discSpaceMegaBytes / (databaseGrowthInMegaBytes / (float32(hours) / 24))
}

func calculateFreeDiscSpace() float32 {
	logInfo("MAIN", "Getting free disk space")
	timer := time.Now()
	usage := du.NewDiskUsage("/")
	logInfo("MAIN", "Disc free space calculated in "+time.Since(timer).String())
	return float32(usage.Available() / (1024 * 1024))
}

func readLastSystemRecord() database.SystemRecord {
	logInfo("MAIN", "Reading last system record")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	var systemRecord database.SystemRecord
	if err != nil {
		logError("MAIN", "Problem opening database: "+err.Error())
		return systemRecord
	}
	db.Last(&systemRecord)
	logInfo("MAIN", "Last system record read in "+time.Since(timer).String())
	return systemRecord
}

func readDatabaseSize() float32 {
	logInfo("MAIN", "Reading database size")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	if err != nil {
		logError("MAIN", "Problem opening database: "+err.Error())
		return 0
	}
	var size float32
	db.Raw("SELECT pg_database_size('system')/1000000 as size;").Scan(&size)
	logInfo("MAIN", "Database size read in "+time.Since(timer).String())
	return size
}

func checkDatabase() {
	logInfo("MAIN", "Complete database check started")
	start := time.Now()
	firstRunCheckComplete := false
	for firstRunCheckComplete == false {
		processRunning = true
		databaseOk := checkDatabaseOnly()
		if databaseOk {
			tablesOk := checkTablesOnly()
			if tablesOk {
				updateProgramVersion()
				firstRunCheckComplete = true
			} else {
				time.Sleep(10 * time.Second)
			}
		} else {
			time.Sleep(10 * time.Second)
		}
		processRunning = false
	}
	logInfo("MAIN", "Database completely checked in "+time.Since(start).String())
}

func checkTablesOnly() bool {
	logInfo("MAIN", "Checking tables")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	if err != nil {
		logError("MAIN", "Problem opening database: "+err.Error())
		return false
	}
	// DEVICE TYPE
	if !db.Migrator().HasTable(&database.DeviceType{}) {
		logInfo("MAIN", "DeviceTypeId table not exists, creating")
		err := db.Migrator().CreateTable(&database.DeviceType{})
		if err != nil {
			logError("MAIN", "Cannot create devicetype table")
		}
		zapsi := database.DeviceType{Name: "Zapsi"}
		db.Create(&zapsi)
		zapsiTouch := database.DeviceType{Name: "Zapsi Touch"}
		db.Create(&zapsiTouch)
		siemens := database.DeviceType{Name: "S7 based Siemens PLCs"}
		db.Create(&siemens)
		printer := database.DeviceType{Name: "Datamax printers"}
		db.Create(&printer)
		file := database.DeviceType{Name: "File based devices"}
		db.Create(&file)
		network := database.DeviceType{Name: "Network based devices"}
		db.Create(&network)
	} else {
		err := db.Migrator().AutoMigrate(&database.DeviceType{})
		if err != nil {
			logError("MAIN", "Cannot update devicetype table")
		}
	}

	// DEVICE
	if !db.Migrator().HasTable(&database.Device{}) {
		logInfo("MAIN", "Device table not exists, creating")
		err := db.Migrator().CreateTable(&database.Device{})
		if err != nil {
			logError("MAIN", "Cannot create device table")
		}
		db.Exec("ALTER TABLE devices ADD CONSTRAINT fk_devices_device_type FOREIGN KEY (device_type_id) REFERENCES device_types(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.Device{})
		if err != nil {
			logError("MAIN", "Cannot update device table")
		}
	}

	// SETTING
	if !db.Migrator().HasTable(&database.Setting{}) {
		logInfo("MAIN", "Setting table not exists, creating")
		err := db.Migrator().CreateTable(&database.Setting{})
		if err != nil {
			logError("MAIN", "Cannot create setting table")
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
		software := database.Setting{Name: "software", Value: "Zapsi"}
		db.Create(&software)
	} else {
		err := db.Migrator().AutoMigrate(&database.Setting{})
		if err != nil {
			logError("MAIN", "Cannot update setting table")
		}
	}

	// DEVICE PORT TYPE
	if !db.Migrator().HasTable(&database.DevicePortType{}) {
		logInfo("MAIN", "DevicePortType table not exists, creating")
		err := db.Migrator().CreateTable(&database.DevicePortType{})
		if err != nil {
			logError("MAIN", "Cannot create deviceporttype table")
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
			logError("MAIN", "Cannot update deviceporttype table")
		}
	}

	// DEVICE PORT
	if !db.Migrator().HasTable(&database.DevicePort{}) {
		logInfo("MAIN", "DevicePort table not exists, creating")
		err := db.Migrator().CreateTable(&database.DevicePort{})
		if err != nil {
			logError("MAIN", "Cannot create deviceport table")
		}
		db.Exec("ALTER TABLE device_ports ADD CONSTRAINT fk_device_ports_device FOREIGN KEY (device_id) REFERENCES devices(id)")
		db.Exec("ALTER TABLE device_ports ADD CONSTRAINT fk_device_ports_device_port_type FOREIGN KEY (device_port_type_id) REFERENCES device_port_types(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.DevicePort{})
		if err != nil {
			logError("MAIN", "Cannot update deviceport table")
		}
	}

	// DEVICE PORT ANALOG RECORD
	if !db.Migrator().HasTable(&database.DevicePortAnalogRecord{}) {
		logInfo("MAIN", "DevicePortAnalogRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.DevicePortAnalogRecord{})
		if err != nil {
			logError("MAIN", "Cannot create deviceportanalogrecord table")
		}
		db.Exec("ALTER TABLE device_port_analog_records ADD CONSTRAINT fk_device_port_analog_records_device_port FOREIGN KEY (device_port_id) REFERENCES device_ports(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.DevicePortAnalogRecord{})
		if err != nil {
			logError("MAIN", "Cannot update deviceportanalogrecord table")
		}
	}

	// DEVICE PORT DIGITAL RECORD
	if !db.Migrator().HasTable(&database.DevicePortDigitalRecord{}) {
		logInfo("MAIN", "DevicePortDigitalRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.DevicePortDigitalRecord{})
		if err != nil {
			logError("MAIN", "Cannot create deviceportdigitalrecord table")
		}
		db.Exec("ALTER TABLE device_port_digital_records ADD CONSTRAINT fk_device_port_digital_records_device_port FOREIGN KEY (device_port_id) REFERENCES device_ports(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.DevicePortDigitalRecord{})
		if err != nil {
			logError("MAIN", "Cannot update deviceportdigitalrecord table")
		}
	}

	// DEVICE PORT SERIAL RECORD
	if !db.Migrator().HasTable(&database.DevicePortSerialRecord{}) {
		logInfo("MAIN", "DevicePortSerialRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.DevicePortSerialRecord{})
		if err != nil {
			logError("MAIN", "Cannot create deviceportserialrecord table")
		}
		db.Exec("ALTER TABLE device_port_serial_records ADD CONSTRAINT fk_device_port_serial_records_device_port FOREIGN KEY (device_port_id) REFERENCES device_ports(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.DevicePortSerialRecord{})
		if err != nil {
			logError("MAIN", "Cannot UPDATE deviceportserialrecord table")
		}
	}

	// PACKAGE TYPE
	if !db.Migrator().HasTable(&database.PackageType{}) {
		logInfo("MAIN", "PackageType table not exists, creating")
		err := db.Migrator().CreateTable(&database.PackageType{})
		if err != nil {
			logError("MAIN", "Cannot create packagetype table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.PackageType{})
		if err != nil {
			logError("MAIN", "Cannot UPDATE packagetype table")
		}
	}

	// PRODUCT
	if !db.Migrator().HasTable(&database.Product{}) {
		logInfo("MAIN", "Product table not exists, creating")
		err := db.Migrator().CreateTable(&database.Product{})
		if err != nil {
			logError("MAIN", "Cannot create product table")
		}
		product := database.Product{Name: "Testovací výrobek"}
		db.Create(&product)
	} else {
		err := db.Migrator().AutoMigrate(&database.Product{})
		if err != nil {
			logError("MAIN", "Cannot update product table")
		}
	}

	// WORKPLACE SECTION
	if !db.Migrator().HasTable(&database.WorkplaceSection{}) {
		logInfo("MAIN", "WorkplaceSection table not exists, creating")
		err := db.Migrator().CreateTable(&database.WorkplaceSection{})
		if err != nil {
			logError("MAIN", "Cannot create workplacesection table")
		}
		machines := database.WorkplaceSection{Name: "Hlavní stroje"}
		db.Create(&machines)
		machinesTwo := database.WorkplaceSection{Name: "Pomocné stroje"}
		db.Create(&machinesTwo)
	} else {
		err := db.Migrator().AutoMigrate(&database.WorkplaceSection{})
		if err != nil {
			logError("MAIN", "Cannot update workplacesection table")
		}
	}

	// WORKPLACE MODE
	if !db.Migrator().HasTable(&database.WorkplaceMode{}) {
		logInfo("MAIN", "Workplacemode table not exists, creating")
		err := db.Migrator().CreateTable(&database.WorkplaceMode{})
		if err != nil {
			logError("MAIN", "Cannot create workplacemode table")
		}
		mode := database.WorkplaceMode{Name: "Produkce", DowntimeDuration: time.Second * 300, PoweroffDuration: time.Second * 300}
		db.Create(&mode)
	} else {
		err := db.Migrator().AutoMigrate(&database.WorkplaceMode{})
		if err != nil {
			logError("MAIN", "Cannot update workplacemode table")
		}
	}

	// WORKPLACE
	if !db.Migrator().HasTable(&database.Workplace{}) {
		logInfo("MAIN", "Workplace table not exists, creating")
		err := db.Migrator().CreateTable(&database.Workplace{})
		if err != nil {
			logError("MAIN", "Cannot create workplace table")
		}
		db.Exec("ALTER TABLE workplaces ADD CONSTRAINT fk_workplaces_workplace_mode FOREIGN KEY (workplace_mode_id) REFERENCES workplace_modes(id)")
		db.Exec("ALTER TABLE workplaces ADD CONSTRAINT fk_workplaces_workplace_section FOREIGN KEY (workplace_section_id) REFERENCES workplace_sections(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.Workplace{})
		if err != nil {
			logError("MAIN", "Cannot update workplace table")
		}
	}

	// ORDER
	if !db.Migrator().HasTable(&database.Order{}) {
		logInfo("MAIN", "Order table not exists, creating")
		err := db.Migrator().CreateTable(&database.Order{})
		if err != nil {
			logError("MAIN", "Cannot create order table")
		}
		db.Exec("ALTER TABLE orders ADD CONSTRAINT fk_orders_product FOREIGN KEY (product_id) REFERENCES products(id)")
		db.Exec("ALTER TABLE orders ADD CONSTRAINT fk_orders_product FOREIGN KEY (product_id) REFERENCES products(id)")
		order := database.Order{Name: "Nespecifikovaná zakázka", ProductID: sql.NullInt32{Int32: 1, Valid: true}, DateTimeRequest: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}}
		db.Create(&order)
		left := database.Order{Name: "Levá mechanika", ProductID: sql.NullInt32{Int32: 1, Valid: true}, DateTimeRequest: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}}
		db.Create(&left)
		right := database.Order{Name: "Pravá mechanika", ProductID: sql.NullInt32{Int32: 1, Valid: true}, DateTimeRequest: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}}
		db.Create(&right)
		upper := database.Order{Name: "Horní uzávěr", ProductID: sql.NullInt32{Int32: 1, Valid: true}, DateTimeRequest: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}}
		db.Create(&upper)
		handler := database.Order{Name: "Držák", ProductID: sql.NullInt32{Int32: 1, Valid: true}, DateTimeRequest: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		}}
		db.Create(&handler)
	} else {
		err := db.Migrator().AutoMigrate(&database.Order{})
		if err != nil {
			logError("MAIN", "Cannot update order table")
		}
	}

	// PACKAGE
	if !db.Migrator().HasTable(&database.Package{}) {
		logInfo("MAIN", "Package table not exists, creating")
		err := db.Migrator().CreateTable(&database.Package{})
		if err != nil {
			logError("MAIN", "Cannot create package table")
		}
		db.Exec("ALTER TABLE packages ADD CONSTRAINT fk_packages_order FOREIGN KEY (order_id) REFERENCES orders(id)")
		db.Exec("ALTER TABLE packages ADD CONSTRAINT fk_packages_package_type FOREIGN KEY (package_type_id) REFERENCES package_types(id)")

	} else {
		err := db.Migrator().AutoMigrate(&database.Package{})
		if err != nil {
			logError("MAIN", "Cannot update package table")
		}
	}

	// FAULT TYPE
	if !db.Migrator().HasTable(&database.FaultType{}) {
		logInfo("MAIN", "FaultType table not exists, creating")
		err := db.Migrator().CreateTable(&database.FaultType{})
		if err != nil {
			logError("MAIN", "Cannot create faulttype table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.FaultType{})
		if err != nil {
			logError("MAIN", "Cannot update faulttype table")
		}
	}

	// FAULT
	if !db.Migrator().HasTable(&database.Fault{}) {
		logInfo("MAIN", "Fault table not exists, creating")
		err := db.Migrator().CreateTable(&database.Fault{})
		if err != nil {
			logError("MAIN", "Cannot create fault table")
		}
		db.Exec("ALTER TABLE faults ADD CONSTRAINT fk_faults_fault_type FOREIGN KEY (fault_type_id) REFERENCES fault_types(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.Fault{})
		if err != nil {
			logError("MAIN", "Cannot update fault table")
		}

	}

	// BREAKDOWN TYPE
	if !db.Migrator().HasTable(&database.BreakdownType{}) {
		logInfo("MAIN", "BreakdownType table not exists, creating")
		err := db.Migrator().CreateTable(&database.BreakdownType{})
		if err != nil {
			logError("MAIN", "Cannot create breakdowntype table")
		}
		electrical := database.BreakdownType{Name: "Elektrická porucha"}
		db.Create(&electrical)
		mechanical := database.BreakdownType{Name: "Mechanická porucha"}
		db.Create(&mechanical)
	} else {
		err := db.Migrator().AutoMigrate(&database.BreakdownType{})
		if err != nil {
			logError("MAIN", "Cannot create update table")
		}
	}

	// BREAKDOWN
	if !db.Migrator().HasTable(&database.Breakdown{}) {
		logInfo("MAIN", "Breakdown table not exists, creating")
		err := db.Migrator().CreateTable(&database.Breakdown{})
		if err != nil {
			logError("MAIN", "Cannot create breakdown table")
		}
		db.Exec("ALTER TABLE breakdowns ADD CONSTRAINT fk_breakdowns_breakdown_type FOREIGN KEY (breakdown_type_id) REFERENCES breakdown_types(id)")
		switches := database.Breakdown{
			Name:            "Vyhozené pojistky",
			BreakdownTypeID: 1,
		}
		db.Create(&switches)
		panel := database.Breakdown{
			Name:            "Nefunguje panel",
			BreakdownTypeID: 1,
		}
		db.Create(&panel)
		button := database.Breakdown{
			Name:            "Nefunguje tlačítko",
			BreakdownTypeID: 2,
		}
		db.Create(&button)
		axis := database.Breakdown{
			Name:            "Drhne osa",
			BreakdownTypeID: 2,
		}
		db.Create(&axis)
	} else {
		err := db.Migrator().AutoMigrate(&database.Breakdown{})
		if err != nil {
			logError("MAIN", "Cannot update breakdown table")
		}

	}

	// DOWNTIME TYPE
	if !db.Migrator().HasTable(&database.DowntimeType{}) {
		logInfo("MAIN", "DowntimeType table not exists, creating")
		err := db.Migrator().CreateTable(&database.DowntimeType{})
		if err != nil {
			logError("MAIN", "Cannot create downtimetype table")
		}
		system := database.DowntimeType{Name: "Systémové"}
		db.Create(&system)
		user := database.DowntimeType{Name: "Uživatelské"}
		db.Create(&user)
	} else {
		err := db.Migrator().AutoMigrate(&database.DowntimeType{})
		if err != nil {
			logError("MAIN", "Cannot update downtimetype table")
		}
	}

	// DOWNTIME
	if !db.Migrator().HasTable(&database.Downtime{}) {
		logInfo("MAIN", "Downtime table not exists, creating")
		err := db.Migrator().CreateTable(&database.Downtime{})
		if err != nil {
			logError("MAIN", "Cannot create downtime table")
		}
		db.Exec("ALTER TABLE downtimes ADD CONSTRAINT fk_downtimes_downtime_type FOREIGN KEY (downtime_type_id) REFERENCES downtime_types(id)")
		system := database.DowntimeType{}
		db.Where("Name = ?", "System").Find(&system)
		noReasonDowntime := database.Downtime{
			Name:           "Nespecifikovaný prostoj",
			DowntimeTypeID: 1,
		}
		db.Create(&noReasonDowntime)
		smoking := database.Downtime{
			Name:           "Kouření",
			DowntimeTypeID: 2,
		}
		db.Create(&smoking)
		cleaning := database.Downtime{
			Name:           "Uklízení",
			DowntimeTypeID: 2,
		}
		db.Create(&cleaning)
		preparing := database.Downtime{
			Name:           "Chystání materiálu",
			DowntimeTypeID: 2,
		}
		db.Create(&preparing)
		hand := database.Downtime{
			Name:           "Ruční operace",
			DowntimeTypeID: 2,
		}
		db.Create(&hand)
	} else {
		err := db.Migrator().AutoMigrate(&database.Downtime{})
		if err != nil {
			logError("MAIN", "Cannot update downtime table")
		}
	}

	// USER TYPE
	if !db.Migrator().HasTable(&database.UserType{}) {
		logInfo("MAIN", "UserType table not exists, creating")
		err := db.Migrator().CreateTable(&database.UserType{})
		if err != nil {
			logError("MAIN", "Cannot create usertype table")
		}
		operator := database.UserType{Name: "Operator"}
		db.Create(&operator)
		admin := database.UserType{Name: "Admin"}
		db.Create(&admin)
	} else {
		err := db.Migrator().AutoMigrate(&database.UserType{})
		if err != nil {
			logError("MAIN", "Cannot update usertype table")
		}
	}

	// USER ROLE
	if !db.Migrator().HasTable(&database.UserRole{}) {
		logInfo("MAIN", "UserRole table not exists, creating")
		err := db.Migrator().CreateTable(&database.UserRole{})
		if err != nil {
			logError("MAIN", "Cannot create userrole table")
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
			logError("MAIN", "Cannot update userrole table")
		}
	}

	//USER
	if !db.Migrator().HasTable(&database.User{}) {
		logInfo("MAIN", "User table not exists, creating")
		err := db.Migrator().CreateTable(&database.User{})
		if err != nil {
			logError("MAIN", "Cannot create user table")
		}
		db.Exec("ALTER TABLE users ADD CONSTRAINT fk_users_user_role FOREIGN KEY (user_role_id) REFERENCES user_roles(id)")
		db.Exec("ALTER TABLE users ADD CONSTRAINT fk_users_user_type FOREIGN KEY (user_type_id) REFERENCES user_types(id)")
		userRole := database.UserRole{}
		db.Where("Name = ?", "Administrator").Find(&userRole)
		userType := database.UserType{}
		db.Where("Name = ?", "Admin").Find(&userType)
		password := hashPasswordFromString([]byte("admin"))
		adminUser := database.User{
			FirstName:  "Admin",
			SecondName: "Admin",
			Email:      "admin@admin.com",
			Password:   password,
			UserRoleID: 1,
			UserTypeID: 2,
			Locale:     "EnUS",
		}
		db.Create(&adminUser)
		password = hashPasswordFromString([]byte("poweruser"))
		powerUser := database.User{
			FirstName:  "Power",
			SecondName: "User",
			Email:      "power@user.com",
			Password:   password,
			UserRoleID: 2,
			UserTypeID: 2,
			Locale:     "DeDE",
		}
		db.Create(&powerUser)
		password = hashPasswordFromString([]byte("user"))
		user := database.User{
			FirstName:  "User",
			SecondName: "User",
			Email:      "user@user.com",
			Password:   password,
			UserRoleID: 3,
			UserTypeID: 2,
			Locale:     "CsCZ",
		}
		db.Create(&user)
		password = hashPasswordFromString([]byte("user"))
		demoUserOne := database.User{
			FirstName:  "Adam",
			SecondName: "Dub",
			Email:      "adam@company.com",
			Password:   password,
			UserRoleID: 3,
			UserTypeID: 1,
			Locale:     "CsCZ",
		}
		db.Create(&demoUserOne)
		demoUserTwo := database.User{
			FirstName:  "David",
			SecondName: "Buk",
			Email:      "david@company.com",
			Password:   password,
			UserRoleID: 3,
			UserTypeID: 1,
			Locale:     "CsCZ",
		}
		db.Create(&demoUserTwo)
		demoUserThree := database.User{
			FirstName:  "Johan",
			SecondName: "Smrk",
			Email:      "johan@company.com",
			Password:   password,
			UserRoleID: 3,
			UserTypeID: 1,
			Locale:     "CsCZ",
		}
		db.Create(&demoUserThree)
		demoUserFour := database.User{
			FirstName:  "Cecil",
			SecondName: "Jasan",
			Email:      "cecil@company.com",
			Password:   password,
			UserRoleID: 3,
			UserTypeID: 1,
			Locale:     "CsCZ",
		}
		db.Create(&demoUserFour)
	} else {
		err := db.Migrator().AutoMigrate(&database.User{})
		if err != nil {
			logError("MAIN", "Cannot update user table")
		}
	}

	// WORKSHIFT
	if !db.Migrator().HasTable(&database.Workshift{}) {
		logInfo("MAIN", "Workshift table not exists, creating")
		err := db.Migrator().CreateTable(&database.Workshift{})
		if err != nil {
			logError("MAIN", "Cannot create workshift table")
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
			logError("MAIN", "Cannot update workshift table")
		}
	}

	// PART
	if !db.Migrator().HasTable(&database.Part{}) {
		logInfo("MAIN", "Part table not exists, creating")
		err := db.Migrator().CreateTable(&database.Part{})
		if err != nil {
			logError("MAIN", "Cannot create part table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.Part{})
		if err != nil {
			logError("MAIN", "Cannot update part table")
		}
	}

	// STATE
	if !db.Migrator().HasTable(&database.State{}) {
		logInfo("MAIN", "State table not exists, creating")
		err := db.Migrator().CreateTable(&database.State{})
		if err != nil {
			logError("MAIN", "Cannot create state table")
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
			logError("MAIN", "Cannot update state table")
		}
	}

	// WORKPLACE WORKSHIFT
	if !db.Migrator().HasTable(&database.WorkplaceWorkshift{}) {
		logInfo("MAIN", "WorkplaceWorkshift table not exists, creating")
		err := db.Migrator().CreateTable(&database.WorkplaceWorkshift{})
		if err != nil {
			logError("MAIN", "Cannot create workplaceworkshift table")
		}
		db.Exec("ALTER TABLE workplace_workshifts ADD CONSTRAINT fk_workplace_workshifts_workplace FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")
		db.Exec("ALTER TABLE workplace_workshifts ADD CONSTRAINT fk_workplace_workshifts_workshift FOREIGN KEY (workshift_id) REFERENCES workshifts(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.Workshift{})
		if err != nil {
			logError("MAIN", "Cannot update workplaceworkshift table")
		}
	}

	// WORKPLACE PORT
	if !db.Migrator().HasTable(&database.WorkplacePort{}) {
		logInfo("MAIN", "WorkplacePort table not exists, creating")
		err := db.Migrator().CreateTable(&database.WorkplacePort{})
		if err != nil {
			logError("MAIN", "Cannot create workplaceport table")
		}
		db.Exec("ALTER TABLE workplace_ports ADD CONSTRAINT fk_workplace_ports_device_port FOREIGN KEY (device_port_id) REFERENCES device_ports(id)")
		db.Exec("ALTER TABLE workplace_ports ADD CONSTRAINT fk_workplace_ports_state FOREIGN KEY (state_id) REFERENCES states(id)")
		db.Exec("ALTER TABLE workplace_ports ADD CONSTRAINT fk_workplace_ports_workplace FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.WorkplacePort{})
		if err != nil {
			logError("MAIN", "Cannot update workplaceport table")
		}
	}

	// STATE RECORD
	if !db.Migrator().HasTable(&database.StateRecord{}) {
		logInfo("MAIN", "StateRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.StateRecord{})
		if err != nil {
			logError("MAIN", "Cannot create staterecord table")
		}
		db.Exec("ALTER TABLE state_records ADD CONSTRAINT fk_state_records_state FOREIGN KEY (state_id) REFERENCES states(id)")
		db.Exec("ALTER TABLE state_records ADD CONSTRAINT fk_state_records_workplace FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.StateRecord{})
		if err != nil {
			logError("MAIN", "update create staterecord table")
		}
	}

	// OPERATION
	if !db.Migrator().HasTable(&database.Operation{}) {
		logInfo("MAIN", "Operation table not exists, creating")
		err := db.Migrator().CreateTable(&database.Operation{})
		if err != nil {
			logError("MAIN", "Cannot create operation table")
		}
		db.Exec("ALTER TABLE operations ADD CONSTRAINT fk_operations_order FOREIGN KEY (order_id) REFERENCES orders(id)")
		mainOne := database.Operation{
			Name:    "Hlavní výroba",
			OrderID: 2,
			Barcode: "1",
			Note:    "",
		}
		db.Create(&mainOne)
		mainTwo := database.Operation{
			Name:    "Hlavní výroba",
			OrderID: 3,
			Barcode: "1",
			Note:    "",
		}
		db.Create(&mainTwo)
		mainThree := database.Operation{
			Name:    "Hlavní výroba",
			OrderID: 4,
			Barcode: "1",
			Note:    "",
		}
		db.Create(&mainThree)
		mainFour := database.Operation{
			Name:    "Hlavní výroba",
			OrderID: 5,
			Barcode: "1",
			Note:    "",
		}
		db.Create(&mainFour)
		cleaning := database.Operation{
			Name:    "Čištění",
			OrderID: 2,
			Barcode: "2",
			Note:    "",
		}
		db.Create(&cleaning)
		brushingOne := database.Operation{
			Name:    "Broušení",
			OrderID: 4,
			Barcode: "1",
			Note:    "",
		}
		db.Create(&brushingOne)
		brushingTwo := database.Operation{
			Name:    "Broušení",
			OrderID: 5,
			Barcode: "1",
			Note:    "",
		}
		db.Create(&brushingTwo)
	} else {
		err := db.Migrator().AutoMigrate(&database.Operation{})
		if err != nil {
			logError("MAIN", "Cannot update operation table")
		}
	}

	// ORDER RECORD
	if !db.Migrator().HasTable(&database.OrderRecord{}) {
		logInfo("MAIN", "OrderRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.OrderRecord{})
		if err != nil {
			logError("MAIN", "Cannot create orderrecord table")
		}
		db.Exec("ALTER TABLE order_records ADD CONSTRAINT fk_order_records_operation FOREIGN KEY (operation_id) REFERENCES operations(id)")
		db.Exec("ALTER TABLE order_records ADD CONSTRAINT fk_order_records_order FOREIGN KEY (order_id) REFERENCES orders(id)")
		db.Exec("ALTER TABLE order_records ADD CONSTRAINT fk_order_records_workplace FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")
		db.Exec("ALTER TABLE order_records ADD CONSTRAINT fk_order_records_workplace_mode FOREIGN KEY (workplace_mode_id) REFERENCES workplace_modes(id)")
		db.Exec("ALTER TABLE order_records ADD CONSTRAINT fk_order_records_workshift FOREIGN KEY (workshift_id) REFERENCES workshifts(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.OrderRecord{})
		if err != nil {
			logError("MAIN", "Cannot update orderrecord table")
		}
	}

	// PART RECORD
	if !db.Migrator().HasTable(&database.PartRecord{}) {
		logInfo("MAIN", "PartRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.PartRecord{})
		if err != nil {
			logError("MAIN", "Cannot create partrecord table")
		}
		db.Exec("ALTER TABLE part_records ADD CONSTRAINT fk_part_records_order_record FOREIGN KEY (order_record_id) REFERENCES order_records(id)")
		db.Exec("ALTER TABLE part_records ADD CONSTRAINT fk_part_records_part FOREIGN KEY (part_id) REFERENCES parts(id)")
		db.Exec("ALTER TABLE part_records ADD CONSTRAINT fk_part_records_user FOREIGN KEY (user_id) REFERENCES users(id)")
		db.Exec("ALTER TABLE part_records ADD CONSTRAINT fk_part_records_workplace FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.PartRecord{})
		if err != nil {
			logError("MAIN", "Cannot update partrecord table")
		}
	}

	// PACKAGE RECORD
	if !db.Migrator().HasTable(&database.PackageRecord{}) {
		logInfo("MAIN", "PackageRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.PackageRecord{})
		if err != nil {
			logError("MAIN", "Cannot create packagerecord table")
		}
		db.Exec("ALTER TABLE package_records ADD CONSTRAINT fk_package_records_order_record FOREIGN KEY (order_record_id) REFERENCES order_records(id)")
		db.Exec("ALTER TABLE package_records ADD CONSTRAINT fk_package_records_package FOREIGN KEY (package_id) REFERENCES packages(id)")
		db.Exec("ALTER TABLE package_records ADD CONSTRAINT fk_package_records_user FOREIGN KEY (user_id) REFERENCES users(id)")
		db.Exec("ALTER TABLE package_records ADD CONSTRAINT fk_package_records_workplace FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.PackageRecord{})
		if err != nil {
			logError("MAIN", "Cannot update packagerecord table")
		}
	}

	// FAULT RECORD
	if !db.Migrator().HasTable(&database.FaultRecord{}) {
		logInfo("MAIN", "FaultRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.FaultRecord{})
		if err != nil {
			logError("MAIN", "Cannot create faultrecord table")
		}
		db.Exec("ALTER TABLE fault_records ADD CONSTRAINT fk_fault_records_fault FOREIGN KEY (fault_id) REFERENCES faults(id)")
		db.Exec("ALTER TABLE fault_records ADD CONSTRAINT fk_fault_records_order_record FOREIGN KEY (order_record_id) REFERENCES order_records(id)")
		db.Exec("ALTER TABLE fault_records ADD CONSTRAINT fk_fault_records_user FOREIGN KEY (user_id) REFERENCES users(id)")
		db.Exec("ALTER TABLE fault_records ADD CONSTRAINT fk_fault_records_workplace FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.FaultRecord{})
		if err != nil {
			logError("MAIN", "Cannot update faultrecord table")
		}
	}

	// BREAKDOWN RECORD
	if !db.Migrator().HasTable(&database.BreakdownRecord{}) {
		logInfo("MAIN", "BreakdownRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.BreakdownRecord{})
		if err != nil {
			logError("MAIN", "Cannot create breakdownrecord table")
		}
		db.Exec("ALTER TABLE breakdown_records ADD CONSTRAINT fk_breakdown_records_breakdown FOREIGN KEY (breakdown_id) REFERENCES breakdowns(id)")
		db.Exec("ALTER TABLE breakdown_records ADD CONSTRAINT fk_breakdown_records_user FOREIGN KEY (user_id) REFERENCES users(id)")
		db.Exec("ALTER TABLE breakdown_records ADD CONSTRAINT fk_breakdown_records_workplace FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.BreakdownRecord{})
		if err != nil {
			logError("MAIN", "Cannot update breakdownrecord table")
		}
	}

	// DOWNTIME RECORD
	if !db.Migrator().HasTable(&database.DowntimeRecord{}) {
		logInfo("MAIN", "DownTimeRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.DowntimeRecord{})
		if err != nil {
			logError("MAIN", "Cannot create downtimerecord table")
		}
		db.Exec("ALTER TABLE downtime_records ADD CONSTRAINT fk_downtime_records_downtime FOREIGN KEY (downtime_id) REFERENCES downtimes(id)")
		db.Exec("ALTER TABLE downtime_records ADD CONSTRAINT fk_downtime_records_workplace FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.DowntimeRecord{})
		if err != nil {
			logError("MAIN", "Cannot update downtimerecord table")
		}
	}

	// USER RECORD
	if !db.Migrator().HasTable(&database.UserRecord{}) {
		logInfo("MAIN", "UserRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.UserRecord{})
		if err != nil {
			logError("MAIN", "Cannot create userrecord table")
		}
		db.Exec("ALTER TABLE user_records ADD CONSTRAINT fk_user_records_order_record FOREIGN KEY (order_record_id) REFERENCES order_records(id)")
		db.Exec("ALTER TABLE user_records ADD CONSTRAINT fk_user_records_user FOREIGN KEY (user_id) REFERENCES users(id)")
		db.Exec("ALTER TABLE user_records ADD CONSTRAINT fk_user_records_workplace FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.UserRecord{})
		if err != nil {
			logError("MAIN", "Cannot update userrecord table")
		}
	}

	// ALARM
	if !db.Migrator().HasTable(&database.Alarm{}) {
		logInfo("MAIN", "Alarm table not exists, creating")
		err := db.Migrator().CreateTable(&database.Alarm{})
		if err != nil {
			logError("MAIN", "Cannot create alarm table")
		}

		db.Exec("ALTER TABLE alarms ADD CONSTRAINT fk_alarms_workplace FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")

	} else {
		err := db.Migrator().AutoMigrate(&database.Alarm{})
		if err != nil {
			logError("MAIN", "Cannot update alarm table")
		}
	}

	// ALARM RECORDS
	if !db.Migrator().HasTable(&database.AlarmRecord{}) {
		logInfo("MAIN", "AlarmRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.AlarmRecord{})
		if err != nil {
			logError("MAIN", "Cannot create alarmrecord table")
		}
		db.Exec("ALTER TABLE alarm_records ADD CONSTRAINT fk_alarm_records_alarm FOREIGN KEY (alarm_id) REFERENCES alarms(id)")
		db.Exec("ALTER TABLE alarm_records ADD CONSTRAINT fk_alarm_records_workplace FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")
	} else {
		err := db.Migrator().AutoMigrate(&database.AlarmRecord{})
		if err != nil {
			logError("MAIN", "Cannot update alarmrecord table")
		}
	}

	// SYSTEM RECORD
	if !db.Migrator().HasTable(&database.SystemRecord{}) {
		logInfo("MAIN", "SystemRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.SystemRecord{})
		if err != nil {
			logError("MAIN", "Cannot create systemrecord table")
		}
	} else {
		err := db.Migrator().AutoMigrate(&database.SystemRecord{})
		if err != nil {
			logError("MAIN", "Cannot update systemrecord table")
		}
	}

	// DEVICE WORKPLACE RECORD
	if !db.Migrator().HasTable(&database.DeviceWorkplaceRecord{}) {
		logInfo("MAIN", "DeviceWorkplaceRecord table not exists, creating")
		err := db.Migrator().CreateTable(&database.DeviceWorkplaceRecord{})
		if err != nil {
			logError("MAIN", "Cannot create deviceworkplacerecord table")
		}
		db.Exec("ALTER TABLE device_workplace_records ADD CONSTRAINT fk_device_workplace_records_device FOREIGN KEY (device_id) REFERENCES devices(id)")
		db.Exec("ALTER TABLE device_workplace_records ADD CONSTRAINT fk_device_workplace_records_workplaces FOREIGN KEY (workplace_id) REFERENCES workplaces(id)")

	} else {
		err := db.Migrator().AutoMigrate(&database.DeviceWorkplaceRecord{})
		if err != nil {
			logError("MAIN", "Cannot update deviceworkplacerecord table")
		}
	}

	// LOCALE
	if !db.Migrator().HasTable(&database.Locale{}) {
		logInfo("MAIN", "Locale table not exists, creating")
		err := db.Migrator().CreateTable(&database.Locale{})
		if err != nil {
			logError("MAIN", "Cannot create locale table")
		}
		createLocales(db)

	} else {
		err := db.Migrator().AutoMigrate(&database.Locale{})
		if err != nil {
			logError("MAIN", "Cannot update locale table")
		}
		createLocales(db)
	}
	logInfo("MAIN", "Tables checked in "+time.Since(timer).String())

	if !db.Migrator().HasTable(&database.PageCount{}) {
		logInfo("MAIN", "PageCount table not exists, creating")
		err := db.Migrator().CreateTable(&database.PageCount{})
		if err != nil {
			logError("MAIN", "Cannot create page count table")
		}

	} else {
		err := db.Migrator().AutoMigrate(&database.PageCount{})
		if err != nil {
			logError("MAIN", "Cannot update page count table")
		}
	}

	if !db.Migrator().HasTable(&database.WebUserRecord{}) {
		logInfo("MAIN", "Web user record table not exists, creating")
		err := db.Migrator().CreateTable(&database.WebUserRecord{})
		if err != nil {
			logError("MAIN", "Cannot create web user records table")
		}

	} else {
		err := db.Migrator().AutoMigrate(&database.WebUserRecord{})
		if err != nil {
			logError("MAIN", "Cannot update web user table")
		}
	}
	logInfo("MAIN", "Tables checked in "+time.Since(timer).String())

	return true
}

func checkDatabaseOnly() bool {
	logInfo("MAIN", "Checking database")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	if err != nil {
		logWarning("MAIN", "Database does not exist")
		db, err := gorm.Open(postgres.Open(postgresConfig), &gorm.Config{})
		sqlDB, _ := db.DB()
		defer sqlDB.Close()
		if err != nil {
			logError("MAIN", "Problem opening database: "+err.Error())
			return false
		}
		db = db.Exec("CREATE DATABASE system;")
		if db.Error != nil {
			logError("MAIN", "Cannot create database system: "+err.Error())
		}
		logInfo("MAIN", "Database created")
		return true
	}
	logInfo("MAIN", "Database checked in "+time.Since(timer).String())
	return true
}

func updateProgramVersion() {
	logInfo("MAIN", "Writing program version into settings")
	timer := time.Now()
	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	if err != nil {
		logError("MAIN", "Problem opening database: "+err.Error())
		return
	}
	var existingSettings database.Setting
	db.Where("name=?", serviceName).Find(&existingSettings)
	existingSettings.Name = serviceName
	existingSettings.Value = version
	db.Save(&existingSettings)
	logInfo("MAIN", "Program version written into settings in "+time.Since(timer).String())
}

func hashPasswordFromString(pwd []byte) string {
	logInfo("MAIN", "Hashing password")
	timer := time.Now()
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		logError("MAIN", "Cannot hash password: "+err.Error())
		return ""
	}
	logInfo("MAIN", "Password hashed in  "+time.Since(timer).String())
	return string(hash)
}
