# PalGate Experiment

A little experimental tool (and [SDK](https://github.com/zcahana/palgate-sdk)) to tinker around with [PalGate](https://www.pal-es.com/) devices (parking gates, etc).

## Install

```sh
go get github.com/zcahana/palgate-cli
```

Requires any recent version of [Go](https://golang.org/).

## Usage

Currently, the tool only supports retrieving the log history of gate operations.  
Simply run the following command:

```sh
$ palgate-cli
```

## Config

The tool currently supports several configuration options, specified via environment variables
or a YAML config file located at `$HOME/.palgate`.

- `PALGATE_SERVER_ADDRESS`, optional, the address of the PalGate server, defaults to `api1.pal-es.com`.
- `PALGATE_GATE_ID`, required, the ID of the gate device to be controlled.
- `PALGATE_AUTH_TOKEN`, admin authentication token

Example YAML config file:
```yaml
gateID: <gate id>
authToken: <auth token>
```

More info will be added later on how to obtain the gate ID and authentication token.

## Future plans

- Additional commands to list, add, remove and update users and remote controls.
- Colored log history
- Archive log history (see [palgate-log-archiver](https://github.com/zcahana/palgate-log-archiver))

