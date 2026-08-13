FROM python:3.13-slim

# Create a non-root sandbox user
RUN useradd -m -u 1000 sandbox

WORKDIR /workspace
RUN chown sandbox:sandbox /workspace

USER sandbox