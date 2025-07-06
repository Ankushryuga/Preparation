
## What is Deployment?
    =>  A deployment in k8s is a controller that:
        1. Manages a set of pods (instances of your app).
        2. Ensures the correct number of pods are running.
        3. Provides self-healing 
        4. Enables Rolling updates
        5. Allow easy scalling up/down
        
        ## example:
        1. kubectl create deployment myapp --image=nginx        //its imperative and it creates one pod


        ## delcarative way (using yaml file):
        1. kubectl create -f deploy-example-declarative.yaml
        
        ## To increase replicas use scale up:
        1. kubectl scale deployment myapp --replicas = 4    //4 pods.

        ## verify: 
        1. kubectl get deployment myapp
        2. kubectl get pods
        

        ## Cleanup:

        kubectl delete deployment myginx1
        # for yaml file..
        kubectl delete deploy mynginx2
        
