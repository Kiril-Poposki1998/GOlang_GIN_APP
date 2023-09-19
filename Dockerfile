FROM golang:latest AS build

LABEL maintainer="Kiril Poposki"
ARG CGO_ENABLED=0
ARG GOOS=linux
WORKDIR /app
COPY go.mod .
RUN go get github.com/gin-gonic/gin
RUN go mod download

COPY . .


RUN go build -a -ldflags '-extldflags "-static"' -o gin
RUN chmod +x gin

FROM scratch

COPY --from=build /app/gin /usr/local/bin/gin
COPY --from=build /app/templates /app/templates
COPY --from=build /app/service /app/service
COPY --from=build /app/middleware /app/middleware
COPY --from=build /app/entity /app/entity
COPY --from=build /app/controller /app/controller
ENV PORT 8080

EXPOSE ${PORT}
CMD ["gin"]