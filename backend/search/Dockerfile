FROM golang:stretch as builder


ARG gitlab_token
RUN git config --global url."https://oauth2:$gitlab_token@gitlab.mi.hdm-stuttgart.de/".insteadOf "https://gitlab.mi.hdm-stuttgart.de/"
RUN apt-get update && update-ca-certificates

COPY . /backend/search
WORKDIR /backend/search

RUN go mod download
#Build as small as possible
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

RUN chmod +x /backend/search/main
#Stage 2
FROM scratch

COPY --from=builder /backend/search/main .
COPY --from=builder /backend/search/.env .
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
EXPOSE 8083

#Start:
CMD ["./main"]

