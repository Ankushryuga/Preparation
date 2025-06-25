# All necessary concept of java, golang microservices.
# What is microservices?
    => These are small, loosely coupled service that is designed to perform a specific business function and each microservices can be developed, deployed, and scaled independently.


# How do microservices work?
    =>
    Microservices break complex applications into smaller, independent services that work together, enhancing scalability, and maintenance.
    # How it works:
    1. Applications are divided into self-contained services, each focused on a specific function, simplifying development and maintenance.
    2. Each microservice handles a particular business feature, like user authentication or product management, allowing for specialized development.
    3. Services interact via APIs, facilitating standardized information exchange and integration.
    4. Microservices can be updated independently, reducing risks during changes and enhancing system resiliences.

# Microservices Design Pattern:
    =>
    1. API Gateway Pattern: API Gateways act as a traffic controller for your microservices, its like having one main door to enter a building instead of using different doors for different rooms.
        1.1. It gives one entry point for all client request
        1.2. Handle cross-cutting concerns like security and monitoring
        1.3. Can transform request and response.

        Short: its a single entry point for all client requests to microservices, Hides internl architecture, enables rounting, aggregation, authentication, throttling etc.

        Tools: Kong, Nginx, AWS API Gateway.
        It avoid tight coupling b/w clients and services.
    

    2. Circuit Breaker Pattern: It's like having a backup plan when something goes wrong, if service keeps failing, the circuit breaker stops trying to use it and returns a backup response instead.
        Short: Prevents a network/service failure from cascading, detects failures and stops calling a failing service temporarily.
        It improves fault tolerance and resiliency.

    3. Service Discovery Pattern: It enables services to find and communicate with each other dynamically, useful in dynamic environment like kubernetes.
        It avoids hardcoding service locations.

    4. Sidecar Pattern: Deploys helper components alongside your service., offloads non-core features from main service. 
        It promotes separation of concerns.

    5. Strangler Fig Pattern: Gradually replaces a legacy monolith with microservices., it allows safe migration without a full rewrite.

    6. CQRS (Command Query Responsibility Segregation): Separate read and write operations using different models. improves performance, scalablity, and complexity handling in microservices.
    Optimizes data operations per use case.

    7. Saga Pattern: Manages distributed transaction across services via a sequence of local transactions.
    maintains consistency without two-phase commit.

    8. Event sourcing pattern: stores state changes as a sequence of events, supports temporal state and replayability.

    9. Database per Service pattern:
        => Each service has its own database, improves service autonomy.

    10. Bulkhead Pattern:
        => Isolates resources per service or function to prevent cascading failures.
        
        
