# argparse
## Project Status
[ ![Codeship Status for clagraff/argparse](https://codeship.com/projects/68eb7800-af6b-0133-1b97-3e80188314d9/status?branch=master)](https://codeship.com/projects/132507)
[![GoDoc](https://godoc.org/github.com/clagraff/argparse?status.svg)](https://godoc.org/github.com/clagraff/argparse)

## Description
A Golang flag & argument parser for program arguments. The goal of argparse (program-arguments) is to emulate the usability and functionallity of Python's [argparse](https://docs.python.org/dev/howto/argparse.html#the-basics) package, where setting up a parser and arguments is both easy and trivial.

# Install
Installing is simple, as with most other Golang packages:

```bash
$ go get github.com/clagraff/argparse
```

Boom! All set! Feel free to read on for examples on getting started and using this package.

# The Basics
## Create a parser
Imagine we have a basic program which outputs text to the user depending on which flags/arguments they provide. We need create a parse, create and include our various options,
and then parse our programs arguments. It would be something like:

```go
package main

import (
    "fmt"
    "os"

    "github.com/clagraff/argparse"
)

func main() {
    p := argparse.NewParser("Output text based on user input").Version("1.3.0a")
    p.AddHelp().AddVersion() // Enable help and version flags

    dry_run := argparse.NewFlag("n dry-run", "dry", "Enable dry-run mode").Default("false")
    max := argparse.NewArg("m max", "max", "Max number of widgets").Default("0")

    p.AddOptions(dry_run, max)

    // Parse all available program arguments (except for the program path).
    
    if ns, err := p.Parse(os.Args[1:]...); err != nil {
        switch err.(type) {
        case argparse.ShowHelpErr, argparse.ShowVersionErr:
            return
        default:
            fmt.Println(err, "\n")
            p.ShowHelp()
        }
    } else {
        // ... do stuff    
    }
}
```

A parser can provide a version number, automatically add `-h --help` and `-v --version`
options, and is used to contain a collection of Options.

## Options
Options are expected parameters to your parser. `argparse.Option` is the base
struct for parseable options. 

Options can represent a variety of input types, but are always serialized to
a `string`. Options may be required or not, may be positional or not, can have
a variable number of parameters, and can operate in a variety of manners.

That said, there are two common types of Options. As such, help functions are 
included to make it easier to create them: Flags and Positional Arguments.

### Flags
Flags represent non-positional, boolean arguments. These Options are `"false"` 
by default, will utilize the `argparse.StoreTrue` action. They do not any parameters. You can create a new flag using:

```go
// Create a -d --default flag, called "use-default", with a provided description.
use_default := argparse.NewFlag("d default", "use-default", "Enable the default mode")
```

### Arguments
Arguments represent positional, 1-parameter options. They will store the
value associated with the Option (to be retrieved & used later).

```go
// Create a --foobar arg, called "foobar", with a provided description.
foobar := argparse.NewFlag("foobar", "foobar", "A string to foobar with")
```

### Options
A option is a representation of a value, usually provided to a program via its command line parameters. An option is indicated either by
its identifying qualifier (e.g.: `-f`, `--display`, `-V`, etc), or by its 
position (for positional options).

### Methods
Options can be configured in a variety of ways. Therefore, method-chaining is
heavily used to quickly create and setup an option. Consider the following example:

```go
// Create a required Option -foo, with a default value of "bar", that expects and stores 1 argument.
f := argparse.NewOption("-f --foo", "foo", "A foo option").Default("bar").Nargs("1").Required().Action(argparse.Store)
```

Options can have the following attributes:
* Is required or not required
* Is positional or not positional
* Has a default value
* Has a constant value
* Expects a specified number of arguments (or no arguments)
* Is identified by one or more public qualifiers (e.g.: `-f` or `--foo`)
* Can require arguments to match specified choices

#### Nargs
Nargs, a shortening of "numer of arguments", represents the number of arguments a flag expects after its presence in a programs complete list of parameters. This could be an actual number, such as `0` or `5`, or it could be any of the following characters: `*+?`. 

The `*` character represents "any and all arguments" following the flag.

The `+` character represents "one or more arguments" following the flag.

The `?` character represents "no arguments or only one argument" following the flag.

The `r` or `R` characters represent "all remaining arguments" that were not consumed during parsing. These narg choices do not consume the parse arguments they are applicable to.

#### Actions
A flag's `action` defines what should occur when a flag is parsed. All flags must have an action. By default, a flag will store `true` in the parser when present, and `false` when not. The following are the currently available actions:

__argparse.StoreTrue__ will store `true` in the parser when the flag is present.
__argparse.StoreFalse__ will store `false` in the parser when the flag is present.
__argparse.StoreConst__ will store the flag's constant value into the parser when the flag is present.
__argparse.Store__ will store the appropriate number of arguments into the parser when the flag & arguments are present.
__argparse.AppendConst__ will append the flag's constant to the flag's slice within the parser.
__argparse.Append__ will append the appropriate number of arguments into the flag's slice within the parser.
__argparse.ShowHelp__ will print the parser's generate help text to `stdout`.
