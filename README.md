<p align="center">
<img
    src="http://res.cloudinary.com/vidsy/image/upload/v1503160820/CoZ_Icon_DARKBLUE_200x178px_oq0gxm.png"
    width="125px"
  >
</p>

<h1 align="center">neo-go-sc</h1>

<p align="center">
    Smart contract framework for the NEO smart economy written in the Go programming language.
</p>

<p align="center">
  <a href="https://github.com/CityOfZion/neo-go-sc/releases">
    <img src="https://img.shields.io/github/tag/CityOfZion/neo-go-sc.svg?style=flat">
  </a>
  <a href="https://circleci.com/gh/CityOfZion/neo-go-sc/tree/master">
    <img src="https://circleci.com/gh/CityOfZion/neo-go-sc/tree/master.svg?style=shield">
  </a>
</p>

# Overview
- Golang to NVM bytecode compiler
- NEO Virtual machine
- Smart contract debugger
- Private network for quickly deploying and testing smart contracts
- Tooling for deploying smart contracts in production environments
- Package manager for smart contract modules that are written in Go

# Installation macOS and unix
`neo-go-sc` uses [dep](https://github.com/golang/dep) as its dependency manager. After installing `dep` you can run:
```
make install
```

After the installation is completed, you can find the binary in `bin/neo-go-sc` or globally use `neo-go-sc`. 

# Getting started

### Compiling smart contracts
To compile a smart contract run the following:
```
neo-go-sc compile -i path/to/file.go 
```
This will output an `.avm` file in the same directory you executed this command in. 

You can change location directory of the output file by adding the `-o, --out` flag.
```
neo-go-sc compile -i path/to/file.go -o path/to/file.avm 
```

# Contributing
Feel free to contribute to this project after reading the
[contributing guidelines](https://github.com/CityOfZion/neo-go-sc/blob/master/CONTRIBUTING.md).

# Contact
- [@anthdm](https://github.com/anthdm) on Github
- [@anthdm](https://twitter.com/anthdm) on Twitter
- Reach out to me on the [NEO Discord](https://discordapp.com/invite/R8v48YA) channel
- Send me an email anthony@cityofzion.io

# Licence
- Open-source MIT
