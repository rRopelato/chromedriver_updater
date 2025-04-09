
# ChromeDriver Updater

ChromeDriver Updater is a tool written in Go that checks, downloads, and automatically updates the latest version of ChromeDriver in the [SeleniumBasic](https://florentbr.github.io/SeleniumBasic/ "Selenium Basic by Florentbr").

## Features

* Checks the installed version of ChromeDriver.
* Fetches the latest available version of ChromeDriver.
* Automatically downloads and replaces the file in the SeleniumBasic folder.
* Simple command-line interface.

## Requirements

* Windows
* Internet connection
* SeleniumBasic installed at `C:\Users\<your_user>\AppData\Local\SeleniumBasic`

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/rRopelato/chromedriver-updater.git
   ```
2. Enter the project folder:
   ```
   cd chromedriver-updater
   ```
3. Compile the program (requires [Go](https://go.dev/ "Go language") to be installed):
   ```
   go build -o chromedriver-updater.exe
   ```

## Usage

1. Run the program:
   ```
   chromedriver-updater.exe
   ```
2. Choose an option from the menu:
   * `1`: Atualizar o ChromeDriver. (Update ChromeDriver.)
   * `2`: Sair. (Exit)

The program will fetch the latest version of ChromeDriver, download it, and automatically replace the old file.

## Execution Example

```
Atualizador ChromeDriver

https://github.com/rRopelato

Versão disponível do ChromeDriver: 123.0.0
Versão instalada do ChromeDriver: 122.0.0

1. Atualizar ChromeDriver
2. Sair
```

## Author

If you have any questions or suggestions, feel free to get in touch:

| Platform     | Contact              |
| ------------ | -------------------- |
| Discord User | ropelato             |
| Discord ID   | 220701036929613825   |
| E-mail       | r.ropelato@proton.me |

Made with ❤️ by [@rRopelato](https://github.com/rRopelato).
