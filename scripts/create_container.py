import docker, sys

client = docker.from_env()

image = sys.argv[1]
name = sys.argv[2]

client.images.pull(image)

container = client.containers.run(image, name=name, detach=True)
print("Container run with image {}".format(image))