# Build stage
FROM golang:1.23.2 AS builder

ENV WORKDIR=/app

WORKDIR ${WORKDIR}

COPY go.mod ${WORKDIR}/go.mod
COPY go.sum ${WORKDIR}/go.sum

COPY . ${WORKDIR}

RUN go mod download
RUN go build -o ${WORKDIR}/main ${WORKDIR}/cmd/server/main.go 

# Runtime stage
FROM golang:1.23.2

ENV WORKDIR=/app

WORKDIR ${WORKDIR} 

COPY --from=builder ${WORKDIR}/ ${WORKDIR}/

RUN chmod +x ./backend-entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["./backend-entrypoint.sh"]
