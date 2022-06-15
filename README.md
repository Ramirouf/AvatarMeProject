# AvatarMeProject

Consigna:
Dada una información personal, como una dirección de correo electrónico, una dirección IP o una clave pública, se debe desarrollar un package que permita generar un avatar único.
Imagine, por ejemplo, que está creando una nueva aplicación y desea que todos sus usuarios tengan un avatar único y predeterminado. Para ello, debe desarrollar y publicar el package que escribirá permitirá la generación de dichos avatares. Por ejemplo, GitHub utiliza este enfoque y genera un idéntico para todos los usuarios nuevos que no tienen una imagen de avatar seleccionada.

Code to be executed in main.go

info1 := Information{name: "Ramiro111111"} \
service1 := AvatarGenerator() \
service1.GenerateAndSaveAvatar(info1) \
