# Stage 1: Build the application
FROM node:16 AS build

# Set the working directory
WORKDIR /frontend

# Copy the rest of the application code
COPY . .

# Install dependencies
RUN npm install

# Build the application
RUN npm run build

# Stage 2: Serve the application
FROM node:16

# Set the working directory
WORKDIR /frontend

# Copy the built application from the previous stage
COPY --from=build /frontend/build ./build

# Install serve to serve the built application
RUN npm install -g serve

# Expose the port the app runs on
EXPOSE 5173

# Command to run the application
CMD ["serve", "-s", "build"]