FROM golang:alpine AS builder

# Set necessary environment variables needed for our image
ENV GO111MODULE=on

# Get build arguments
ARG BUILD_DATE
ARG BUILD_HOST
ARG GIT_URL
ARG BRANCH
ARG VERSION

# Install Tools and dependencies
RUN apk add --update --no-cache openssl-dev musl-dev zlib-dev curl tzdata

# Move to working directory /build
WORKDIR /build

# Copy the code into the container
COPY . .

# Build the application
RUN go build -ldflags "\
    -X main.buildDate=$BUILD_DATE \
    -X main.buildHost=$BUILD_HOST \
    -X main.gitURL=$GIT_URL \
    -X main.branch=$BRANCH \
    -X main.version=$VERSION" \
    -o main ./cmd/main.go;

FROM alpine:latest

COPY --from=builder /build/main .

ENTRYPOINT ["/main"]

# Command to run when starting the container
CMD ["/main"]
