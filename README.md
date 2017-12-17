# kubely-services
This repository holds all services needed by Kubely

### test:

```
curl -d '{"kedge":[{"filename":"httpd","data":{"name":"httpd","containers":[{"image":"centos/httpd"}],"services":[{"name":"httpd","type":"NodePort","portMappings":["8080:80"]}],"routes":[{"name":"httpd","to":{"kind":"Service","name":"httpd"}}]}}]}' -X POST http://localhost:9999/api/v1/generate
```