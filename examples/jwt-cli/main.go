package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/clagraff/argparse"
)

func callback(p *argparse.Parser, ns *argparse.Namespace, leftovers []string, err error) {
	if err != nil {
		switch err.(type) {
		case argparse.ShowHelpErr, argparse.ShowVersionErr:
			// For either ShowHelpErr or ShowVersionErr, the parser has already
			// displayed the necessary text to the user. So we end the program
			// by returning.
			return
		default:
			fmt.Println(err, "\n")
			p.ShowHelp()
		}

		return // Exit program
	}

	name := ns.Get("name").(string)
	upper := ns.Get("upper").(string) == "true"

	if upper == true {
		name = strings.ToUpper(name)
	}

	fmt.Printf("Hello, %s!\n", name)
	if len(leftovers) > 0 {
		fmt.Println("\nUnused args:", leftovers)
	}
}

func createToken(p *argparse.Parser, ns *argparse.Namespace, args []string, err error) {
	if err != nil {
		switch err.(type) {
		case argparse.ShowHelpErr, argparse.ShowVersionErr:
			return
		default:
			fmt.Println(err, "\n")
			p.ShowHelp()
			return
		}
	}

	var (
		data   []string
		secret string
		ok     bool
	)

	secret = ns.Get("secret").(string)
	if data, ok = ns.Get("data").([]string); ok {
		fmt.Println(data)
	}
	fmt.Println("secret:", secret)
}

func addEncryptParser(parent *argparse.Parser) {
	p := argparse.NewParser("Generate JSON Web Tokens", createToken)
	p.AddHelp()

	s := argparse.NewArg("s secret", "secret", "Secret used for generating the signature").Required()
	d := argparse.NewArg("d data", "data", "key=value pairs to include in the token Payload").Action(argparse.Append)

	p.AddOptions(s, d)

	parent.AddParser("enc", p)
}

func main() {
	p := argparse.NewParser("Generate or validate JSON Web Tokens (JWTs)", callback).Version("1.0.0")
	p.AddHelp().AddVersion() // Enable help and version flags

	addEncryptParser(p)

	// Parse all available program arguments (except for the program path).
	p.Parse(os.Args[1:]...)
}
