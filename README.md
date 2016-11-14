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

## Uso de la API
La API que se incluye en el repositorio es una muestra de una api CRUD tipo 
como la que se utiliza para crear, modificar y eliminar usuarios de una aplicación.

**Insert or update**  
`curl -k -H "Content-Type: application/json" -X POST -d \
'{"email":"foo@bar.com", "fullname": "Tito Bbva", "password": "insecure1"}' \ 
https://proxy/users`  

**Get user**  
`curl -k -X GET https://localhost/users/hash_id`  

**Delete user**  
`curl -k -H "Content-Type: application/json" -X DELETE https://proxy/users/hash_id`  

