# Get Namespaces:
    => 
    1. kubectl get namespaces
    2. kubectl get ns

# Get the pods list
    =>
    1. kubectl get pods  //it will get pods from the default namespaces.
    2. kubectl get pods --namespace=kube-system
    3. kubectl get pods -n kube-system


# Change namespace:
    => change the namespace to the docker one and get the pods list.
    kubectl config set-context --current --namespace=kube-system
    kubectl get pods
    
    => change to default
    kubectl config set-context --current --namespace=default
    kubectl get pods


# Create and delete namespace
    =>
    kubectl create ns [name]
    kubectl get ns
    kubectl delete ns [name]

    
