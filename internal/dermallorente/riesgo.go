package dermallorente

import "errors"

type Temperatura float64
type HumedadRelativa float64
type VelocidadViento float64
type NivelRiesgoCutaneo string

const (
	MinimoHumedad       HumedadRelativa = 0
	MaximoHumedad       HumedadRelativa = 100
	RiesgoExtremoTemp   Temperatura     = 10
	RiesgoExtremoHum    HumedadRelativa = 20
	RiesgoExtremoViento VelocidadViento = 20
	RiesgoAltoHum       HumedadRelativa = 30
	RiesgoAltoTemp      Temperatura     = 5
	RiesgoModHum        HumedadRelativa = 45
	RiesgoModTemp       Temperatura     = 12
)

const (
	RiesgoBajo     NivelRiesgoCutaneo = "Bajo"
	RiesgoModerado NivelRiesgoCutaneo = "Moderado"
	RiesgoAlto     NivelRiesgoCutaneo = "Alto"
	RiesgoExtremo  NivelRiesgoCutaneo = "Extremo"
)

func EvaluarRiesgoCutaneo(temp Temperatura, humedad HumedadRelativa, viento VelocidadViento) (NivelRiesgoCutaneo, error) {
	if humedad < MinimoHumedad || humedad > MaximoHumedad {
		return "", errors.New("humedad relativa fuera de rango (0-100%)")
	}

	if temp <= RiesgoExtremoTemp && humedad <= RiesgoExtremoHum && viento >= RiesgoExtremoViento {
		return RiesgoExtremo, nil
	}

	if humedad <= RiesgoAltoHum || temp <= RiesgoAltoTemp {
		return RiesgoAlto, nil
	}

	if humedad <= RiesgoModHum || temp <= RiesgoModTemp {
		return RiesgoModerado, nil
	}

	return RiesgoBajo, nil
}
