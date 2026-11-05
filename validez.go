package dermallorente

import "fmt"

type validez struct{
	hora int
}

const(
	MinimoHora int = 0
	MaximoHora int = 23
)


func NewValidez(hora, int) (Validez, error){
	
	if hora < MinimoHora || hora > MaximoHora {
		return Validez{}, fmt.Errorf("hora %d fuera de rango [%d, %d]", hora, MinimoHora, MaximoHora)
	}
	return Validez{hora: hora}, nil
}
