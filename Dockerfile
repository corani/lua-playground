# ---- Build Stage ----------------------------------------
FROM golang:1.20.3-bullseye AS builder

WORKDIR /build

COPY . /build

ARG GOPROXY
RUN CGO_ENABLED=0 go build -o lua-playground .

# ---- Final Stage ----------------------------------------
FROM scratch

COPY static /static
COPY views /views
COPY --from=builder /build/lua-playground /

EXPOSE 6969

ENTRYPOINT [ "/lua-playground" ]
