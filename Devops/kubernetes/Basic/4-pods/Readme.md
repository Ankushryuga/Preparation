# Create the pod.
    => kubectl run mynginx --image=nginx

# Get a list of running pods.
    => kubectl get pods

# Get more info
    => kubectl get pods -o wide
    => kubectl describe pod mynginx

# Delete the pod:
    => kubectl delete pod mynginx

## Create a pod using declarative way:
    => kubectl create -f myapp.yaml

## Get some infor:
    => 
    1. kubectl get pods -o wide
    2. kubectl describe pod myapp-pod

## Attack terminal:
    =>
    kubectl exec -it myapp-pod -- bash

# cleanup:
    =>
    kubectl delete -f myapp.yaml
