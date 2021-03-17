FROM golang:1.15.6-alpine3.13
WORKDIR /src
COPY . .
RUN go build -o ../bin/calculator ./src
CMD ["./bin/calculator"] 