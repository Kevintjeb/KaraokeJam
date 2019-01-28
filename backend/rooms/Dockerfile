FROM golang:stretch as builder

ARG gitlab_token
RUN git config --global url."https://oauth2:$gitlab_token@gitlab.mi.hdm-stuttgart.de/".insteadOf "https://gitlab.mi.hdm-stuttgart.de/"

COPY . /backend/rooms
WORKDIR /backend/rooms

RUN go mod download
#Build as small as possible
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

RUN chmod +x /backend/rooms/main
#Stage 2
FROM scratch
COPY --from=builder /backend/rooms/main .
COPY --from=builder /backend/rooms/.env .

EXPOSE 8081

#Start:
CMD ["./main"]
