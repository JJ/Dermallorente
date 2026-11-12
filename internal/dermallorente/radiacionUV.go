package dermallorente

type IndiceUV uint8

type Ciudad string

type RadiacionUVCiudad struct {
	Ciudad Ciudad
	Indice IndiceUV
}

type PrediccionUVAEMET struct {
	Elaborado   Momento
	Validez     Momento
	Radiaciones []RadiacionUVCiudad
}
