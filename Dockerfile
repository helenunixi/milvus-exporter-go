FROM golang:1.20-alpine

WORKDIR /app
COPY . .
RUN go mod download && go build -o milvus-exporter .

EXPOSE 9100
ENTRYPOINT ["./milvus-exporter"]
