FROM registry.datacommand.co.kr/golang:1.14 as builder

WORKDIR /build
COPY . .

ARG VERSION

RUN make clean build

FROM scratch

COPY --from=builder /build/snapshot /

ENTRYPOINT ["/snapshot"]
