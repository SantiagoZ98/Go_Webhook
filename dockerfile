# Usa la imagen oficial de Go como base
FROM golang:1.20-alpine AS builder

# Establece el directorio de trabajo en el contenedor
WORKDIR /app

# Copia el archivo go.mod y go.sum
COPY go.mod ./


# Copia el código fuente de tu aplicación al contenedor
COPY . .

# Compila la aplicación Go
RUN go build -o main .

# Usa una imagen más ligera para la ejecución del contenedor
FROM alpine:latest

# Copia el binario compilado desde la etapa anterior
COPY --from=builder /app/main /main

# Expone el puerto en el que corre el servidor
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["/main"]
