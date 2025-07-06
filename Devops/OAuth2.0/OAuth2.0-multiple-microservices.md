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

