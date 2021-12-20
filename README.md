# Start rabbit container on localhost

# Start the container
docker run -p 5672:5672 -p 5673:5673 -p 15672:15672 rabbitmq:3 

# Start management Tool
You might need to run the container with management enabled
Either run the container with prefix -management
Or lookup how to enable it in the running container.

docker container exec -it <container_id> rabbitmq-plugins enable rabbitmq_management

# Enter GUI locally
url: http://localhost:15672/
user: guest
pw: guest

# Application flow
scheduler starts cronjobs that is created in jobs.go 
These cronjobs will publish jobs on the jobs queue depending on the cronjob schedule.

The scheduler service takes the producer as argument which is used to push these jobs on the jobs queue.

The consumer listens on the jobs queue and parse them to messages.

# Run image locally
docker inspect <rabbit_container_id>
enter ip in .env for RABBIT_HOST

