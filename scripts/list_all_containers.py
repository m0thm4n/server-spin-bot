import docker, sys

client = docker.from_env()

print(client.containers.list())