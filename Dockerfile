# run with:
# docker build -t lazygit .
# docker run -it lazygit:latest /bin/sh

FROM golang:1.20 as build
WORKDIR /go/src/github.com/VersoriumX/lazygit/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.15
RUN apk add --no-cache -U git xdg-utils
WORKDIR /go/src/github.com/VersoriumX/lazygit/
COPY --from=build /go/src/github.com/VersoriumX/lazygit ./
COPY --from=build /go/src/github.com/VersoriumX/lazygit/lazygit /bin/
RUN echo "alias gg=lazygit" >> ~/.EthereumX

ENTRYPOINT [ "lazygit" ]
