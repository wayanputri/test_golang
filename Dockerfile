FROM golang:1.21-alpine

# create directory folder
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable file with name "test_golang"
RUN go build -o test_golang

# run executable file
CMD ["./test_golang"]