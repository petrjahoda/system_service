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

## [2021.2.1.6] - 2021-04-06

### Changed
- updated to latest database structure

## [2021.1.3.30] - 2021-03-30

### Changed
- updating state_records table

## [2021.1.3.25] - 2021-03-25

### Added
- locales for states

## [2021.1.3.23] - 2021-03-23

### Fixed
- typos in locales

## [2021.1.3.22] - 2021-03-22

### Added
- locales for settings workplace-port details

## [2021.1.3.19] - 2021-03-19

### Added
- locales for settings workplace details

## [2021.1.3.17] - 2021-03-17

### Added
- locales for settings device details

## [2021.1.3.16] - 2021-03-16

### Added
- locales for settings user details

## [2021.1.3.15] - 2021-03-15

### Added
- locales for settings workshifts details
- locales for settings products details
- locales for settings states details

## [2021.1.3.11] - 2021-03-11

### Added
- locales for settings operation details

## [2021.1.3.9] - 2021-03-09

### Added
- locales for settings page
- locales for settings alarm details


## [2021.1.3.4] - 2021-03-04

### Changed
- capitalized locales for data page

## [2021.1.3.3] - 2021-03-03

### Added
- translation for charts selection


## [2021.1.3.2] - 2021-03-02

### Added
- translation for alarm table
- translation for breakdown table
- translation for downtime table
- translation for fault table
- translation for package table
- translation for part table
- translation for state table
- translation for user table
- translation for system stats table

## Changed
- interval for processing calculating disc size

## [2021.1.2.25] - 2021-02-25

### Added
- translation for tables
- translation for order table

## [2021.1.2.23] - 2021-02-23

### Added
- translation for data menu

## [2021.1.2.21] - 2021-02-21

### Changed
- foreign keys made from code, not from database structs

## [2021.1.2.13] - 2021-02-13

### Added
- locales for new menu

## [2021.1.2.9] - 2021-02-09

### Added
- locales for live workplace

## [2021.1.1.28] - 2021-01-28

### Added
- locales for production-downtime-poweroff

## [2021.1.1.21] - 2021-01-21

### Added
- locales for live menu

## [2020.4.3.14] - 2020-12-14

### Changed
- updated to latest go
- updated to latest libraries

## [2020.4.2.24] - 2020-11-24

### Added
- menu items for locale table

## [2020.4.2.17] - 2020-11-17

### Changed
- updated all go libraries 
- updated Dockerfile
- updated create.sh
- added locale table

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
