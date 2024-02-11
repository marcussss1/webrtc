### syntax=docker/dockerfile:1
##FROM golang:1.21-alpine AS base
##WORKDIR /src
##COPY . .
##RUN --mount=type=cache,target=/go/pkg/mod/ \
##--mount=type=bind,source=go.sum,target=go.sum \
##--mount=type=bind,source=go.mod,target=go.mod \
##go mod download -x
##
##FROM base AS build-server
##RUN --mount=type=cache,target=/go/pkg/mod/ \
##--mount=type=bind,target=. \
##go build -o /bin/server .
##
##FROM scratch AS server
##COPY --from=build-server /bin/server /bin/
##ENTRYPOINT [ "/bin/server" ]
#
# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS base
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download -x

FROM base AS build-server
COPY server .
RUN go build -o /bin/server .

FROM scratch AS server
COPY --from=build-server /bin/server /bin/
ENTRYPOINT [ "/bin/server" ]


#FROM golang:1.21.4-alpine
#
#WORKDIR /build
#
#COPY go.mod .
#
#RUN go mod download
#
#COPY . .
#
#RUN go build -o /api server/main.go
#
#EXPOSE 8000
#
#ENTRYPOINT ["/api"]
