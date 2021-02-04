# Go Quote API
A simple golang quote API server (Docker and Kubernetes)

## Run
- To run the api server, run
```sh
docker-compose up
```

## Test
- To test the API use curl or postman.

### POST Quote
- URL: http://localhost:8080/quote
- Body
```json
{
    "author": "Jalal ad-Din Muhammad Rumi",
    "content": "What you seek is seeking you."
}
```

### GET Quotes
- URL: http://localhost:8080/quotes
```json
[{
    "author": "Jalal ad-Din Muhammad Rumi",
    "content": "What you seek is seeking you."
}]
```

## Deployment (Kubernetes cluster)
- To deploy rum
```sh
chmod +x kubectl.sh
./kubectl.sh
```