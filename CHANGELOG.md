# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/) 
and this project adheres to [Semantic Versioning](http://semver.org/)

## [v1.0.2]
### Added
- Added support for using environmental variable values for the default value
of options.
- Parsers can have epilog text, which is a string that can be displayed at the 
end of parser help-text.

### Changed
- Changed the CHANGELOG.md format.
- Changed the type of a Namespace from a struct to a `map[string]interface{}`.
- Updated main README.md to include proper links for getting this package by 
stable or develoment versions.
- Fixing vetting issues in the codebase, including simplifying the code, adding
comments, and formatting issues.

### Fixed
- Fixed case-sensitive bug with `Option.DisplayName()`.

## [v1.0.1]
## Fixed 
- Added missing return after executing callback for options.
- Fixed conditions where the callback is not called if an error occurred.
