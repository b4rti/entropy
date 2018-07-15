FROM golang

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /ddvs .

######################################################################

FROM alpine

COPY --from=0 /ddvs /bin/ddvs
#COPY /config.json /config.json

ENTRYPOINT ["/bin/ddvs"]