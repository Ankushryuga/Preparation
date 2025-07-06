# Kubernetes:
    => its a container orchestration tool, that allow you to more than just containerization like, self healing, scalling, load balancing, monitoring etc, with minimal efforts.

# Main Features of K8S:
    =>
    1. Automated scheduling:  k8s automatically schedules containers to run on the available resources in the cluster.
    2. Self-healing: k8s automatically replaces failed containers and reschedules them on healthy nodes in the cluster.
    3. Automated rollouts and rollback: rollouts meaning enabling new version of s/w easily and roll them back in case of any issues.
    4. Horizontal scaling and load balancing
    5. configuration management
    6. Service discovery and networking.

  
# K8s components:
    => k8s have 2 main component:
        1. Master Nodes: it is responsible for managing worker node or minions.
            Master Node contains:
                1. Api Server: it acts as the entry point for interactions with the kubernetes cluster. it exposes the kubernetes API,
                which allows users, administrators, and other components to communicate with the cluster.
                2. etcd: it's a distributed key-value store that serves as kubernetes backing store for all cluster data, it holds the configuration data and the state of the entire cluster.
                this includes informations about pods, services, replication settings and more.
                
                3. Controller Manager: it is responsible for monitoring the state of various objects in the cluster and taking corrective actions to ensure the desired state in maintained. 
                It includes several built-in controllers, such as Replication controller, Deployment Controller, and StatefulSet Controller.

                4. Scheduler: It is responsible for placing the pods onto suitable worker nodes based on the requirement of pods (like resource availability , etc).

                
        2. Worker Nodes (minions): it is the heart of k8s cluster, they are responsible for running containers and  executing the actual workloads of your applications.
                1. Kubelet: its an agent that runs on each worker node and communicate with the master node, its primary resposibility is to ensure that containers within the pods are running and healthy as per the desired state defined in the cluster's configuration.
                    kubelet works closely with the master control plane to start, stop and manage containers based on pod specification.
                
                2. Kube Proxy: it set up routing and load balancing so that applications can seamlessly communicate with each other and external resources (incoming-request-> kube-proxy->send to available pods).

                3. Container Runtime: its a s/w resposible for running container on the worker nodes.

                4. Container Storage interfaces (CSI): worker nodes need to provide storage for persistent data. 


## Pods:
        => these are the building block in K8s that group one or more containers together and provide a shared environment for them to run within the same network and storage context.
            1. Container Co-location: pods allow you to colocate containers that need to work closely together within the same network namespace, this means they can communicate using localhost and share the name IP address and port space.

            2. Shared Storage and volume:Containers within a Pod share the same storage volumes, which allows them to easily exchange data and files. Volumes are attached to the Pod and can be used by any of the containers within it.

            3. Single Unit of Deployment: Kubernetes schedules Pods as the smallest deployable unit. If you want to scale or manage your application, you work with Pod replicas, not individual containers.

            4. Init Containers: A pod can include init container, which are containers that run before the main application containers.


## Controller:
        => it is responsible for maintaining the desired state of resources in the cluster, they monitor changes to resources and ensures that the actual state matches the intended state specified in the cluster's configuration.
        Controller automate tasks like scaling, self-healing, and application management.
       
        1. Replication Controller: The Replication Controller ensures a specified number of identical replica Pods are running at all times. If a Pod fails or is deleted, the Replication Controller creates a new one to maintain the desired number.

        2. Deployment: It provides RollOut and Rollback.

        3. StatefulSet: It is worker API that manages stateful application that requiers unique identities and persistent storage. They ensure that each Pod gets a stable, unique hostname and are created and scaled in a predictable order.

        4. DaemonSet: It is responsible for background process and ensures that a specific pod runs on every node in the cluster. they're useful for deploying monitoring agents, log collectors etc.

        5. Scalling (Horizontal, Vertical Pod Autoscaler)


## Services: 
        => 
        Services are a fundamental concepts that enables communciation and load balancing b/w different set of pods
        1. Service Type:
            - Cluster Ip: creates a internal virutal IP that exposes the service within the cluster, through cluster ip application can communucate within the cluster.
            - Node port: exposes an external load balancer that distributes traffic to the service across multiple nodes.
            - Load Balancer: Creates an external load balancer that distributes traffic to the Service across multiple nodes.
            - ExternalName: Maps the service to a DNS name, allowing you to reference services external to the cluster.

        2. Selectors and Labels:
            services uses selectors and labels to identify the pods they should target, labels are key-value pairs attached to Pods, and selectors define which Pods the service should include
            example: A service with a selector might target all Pods with the label: "app=web".

        3. Load Balancing: Services provide load balancing across multiple Pods with the same label. When you send traffic to a Service, it distributes the traffic evenly among the available Pods. This ensures that no single Pod gets overwhelmed with requests.

        4. Headless Services: A Headless Service is a special type of Service that doesn’t load balance or provide a stable IP. Instead, it allows you to access individual Pods directly using their IPs or DNS names. This is useful for applications that require direct communication with specific Pods.

## Volumes:
    => for persistant data:
     Types of Volumes: Kubernetes supports various types of volumes to accommodate different storage needs:
     1. EmptyDir: A Temp storage that's created when a pod is assigned to a node and deleted when the pod removed or reschedule.
     2. HostPath: Mount file or directory from the host machine into the container.
     3. ConfigMap and Secret Volumes: Special volumes that allow you to inject ConfigMap or Secret data as files into containers.
     4. NFS: Network File System volumes allow you to use network-attached storage in your containers.


## ConfigMap and secrets:
    =>
    configMaps: are text file that contains configuration related to your application.
    secrets: are for sensitive datas.

# Namespace: 
    => its way to organize and partition resources within a cluster. they provide a way to create multiple virtual clusters within the same cluster, allowing you to separate and manage resources.
    ## purpose of Namespace:
        - Resource isolation: Namespaces create isolated environments, so resources in one namespace are distinct from those in another.
        - Resource management:  Namespaces help organize and manage resources more effectively.
        - Access Control: Namespaces enable fine-grained access control.

    ## Default Namespace: When you create a Kubernetes cluster, there’s a default namespace where resources are created if you don’t specify a namespace explicitly.

    ## Namespace scope: Certain resources, like Nodes and PersistentVolumes, are not tied to a specific namespace and are accessible from all namespaces. Other resources, such as Pods, Services, ConfigMaps, and Secrets, belong to a specific namespace.

    ## Access Control and Quotas: Namespaces allow you to implement Role-Based Access Control (RBAC) to manage who can access or modify resources within a namespace. Additionally, you can set resource quotas and limits on namespaces to prevent resource overuse.


## Ingress:
    => Ingress is a resource that manages external access to services within your cluster. 
    It acts as a way to configure and manage routing rules for incoming traffic, allowing you to 
    expose your services to the outside world, often using HTTP and HTTPS protocols.




