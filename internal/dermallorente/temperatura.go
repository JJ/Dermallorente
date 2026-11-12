package dermallorente

import "fmt"

type temperatura float32

const (
	MinimaTemperatura temperatura = -30.0
	MaximaTemperatura temperatura = 60.0
)

func NuevaTemperatura(grados float32) (temperatura, error) {
	t := temperatura(grados)
	if !esTemperaturaValida(t) {
		return 0, fmt.Errorf("temperatura %.1f fuera de rango permitido [%.1f, %.1f]", grados, MinimaTemperatura, MaximaTemperatura)
	}
	return t, nil
}

func esTemperaturaValida(t temperatura) bool {
	return t >= MinimaTemperatura && t <= MaximaTemperatura
}