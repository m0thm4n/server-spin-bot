import docker, sys

client = docker.from_env()

container_to_get = sys.argv[1]

container = client.containers.get(container_to_get)

print(container.attrs['Config']['Image'])

print(container.logs())

container.start()