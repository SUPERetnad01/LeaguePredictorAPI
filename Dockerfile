FROM golang:1.17.3-alpine3.14

# Git is required for fetching the dependencies.
RUN apk add --no-cache git


WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
#RUN go get github.com/SUPERetnad01/LeaguePredictorAPI/proto/predictor
COPY . .
RUN go build -o /server

EXPOSE 8080
EXPOSE 50051

CMD [ "/server" ]
