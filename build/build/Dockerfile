FROM "golang:alpine"

# Copy Go files
WORKDIR /go/src/app
COPY . .

# Run project
CMD ["./app", "../../configs/config.prod.json" ]
