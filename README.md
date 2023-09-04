# gtga – Go TeleGram Alert

[![goreleaser](https://github.com/jtprogru/gtga/actions/workflows/goreleaser.yaml/badge.svg)](https://github.com/jtprogru/gtga/actions/workflows/goreleaser.yaml)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/jtprogru/gtga)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jtprogru/gtga)
[![GitHub issues](https://img.shields.io/github/issues/jtprogru/gtga)](https://github.com/jtprogru/gtga/issues)
[![GitHub](https://img.shields.io/github/license/jtprogru/gtga)](https://github.com/jtprogru/gtga/blob/main/LICENSE)

Simple CLI tool to send Telegram messages from CLI.

## Usage

```bash
gtga --msg "Hello World"
```

## Environment Variables

- `GTGA_BOT_TOKEN` – Bot token from [BotFather](https://t.me/BotFather)
- `GTGA_CHANNEL_ID` – Channel ID for notification

## Installation

For installation you need to load latest version from [Release](https://github.com/jtprogru/gtga/releases) page and download version for you platform.

Another way is usage `go install`:

```shell
# Get latest version from CLI
VERSION=`curl -sSL https://api.github.com/repos/jtprogru/gtga/releases/latest -s | jq .name -r`
go install github.com/jtprogru/gtga@$VERSION
```

## License

[MIT](https://github.com/jtprogru/gtga/blob/main/LICENSE)

