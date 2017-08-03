FROM alpine:latest

# Copy the server executable to /opt/server on container
COPY server /opt/server

ENTRYPOINT /opt/server
