package scraper

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gocolly/colly"
)

func GetCompanyDataByRNC(fiscalIdentity string) (string, error) {
	// Create a collector to set initial values
	c := colly.NewCollector(
		colly.AllowedDomains("www.dgii.gov.do"),
	)

	var viewState, eventValidation, companyName string

	// Get __VIEWSTATE value
	c.OnHTML("input[name='__VIEWSTATE']", func(e *colly.HTMLElement) {
		viewState = e.Attr("value")
		fmt.Println("__VIEWSTATE finded:", viewState)
	})

	// Get __EVENTVALIDATION value
	c.OnHTML("input[name='__EVENTVALIDATION']", func(e *colly.HTMLElement) {
		eventValidation = e.Attr("value")
		fmt.Println("__EVENTVALIDATION finded:", eventValidation)
	})

	// Evento después de obtener los valores dinámicos
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Sending form with values...")

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)

		// Add fields to form
		writer.WriteField("__VIEWSTATE", viewState)
		writer.WriteField("__EVENTVALIDATION", eventValidation)
		writer.WriteField("ctl00$cphMain$txtRNCCedula", fiscalIdentity)
		writer.WriteField("ctl00$cphMain$btnBuscarPorRNC", "Buscar")

		err := writer.Close()
		if err != nil {
			return
		}

		// Create a new collector to send form
		postCollector := colly.NewCollector(
			colly.AllowedDomains("www.dgii.gov.do"),
		)

		postCollector.OnHTML("#ctl00_cphMain_dvDatosContribuyentes", func(e *colly.HTMLElement) {
			goquerySelection := e.DOM
			companyName = goquerySelection.Find("tr").Eq(1).Find("td").Eq(1).Text()
		})

		// Set HTTP header on multipart/form-data
		headers := http.Header{}
		headers.Set("Content-Type", writer.FormDataContentType())

		// Send form with POST
		err = postCollector.Request(
			"POST",
			"https://www.dgii.gov.do/app/WebApps/ConsultasWeb/consultas/rnc.aspx",
			body,
			nil,
			headers,
		)
		if err != nil {
			log.Fatal("Error trying to send form", err)
		}
	})

	// Visit page to get initial values
	err := c.Visit("https://www.dgii.gov.do/app/WebApps/ConsultasWeb/consultas/rnc.aspx")
	if err != nil {
		return "", err
	}

	return companyName, err
}
