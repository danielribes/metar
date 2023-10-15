# METAR

METAR és un projecte personal per practicar Go i alhora una futura app web per mostrar informació METAR d'un aeroport de manera visual i descodificada en diferents idiomes

## Enfoc

Comença com un app de consola que evolucionara a un llibreria que donara suport a una aplicació web amb Go.

## Utilització

Ara simplement s'executa en consola. Clonar el repo i fer un:

```
go build metar.go
./metar LELL
```

## TODO

* Filtrar que els codis OACI rebuts siguin d'aeroports reals
* Parser de la informació METAR i traducció a un primer idioma
* Suport multi idioma

## API METAR

Aquesta aplicació utilitza l'API de [https://tgftp.nws.noaa.gov/data/observations/metar/stations/](https://tgftp.nws.noaa.gov/data/observations/metar/stations/) per obtenir la informació METAR d'un aeroport concret.

## Objectiu del parser

Es tracta de descodificar la informació d'un METAR i traduir-la a un format huma, per exemple un METAR com el següent:

```
LELL 282230Z AUTO VRB02KT CAVOK 15/09 Q1017
```

Ha de generar la traducció:

```
Previsió per Aeroport de Sabadell. Dia 28 a les 22:30 UTC. Observació generada automàticament
Vent variable de 2kt. Visibilitat de més de 10km i cap boira per sota de 5000ft (1500m)
Temperatura de 15º punt de Rocio a 9º
QNH 1017

```

## Informació de referencia

* [Decodificar un METAR](ME030 5.1. informes meteo.pdf)
* [Aiport codes](https://github.com/datasets/airport-codes)
* [Open Data Airports](https://ourairports.com/data/)
* [Así es como se lee un METAR](https://metar-taf.com/es/explanation)