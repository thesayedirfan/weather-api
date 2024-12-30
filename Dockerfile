FROM golang
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .