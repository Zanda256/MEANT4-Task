# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

# Set the working directory to /factorial
WORKDIR /factorial

# Copy the go.mod and go.sum 
COPY go.mod ./
COPY go.sum ./

# Install any needed packages specified in requirements.txt
RUN go mod download

#Copy everything in the current folder to the image 
COPY . ./

# Make port 5100 available to the world outside this container
EXPOSE 5100

# Define environment variable
ENV FLOAT_PRECISION 15

#Compile our app and start the server binary
RUN make server
RUN make client

# Run app.py when the container launches
CMD [ "/serverbin" ]