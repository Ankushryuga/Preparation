# Production level K8s yaml
    =>
    | Feature                       | Benefit                      |
    | ----------------------------- | ---------------------------- |
    | **Replicas = 3**              | High availability            |
    | **Resource requests/limits**  | Prevents overuse             |
    | **Liveness/readiness probes** | Health checks                |
    | **Ingress + TLS**             | Secure public access         |
    | **Labels**                    | Organized, manageable        |
    | **ClusterIssuer Annotation**  | Auto TLS (with cert-manager) |


## How to deploy:
    => 
    kubectl apply -f nginx-prod-deployment.yaml
    kubectl apply -f nginx-prod-service.yaml
    kubectl apply -f nginx-prod-ingress.yaml
