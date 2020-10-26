package main

import (
	"github.com/kardianos/service"
	"time"
)

const version = "2020.4.1.26"
const serviceName = "System Service"
const serviceDescription = "Creates database and checks system data"
const config = "user=postgres password=Zps05..... dbname=version3 host=database port=5432 sslmode=disable"
const postgresConfig = "user=postgres password=Zps05..... dbname=postgres host=database port=5432 sslmode=disable"
const downloadInSeconds = 86400

var serviceRunning = false
var processRunning = false

type program struct{}

func main() {
	logInfo("MAIN", serviceName+" ["+version+"] starting...")
	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	program := &program{}
	s, err := service.New(program, serviceConfig)
	if err != nil {
		logError("MAIN", "Cannot start: "+err.Error())
	}
	err = s.Run()
	if err != nil {
		logError("MAIN", "Cannot start: "+err.Error())
	}
}

func (p *program) Start(service.Service) error {
	logInfo("MAIN", serviceName+" ["+version+"] started")
	serviceRunning = true
	go p.run()
	return nil
}

func (p *program) Stop(service.Service) error {
	serviceRunning = false
	if processRunning {
		logInfo("MAIN", serviceName+" ["+version+"] stopping...")
		time.Sleep(1 * time.Second)
	}
	logInfo("MAIN", serviceName+" ["+version+"] stopped")
	return nil
}

func (p *program) run() {
	for serviceRunning {
		processRunning = true
		start := time.Now()
		logInfo("MAIN", serviceName+" ["+version+"] running")
		checkDatabase()
		databaseSizeMegaBytes := readDatabaseSize()
		lastSystemRecord := readLastSystemRecord()
		databaseGrowthInMegaBytes := databaseSizeMegaBytes - lastSystemRecord.DatabaseSizeInMegaBytes
		if databaseGrowthInMegaBytes > 0 {
			logInfo("MAIN", "Database is larger than before")
			discSpaceMegaBytes := calculateFreeDiscSpace()
			estimatedDiscSpaceDays := calculateEstimatedDiscSpaceInDays(databaseGrowthInMegaBytes, discSpaceMegaBytes)
			createNewSystemRecord(databaseSizeMegaBytes, databaseGrowthInMegaBytes, discSpaceMegaBytes, estimatedDiscSpaceDays)
			if estimatedDiscSpaceDays < 30 {
				sendEmail("Low Disc Space")
			}
		} else {
			logInfo("MAIN", "No change in database size")
		}
		sleepTime := downloadInSeconds*time.Second - time.Since(start)
		logInfo("MAIN", "Sleeping for "+sleepTime.String())
		time.Sleep(sleepTime)
		processRunning = false
	}
}
