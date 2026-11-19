package dermallorente

import "errors"

type Temperatura float64
type HumedadRelativa float64
type VelocidadViento float64
type NivelRiesgoCutaneo string

const MinimoHumedad int = 0
const MaximoHumedad int = 100
const RiesgoExtremoTemp int = 10
const RiesgoExtremoHum int  = 20
const RiesgoExtremoViento int = 20.0
const RiesgoAltoHum int = 30
const RiesgoAltoTemp int = 5
const RiesgoModHum int = 45
const RiesgoModTemp int = 12


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

	if temp <= RiesgoExtremoTemp && humedad <= RiegoExtremoHum && viento >= RiesgoExtremoViento {
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
