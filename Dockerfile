FROM golang

WORKDIR /go/src/github.com/b4rti/entropy/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /entropy .

######################################################################

FROM alpine

COPY --from=0 /entropy /bin/entropy

ENTRYPOINT ["/bin/entropy"]