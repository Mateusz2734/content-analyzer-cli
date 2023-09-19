
# content-analyzer

content-analyzer (ca) is a versatile command-line tool designed for analyzing word frequency and letter count within files. It offers a wide range of configurable options through command-line flags.


## Installation

To build and install the project, follow these steps:

1. Clone the repository:
```bash
git clone https://github.com/Mateusz2734/content-analyzer-cli.git
cd content-analyzer-cli
```
2. Build the project using `go build`:
```bash
go build
```
3. Optionally, you can copy the resulting binary to a directory included in your `$PATH`. For example:
```bash
sudo cp ca /bin
```
Feel free to replace `/bin` with any directory included in your `$PATH`. This step ensures that you can run the `ca` command from anywhere in your terminal.

4. To verify that `ca` is working correctly, check the help page:
```bash
ca -h
```
If everything is set up properly, the output should resemble the following:
```bash
NAME:
   ca - Analyze word frequency and letter count of a file

USAGE:
   ca [global options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -n n                     output only n most frequent words, 0 means all words (default: 15)
   --pretty, -p             apply pretty print (default: false)
   --file FILE, -f FILE     read from the given FILE
   --out FILE, -o FILE      output to the given FILE
   --letter-percentage, -l  enhance output with letter percentage statistics (default: false)
   --help, -h               show help
```