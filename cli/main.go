package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/CityOfZion/neo-go/pkg/rpc"
	"github.com/CityOfZion/neo-storm/compiler"
	"github.com/urfave/cli"
)

var (
	errNoInput             = errors.New("No input file was found, specify an input file with the '--in or -i' flag")
	errNoSmartContractName = errors.New("No name was provided, specify the '--name or -n' flag")
	errFileExist           = errors.New("A file with given smart-contract name already exists")
)

var (
	// smartContractTmpl is written to a file when used with `init` command.
	// %s is parsed to be the smartContractName
	smartContractTmpl = `package %s

import "github.com/CityOfZion/neo-storm/interop/runtime"

func Main(op string, args []interface{}) {
    runtime.Notify("Hello world!")
}`
)

func main() {
	ctl := cli.NewApp()
	ctl.Name = "neo-storm"
	ctl.Usage = "Neo smart contract framework for the Go programming language"
	ctl.Commands = []cli.Command{
		{
			Name:   "compile",
			Usage:  "compile a smart contract to an .avm file",
			Action: contractCompile,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "in, i",
					Usage: "input file to be compiled",
				},
				cli.StringFlag{
					Name:  "out, o",
					Usage: "output destination of the compiled contract",
				},
				cli.BoolFlag{
					Name:  "debug, d",
					Usage: "compile the contract in debug mode for additional compile information",
				},
			},
		},
		{
			Name:   "testinvoke",
			Usage:  "testinvoke a smart contract against a remote NEO RPC node",
			Action: contractTestInvoke,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "in, i",
					Usage: "input file to be compiled",
				},
			},
		},
		{
			Name:   "init",
			Usage:  "initialize a new smart-contract in a directory with boiler plate code",
			Action: initSmartContract,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name, n",
					Usage: "name of the smart-contract to be initialized",
				},
				cli.BoolFlag{
					Name:  "skip-details, skip",
					Usage: "skip filling in the projects and contract details",
				},
			},
		},
	}
	ctl.Run(os.Args)
}

// initSmartContract initializes a given directory with some boiler plate code.
func initSmartContract(ctx *cli.Context) error {
	contractName := ctx.String("name")
	if contractName == "" {
		return cli.NewExitError(errNoSmartContractName, 1)
	}

	// Check if the file already exists, if yes, exit
	if _, err := os.Stat(contractName); err == nil {
		return cli.NewExitError(errFileExist, 1)
	}

	basePath := contractName
	fileName := "main.go"

	// create base directory
	if err := os.Mkdir(basePath, os.ModePerm); err != nil {
		return cli.NewExitError(err, 1)
	}

	// Ask contract information and write a storm.yml file unless the -skip-details flag is set.
	// TODO: Fix the missing storm.yml file with the `init` command when the package manager is in place.
	if !ctx.Bool("skip-details") {
		details := parseContractDetails()
		if err := ioutil.WriteFile(filepath.Join(basePath, "storm.yml"), details.toStormFile(), 0644); err != nil {
			return cli.NewExitError(err, 1)
		}
	}

	data := []byte(fmt.Sprintf(smartContractTmpl, contractName))
	if err := ioutil.WriteFile(filepath.Join(basePath, fileName), data, 0644); err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Printf("Successfully initialized smart contract [%s]\n", contractName)

	return nil
}

func contractCompile(ctx *cli.Context) error {
	src := ctx.String("in")
	if len(src) == 0 {
		return cli.NewExitError(errNoInput, 1)
	}

	o := &compiler.Options{
		Outfile: ctx.String("out"),
		Debug:   ctx.Bool("debug"),
	}

	if err := compiler.CompileAndSave(src, o); err != nil {
		return cli.NewExitError(err, 1)
	}

	return nil
}

func contractTestInvoke(ctx *cli.Context) error {
	src := ctx.String("in")
	if len(src) == 0 {
		return cli.NewExitError(errNoInput, 1)
	}

	b, err := ioutil.ReadFile(src)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	// FIXME: Make this configurable (probably in the general storm.yml file/config)
	endpoint := "http://seed3.ngd.network:10332"
	client, err := rpc.NewClient(context.TODO(), endpoint, rpc.ClientOptions{})
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	scriptHex := hex.EncodeToString(b)
	resp, err := client.InvokeScript(scriptHex)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	b, err = json.MarshalIndent(resp.Result, "", "  ")
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Println(string(b))

	return nil
}

type ContractDetails struct {
	Author      string
	Email       string
	Version     string
	ProjectName string
	Description string
}

func (d ContractDetails) toStormFile() []byte {
	buf := new(bytes.Buffer)

	buf.WriteString("# Storm specific configuration. Do not modify this unless you know what you are doing!\n")
	buf.WriteString("storm:\n")
	buf.WriteString("  version: 1.0\n")

	buf.WriteString("\n")

	buf.WriteString("# Project section contains information about your smart contract\n")
	buf.WriteString("project:\n")
	buf.WriteString("  author: " + d.Author)
	buf.WriteString("  email: " + d.Email)
	buf.WriteString("  version: " + d.Version)
	buf.WriteString("  name: " + d.ProjectName)
	buf.WriteString("  description: " + d.Description)

	buf.WriteString("\n")

	buf.WriteString("# Module section contains a list of imported modules\n")
	buf.WriteString("# This will be automatically managed by the neo-storm package manager\n")
	buf.WriteString("modules: \n")
	return buf.Bytes()
}

func parseContractDetails() ContractDetails {
	details := ContractDetails{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Author: ")
	details.Author, _ = reader.ReadString('\n')

	fmt.Print("Email: ")
	details.Email, _ = reader.ReadString('\n')

	fmt.Print("Version: ")
	details.Version, _ = reader.ReadString('\n')

	fmt.Print("Project name: ")
	details.ProjectName, _ = reader.ReadString('\n')

	fmt.Print("Description: ")
	details.Description, _ = reader.ReadString('\n')

	return details
}
