FROM golang:1.16-alpine AS Builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker-user-api .
#EXPOSE 8080
#CMD [ "/docker-user-api" ]

FROM alpine:latest 
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=Builder /app/docker-user-api ./
EXPOSE 8080
ENV ELASTICSEARCH_URL=http://es01:9200
CMD ["./docker-user-api"]
