FROM golang:1.23.1-alpine AS builder

RUN apk add --no-cache git gcc musl-dev make curl sqlite

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.3.5/tailwindcss-linux-x64 \
  && chmod +x tailwindcss-linux-x64 \
  && mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make css
RUN make build

FROM alpine:3.19

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/db ./db
COPY --from=builder /app/css ./css
COPY --from=builder /app/public ./public
COPY --from=builder /app/static ./static
COPY --from=builder /app/templates ./templates

EXPOSE 8080

ENV GO_ENV=production

CMD ["./app"]
