FROM golang:stretch as builder

ARG gitlab_token
RUN git config --global url."https://oauth2:$gitlab_token@gitlab.mi.hdm-stuttgart.de/".insteadOf "https://gitlab.mi.hdm-stuttgart.de/"

COPY . /backend/queue
WORKDIR /backend/queue

RUN go mod download

#Build as small as possible
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

RUN chmod +x /backend/queue/main
#Stage 2
FROM scratch
COPY --from=builder /backend/queue/main .
COPY --from=builder /backend/queue/.env .


EXPOSE 8082
#Start:


CMD ["./main"]
