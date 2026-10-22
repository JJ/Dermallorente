package dermallorente

import "errors"

type Temperatura float64
type HumedadRelativa float64
type VelocidadViento float64

type NivelRiesgoCutaneo string

const (
	RiesgoBajo     NivelRiesgoCutaneo = "Bajo"
	RiesgoModerado NivelRiesgoCutaneo = "Moderado"
	RiesgoAlto     NivelRiesgoCutaneo = "Alto"
	RiesgoExtremo  NivelRiesgoCutaneo = "Extremo"
)

func EvaluarRiesgoCutaneo(temp Temperatura, humedad HumedadRelativa, viento VelocidadViento) (NivelRiesgoCutaneo, error) {
	if humedad < 0 || humedad > 100 {
		return "", errors.New("humedad relativa fuera de rango (0-100%)")
	}

	if temp <= 10.0 && humedad <= 20.0 && viento >= 20.0 {
		return RiesgoExtremo, nil
	}

	if humedad <= 30.0 || temp <= 5.0 {
		return RiesgoAlto, nil
	}

	if humedad <= 45.0 || temp <= 12.0 {
		return RiesgoModerado, nil
	}

	return RiesgoBajo, nil
}
