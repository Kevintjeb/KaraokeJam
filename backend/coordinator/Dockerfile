FROM golang:stretch AS builder

ARG gitlab_token
RUN git config --global url."https://oauth2:$gitlab_token@gitlab.mi.hdm-stuttgart.de/".insteadOf "https://gitlab.mi.hdm-stuttgart.de/"


COPY . /backend/coordinator

WORKDIR /backend/coordinator
RUN go mod download
#Build as small as possible
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
#Stage 2
RUN chmod +x /backend/coordinator/main
FROM scratch

EXPOSE 8080

# Copy our static executable.
COPY --from=builder /backend/coordinator/main .
COPY --from=builder /backend/coordinator/index.html .
COPY --from=builder /backend/coordinator/.env .


CMD ["./main"]
