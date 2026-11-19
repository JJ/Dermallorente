package dermallorente

import "time"

type CodigoMunicipio string

type Momento struct {
	Año uint16
	Mes  time.Month
	Dia  uint8
	Hora validezMáxima
}

type ClaveSituacion struct {
	Zona    CodigoMunicipio
	Momento Momento
}

type SituacionClimatica struct {
	Clave       ClaveSituacion
	Temperatura Temperatura
	Humedad     HumedadRelativa
	Viento      VelocidadViento
}
