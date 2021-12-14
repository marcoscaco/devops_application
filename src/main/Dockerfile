# Iniciamos com a default GoLang image 

FROM golang:1.13-alpine AS build

# Copiamos os arquivos necessários do programa
COPY main.go /src/main/main.go
WORKDIR /src/main

# executamos o Build
RUN CGO_ENABLED=0 go build -o /bin/main

# Iniciamos o container com uma nova imagem linux sem nenhuma dependencia 
# para diminuir o tamnho do build final visto que só precisamos do binário compilado
FROM scratch
COPY --from=build /bin/main /bin/main
ENTRYPOINT ["/bin/main"]