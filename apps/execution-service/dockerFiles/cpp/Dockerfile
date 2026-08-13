FROM alpine:3.20

# Install GCC, G++, and basic build tools
RUN apk add --no-cache gcc g++ make musl-dev

# Create non-root user for execution
RUN adduser -D -u 1000 sandbox

WORKDIR /workspace
RUN chown sandbox:sandbox /workspace

USER sandbox