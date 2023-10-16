// ----------------------------------------------------------------------------
//
// metar.go
// Mostra el METAR actual de l'aeroport indicat amb el codi OACI
// a partir de l'API de https://tgftp.nws.noaa.gov/data/observations/metar
//
// ----------------------------------------------------------------------------
package metar

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/danielribes/metar/airports"
)

const urlapi string = "https://tgftp.nws.noaa.gov/data/observations/metar/stations/"

func GetMETAR(oaci string) (string, string, string) {
	var metar []string

	oacicode := strings.ToUpper(oaci)
	urlmetar := urlapi + oacicode + ".TXT"

	resp, err := http.Get(urlmetar)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Get %s: %v\n", oaci, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		metarraw, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Reading %s: %v\n", urlmetar, err)
			os.Exit(1)
		}

		metarinfo := string(metarraw)
		metar = strings.Split(metarinfo, "\n")
	}

	return metar[1], metar[0], airports.AirportData[oacicode]
}
