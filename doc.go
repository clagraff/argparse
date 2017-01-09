// Package argparse is a Golang command line argument parsing library, taking heavy influance from Python's argparse module.
//
// Using argparse, it is possible to easily create command-line interfaces, such
// as:
//
//		> exc --help
//
//		usage: main [-h] [-v] [-e] [-x  ...] [-n] [-f] [-k] [p PATTERN] [s SPLIT] [c CHAR]
//
//		Construct and execute arguments from Stdin
//
//		positional arguments:
//		  [p PATTERN]         Stdin regex grouping pattern
//		  [s SPLIT]           Delimiting regex for Stdin
//		  [c CHAR]            Replacement string for argument parsing
//
//		optional arguments:
//		  -h, --help          Show program help
//		  -v, --version       Show program version
//		  -e, --empty         Allow empty text
//		  -x, --exec          Pasrable command string
//		  -n, --dry-run       Output commands instead of executing
//		  -f, --force         Force continue command execution upon errored commands
//		  -k, --keep-newline  Allow trailing newline from Stdin
//
// Much of the heavy lifting for creating a cmd-line interface is managed by argparse,
// so you can focus on getting your program created and running.
//
// For example, the code required to create the above interface is as follows:
//
//		import (
//			"github.com/clagraff/argparse"
//		)
//
//		func main() {
//			p := argparse.NewParser("Construct and execute arguments from Stdin").Version("0.0.0")
//			p.AddHelp().AddVersion() // Enable `--help` & `-h` to display usage text to the user.
//
//			pattern := argparse.NewArg("p pattern", "pattern", "Stdin regex grouping pattern").Default(".*")
//			split := argparse.NewArg("s split", "split", "Delimiting regex for Stdin").Default("\n")
//			nonEmpty := argparse.NewOption("e empty", "empty", "Allow empty text")
//			keepNewline := argparse.NewFlag("k keep-newline", "keep-newline", "Allow trailing newline from Stdin").Default("false")
//			command := argparse.NewOption("x exec", "exec", "Pasrable command string").Nargs("r").Action(argparse.Store)
//			replacementChar := argparse.NewArg("c char", "char", "Replacement string for argument parsing").Default("%")
//			dryRun := argparse.NewFlag("n dry-run", "dry", "Output commands instead of executing")
//			ignoreErrors := argparse.NewFlag("f force", "force", "Force continue command execution upon errored commands")
//
//			p.AddOptions(pattern, split, nonEmpty, command, replacementChar, dryRun, ignoreErrors, keepNewline)
//
//			ns, _, err := p.Parse(os.Args[1:]...)
//			switch err.(type) {
//			case argparse.ShowHelpErr:
//				return
//			case error:
//				fmt.Println(err, "\n")
//				p.ShowHelp()
//				return
//			}
//
//
// To get started, all you need is a Parser and a few Options!
package argparse
