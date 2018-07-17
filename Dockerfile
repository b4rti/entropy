FROM golang

WORKDIR /go/src/github.com/b4rti/entropy/
COPY . .

RUN go build -a -o /entropy .

######################################################################

FROM alpine

COPY --from=0 /entropy /bin/entropy

ENTRYPOINT ["/bin/entropy"]