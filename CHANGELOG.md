# Alarm Service Changelog

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/).

Please note, that this project, while following numbering syntax, it DOES NOT
adhere to [Semantic Versioning](http://semver.org/spec/v2.0.0.html) rules.

## Types of changes

* ```Added``` for new features.
* ```Changed``` for changes in existing functionality.
* ```Deprecated``` for soon-to-be removed features.
* ```Removed``` for now removed features.
* ```Fixed``` for any bug fixes.
* ```Security``` in case of vulnerabilities.

## [2020.4.1.26] - 2020-10-26

### Fixed
- fixed leaking goroutine bug when opening sql connections, the right way is this way

## [2020.3.2.22] - 2020-08-29

### Changed
- functions naming changed to idiomatic go (ThisFunction -> thisFunction)

## [2020.3.2.22] - 2020-08-22

### Added
- automatic go get -u all when creating docker image

## [2020.3.2.13] - 2020-08-13

### Changed
- updated to latest libraries
- updated to go 1.15 -> final program size reduced to 90%

## [2020.3.2.4] - 2020-08-04

### Fixed
- proper creating operation table

### Changed
- updated to all latest libraries 
- functions and method renamed
- logging improved
- behavior improved
- code moved into multiple files

### Added
- MIT license


## [2020.3.1.30] - 2020-07-30

### Fixed
- proper closing database connections with sqlDB, err := db.DB() and defer sqlDB.Close()

### Added
- creating default operation

### Changed
- added tzdata to dockerfile

## [2020.3.1.28] - 2020-07-28

### Added
- creating timezone settings

## [2020.3.1.26] - 2020-07-26

### Changed
- for updated database structure

## [2020.3.1.22] - 2020-07-22

### Changed
- complete change to gorm v2

### Removed
- all about config files
- writing to log files

## [2020.2.2.18] - 2020-05-18

### Added
- init for actual service directory
- db.logmode(false)

## [2020.1.3.31] - 2020-03-31

### Added
- updated create.sh for uploading proper docker version automatically

## [2020.1.2.29] - 2020-02-29

### Removed
- MySQL database

### Fixed
- added sleep when creating tables and database is not available

## [2020.1.1.1] - 2020-01-01

### Added
- creating properly mariadb, mysql, postgresql and mssql database
- checking database size for all four databases
- checking daily database growth rate
- checking disc space
- calculating estimation in days for new data
- sending email, when this estimation get lower than 30 days
