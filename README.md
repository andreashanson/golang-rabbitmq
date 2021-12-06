# Start rabbit container on localhost

# Start the container
docker run -p 5672:5672 -p 5673:5673 -p 15672:15672 rabbitmq:3 

# Start management Tool
You might need to run the container with management enabled
Either run the container with prefix -management
Or lookup how to enable it in the running container.

docker container exec -it some-rabbit5 rabbitmq-plugins enable rabbitmq_management


# Enter locally
url: http://localhost:15672/
user: guest
pw: guest

