# requestbot
A bot to send http requests per second. Runs via a Docker image.

Requires task, docker, docker-compose

# Run
Run docker container
`task run`

# Usage
Set `command` in docker-compose to your desired RPS and URL.
`requestbot -c 10 -u https://google.com`

# Build binary
Build the Golang binary locally
`task build_bin`

# Clean
Removes local binary and docker image
`task clean`
