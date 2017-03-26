```

curl http://localhost:8080/healthy -i

curl -X POST --data 0  http://localhost:8080/healthy -i

docker run -p 8080:8080 -d --name hc --rm hc           

docker inspect hc -f '{{ json .State }}' | jq

```
