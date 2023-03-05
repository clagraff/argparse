# ROADMAP
## Overview
This document defines the entire current roadmap for the project; itemizing
all current & target features. The features are segragated into various
categories of similar functionallity or sane heirachal structures.

Last Updated: 22 FEB 2016

## Description
The map below illistrates the main goals and milestones of the project. Once 
each level of objectives has been meet, that level's parent can be marked-off.

The main goal of this project is to provide an interface similar to Python's
argparse pacakge. As such, code functionallity & attributes found in the
argparse packaged are included.

## The Map
- [ ] Parser
    - [ ] Support ArgumentParser attribute functionallity
        - [x] prog
        - [x] usage
        - [x] description
        - [x] epilog
        - [ ] argument_default
        - [ ] conflict_handler
        - [ ] add_help
    - [x] Auto-determine Program name
    - [x] Output entire program usage
    - [ ] Support parent parsers
    - [ ] Support multiple prefix characters
    - [ ] Determine & display conflicting options
    - [x] Parse multiple short-arguments in single argument flag
    - [x] Parse from sys.Args by default
    - [x] Support sub-parsers / commands
- [ ] Argument
    - [x] Support Argument attribute functionallity
        - [x] name
        - [x] action
        - [x] nargs
        - [x] const
        - [x] default
        - [x] type
        - [x] choices
        - [x] required
        - [x] help
        - [x] metavar
        - [x] dest
    - [x] Support short & long named options
    - [x] Associate short & long named options as single option
    - [x] Support Nargs options
        - [x] Any positive integer
        - [x] "?" - One argument or none
        - [x] "*" - Any arguments or none
        - [x] "+" - One or more arguments
        - [x] "rR" - Remaining arguments
    - [x] Support argument type-asserting
    - [x] Support limiting to available argument Choices
    - [ ] Allow for mutually-exclusive arguments
    - [ ] Provide validity checking for Option based on provided arguments
- [x] Namespace
    - [x] Contain parsed values for arguments
- [ ] Actions
    - [x] store
    - [x] store_const
    - [x] store_true
    - [x] store_false
    - [x] append
    - [x] append_const
    - [ ] count
    - [x] help
    - [x] version
- [ ] Project / General milestones
    - [ ] Documentation / examples
    - [x] Unit tests
    - [X] Strong code coverage
    - [X] Comprehensive & cohesive error messages
