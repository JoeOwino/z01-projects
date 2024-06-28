# ASCII Art Color

## Overview
ASCII Art Generator with Color Support is a Go program designed to convert input strings into graphical representations using ASCII characters, with the added feature of supporting colored text. This project aims to provide a fun and creative way to display text-based graphics with customizable colors.

## Features
```
Parses input strings containing numbers, letters, spaces, special characters, and newline characters ('\n').
* Supports color customization using the --color=<color> flag.
* Allows coloring of a single letter or a set of letters within the input string.
* Utilizes a predefined set of ASCII characters to render the input string.
* Supports multiple banner formats, including "shadow," "standard," and "thinkertoy."
* Implements error handling for unexpected input.
```
## Usage
To run the program, execute the following command in your terminal:
bash
```
go run . [OPTION] [STRING]
```

Replace [OPTION] with the desired color option (--color=<color> <letters_to_be_colored>), and [STRING] with the text you want to convert into ASCII art.

### Example:
```
go run . --color=red H "Hello, World!"
```
### Color Options
You can specify the color using different notations, such as RGB, HSL, or ANSI.

## Banner Formats
Each character in the ASCII art representation has a height of 8 lines and is separated by a newline character      ('\n'). The following banner formats are supported:
```
* shadow
* standard
* thinkertoy
```
## Project Structure
The project consists of the following files:
```
* main.go: Contains the main program logic.
* banners/: Directory containing banner files for 
```
different formats.

## Development
### Prerequisites
Ensure you have Go installed on your system.
Setup
```
Clone the repository to your local machine.
Navigate to the project directory.
```
## Running Tests
To run unit tests, execute the following command:
```
go test ./...
```
## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
Special thanks to the contributors and maintainers of the Go programming language.
