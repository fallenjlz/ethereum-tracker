## Usage

To run Ethereum Tracker, use the following command:

```shell
docker-compose up --build
```

This command builds the Docker image and starts the service. The application will be accessible at `http://localhost:8080`.

The application exposes an API endpoint for fetching the latest 5 logged Ethereum events.

- **Endpoint:** `/`
- **Method:** `GET`
- **Description:** Returns the latest Ethereum events logged by the application.
- **Response Format:** JSON
- **Sample Response:**

```shell
curl 127.0.0.1:8080/ | python3 -m json.tool
```

```
[
    {
        "name": "Transfer",
        "block_number": 1234567,
        "index": 1,
        "data": {
            "from": "0x...",
            "to": "0x...",
            "tokens": "100"
        }
    },
    ...
]
```

## Project Structure

- `main.go`: Entry point of the application.
- `listener/`: Contains logic for listening to Ethereum blockchain events.
- `handler/`: Handles processing and logging of blockchain events.
- `monitor/`: Manages Prometheus metrics.
- `Dockerfile`: Defines the Docker setup for the application.
- `docker-compose.yml`: Docker Compose configuration for running the application and Grafana Agent.
- `config.yaml`: Configuration file for Ethereum node URL and contract address.
- `agent.yaml`: Configuration file for Grafana Agent.

