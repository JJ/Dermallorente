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
Fuente: [Tabla UVI AEMET](https://www.aemet.es/es/eltiempo/observacion/radiacion/ultravioleta?datos=tabla)
### Temperatura y Humedad – IFAPA: 
El frío y el aire seco son de los factores que más empeoran los brotes de dermatitis y psoriasis. Estos datos los tomamos de las estaciones de la Red de Información Agroclimática de Andalucía (IFAPA), que ofrece valores diarios de humedad y temperatura por provincia a través de su web.
Fuente: [Red de estaciones IFAPA](https://www.juntadeandalucia.es/agriculturaypesca/ifapa/riaweb/web/estacion/18/1) (ejemplo de estación en Granada)

## Predicciones semanales y horarias – AEMET:
Añadimos el análisis del día anterior para poder planificar salidas o recados, usamos las predicciones en formato XML que la AEMET ofrece para cada municipio andaluz (tanto la predicción horaria como la semanal). 
Fuente: Archivos XML por municipio (ejemplo Granada: [predicción horaria](https://www.aemet.es/xml/municipios_h/localidad_h_18087.xml) y [predicción a 7 días](https://www.aemet.es/xml/municipios/localidad_18087.xml)).

