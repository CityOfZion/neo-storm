package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/urfave/cli"
)

var (
	lastExitCode = 0
	fakeOsExiter = func(rc int) {
		lastExitCode = rc
	}
	fakeErrWriter = &bytes.Buffer{}
)

func init() {
	cli.OsExiter = fakeOsExiter
	cli.ErrWriter = fakeErrWriter
}

func TestCommand_InitSmartContract(t *testing.T) {
	app := cli.NewApp()
	app.Name = "neo-storm"
	app.Commands = []cli.Command{cli.Command{
		Name:   "init",
		Usage:  "initialize a new smart-contract in a directory with boiler plate code",
		Action: initSmartContract,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "name, n",
				Usage: "name of the smart-contract to be initialized",
			},
		},
	},
	}

	smartContractName := "testcontract"
	// Setup
	os.RemoveAll(smartContractName)
	// Cleanup
	defer os.RemoveAll(smartContractName)

	err := app.Run([]string{"neo-storm", "init", "--name=testcontract"})
	if err != nil {
		t.Errorf("Did not expect to receive error on first-time contract-creation")
	}

	// Check if the file already exists, after invoking the init command
	if _, err := os.Stat("testcontract/main.go"); os.IsNotExist(err) {
		t.Errorf("Smartcontract directory should have been initialized and must contain main.go file.")
	}

	data, _ := ioutil.ReadFile("testcontract/main.go")
	if fmt.Sprintf(smartContractTmpl, smartContractName) != string(data) {
		t.Errorf("Should have generated boilerplate code in main.go file")
	}

	err = app.Run([]string{"neo-storm", "init", "--name=testcontract"})
	if err == nil {
		t.Errorf("Should have thrown error in case the directory already exists")
	}

}
