ARG GOVERSION=1.18
ARG BINDPORT=8080
ARG BINDADDR=127.0.0.1

FROM golang:$GOVERSION AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o timestamps

FROM gcr.io/distroless/base-debian11 AS final

WORKDIR /
COPY --from=build /app/timestamps /timestamps

EXPOSE $BINDPORT
ENV GIN_MODE=release
USER nonroot:nonroot

ENTRYPOINT ["/timestamps"]
