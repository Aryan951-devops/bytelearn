FROM node:22-slim

# Install TypeScript compiler globally
RUN npm install -g typescript ts-node

WORKDIR /workspace
# Give ownership of /workspace to the built-in node user
RUN chown node:node /workspace

# Switch to the existing default user
USER node