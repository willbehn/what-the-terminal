# wtt — What The Terminal 

What The Terminal is a command-line tool that records shell commands in a local SQLite database, so you can review, search, and analyze your terminal history.


## Prerequisites
- Go
- Zsh

## Install locally

```bash
git clone https://github.com/willbehn/what-the-terminal
cd what-the-terminal
mkdir -p bin database
go build -o bin/wtt .
```

## Enable shell logging
Add this to your `~/.zshrc`:

```bash
# what-the-terminal
export WTT_BIN="/absolute/path/to/what-the-terminal/bin/wtt"
export WTT_DB="/absolute/path/to/what-the-terminal/database/history.db"

export PATH="/absolute/path/to/what-the-terminal/bin:$PATH"
source "/absolute/path/to/what-the-terminal/scripts/wtt-preexec.zsh"

```

Reload shell:

```bash
source ~/.zshrc
```

Initialize DB schema:

```bash
wtt init
```

## Supported commands

### `wtt` (root)
Base command that exposes all subcommands below.

Global flags:
- `-l, --long` show long output for commands that print command history (`recent`, `search`)

### `wtt init`
Initializes the SQLite database schema.

### `wtt recent [count]`
Shows the most recent commands from the database.
- Default count is `10`
- Pass a number to change result count, for example `wtt recent 25`
- Use `--long` for detailed output

### `wtt record`
Inserts one command event into the database (used by the shell hook).

### `wtt search <term...>`
Searches recorded commands by one or more terms.

### `wtt ask <task...>`
Asks a local chat model for shell command guidance and prints:
- Summary
- Suggested commands
- Risk level (`safe`, `caution`, `dangerous`)
- Notes

Flags:
- `--model` model name (default: `mistral`)
- `--endpoint` chat API endpoint (default: `http://localhost:11434/api/chat`)

### `wtt stats`
Shows top `10` most-used commands from the last `7` days.

## Security note
- Database at `$WTT_DB` is unencrypted.
- To avoid logging a command, start it with a leading space.
- Add ignore patterns in `scripts/wtt-preexec.zsh`.

## TODO
