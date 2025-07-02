#  🧱 1. Common Structure in All Kubernetes YAML
    =>
    apiVersion: <API version>   # Defines which version of the Kubernetes API to use
    kind: <Resource type>       # e.g., Deployment, Service, Pod
    metadata:                   # Metadata about the object
      name: <name>              # Unique name for the resource
      labels:                   # Key-value pairs to organize and select resources
        <key>: <value>
      annotations:              # Optional metadata (not used for selection)
        <key>: <value>
# 🚀 2. Deployment Resource  
    =>
    kind: Deployment
    spec:
      replicas: <number>                # Number of desired pod replicas
      selector:
        matchLabels:                    # Labels used to match pods controlled by this Deployment
          app: my-app
      template:                         # Template used to create pods
        metadata:
          labels:                       # Labels assigned to created pods
            app: my-app
        spec:                           # Pod specification
          containers:                   # List of containers in the pod
          - name: <container-name>
            image: <image-name>         # Docker image to use
            ports:
            - containerPort: <port>     # Port your app listens on inside the container
            env:                        # (Optional) Environment variables
            - name: ENV_VAR_NAME
              value: "some-value"
            resources:                  # (Optional) Resource requests and limits
              requests:
                cpu: "100m"
                memory: "128Mi"
              limits:
                cpu: "500m"
                memory: "256Mi"
            volumeMounts:               # (Optional) Mount volumes into the container
            - name: config-volume
              mountPath: /app/config
          volumes:                      # (Optional) Define volumes used by the pod
          - name: config-volume
            configMap:
              name: app-config
          imagePullPolicy: Always|IfNotPresent|Never  # Image pull behavior
          restartPolicy: Always|OnFailure|Never       # Pod restart behavior

# 🌐 3. Service Resource
    =>
    kind: Service
    spec:
      type: ClusterIP|NodePort|LoadBalancer|ExternalName
      selector:
        app: my-app                # Targets pods with this label
      ports:
      - protocol: TCP|UDP
        port: <service-port>       # Port exposed by the service
        targetPort: <pod-port>     # Port on the container
        nodePort: <port>           # (Only for NodePort type)


# 📦 4. ConfigMap Resource
    =>
    kind: ConfigMap
    data:
      <key>: <value>               # Key-value pairs accessible by pods


# 🔑 5. Secret Resource
    =>
    kind: Secret
    type: Opaque                  # Basic key-value secret (base64 encoded)
    data:
      username: dXNlcm5hbWU=      # Base64 encoded value
      password: cGFzc3dvcmQ=


# 📄 6. Ingress Resource (Optional for external routing)
    =>
    kind: Ingress
    spec:
      rules:
      - host: myapp.local
        http:
          paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: my-service
                port:
                  number: 80

                  
# 📊 Summary Table
| Tag/Field                  | Description                                    |
| -------------------------- | ---------------------------------------------- |
| `apiVersion`               | Kubernetes API version used                    |
| `kind`                     | Type of object: Pod, Service, Deployment, etc. |
| `metadata.name`            | Name of the object                             |
| `metadata.labels`          | Key-value labels to identify and group objects |
| `spec.replicas`            | Number of Pods to run in Deployment            |
| `spec.selector`            | Match labels for selecting Pods                |
| `template.spec.containers` | Container config for each Pod                  |
| `image`                    | Docker image used in container                 |
| `ports.containerPort`      | Container port the app listens on              |
| `env`                      | Environment variables in container             |
| `resources`                | CPU/memory requests and limits                 |
| `volumeMounts`             | Mount volumes inside container                 |
| `volumes`                  | Volumes attached to the Pod                    |
| `service.type`             | Access type: `ClusterIP`, `NodePort`, etc.     |
| `service.selector`         | Selects Pods for the service                   |
| `targetPort`               | Port on the container to forward traffic       |
| `nodePort`                 | External port on the node (for NodePort)       |
