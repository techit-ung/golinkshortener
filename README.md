# A Simple Link Shortener
A simple link shortener. Inspired by Google's go link story.

This app is build to be lightweight in-company url shortener.
The reason why YAML is because of it's kinda straightforward to do version control of the YAML file.
In case we want a better performance, we can load YAML data into memcache and query for link through that instead.

Written in Go. You can use build it locally or use Docker.

## Running using Docker
### Build docker image with command below:
```shell
docker build -t linkshortener .
```

### Run docker using this:
```shell
docker run \
-it -p 9091:9090 \
--rm \
--name myshortener \
-v "$(pwd)"/new_link.yaml:/root/new_link.yaml \
linkshortener --path=/root/new_link.yaml
```

## Links YAML format
```yaml
links:
  link1: https://www.github.com
  link2: https://www.google.com
```