# PMT - Password Management Tool

## What is PMT?

Pmt is a tool where you can It can be used to save and manage passwords, private notes, burglar alarm codes, credit and debit card details, PINs, software keys, …
Pmt use language Go and storage data in sqlite. Data is encrypted before being storing database with password to protect your data.
## Install
- Install PMT via `go get`
```
go get github.com/khainguyen95/pmt
```
## Console Application Usage
Run `$pmt` to view list command support:
```
Pmt is a tool to manage your application with sensitive data.

Usage:
  pmt [command]

Available Commands:
  config      Config User for application.
  create      Create new info of application.
  delete      Delete info of application.
  help        Help about any command
  list        List info of application.
  show        Show info of application.
  version     Print the version number of pmt.

Flags:
  -h, --help   help for pmt

Use "pmt [command] --help" for more information about a command.
```
### Manage User:
- Add new user:
```
$ pmt config create
```
- Switch between different users: 
```
$ pmt config change <username>
```
- View list user:
```
$ pmt config list <username>
```
### Manage application:

- Add new application:
```
$ pmt create --application=<name> --password=<password>
```
- Or you can add many other sensitive information.
```
$ pmt create --application=<name> --password=<password> --other=<field1>:value1|<field2>:value2
Ex: pmt create --application=foo --password=xxx --other=public_key:xxx|private_key:xxx
```
- View list application of current user (require password current user):
```
$ pmt list
```
- Show sensitive data of application (require password current user):
```
$ pmt show <application name>
```
- Delete application (require password current user):
```
$ pmt delete <application name>
```
## Contributing

Contributions are greatly appreciated. The maintainers actively manage the issues list, and try to highlight issues suitable for newcomers. The project follows the typical GitHub pull request model.  Before starting any work, please either comment on an existing issue, or file a new one.
