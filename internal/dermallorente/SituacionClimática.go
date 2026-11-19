package dermallorente

import "time"

type CodigoMunicipio string

type Momento struct {
	Año  uint16
	Mes  time.Month
	Dia  uint8
	Hora uint8
}

type ClaveSituacion struct {
	Zona    CodigoMunicipio
	Momento Momento
}

type SituacionClimatica struct {
	Clave       ClaveSituacion
	Temperatura temperatura
	Humedad     Humedad
	Viento      VelocidadViento
	RadiacionUV IndiceUV
}
