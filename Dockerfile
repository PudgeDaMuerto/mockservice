FROM --platform=${BUILDPLATFORM} golang:1.26-alpine AS builder

ARG TARGETOS
ARG TARGETARCH

ENV CGO_ENABLED=0
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN apk add --no-cache make
RUN make build_release


FROM scratch

COPY --from=builder /app/mockservice /mockservice

ENTRYPOINT ["/mockservice"]
