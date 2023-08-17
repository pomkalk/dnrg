FROM golang as build
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/main.go

FROM alpine as prod
COPY --from=build /src/app /app/srv
COPY --from=build /src /app/data
EXPOSE 4000
ENTRYPOINT /app/srv