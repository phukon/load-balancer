# A simple Load Balancer
![golang](https://img.shields.io/badge/Golang-blue)
> Least Connections

This Go program demonstrates a basic load balancer that distributes incoming HTTP requests among a set of backend servers. The load balancer checks the health of each server through periodic health checks and adjusts the load distribution accordingly.

> I've kept the files s1, s2 and s3 to quickly spin up test servers on ports 8081, 8082 and 8083 respectively.

---

## Table of Contents

- Overview
- Features
- Getting Started
- Configuration
- Health Checks
- Load Balancing Strategy
- Customization
- Running the Program
- Contributing

## Overview
The code defines a Server struct representing a backend server and a Config struct for configuring the load balancer. The main program initializes servers, performs health checks, and distributes incoming requests based on the least active connections.

## Features
Periodic health checks to determine server health.
Round-robin load balancing strategy among healthy servers.
Configurable parameters via a JSON configuration file.

## Getting Started
1. Clone the repository:

```bash
git clone https://github.com/yourusername/load-balancer-go.git
```
2. Navigate to the project directory:

```bash
cd load-balancer-go
```

3. Run the program:

```bash
go run main.go
```

### Configuration
The load balancer reads its configuration from a config.json file. Example configuration:

```json
{
  "healthCheckInterval": "5s",
  "servers": ["http://localhost:8081", "http://localhost:8082"],
  "listenPort": ":8080"
}
```

- healthCheckInterval: The interval between health checks.
- servers: An array of backend server URLs.
- listenPort: The port on which the load balancer listens.

## Health Checks
Health checks are performed at regular intervals to assess the status of each backend server. Unhealthy servers (status code >= 500 or errors during health check) are temporarily excluded from load balancing.

## Load Balancing Strategy
The load balancer uses a round-robin strategy among healthy servers to distribute incoming requests. The server with the least active connections is chosen for each request.

## Customization
Feel free to customize the code according to your requirements. You can modify health check intervals, add more backend servers, or implement different load balancing strategies.

## Running the Program
Ensure Go is installed on your machine.
Execute go run main.go in the project directory.
Access the load balancer at http://localhost:8080 (or the specified listenPort).
Contributing
Contributions are welcome! Fork the repository, make your changes, and submit a pull request.