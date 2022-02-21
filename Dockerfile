# docker image build -t vvd-golang:1.0.0 .
# docker container run -d --name vvd-golang-app -e VVD_PORT=3000 -p 3000:3000 vvd-golang:1.0.0
# docker container rm -f vvd-golang-app
# docker image rm -f vvd-golang:1.0.0 
FROM golang:1.17.7-alpine3.15
RUN apk add build-base
WORKDIR /usr/src/app
COPY go.mod go.sum* ./
RUN go mod download
EXPOSE 3000
COPY . .
RUN go build main.go patient-info.go schedule-info.go create.go cancel.go parse-soap-envelope.go
CMD ["./main"]