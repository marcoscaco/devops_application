# Iniciamos com a default GoLang image 

FROM golang:1.13-alpine AS build

# Copiamos os arquivos necessários do programa
COPY main.go /src/main/main.go
COPY go.mod /src/main/go.mod

WORKDIR /src/main

RUN go mod download

# executamos o Build
RUN CGO_ENABLED=0 go build -o /bin/main

# Iniciamos o container com uma nova imagem linux sem nenhuma dependencia 
# para diminuir o tamnho do build final visto que só precisamos do binário compilado
FROM scratch
EXPOSE 5001
COPY --from=build /bin/main /bin/main
ENTRYPOINT ["/bin/main"]