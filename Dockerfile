FROM golang

WORKDIR /go/src/github.com/b4rti/ddvs/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /ddvs .

######################################################################

FROM alpine

COPY --from=0 /ddvs /bin/ddvs

ENTRYPOINT ["/bin/ddvs"]