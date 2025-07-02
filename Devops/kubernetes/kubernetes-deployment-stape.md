# step 1: Building and containerizing the microservice:
  ## for single module or single microservice::
    =>https://github.com/OpenLiberty/guide-kubernetes-intro/tree/prod
    1. docker build -t system:1.0-SNAPSHOT system/.

    The -t flag in the docker build command allows the Docker image to be labeled (tagged) in the name[:tag] format. 
    The tag for an image describes the specific image version. If the optional [:tag] tag is not specified, the latest tag is created by default.

    2. docker images

  # Make sure images are up using docker images command.

## Deploying the microservices:
    => now that Docker images are built, deploy them using a kubernetes resource definition.
    A Kubernetes resource definition is a yaml file that contains a description of all your deployments, services, or any other resources that you want to deploy. 
    All resources can also be deleted from the cluster by using the same yaml file that you used to deploy them.

# Create the kubernetes configuratoin file in start directory:
    => kubernetes.yaml
    

# Ensure your local cluster is running:
    =>
    1. Minikube: minikube start
    2. Make sure your image is available to the cluster:
        docker build -t system:1.0-SNAPSHOT system/.

    3. Apply the kubernetes config:
        kubectl apply -f kubernetes.yaml
    4. check deployment status:
        kubectl get pods
        kubectl get services
    5. Access your app:
        on web:=> http://localhost:30080
        
