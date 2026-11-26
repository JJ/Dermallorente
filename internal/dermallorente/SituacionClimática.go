package dermallorente

type CodigoMunicipio string

type ClaveSituacion struct {
	Zona    CodigoMunicipio
	Momento Momento
}

type SituacionClimatica struct {
    Temperatura int16
    Humedad uint8
    CalidadAire uint8
    IndiceUV uint8
}
