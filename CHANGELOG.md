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
