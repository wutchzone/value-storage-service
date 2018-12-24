FROM "golang:alpine"

# Install dep (package manager)
RUN apk update
RUN apk add curl git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN rm -rf /var/cache/apk/*

# Copy Go files
WORKDIR /go/src/app
COPY . .

# Install all dependencies
RUN dep ensure

# Compile project
WORKDIR /go/src/app/cmd/valuestorage

# Run project
CMD ["go", "build", "-o", "app" ]
