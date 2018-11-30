# To run this dockerfile:
# Make sure you have docker installed lol
# 1. docker build --rm -f Dockerfile -t "quizkiller:v1" .
# Once its built
# 2. docker run --rm -d -p 8080:8080 "quizkiller:v1"
# you can rm -d to see outputs to terminal and can change the -p args for different ports
# Excessively important for the ubuntu runs to run, if you don't install the certificates then
# http.Get() won't run. (x509: certificate signed by unknown authority) error.

FROM golang:latest
WORKDIR /go/src/QuizKiller
ADD . /go/src/QuizKiller
RUN go build

# Using ubuntu cuz alpine throws an error
FROM ubuntu:latest
RUN apt-get update
RUN apt-get install -y ca-certificates
COPY --from=0 /go/src/QuizKiller .
EXPOSE 8080
CMD ["./QuizKiller"]