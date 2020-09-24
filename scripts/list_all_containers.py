import docker, sys

client = docker.from_env()

containers = client.containers.list(all=True)
for container in containers:
    print(container.attrs['Name'] + " " + str(container) + " " + container.attrs['Id'])