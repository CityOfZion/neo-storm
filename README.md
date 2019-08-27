<p align="center">
<img
    src="https://raw.githubusercontent.com/CityOfZion/visual-identity/develop/_CoZ%20Branding/_Logo/_Logo%20icon/_PNG%20200x178px/CoZ_Icon_DARKBLUE_200x178px.png"
    width="125px"
  >
</p>

<h1 align="center">neo-storm</h1>

<p align="center">
    Smart contract framework for the NEO smart economy written in the Go programming language.
</p>

<p align="center">
  <a href="https://github.com/CityOfZion/neo-storm/releases">
    <img src="https://img.shields.io/github/tag/CityOfZion/neo-storm.svg?style=flat">
  </a>
  <a href="https://circleci.com/gh/CityOfZion/neo-storm/tree/master">
    <img src="https://circleci.com/gh/CityOfZion/neo-storm/tree/master.svg?style=shield">
  </a>
</p>

# Overview
- Golang to NVM bytecode compiler
- NEO Virtual machine
- Smart contract debugger
- Private network for quickly deploying and testing smart contracts
- Tooling for deploying smart contracts in production environments
- Package manager for smart contract modules that are written in Go

# Installation
The following section will help you with installing neo-storm and it's dependencies. 

[**A very in-depth tutorial about how to get started with neo-storm can be found here**](https://medium.com/@likkee.chong/neo-token-contract-nep-5-in-go-f6b0102c59ee)

## Project dependencies
### Golang
neo-storm requires a working and proper ***Golang*** installation. To install Golang you can check out these [installation instructions](https://golang.org/doc/install).

### Godep
For package management neo-storm uses ***dep***. To install dep you can check out these [installations instructions](https://github.com/golang/dep).

# Installing the neo-storm framework
### Unix
`neo-storm` uses [dep](https://github.com/golang/dep) as its dependency manager. After installing `dep` you can run:
```
make install
```

After the installation is completed, you can find the binary in `bin/neo-storm` or globally use `neo-storm`.

# Getting started
Lot's of **examples contracts** can be found in the [examples folder](https://github.com/CityOfZion/neo-storm/tree/master/examples).

### Create a new smart contract
To create a new smart contract you can run the `init` command:
```
neo-storm init --name mycontract
```

This will generate a folder called `mycontract` with a `main.go` file in the root directory.

The folder structure will look like this:
```
- mycontract
    - main.go
```

And will produce the following `main.go` file in the root of the directory:
```
package mycontract

import "github.com/CityOfZion/neo-storm/interop/runtime"

func Main(op string, args []interface{}) {
    runtime.Notify("Hello world!")
}
```

### Compiling smart contracts
To compile a smart contract you can run the `compile` command:
```
neo-storm compile -i path/to/file.go
```
This will output an `.avm` file in the same directory you executed this command in.

You can change location directory of the output file by adding the `-o, --out` flag.
```
neo-storm compile -i path/to/file.go -o path/to/file.avm
```

# Tutorials
- [Step-by-step guide on issuing your NEP-5 token on NEO’s Private net using Go](https://medium.com/@likkee.chong/neo-token-contract-nep-5-in-go-f6b0102c59ee)

# Contributing
Feel free to contribute to this project after reading the
[contributing guidelines](https://github.com/CityOfZion/neo-storm/blob/master/CONTRIBUTING.md).

# Contact
- [@anthdm](https://github.com/anthdm) on Github
- [@anthdm](https://twitter.com/anthdm) on Twitter
- [@jeroenptrs](https://github.com/jeroenptrs) on Github
- [@_jptrs](https://twitter.com/_jptrs) on Twitter
- Reach out to us on the [NEO Discord](https://discordapp.com/invite/R8v48YA) channel
- Send us an email: anthony@cityofzion.io

# Licence
- Open-source MIT
