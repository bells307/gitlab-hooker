FROM golang:alpine

WORKDIR /usr/src/gitlab-hooker

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/gitlab-hooker ./cmd
WORKDIR /opt/gitlab-hooker

CMD ["gitlab-hooker"]