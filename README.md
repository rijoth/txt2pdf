# Txt2PDF

**Txt2PDF** is a simple utility written in golang to convert text files to pdf.
 Currently binary file is available for Linux only.
 You can build for other platform using **go build** command.

## Installation

Currently build is only available for Linux systems, download the compiled package and run by providing required parameters.
You can get the binary from [releases](https://github.com/rijoth/txt2pdf/releases/tag/v0.1) page.

**Other Systems:**
clone the repository and run go build command
```go
go build txt2pdf.go
```
# Flags
**-i**   -  This is a required flag which is used to provide input filename.

**-o** - This is an optional flag which is used to provide output filename.

## Usage

```bash
# Example
sh txt2pdf -i <filename> -o <filename> 

# Demo
sh txt2pdf -i testfile.txt -o testfile.pdf
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
