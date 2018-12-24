FROM "golang:alpine"

# Install dep (package manager)
RUN apk update
RUN apk add curl git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy Go files
WORKDIR /go/src/app
COPY . .

# Compile project
WORKDIR /go/src/app/cmd/valuestorage
RUN ls
RUN go build -o "app"

# Run project
CMD ["./app", "../../configs/config.prod.json" ]
