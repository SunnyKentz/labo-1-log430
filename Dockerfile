FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY out/* .


# Run the application
CMD ["./caisse-app"] 