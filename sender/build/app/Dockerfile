FROM golang:1.18 as build

ENV CODE_DIR /go/src/
ENV BIN_FILE /go/src/app
ENV ENV_FILE /go/src/.env

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS

RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/app/main.go

FROM alpine:3.9

LABEL MAINTAINERS="arthurshafikov"

ENV BIN_FILE /go/src/app
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV ENV_FILE /go/src/.env
COPY --from=build ${ENV_FILE} ${ENV_FILE}

CMD ${BIN_FILE} -env ${ENV_FILE}
