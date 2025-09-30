package scraper

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func GetCompanyDataByRNC(fiscalIdentity string) (string, error) {
	url := "https://www.dgii.gov.do/app/WebApps/ConsultasWeb2/ConsultasWeb/consultas/rnc.aspx"
	var companyName string
	var viewState, eventValidation string

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
			"AppleWebKit/537.36 (KHTML, like Gecko) "+
			"Chrome/120.0.0.0 Safari/537.36"),
		colly.AllowURLRevisit(),
		colly.MaxDepth(1),
	)

	c.OnHTML("input[name=__VIEWSTATE]", func(e *colly.HTMLElement) {
		viewState = e.Attr("value")
	})
	c.OnHTML("input[name=__EVENTVALIDATION]", func(e *colly.HTMLElement) {
		eventValidation = e.Attr("value")
	})

	c.OnScraped(func(r *colly.Response) {
		postCollector := colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
				"AppleWebKit/537.36 (KHTML, like Gecko) " +
				"Chrome/120.0.0.0 Safari/537.36"),
		)

		postCollector.OnHTML("#cphMain_dvDatosContribuyentes", func(e *colly.HTMLElement) {
			rows := e.DOM.Find("tr")
			if rows.Length() > 1 {
				companyName = strings.TrimSpace(rows.Eq(1).Find("td").Eq(1).Text())
			}
		})

		err := postCollector.Post(url, map[string]string{
			"__VIEWSTATE":                   viewState,
			"__EVENTVALIDATION":             eventValidation,
			"ctl00$cphMain$txtRNCCedula":    fiscalIdentity,
			"ctl00$cphMain$btnBuscarPorRNC": "Buscar",
		})
		if err != nil {
			log.Println("Error en POST:", err)
		}
	})

	err := c.Visit(url)
	if err != nil {
		return "", err
	}

	return companyName, nil
}
