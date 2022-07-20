FROM golang AS build
WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .
RUN go build main.go

FROM gcr.io/distroless/base-debian10
WORKDIR /app
EXPOSE 8080
USER nonroot:nonroot

COPY --from=build app/main .
ENTRYPOINT ["./main"]