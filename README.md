# [genDevOps] Docker workshop

El propósito de esta práctica final es deplegar un stack completo bastante 
cercano a un entorno real, poniendo en práctica todos los conceptos vistos 
hasta ahora.

El stack se compone de un proxy, una api, un mongo y un clave valor. Para el deploy 
usaremos una de las herramientas de orquestación más de moda actualmente: *Rancher*

Este repositorio contiene diversas pistas, archivos a medio configurar y todos 
los archivos accesorios necesarios.

## Primera parte
Desplegar la base de datos, el proxy y la apiv1.

## Bonus
Desplegar la apiv2, una instancia de etcd y browser de etcd.

## Tips
Antes de poder hacer `docker push` es necesario construir la imagen como se 
indica a continuación.

`docker build -t registry-url/org/repo:tag`   
`docker push registry-url/org/repo:tag`
