# GoETHWallet
 
This is a CLI App example using the [Cobra](https://github.com/spf13/cobra). This CLI app generate the ethereum wallet randomlly,
as well as extract the wallet from ethereum private key

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  eth         extract ethereum wallet address form private key
  help        Help about any command
  version     print the version of cli

Flags:
  -h, --help   help for cli

main commands are under eth command

Available Commands:
  ethwallet   get ethereum wallet form private key usgin flag --key [privatekey]
  generate    generate ethereum base wallet address

Flags:
  -h, --help   help for eth