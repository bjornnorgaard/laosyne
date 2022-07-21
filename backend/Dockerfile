FROM golang AS build
WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .
RUN go build main.go

FROM gcr.io/distroless/base-debian10

COPY --from=build /app/main .
USER nonroot:nonroot

ENTRYPOINT ["./main"]
