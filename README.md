# RNC Scraper - DGII

A Go-based web scraper designed to query the RNC or ID card number of taxpayers on the Dominican Republic's DGII (Dirección General de Impuestos Internos) website. This tool uses the colly library to send POST requests with dynamic parameters such as __VIEWSTATE and __EVENTVALIDATION, simulating form submissions to retrieve taxpayer data.

### Features:
Scrapes RNC or ID card information from the DGII website.
Handles dynamic form parameters like __VIEWSTATE and __EVENTVALIDATION.
Built using Go and the colly library for efficient scraping.
Prerequisites:
Go 1.18+ installed on your machine.
A working internet connection to access the DGII website.
Installation Steps:
Clone the repository:

Open your terminal and clone the repository to your local machine:

```bash
git clone https://github.com/your-username/dgii-scraper.git
```

Navigate to the project directory:

```bash
cd dgii-scraper
```

### Install dependencies:

If you're using Go modules (Go 1.11+), run the following command to install the required dependencies:

```bash
go mod tidy
```

Modify the main.go file:

Open the main.go file, and update the RNC or ID number you want to query. You can modify this dynamically by passing values to the function or hardcoding them into the script.

### Run the scraper:

Execute the following command to start the scraper:

```bash
go run main.go
```

The scraper will fetch the data for the specified RNC or ID number.

### How It Works:
The scraper sends a POST request to the DGII website's form, passing the required form parameters, including __VIEWSTATE and __EVENTVALIDATION.
The request is processed, and the data for the specified RNC is retrieved.
The scraper extracts the relevant information from the response.

#### Example Usage:
Once you run the program, the scraper will simulate a form submission and scrape the data for the specified RNC or ID. For example:

```bash go
package main
```

```go
import (
"fmt"
"dgiiScraper/scraper"
)

func main() {
    rnc := "132470192" // Replace this with the actual RNC
    scraper.GetCompanyDataByRNC(rnc)
    fmt.Println("Scraping complete!")
}
```

Contributing:
Feel free to fork this repository, submit issues, or create pull requests to improve the project.

License:
This project is licensed under the MIT License - see the LICENSE file for details.