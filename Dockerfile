FROM registry.access.redhat.com/ubi8/ubi:8.4 AS builder
RUN dnf install -y golang && dnf clean all
WORKDIR /go/src/github.com/vrutkovs/k8s-podhunt
COPY . .
RUN go mod vendor && go build -o ./k8s-podhunt .


FROM registry.access.redhat.com/ubi8/ubi-minimal:8.4
COPY --from=builder /go/src/github.com/vrutkovs/k8s-podhunt/k8s-podhunt /bin/k8s-podhunt
WORKDIR /srv
ENTRYPOINT ["/bin/k8s-podhunt"]
