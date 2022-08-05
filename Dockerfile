FROM golang:latest

LABEL maintainer="Kiril Poposki"

WORKDIR /app
COPY go.mod .
RUN go get github.com/gin-gonic/gin
RUN go mod download

COPY . .

ENV PORT 8000
RUN go build

RUN find . -name "*.go" -type f -delete

EXPOSE ${PORT}
CMD ["./smidGIN"]