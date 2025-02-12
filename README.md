# mdout

**mdout** is a lightweight CLI tool to capture piped command output and save it into a Markdown file. It prints the output to the terminal and appends it as a formatted Markdown snippet (with syntax highlighting) to a Markdown file.

## Install

### Using `go install`

Make sure you have Go installed. Then run:

```bash
go install github.com/pedro-coelho-dr/mdout@latest
```

This installs mdout into your `$GOPATH/bin` (or `$HOME/go/bin`) directory, which should be in your system's `PATH`.

### Building from Source
Clone the repository and build the executable:

```bash
git clone https://github.com/pedro-coelho-dr/mdout.git
cd mdout
go build -o mdout
```

This produces an executable named `mdout` in the current directory.

## Usage

Use `mdout` as a filter to capture command output. For example: 

```bash
ls | mdout
```

Which writes to the Markdown file and still displays on terminal:

```bash
ls | mdout

LICENSE
README.md
cmd
config
go.mod
go.sum
main.go
mdout
```

## Commands

```bash
mdout [command]

Available Commands:
  config      View or update the current mdout configuration
  help        Help about any command
  version     Print the version of mdout
```

```bash
mdout config [flags]

-c, --capture string    Set command capture method (none, zsh)
-l, --language string   Set snippet language for Markdown output (e.g., bash, console, go, etc.)
-o, --output string     Set output file name (e.g., mdout.md or mdout)
```
Example:
```bash
mdout config --capture=zsh --language=bash --output=mylog
```


## Configuration
The configuration is stored in `$HOME/.config/mdout.yaml`. You can update the following settings via the config command:

- **Command Capture:** Method for capturing commands (default is `none` or set to `zsh`, ). 
- **Snippet Language:** The language identifier used for syntax highlighting in the Markdown snippet. 
- **Output File:** The file where captured output is saved (if the provided name lacks a .md extension, it will be appended automatically).


## Disclaimer

**Beta Release:** This is an early beta version of `mdout`. It is still under development and may contain bugs or incomplete features. Use at your own risk and feel free to contribute improvements or report issues.  


## Credits

Developed by Corisco (2025) – Recife, Brasil.  

Contributions are welcome! This project is released under the GNU General Public License.