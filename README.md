# ElasticSearch Curator

This program for batch processing to clean up ElasticSearch indexes. [curator](https://github.com/elastic/curator) is executed on a regular basis. **It was created with the assumption that CronJob is not available in Kubernetes, and uses Deployment to perform periodic execution in the process.**

## Usage

### Local

```
go run main.go
```

### Docker

```
docker build -o es-curator .
docker run -it --rm es-curator
```

### Kubernetes

```
kubectl apply -f k8s/deployment.yaml
```
