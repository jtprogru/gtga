# gtga – Go TeleGram Alert

> **Superseded by [notiflow](https://github.com/jtprogru/notiflow).**
>
> notiflow does everything gtga did and rather more: templates, retries, edit-in-place, a
> GitHub Action, and Homebrew/crates.io distribution. This repository is in maintenance and
> will be archived; existing installs keep working, nothing is being taken away.

[![GitHub](https://img.shields.io/github/license/jtprogru/gtga)](https://github.com/jtprogru/gtga/blob/main/LICENSE)

Simple CLI tool to send Telegram messages from the command line.

## Moving to notiflow

Installation options are on the [notiflow installation page](https://jtprogru.github.io/notiflow/install/) — Homebrew, `cargo install`, a release archive, or a one-line installer.

| gtga | notiflow |
|---|---|
| `gtga --msg "text"` | `notiflow send --message "text"` |
| `GTGA_BOT_TOKEN` | `NOTIFLOW_BOT_TOKEN` |
| `GTGA_CHANNEL_ID` | `NOTIFLOW_CHAT_ID` |

```bash
# before
GTGA_BOT_TOKEN=... GTGA_CHANNEL_ID=... gtga --msg "deploy finished"

# after
NOTIFLOW_BOT_TOKEN=... NOTIFLOW_CHAT_ID=... notiflow send --message "deploy finished"
```

There is no `GTGA_*` fallback in notiflow: renaming two environment variables is a smaller
cost than carrying an alternative credential source through the configuration layer forever.

What you get on top:

```bash
./deploy.sh | notiflow send --stdin                    # reads a pipeline
notiflow send --template '✅ {{.Repo}} @ {{.ShortSha}}' # git-aware templates
notiflow whoami                                        # check the token
notiflow send --profile work -m "..."                  # named profiles in a config file
```

See the [documentation](https://jtprogru.github.io/notiflow/) and the
[CLI guide](https://jtprogru.github.io/notiflow/cli/usage/).

## Original usage

```bash
gtga --msg "Hello World"
```

Environment variables:

- `GTGA_BOT_TOKEN` – bot token from [BotFather](https://t.me/BotFather)
- `GTGA_CHANNEL_ID` – channel ID for notifications

## License

[MIT](https://github.com/jtprogru/gtga/blob/main/LICENSE)
