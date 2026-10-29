package dermallorente

struct validez{
	type hora int
}

const(
	MinimoHora int = 0
	MaximoHora int = 23
)


func NuevaValidez(hora, int) (Validez, error){
	if hora < MinimoHora || hora > MaximoHora
		//DEVUELVE ERROR

	return Validez{hora: hora}, nil
}
