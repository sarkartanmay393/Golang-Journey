# Using this minimal base image for our work.
# It will consist of golang tool and nothing much more.
FROM golang:1.19-alpine

# Setting our current directory to /app inside our base image.
WORKDIR /app

# Copying everything from our Web-Server directory to our image's /app directory.
COPY . .

# Now all are necessary files are in our image.
# So we can start our commands to run the go app.
RUN go mod download
RUN go build main.go



# Run the executable.
CMD ./main