# POD vs Service vs Deployment.
    =>
    🚀 1. Pod
    
    The smallest deployable unit in Kubernetes.
      - Represents a single instance of a running container (or multiple containers sharing resources).
      - Not durable — if it dies, it's not automatically restarted unless part of a Deployment or other controller.
    
    ✅ Example Use:
      - Testing a single container.
      - Debugging.
    🔴 Limitations:
      - Manual management.
      - No built-in self-healing or scaling.

      📦 2. Deployment
      
        - A higher-level abstraction that manages Pods.
        - Ensures the desired number of Pods are always running.
        - Supports:
            - Self-healing (restarts Pods on failure)
            - Rolling updates & rollbacks
            - Scaling
        - ✅ Example Use:
            - Running production apps.
            - Auto-updating containers with zero downtime.

      🌐 3. Service

        - A stable network abstraction in front of Pods.
        - Provides a permanent IP + DNS name to reach Pods.
        L- oad balances traffic to matching Pods.
        ✅ Types:
        - ClusterIP (default): Internal-only access
        - NodePort: Exposes on each node’s IP:Port
        - LoadBalancer: External access via cloud provider
        - ExternalName: Maps to external DNS




## Summary:
    =>
    | Feature        | Pod               | Deployment           | Service                  |
    | -------------- | ----------------- | -------------------- | ------------------------ |
    | Purpose        | Runs containers   | Manages Pods         | Exposes Pods on network  |
    | Self-healing   | ❌                 | ✅                    | ❌                        |
    | Scaling        | ❌                 | ✅                    | ❌ (used with Deployment) |
    | Load Balancing | ❌                 | ❌                    | ✅                        |
    | Stable IP/DNS  | ❌ (ephemeral)     | ❌                    | ✅                        |
    | Rolling Update | ❌                 | ✅                    | ❌                        |
    | Use Case       | Testing/debugging | Production workloads | Internal/external access |



🧠 Simple Analogy:

Pod: A single worker
Deployment: A manager that ensures the right number of workers are always present
Service: A receptionist that forwards traffic to available workers
