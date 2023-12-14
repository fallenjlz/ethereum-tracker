## Usage

To run Ethereum Tracker, use the following command:

```shell
docker-compose up --build
```

This command builds the Docker image and starts the service. The application will be accessible at `http://localhost:8080`.

## Project Structure

- `main.go`: Entry point of the application.
- `listener/`: Contains logic for listening to Ethereum blockchain events.
- `handler/`: Handles processing and logging of blockchain events.
- `monitor/`: Manages Prometheus metrics.
- `Dockerfile`: Defines the Docker setup for the application.
- `docker-compose.yml`: Docker Compose configuration for running the application and Grafana Agent.
- `config.yaml`: Configuration file for Ethereum node URL and contract address.
- `agent.yaml`: Configuration file for Grafana Agent.