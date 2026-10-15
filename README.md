# Dermallorente
Análisis de condiciones atmosféricas relacionadas con problemas de piel

## Problema

El problema se enfoca principalmente en las personas que tienen enfermedades en la piel. Ya que dependiendo de la estación del año (sobre todo en verano e invierno)  y de las condiciones meteorológicas que pueden suceder durante estas estaciones, las enfermedades de la piel no solo pueden generar malestar físico en aquellas personas que la padecen, sino que también pueden empeorar dependiendo de la exposición a los distintos factores climatológicos, estos factores también puede verse afectado por la zona geográfica de España, con sus respectivas condiciones atmosféricas.

## Objetivo

El objetivo es analizar las distintas fuentes de información sobre los cambios meteorológicos y los aspectos que afectan a la gente con problemas o enfermedades de piel.

### ¿Qué datos existen ya y de dónde se extraen?
Para este proyecto no vamos a meter datos a mano ni a depender de APIs. Toda la información ya existe de forma pública, lo que hacemos nosotros es recopilarla de diferentes páginas web.
### Radiación ultravioleta – AEMET 
Permite extraer los índices UV por horas del día anterior de la tabla HTML pública que publica la AEMET.
Fuente: [Tabla UVI AEMET](https://www.aemet.es/es/eltiempo/prediccion/radiacionuv?w=0&datos=det)

## Predicciones semanales y horarias – AEMET:   
Añadimos el análisis del día anterior para poder planificar salidas o recados, usamos las predicciones en formato XML que la AEMET ofrece para cada municipio andaluz (tanto la predicción horaria como la semanal). 
Fuente: Archivos XML por municipio (
[predicción a 7 días](https://www.aemet.es/xml/municipios/localidad_18087.xml)).


## Predicción en Granada

Tenemos esta fuente de datos con datos relevantes com NO2 en Granada (en CSV y también en JSON)

https://www.granada.org/inet/calidadaire.nsf/aire.csv
https://www.granada.org/inet/calidadaire.nsf#estado
