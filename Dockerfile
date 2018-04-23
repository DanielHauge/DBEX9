FROM golang:jessie


RUN go get github.com/lib/pq
RUN go get gopkg.in/cq.v1


COPY Experiment.go Experiment.go
COPY main.go main.go
COPY neo4j.go neo4j.go
COPY sql.go sql.go

RUN go build

ENTRYPOINT ./go