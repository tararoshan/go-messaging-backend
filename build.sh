# Adapted from https://stackoverflow.com/questions/43360476/elastic-beanstalk-procfile-for-go
# Stops the process if something fails
set -xe

# All of the dependencies needed/fetched for your project
go get "github.com/gorilla/mux"

# Create the application binary that EB uses
GOOS=linux GOARCH=amd64 go build -o bin/application -ldflags="-s -w"
