FROM golang:1.23-alpine

# Create non-root user
RUN adduser -D -u 1000 sandbox

WORKDIR /workspace
RUN chown sandbox:sandbox /workspace

USER sandbox