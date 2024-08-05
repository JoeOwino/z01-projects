# ascii-art-web

## Description

Ascii-art-web is a web server application that provides a graphical user interface (GUI) for generating ASCII art. Users can enter text and select from different banners to create ASCII art, which will be displayed on the webpage. The supported banners are:
```
- Shadow
- Standard
- Thinkertoy
```
## Usage

### Prerequisites
```
- Go (Golang) installed on your machine
```
### Installation

1. Clone the repository:
    ```bash
    git clone https://learn.zone01kisumu.ke/git/kevwasonga
    cd ascii-art-web
    ```

2. Run the server:
    ```bash
    go run main.go
    ```

3. Open your web browser and navigate to:
    ```
    http://localhost:8080
    ```
### Usage Example
```
1. Enter the text you want to convert into ASCII art.
2. Select the banner (Shadow, Standard, or Thinkertoy).
```
## Implementation Details

### Algorithm

1. **Main Page (GET /)**
    - Sends an HTML response that includes a form with:
    - Text input for the ASCII art text.
    - Radio buttons for banner selection.

2. **ASCII Art Generation (POST /ascii-art)**

   - Receives the text and banner type from the form submission.
   - Processes the input to generate ASCII art using the selected banner.
   - Returns the generated ASCII art as an HTML response to the web page.

### HTTP Status Codes

- **200 OK**: Successful request.
- **404 Not Found**: Templates or banners not found.
- **400 Bad Request**: Incorrect requests eg the text contains no ASCII character(s).
- **500 Internal Server Error**: Unhandled errors.

## Contributing

Something is wrong? Submit an issue or even better, propose a change!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Authors

[kevwasonga](https://learn.zone01kisumu.ke/git/kevwasonga)

[vomolo](https://learn.zone01kisumu.ke/git/vomolo)

[joseowino](https://learn.zone01kisumu.ke/git/joseowino)
