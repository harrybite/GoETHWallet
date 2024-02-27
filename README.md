# GoETHWallet

This is a CLI App example using the [Cobra](https://github.com/spf13/cobra) library. This CLI app generates Ethereum wallets randomly and extracts the wallet address from an Ethereum private key.

## Available Commands

- `completion`: Generate the autocompletion script for the specified shell
- `eth`: Extract Ethereum wallet address from a private key
- `help`: Help about any command
- `version`: Print the version of the CLI

## Flags

- `-h, --help`: Help for the CLI

## Main Commands (Under `eth` Command)

### Available Commands

- `ethwallet`: Get Ethereum wallet from a private key using the flag `--key [privatekey]`
- `generate`: Generate an Ethereum base wallet address

### Flags

- `-h, --help`: Help for `eth` command
