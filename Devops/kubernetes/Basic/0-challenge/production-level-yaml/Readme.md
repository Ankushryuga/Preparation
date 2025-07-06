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



## Prometheus Scraping & Adapters:
    =>
    1. Install Prometheus and Grafana via Helm
    - helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
    - helm install prometheus prometheus-community/kube-prometheus-stack --namespace monitoring --create-namespace

    This deploys Prometheus, Grafana, and the adapter to serve metrics to Kubernetes .

# b) Define custom rules in prometheus-adapter for HPA:
    =>
    rules:
      custom:
      - seriesQuery: 'http_requests_total'
        resources:
          overrides:
            kubernetes_namespace:
              resource: namespace
            kubernetes_pod_name:
              resource: pod
        name:
          matches: "^(.*)_total"
          as: "${1}_rate"
        metricsQuery: 'sum(rate(http_requests_total[2m])) by (pod, namespace)'


    This publishes a metric like pods/http_requests_rate_avg accessible via the custom metrics API


# c) Use a Prometheus-based HPA:
    =>
    apiVersion: autoscaling/v2
    kind: HorizontalPodAutoscaler
    metadata:
      name: nginx-prom-hpa
    spec:
      scaleTargetRef:
        apiVersion: apps/v1
        kind: Deployment
        name: nginx-prod
      minReplicas: 2
      maxReplicas: 10
      metrics:
      - type: Pods
        pods:
          metric:
            name: http_requests_rate
          target:
            type: AverageValue
            averageValue: "10"
    This scales based on average requests per second per Pod 


# 4. 📈 Observability (Prometheus + Grafana)
        
        Your Helm stack includes:
        
        Prometheus: Scrapes metrics via the annotations
        Grafana: Prebuilt dashboards & HPA metrics (e.g., current vs desired replicas)



## summary:

| Feature                                               | YAML snippet                                  |
| ----------------------------------------------------- | --------------------------------------------- |
| 🤖 Base Deployment w/ Probes & Prometheus annotations | `nginx-prod-deployment.yaml`                  |
| 🔁 HPA (built-in CPU)                                 | `nginx-hpa.yaml`                              |
| 📊 Full Prometheus stack                              | Helm install of `kube-prometheus-stack`       |
| 🌐 Custom HPA via Prometheus metrics                  | Custom metrics rules + HPA → `nginx-prom-hpa` |


## deploy:
        =>
        kubectl apply -f nginx-prod-deployment.yaml
        kubectl apply -f nginx-hpa.yaml
