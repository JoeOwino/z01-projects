FROM golang:1.22.2-alpine As jose

WORKDIR /dockerContainer

COPY . .

RUN GOOS=linux go build -o ascii

# Second stage 
FROM alpine:latest

WORKDIR /dockerContainer

COPY --from=jose /dockerContainer/ascii .
https://learn.zone01kisumu.ke/git/joseowino/ascii-art-web-dockerize
COPY --from=jose /dockerContainer/banners /dockerContainer/banners

COPY --from=jose /dockerContainer/static /dockerContainer/static

CMD ["/dockerContainer/ascii"]

# Metadata
LABEL docker-version="27.1.1"
LABEL golang version="1.22.2"
LABEL Contributers="<joseowino> <kewasonga> <vomolo>"


#Docker Prmission
# curl -fsSL https://get.docker.com/rootless | sh
 # export PATH=/home/docker/bin:$PATH
 #  export DOCKER_HOST=unix:///run/user/10531/docker.sock

