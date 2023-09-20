
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

4. To verify that `ca` is working correctly, check the version:
```bash
ca -v
```
If everything is set up properly, the output should resemble the following:
```bash
ca version 1.0.0
```