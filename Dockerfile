#Initial Stage

FROM golang:1.16.3-alpine3.13 as build-env

RUN apk add --no-cache git
RUN apk update && apk upgrade libcurl && apk add git openssh-client curl gcc musl-dev

# Set the Current Working Directory inside the container, to enable module features
ENV GO111MODULE on

ENV WKDIR /app
WORKDIR ${WKDIR}
# Copy go mod and sum files
COPY go.mod ${WKDIR}
COPY go.sum ${WKDIR}
COPY *.go ${WKDIR}/

COPY apiServer ${WKDIR}/apiServer
COPY config ${WKDIR}/config
COPY customLogger ${WKDIR}/customLogger
COPY helper ${WKDIR}/helper
COPY testSDKService ${WKDIR}/testSDKService
COPY cert ${WKDIR}/cert

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

RUN CGO_ENABLED=1 go build -o ${WKDIR}/frankieoneSDK



#Final Stage
FROM alpine:3.13

#create and run go program as seperate user
RUN mkdir -p /data
WORKDIR /
COPY --from=build-env /app /
COPY --from=build-env /app/cert /cert
COPY --from=build-env /app/config /config
#ENV SERVICE_ADDR :8888
EXPOSE 8081

CMD ["/frankieoneSDK"]