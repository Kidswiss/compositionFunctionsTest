FROM golang:1.20-alpine as Build

WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 go build main.go

FROM scratch

COPY --from=Build /app/main /main

CMD [ "/main" ]
