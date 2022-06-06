# Derocli
[![lint](https://github.com/stratumfarm/derocli/actions/workflows/lint.yml/badge.svg)](https://github.com/stratumfarm/derocli/actions/workflows/lint.yml)
[![goreleaser](https://github.com/stratumfarm/derocli/actions/workflows/release.yml/badge.svg)](https://github.com/stratumfarm/derocli/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/stratumfarm/derocli)](https://goreportcard.com/report/github.com/stratumfarm/derocli)

## About

A cli tool to fetch information from a dero rpc node.

## Usage
```
Usage:
  derocli [flags]
  derocli [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  console     Start an interactive console
  height      Get the current height of the blockchain
  help        Help about any command
  info        Get information about the node
  peers       Get the connected peers from the node
  txpool      Get the transaction pool
  version     Print the version info

Flags:
  -c, --config string   path to the config file
  -h, --help            help for derocli
  -r, --rpc string      address of the node (default "localhost:10102")

Use "derocli [command] --help" for more information about a command.
```