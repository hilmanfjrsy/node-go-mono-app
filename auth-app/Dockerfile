# Use an official Node.js runtime as a parent image
FROM node:14-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY package*.json ./
COPY . /app

# Install any needed packages
RUN npm install

# Make port 5001 available to the world outside this container
EXPOSE 5001

# Run the command to start the app
CMD ["npm", "start"]
