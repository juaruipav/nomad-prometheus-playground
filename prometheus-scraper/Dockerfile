FROM golang:1.14.4

RUN go get -u   github.com/prometheus/client_golang/prometheus/promhttp \
                github.com/prometheus/client_golang/prometheus \
	            github.com/prometheus/client_golang/prometheus/promauto \
                github.com/sparrc/go-ping 

COPY main.go .

RUN go build main.go

ENTRYPOINT ./main

