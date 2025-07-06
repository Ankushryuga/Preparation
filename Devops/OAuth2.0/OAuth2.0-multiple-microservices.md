## OAuth2.0:
    =>
    OAuth 2.0 is an authorization framework that allows applications to obtain limited access to user 
    resources without exposing credentials.


## Key concepts of OAuth2.0:
    =>
    1. Resource Owner: user.
    2. Authorization server: Issues access tokesn after authenticating the user.
    3. Client : The application interface.
    4. Resource Server: Hosts the protected APIs.


## OAuth2.0 Flow in in Microservice Architecture:
    => 
    1. You have multiple microservices (e.g., ServiceA, ServiceB, ServiceC).
    2. A Gateway or API Gatway sites in front of them.
    3. An Authorization server(like OAuth0, Okta, Keyclock, or a custom server) is used.
    4. A Client(like a web or mobile app) interact with services.

## Step-by-step OAuth2.0 Flow::
    =>
    1. User Authentication via Client
      1.1. The user logs in via the client app (e.g., a single-page application or mobile app).
      1.2. The client redirects the user to the Authorization Server to authenticate.
    2. Access Token Issuance:
      2.1. After successful authentication, the Authorization Server issues an access token (and optionally a refresh token).
      2.2. This token is usually a JWT or opaque token.
    3. Client Calls API Gateway:
      3.1. The client sends (HTTP, gRPC, or graphQL) requests to the API Gateway, attaching the access token in the Autherization header.
          => Authorization: Bearer <access_token>
    4. API Gateway validate token:
        4.1. API Gateway validates the token, if its a JWT, it verifies the token signature using a public key(e.g, from a JWKS endpoint)
        4.2. If it's opaque, it calls the Authorization server's introspection endpoint.
    5. Gateway Forwards Requests to Microservices:
        5.1. Once validated, the gateway can, forward the request to the appropriate microservice., optionally include user identity and claims (e.g., roles user ID) in a header like X-User-Id.
    6. Microservice Enforce Authorization:
        6.1. Each microservice, extract user identity and claim from the token(or from headers).
        6.2. Enforces fine-grained authorization based on user roles, scopes or permissions.
        6.3. Optinoally re-validates the token if not relying solely on the gateway.



# 🛡️ Security Considerations
    => 
    1. HTTPS: Always use HTTPS to protect tokens in transit.
    2. Token Expiration: Short-lived access tokens + refresh tokens.
    3. Token Storage: Store tokens securely in client apps (e.g., avoid localStorage in web apps).
    4. Scopes and Roles: Use token scopes or roles for fine-grained access control.
    5. Rate Limiting & Throttling: Protect services from abuse.
    


# 📌 Example Use Case
System:
Authorization Server: Keycloak
Client: React app
Microservices: AuthService, ProductService, OrderService
Gateway: NGINX with Lua plugin or API Gateway with token validation

# Flow:
React app redirects user to Keycloak.
Keycloak authenticates and sends back a JWT.
React app sends requests with JWT to API Gateway.
API Gateway validates JWT, forwards request to ProductService.
ProductService uses claims in JWT (e.g., role=admin) to allow/deny operations.






#### NOTE VVI: Service-to-service Communication..
    => Sometime microservices talk to each other, you have 2 options:
    Option 1: Propagate the User's token:
        1.1. Forward the original access token b/w services.
        1.2. Use it to maintain user context and authorization.
    Option 2: Client Credential Flow:
        2.1. Use the OAuth2.0 client credentials grant b/w services.
        2.2. Each service has its own credentials to obtain a service-level token.



| Component     | Responsibility                         |
| ------------- | -------------------------------------- |
| Client        | Gets access token                      |
| Auth Server   | Issues and validates tokens            |
| Gateway       | Authenticates token and routes traffic |
| Microservices | Authorize and process requests         |



## When using a mix of HTTP, gRPC, and GraphQL in a microservices architecture, the OAuth 2.0 flow remains conceptually the same, but how the access token is passed and validated depends on the protocol and client implementation.
    =>
    🎯 High-Level Architecture
        1. Client (e.g., web or mobile app) gets an access token from the Authorization Server.
        2. This token is passed along in requests to:
            2.1. HTTP REST endpoints  
            2.2. gRPC services
            2.3. GraphQL APIs
        3. Each service validates the token (or relies on a gateway to do it) and enforces authorization.

# 🔁 Unified OAuth 2.0 Flow (Protocol-Agnostic)
    1. User logs in via client → gets an access token from the Auth Server.
    2. Client stores token securely (e.g., in memory or secure storage).
    3. Client includes token in requests to your services:
        3.1. HTTP: in Authorization header
        3.2. gRPC: via metadata
        3.3. GraphQL: in headers or variables
    4. Gateway or service validates token using:
        4.1. JWT validation (public key)
        4.2. Or token introspection (for opaque tokens)
    5. Service authorizes request using token claims (roles, scopes, etc.)


## Protocol-Specific Details:
    1. HTTP REST     
        1.1. Token Location: HTTP Authroization header
            GET /orders HTTP/1.1
            Authorization: Bearer <access_token>
        1.2. Validation:
            - At API Gateway
            - Or Each Service(if there's no central gateway)
            
    2. gRPC Services:
        2.1. gRPC uses metadata for sending auth tokens
        2.2. Token Location: Authorization metadata in the gRPC request.
        2.3. Client Example:
            metadata.NewOutgoingContext(ctx, metadata.Pairs("authorization", "Bearer <access_token>"))
        2.4. Server middleware:
            - Interceptor checks the token in the metadata.
            - validates JWT or introspects token
        2.5. Use interceptors/middleware to centralize token validation.
    3. 🧠 GraphQL APIs
        =>GraphQL usually runs over HTTP (though it can be over WebSocket), so you treat it like a REST API.
        3.1 Token Location: Same Authrorization header as REST.
            POST /graphql HTTP/1.1
            Authorization: Bearer <access_token>
        3.2. If using Apollo Server, token extraction is done in the context function:
            const server = new ApolloServer({
              context: ({ req }) => {
                const token = req.headers.authorization || '';
                return { token };
              },
            });
        3.3. You can also pass the token in query variables (less secure):
           # graphql:
            query {
              user(token: "<access_token>") {
                id
              }
            }


| Approach                  | Description                                                | Pros                                         | Cons                                  |
| ------------------------- | ---------------------------------------------------------- | -------------------------------------------- | ------------------------------------- |
| **Centralized (Gateway)** | Gateway validates token and injects identity into requests | Simplifies services, single point of control | Gateway is a bottleneck/failure point |
| **Decentralized**         | Each service validates token                               | Microservices are autonomous                 | Duplicates token logic, more config   |


Best practice: validate at the edge (gateway), pass claims to internal services via headers (e.g., X-User-Id, X-Roles).



🔐 Internal Service-to-Service Calls
When microservices talk to each other (via gRPC, REST, or GraphQL), two options:

1. Propagate User Token
Service A forwards the user’s token to Service B.

Preserves user context across service boundaries.

2. Client Credentials Flow (Service Identity)
Service A uses its own credentials to get an access token.

Useful when Service A is acting on its own behalf, not the user's.

POST /token
grant_type=client_credentials
client_id=service-a
client_secret=...

✅ Summary by Protocol
| Protocol      | Token Location              | Validation                | Notes             |
| ------------- | --------------------------- | ------------------------- | ----------------- |
| **HTTP REST** | `Authorization: Bearer`     | JWT decode or introspect  | Most common       |
| **gRPC**      | `Authorization` in metadata | Interceptor or middleware | More setup needed |
| **GraphQL**   | HTTP header or query var    | Context processing        | Same as REST      |



🚧 Extra Considerations
Token Expiration: Use refresh tokens in the client or handle token renewal.

Scopes/Roles: Include in token and use for fine-grained access control.

Revocation: Short-lived tokens + blacklist support (if needed).

Multi-protocol Auth Middleware: Consider shared libraries across services for token validation.
