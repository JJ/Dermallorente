// Package dermallorente contains the risk-evaluation logic for skin exposure.
package dermallorente

import "errors"

// Temperatura represents ambient temperature in degrees Celsius.
type Temperatura float64

// HumedadRelativa represents relative humidity as a percentage.
type HumedadRelativa float64

// VelocidadViento represents wind speed in km/h.
type VelocidadViento float64

// NivelRiesgoCutaneo describes the assessed cutaneous risk level.
type NivelRiesgoCutaneo string

// Risk thresholds used to evaluate cutaneous exposure.
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
	// RiesgoBajo indicates minimal cutaneous risk.
	RiesgoBajo NivelRiesgoCutaneo = "Bajo"
	// RiesgoModerado indicates moderate cutaneous risk.
	RiesgoModerado NivelRiesgoCutaneo = "Moderado"
	// RiesgoAlto indicates high cutaneous risk.
	RiesgoAlto NivelRiesgoCutaneo = "Alto"
	// RiesgoExtremo indicates extreme cutaneous risk.
	RiesgoExtremo NivelRiesgoCutaneo = "Extremo"
)

// EvaluarRiesgoCutaneo assesses the cutaneous risk based on temperature, humidity, and wind.
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
