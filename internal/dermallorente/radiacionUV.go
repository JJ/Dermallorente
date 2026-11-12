package dermallorente

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
