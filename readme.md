# GO IP converter

CLI Tool to convert ip related values.

# Installation

Currently this application has not been added to any package manager.
If you have go installed on your computer, you can easily build it yourself.

```sh
go build .
sudo mv ipconv /usr/local/bin
```

# Examples

## IP address

Returns information about a specific IP address.

```sh
ipconv -i 192.168.0.1
```

To filter for particular values use -f:

```sh
ipconv -i 192.168.0.50 -f dec
```

## CIDR

Returns informations about a specific CIDR.

```sh
ipconv -i 192.168.0.1/24
```
