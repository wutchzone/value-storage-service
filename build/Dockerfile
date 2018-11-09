FROM "golang:1.11-alpine"

# Install dep (package manager)
RUN apk add curl git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Copy Go files
WORKDIR /go/src/app
COPY . .

# Install all dependencies
RUN dep ensure

# Compile project
WORKDIR /go/src/app/cmd/valuestorage
RUN ls
RUN go build -o "app"

# Run project
CMD ["./app"]