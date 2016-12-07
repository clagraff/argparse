# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/) 
and this project adheres to [Semantic Versioning](http://semver.org/)

## [Unreleased]
### Added
- Added support for using environmental variable values for the default value
of options.

### Changed
- Changed the CHANGELOG.md format.
- Changed the type of a Namespace from a struct to a `map[string]interface{}`.

## [v1.0.1]
## Fixed 
- Added missing return after executing callback for options.
- Fixed conditions where the callback is not called if an error occurred.
