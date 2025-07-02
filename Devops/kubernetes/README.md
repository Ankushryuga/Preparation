# Kubernetes interview questions:
# ✅ Key Production Considerations
    =>
    1. Secrets for credentials
    2. ConfigMaps for environment variables
    3. Liveness & Readiness Probes
    4. Resource limits
    5. Autoscaling (HPA)
    6. Ingress with HTTPS (e.g. via cert-manager + NGINX)
    7. Separate namespaces
    8. Database management via Helm (optional)
    9. Storage class for PVCs
    10. Monitoring & logging (Prometheus, Loki, Grafana, etc.)



    k8s/
    ├── base/
    │   ├── namespace.yaml
    │   ├── config/
    │   │   ├── golang-configmap.yaml
    │   │   ├── spring-configmap.yaml
    │   │   └── db-secrets.yaml
    │   ├── db/
    │   │   ├── postgres.yaml
    │   │   └── mysql.yaml
    │   ├── services/
    │   │   ├── golang-deployment.yaml
    │   │   ├── spring-deployment.yaml
    │   │   ├── golang-service.yaml
    │   │   └── spring-service.yaml
    │   ├── ingress/
    │   │   └── ingress.yaml
