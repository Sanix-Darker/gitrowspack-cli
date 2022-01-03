# ---- 
# bui# ld executable binary
FROM golang:latest as builder

MAINTAINER s4nixd@gmail.com

WORKDIR $GOPATH/src/github.com/sanix-darker/grp
# We only copy our app
COPY . .


WORKDIR $GOPATH/src/github.com/sanix-darker/grp
# We fetch dependencies
RUN go get -d -v

# We compile
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/grp

# ---
# Let's build our small image
FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/grp  /go/bin/grp

# Run the hello binary.
ENTRYPOINT ["/go/bin/grp"]

