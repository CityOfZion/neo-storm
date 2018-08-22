<p align="center">
<img
    src="http://res.cloudinary.com/vidsy/image/upload/v1503160820/CoZ_Icon_DARKBLUE_200x178px_oq0gxm.png"
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
### Unix
`neo-storm` uses [dep](https://github.com/golang/dep) as its dependency manager. After installing `dep` you can run:
```
make install
```

After the installation is completed, you can find the binary in `bin/neo-storm` or globally use `neo-storm`.

# Getting started

### Compiling smart contracts
To compile a smart contract run the following:
```
neo-storm compile -i path/to/file.go
```
This will output an `.avm` file in the same directory you executed this command in.

You can change location directory of the output file by adding the `-o, --out` flag.
```
neo-storm compile -i path/to/file.go -o path/to/file.avm
```

# Contributing
Feel free to contribute to this project after reading the
[contributing guidelines](https://github.com/CityOfZion/neo-storm/blob/master/CONTRIBUTING.md).

# Contact
- [@anthdm](https://github.com/anthdm) on Github
- [@anthdm](https://twitter.com/anthdm) on Twitter
- [@jeroenptrs](https://github.com/jeroenptrs) on Github
- [@_jptrs](https://twitter.com/jptrs) on Twitter
- Reach out to us on the [NEO Discord](https://discordapp.com/invite/R8v48YA) channel
- Send us an email: anthony@cityofzion.io

# Licence
- Open-source MIT
