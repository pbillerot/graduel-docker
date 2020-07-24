# http://golang.io/fr/tutoriels/deployer-un-serveur-go-avec-docker/

# Utilisation de golang comme image de base
# Le GOPATH par défaut de cette image est /go.
FROM golang:alpine

# Copie des sources de notre projet dans le container,
# dans notre cas le main.go
COPY ./src /go/src

# Lancement de la compilation avec go install
RUN go install github.com/pbillerot/hello

# Définissions du point d'entré de notre programme compilé
ENTRYPOINT /go/bin/hello

# Le port sur lequel notre serveur écoute
EXPOSE 8008

# docker image build -t hello:1.0 .
# docker container run --name hello -p 8000:8008 -d hello:1.0
# http://localhost:8000/hello
# docker container ps
# docker container stop hello
# docker container start hello