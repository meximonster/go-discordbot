FROM golang:1.19

WORKDIR /usr/src/go-discordbot

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build .

CMD ["/usr/src/go-discordbot/go-discordbot"]