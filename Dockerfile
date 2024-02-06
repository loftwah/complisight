# Use the official Go image to create a build artifact.
FROM golang:1.19 AS builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o /soc2

# Use the official Alpine image for a lean production container.
FROM alpine:latest  
RUN apk --no-cache add ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /soc2 /soc2

# Run the web service on container startup.
ENTRYPOINT ["/soc2"]
