### docker swaggerapi/swagger-ui

```shell
docker pull swaggerapi/swagger-ui
# ${PWD} return the path in powershell in linux just use $pwd is the same :)
docker run --rm -p 80:8080 -e SWAGGER_JSON=/app/mirrorFinder/openapi.json -v ${PWD}\..\:/app swaggerapi/swagger-ui
```
