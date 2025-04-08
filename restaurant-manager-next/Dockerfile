# Use Node.js LTS version as the base image
FROM node:20-alpine

# Set working directory
WORKDIR /app

# Create a non-root user to own the files and run the application
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Change ownership of the app directory
RUN chown -R appuser:appgroup /app

# Switch to non-root user before installing dependencies
USER appuser

# Copy package.json and package-lock.json with correct ownership
COPY --chown=appuser:appgroup package*.json ./

# Install dependencies as the non-root user
CMD ["ls", "-l", "/app"]
RUN npm install

# Copy the rest of the application with correct ownership
COPY --chown=appuser:appgroup . .

# Build the Next.js application
RUN npm run build

# Expose the port the app will run on
EXPOSE 3000

# Command to start the application
CMD ["npm", "start"]
