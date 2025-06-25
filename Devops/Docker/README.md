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
    FROM golang:1.21
    
    # Set the working directory
    WORKDIR /app
    
    # Copy go files
    COPY go.mod ./
    COPY go.sum ./
    RUN go mod download
    
    COPY . .
    
    # Build the Go app
    RUN go build -o app
    
    # Run the binary
    CMD ["./app"]

# Build Docker Image:
    => docker build -t my-python-app .
# Run Docker Container:
    => docker run -d --name my-running-app my-python-app
# Check Containers and Images:
    =>
    docker ps        # List running containers
    docker images    # List images


# Dockerfile vs Docker-Compose file:
    => 
    Dockerfile: its a script containing instruction to build a Docker image.
    # Dockerfile:
    FROM golang:1.21

    # set the working directory
    WORKDIR /app

    # copy go files.
    COPY go.mod ./
    COPY go.sum ./
    RUN go mod download

    COPY . .

    # Build the go app.
    RUN go build -o app

    # Run the binary.
    CMD ["./app"]
    
    # docker-compose.yml:
    version: '3.9'
    services:
      app:
        build: .
        ports: 
          - "8080:8080"
        depends_on:
          - db
        networks:
          - backend

     db:
      image: postgres:16
      restart: always
      environment: 
        POSTGRES_USER: user
        POSTGRES_PASSWORD: pass
        POSTGRES_DB: mydb
      volumes:
        - db-data:/var/lib/postgressql/data
      networks:
        - backend

    Volumes:
      db-data:

    networks:
      backend:


## Run it:
    docker-compose up --build

    
