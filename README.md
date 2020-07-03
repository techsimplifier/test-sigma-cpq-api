# test-sigma-cpq-api
Tool to investigate Sigma CPQ swagger API

## API Client generation using swagger
The client code was generated using https://github.com/swagger-api/swagger-codegen

## Notes on VPN
Ensure that your VPN does not block access to the endpoint url. This can happen if your VPN DNS does not resolve the endpoint.

## Building and Running the tool using Docker

You must have [Docker](https://www.docker.com/) installed for the commands to work.

### Build it
```bash
docker build -t test-sigma-cpq-api:0.0.0 .
```

## Run it
```bash
docker run test-sigma-cpq-api:0.0.0 quotes get --help
```