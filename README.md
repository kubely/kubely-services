# kubely-services
This repository holds all services needed by Kubely


## Prerequisites
1. [glide](https://github.com/Masterminds/glide/releases)
2. go (>=1.8.4)
3. make


### install:
```
make install
```

### build:
```
make build
```

### run server for local development
```
make run
```

### test:

```
curl -d '{"kedge":[{"filename":"httpd","data":{"name":"httpd","containers":[{"image":"centos/httpd"}],"services":[{"name":"httpd","type":"NodePort","portMappings":["8080:80"]}],"routes":[{"name":"httpd","to":{"kind":"Service","name":"httpd"}}]}}]}' -X POST http://localhost:9999/api/v1/generate
```