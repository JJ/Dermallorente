package dermallorente

import (
	"errors"
	"fmt"
)

const (
	humedadMinima uint8 = 0
	humedadMaxima uint8 = 100
)

var ErrHumedadFueraDeRango = errors.New("dermallorente: humedad fuera de rango")

type Humedad struct {
	porcentaje uint8
}

func NuevaHumedad(porcentaje int) (Humedad, error) {
	if porcentaje < int(humedadMinima) || porcentaje > int(humedadMaxima) {
		return Humedad{}, fmt.Errorf("%w (%d-%d %%): %d", ErrHumedadFueraDeRango, humedadMinima, humedadMaxima, porcentaje)
	}
	return Humedad{porcentaje: uint8(porcentaje)}, nil
}

func (h Humedad) Porcentaje() uint8 {
	return h.porcentaje
}

func (h Humedad) String() string {
	return fmt.Sprintf("%d %%", h.porcentaje)
}
