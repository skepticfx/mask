# Mask

A command-line utility in Go that reads from standard input and replaces specified substrings with their corresponding replacements based on a config file.

## Description

The `mask` utility scans text from standard input and looks for substrings defined in a configuration file, replacing them with the specified replacements. It's not limited to whole words but replaces any matching substrings.

The substring matches and replacements are all lower-cased by default.

The default configuration file is located at `~/.mask.config` and contains the substrings to be masked, along with their replacements. Each line has the format `substring:replacement`.

### Installation
```
go install github.com/skepticfx/mask@latest
```

### Example Config File
~/.mask.config
```
hello:bye
ice:fire
```
Running `mask` for the first time creates a sample config file under `~/.mask.config`

## Usage

1. Ensure that you have a configuration file at `~/.mask.config`.
2. Run the `mask` utility and provide the text through standard input.

```bash
echo "Hello there! A song of ice and fire" | mask
bye there! A song of fire and fire
```

## Development
You can make changes by cloning this repository and building the source code.
```
git clone https://github.com/skepticfx/mask.git
cd mask
go build
```

## Tests
To run the tests, execute the following command in the project directory:
```
go test . -v
```
