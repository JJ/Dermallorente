package dermallorente

import "time"

type Origen string

const (
	Observado  Origen = "Observado"
	Prediccion Origen = "Prediccion"
)

type CodigoMunicipio string

type Momento struct {
	Año int
	Mes  time.Month
	Dia  int
	Hora validezMáxima
}

type ClaveSituacion struct {
	Zona    CodigoMunicipio
	Momento Momento
	Origen  Origen
}

type SituacionClimatica struct {
	Clave       ClaveSituacion
	Temperatura Temperatura
	Humedad     HumedadRelativa
	Viento      VelocidadViento
}
