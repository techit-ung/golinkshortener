Build docker image with command below:
```shell
docker build -t linkshortener .
```

Run docker using this:
```shell
docker run \
-it -p 9091:9090 \
--rm \
--name myshortener \
-v "$(pwd)"/new_link.yaml:/root/new_link.yaml \
linkshortener --path=/root/new_link.yaml
```