FROM golang:1.20 as stage
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM scratch
WORKDIR /app
COPY --from=stage /app/server .
ENTRYPOINT [ "./server" ]