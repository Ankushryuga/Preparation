**Throughput**:
  => It refers to the amount of work a system can process or handle within a given timeframe. It's a measure of how much data, requests or operations a system can handle per unit of time, like requests, or operations a system can handle per unit of time, like requests per second or transactions per minute. High throughput is crucial for systems that need to handle large volumes of data or traffic efficiency.
  

## Breakdown of system design for easy understanding

**Before doing anything clarify the requirements correctly and avoid asking repeated questions**.  
# Step 1: requirement gathering
- Core Features or functionality of systems: (e.g., CRUD operations)
  
- Non Functional Requirement: It define how a system performs rather than what it does.
# Common Non-function requirements:
          | Category              | Description                               | Examples                                            |
      | --------------------- | ----------------------------------------- | --------------------------------------------------- |
      | **Performance**       | How fast the system responds or processes | Response time < 200ms, throughput ≥ 10K req/sec     |
      | **Scalability**       | How well the system handles growth        | Support 10× users without redesign                  |
      | **Availability**      | System uptime over time                   | 99.99% uptime (52 mins/year downtime)               |
      | **Reliability**       | Consistency of system operation           | No crashes in 6 months; retries for failed messages |
      | **Maintainability**   | Ease of updates or bug fixes              | Modular codebase, low coupling                      |
      | **Security**          | Protect data and operations               | Encryption, authentication, audit logging           |
      | **Usability**         | User-friendliness and accessibility       | Intuitive UI, WCAG compliance                       |
      | **Portability**       | Ability to run in different environments  | Deployable on AWS, Azure, and GCP                   |
      | **Testability**       | Ease of testing the system                | Unit tests, CI/CD pipelines, mockable components    |
      | **Compliance**        | Regulatory/legal constraints              | GDPR, HIPAA, PCI-DSS requirements                   |
      | **Capacity**          | Maximum limits of the system              | Supports 5M users, 1TB storage without degradation  |
      | **Disaster Recovery** | Recovery from failures or disasters       | RTO = 5 mins, RPO = 1 min                           |


# Example : URL Shortener.
          - Functional requirement: Shorten a given long URL.
          - Non-functional requirements:
            1. Response time < 100ms
            2. 99.99% availability
            3. Horizontal scalability for 1B urls

          
