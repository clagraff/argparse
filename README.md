# parg
A Golang flag & argument parser for program arguments. The goal of parg (program-arguments) is to emulate the usability and functionallity of Python's [argparse](https://docs.python.org/dev/howto/argparse.html#the-basics) package, where setting up a parser and arguments is easy and trivial.

# Install
Installing is simple, as with most other Golang packages:

```bash
$ go get github.com/clagraff/pargs
```

Boom! All set! Feel free to read on for examples on getting started and using this package.

# The Basics
## Create a parser
Imagine we have a basic program which outputs text to the user depending on which flags/arguments they provide. First, we will need to setup a parser:

```go
package main

import (
    "fmt"
    "os"

    "github.com/parg"
)

func main() {
    p := parg.NewParser("Output text based on user input")
    p.AddHelp() // Enable `--help` & `-h` to display usage text to the user.

    // Parse all available program arguments (except for the program path).
    if err := p.Parse(os.Args[1:]...); err != nil {
        // An error occurred? Print it out, and display the help text!
        fmt.Println(err)
        p.ShowHelp()
    }
}
```

## Adding Flags
Now that we have a parser, lets add two flags and some program functionallity. If
the user provides a `--foo` flag, we will output the string: `foobar!`. In addition,
we will output zero by default, or whatever number the user provides with the `-i I` flag.

```go
package main

import (
    "fmt"
    "os"

    "github.com/parg"
)

func main() {
    p := parg.NewParser("Output text based on user input")
    p.AddHelp() // Enable `--help` & `-h` to display usage text to the user.

    foo := parg.NewFlag("foo", "Enable foobar text output")
    num := parg.NewFlag("i", "Set desired integer number to output").Nargs("1").Action(parg.Store).Default(0)

    p.AddFlag(foo).AddFlag(num)

    // Parse all available program arguments (except for the program path).
    if err := p.Parse(os.Args[1:]...); err != nil {
        // An error occurred? Print it out, and display the help text!
        fmt.Println(err)
        p.ShowHelp()
    } else {
        if p.Values["foo"] == true {
            fmt.Println("foobar!")
        }
        fmt.Println(p.Values["i"])
    }
}
```

# Explanations
## Parser
The parser is a value which stores program-level information, such as a general description. It also contains all possible flags for the program, and is used to parse through a slice of string arguments. It will attempt to parse every flag, returning errors either due to actions taken as specified by an individual flag, because a flag was required but not present, or because a flag was present which was not defined beforehand.

### Usage Text
The parse is able to generate usage-text. This text will include a template for calling the program, the parser's usage text, and a list of all possible flags (if any). For each flag, the flag's identifier and help text is displayed.

The usage text will automatically attempt to "word wrap" to the maximum width of the console/terminal window the program is being executed within.

## Flags
Flags, also known as "switches" or "options", represent actions to take or settings to modify for a parser. Flags, which look like: `-f` and `--foobar`, may also expect arguments to be present immediatly after them. 

### Nargs
Nargs, a shortening of "numer of arguments", represents the number of arguments a flag expects after its presence in a programs complete list of parameters. This could be an actual number, such as `0` or `5`, or it could be any of the following characters: `*+?`. 

The `*` character represents "any and all arguments" following the flag.

The `+` character represents "one or more arguments" following the flag.

The `?` character represents "no arguments or only one argument" following the flag.

### Actions
A flag's `action` defines what should occur when a flag is parsed. All flags must have an action. By default, a flag will store `true` in the parser when present, and `false` when not. The following are the currently available actions:

__parg.StoreTrue__ will store `true` in the parser when the flag is present.
__parg.StoreFalse__ will store `false` in the parser when the flag is present.
__parg.StoreConst__ will store the flag's constant value into the parser when the flag is present.
__parg.Store__ will store the appropriate number of arguments into the parser when the flag & arguments are present.
__parg.AppendConst__ will append the flag's constant to the flag's slice within the parser.
__parg.Append__ will append the appropriate number of arguments into the flag's slice within the parser.
__parg.ShowHelp__ will print the parser's generate help text to `stdout`.
