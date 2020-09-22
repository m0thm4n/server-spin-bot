import docker, sys

client = docker.from_env()

image = sys.argv[1]
name_of_image = sys.argv[2]

client.images.pull(image)

client.containers.run(image, name=name_of_image, detach=True)
print("Container run with image {}".format(image))