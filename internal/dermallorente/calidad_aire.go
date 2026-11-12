package dermallorente

type IndiceCalidadAire uint8

const(
	CalidadBuena IndiceCalidadAire = 1
	CalidadRazonablementeBuena IndiceCalidadAire = 2
	CalidadRegular IndiceCalidadAire = 3 
	CalidadDesfavorable IndiceCalidadAire = 4
	CalidadMuyDesfavorable IndiceCalidadAire = 5
	CalidadExtremadamenteDesfavorable IndiceCalidadAire = 6
)

type Contaminante string

const(
	O3   Contaminante = "O3"
	NO2  Contaminante = "NO2"
	SO2  Contaminante = "SO2"
	PM10 Contaminante = "PM10"
	PM25 Contaminante = "PM2.5"
)

type CalidadAire struct{
	indice IndiceCalidadAire
	debidoA []Contaminante
}
