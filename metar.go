// ----------------------------------------------------------------------------
//
// metar.go
// Mostra el METAR actual de l'aeroport indicat amb el codi OACI
// a partir de l'API de https://tgftp.nws.noaa.gov/data/observations/metar
//
// ----------------------------------------------------------------------------
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"metar/airports"
)

const urlapi string = "https://tgftp.nws.noaa.gov/data/observations/metar/stations/"

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}
	oaci := strings.ToUpper(os.Args[1])
	urlmetar := urlapi + oaci + ".TXT"

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
		metar := strings.Split(metarinfo, "\n")

		// Aquesta API retorna un TXT amb dues línies
		// - data/hora metar  [0]
		// - previsió metar   [1]
		fmt.Printf("Aeroport de %s\n", airports.AirportData[oaci])
		fmt.Printf("%s\n", metar[1])
		fmt.Printf("Ultima actualització: %s\n", metar[0])
	}
}

func help() {
	fmt.Printf("metar 0.1\n")
	fmt.Printf("--> Cal indicar un codi OACI d'aeroport\n\n")
	fmt.Printf("Exemple: metar LELL\n")
}
