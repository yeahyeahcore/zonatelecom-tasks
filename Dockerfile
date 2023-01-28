FROM golang:alpine as builder

#Update depences
RUN apk update && apk add --no-cache git curl
#Create build directory
RUN mkdir /app/bin -p
#Download health check utility
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.6 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
#Set home directory
WORKDIR /app
#Enable go option
ENV GO111MODULE=on
#Copy go.mod
ADD go.mod go.sum /app/
#Change url for private repos in order to use sb-common
RUN git config --global url."https://SB_service_instance:glpat-JMMAEiRUxGDn3Zg2Ba6y@gitlab.doslab.ru".insteadOf "https://gitlab.doslab.ru"
#Download go depences
RUN go mod download
#Copy all local files
ADD . /app
#Build app
RUN GOOS=linux go build -o bin ./...

#Run container
FROM alpine:latest
#Install packages
RUN apk --no-cache add ca-certificates
#Create home directory
WORKDIR /app
#Copy build file
COPY --from=builder /app/bin/app ./app
#Copy migration dir
COPY --from=builder /bin/grpc_health_probe /bin/grpc_health_probe

CMD ["./app"]


