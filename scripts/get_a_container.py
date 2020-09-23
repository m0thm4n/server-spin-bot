import docker, sys

container_to_get = sys.argv[1]

client = docker.from_env()

container = client.containers.get(container_to_get)

print(container.attrs['Config']['Image'])

print(container.logs())