# Docker Interview questions 

# docker vs docker container vs docker image and how to create these
    => Docker: docker is a platform that enables developers to build, ship, and run applications in containers, it includes Docker Engine, Docker CLI, Docker Compose (for multi-container applications)

    => Docker Image: A docker image is a read-only template with the application code, libraries, dependencies, and configurations needed to run the application, built from A **Dockerfile**.
      => docker build -t my-image-name .
        -t: tag (name) for the image.
        .: build context (usually the current directory)


    => Docker container: A docker container is a running instance of an image, it's isolated, has its own filesystem, and runs based on the image.
      => it's like a live application running from the blueprint(image).
        => docker run -d --name my-container my-image


# 🛠️ Summary: Creation Workflow
# Write a Dockerfile – a text file with instructions to build your image.
      =>
      Example:
      # Dockerfile
      FROM python:3.9
      WORKDIR /app
      COPY . .
      RUN pip install -r requirements.txt
      CMD ["python", "app.py"]
# Build Docker Image:
    => docker build -t my-python-app .
# Run Docker Container:
    => docker run -d --name my-running-app my-python-app
# Check Containers and Images:
    =>
    docker ps        # List running containers
    docker images    # List images

    
