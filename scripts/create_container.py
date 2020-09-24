import docker, sys

client = docker.from_env()

image = sys.argv[1]
name = sys.argv[2]
port_number = sys.argv[3]

env = ["EULA=TRUE"]

client.images.pull(image)

port = {port_number+"/tcp": ('127.0.0.1', port_number)}

container = client.containers.run(image, ports=port, environment=env, mem_limit="8g", name=name, detach=True)
print("Container run with image {}".format(image))