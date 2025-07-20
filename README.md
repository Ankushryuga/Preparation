# Interview-Round-Prep

Be a Better Engineer


Practice LLD/HLD/System Design More and More within this given time period.


### VVI MOST Important Think loud your assumtions so that Interview is following you along the way.


## How to answer System Design Interview Questios?
**Step 1** Before designing anything, clairfy the requirements:
      - Core Features or functionality?.
      - Non Functional Requirements (e.g., Latency, Availability etc.)?
      - How many users are expected?
      - Is this a read-heavy or write-heavy system? (You have to identify yourself)

👉 Example: "When you say 'design a URL shortener,' should it support analytics, custom URLs, or expiration dates?"

**Step 2** Define the scope: Don't design the entire internet, Narrow the focus to MVP(Minimu Viable Product) first.
      - What is in scope? What's out of scope?
      - Are we building for scale from day one, or starting small?


**Step 3** Make Assumptions and Estimate:
      -  Traffic volumes (e.g., 10M requests/day)
      -  Data size(e.g., 500 bytes/request)
      -  Throughput, QPS, storage needs, etc.
💡 Tip: Interviewers want to see if you can reason about scale.


**Step 4** High Level Design: Sketch a block diagram with major components:
      - Client
      - Load Balancer
      - Application Servers
      - Database
      - Cache
      - Message Queues (If Applicable)
      - CDN, storage services, etc
Explain:
      - How data flows
      - How users interact with the system
      - key APIs/interfaces
      

**Step 5** Deep Dive on Components: Pick 1-2 components to go deeper into, depending on the interviewer's signals.
      Examples:
      - Database schema and Indexing
      - Consistency and Partitioning strategy
      - API design
      - Caching strategy (e.g., Redis, LRU)
      - Load balancing(round robin, sticky sessions)
      

**Step 6** Address Scalability & Reliability:
      Discus how you'll scale and make the system resilient:
      - Horizontal vs vertical scaling
      - Sharding or partitioning
      - Caching layers
      - Replication & Failover
      - Rate Limiting
      - Monitoring & alerting


**Step 7** Talk about Trade-offs:
      - Acknowledge trade-offs(e.g., SQL vs NoSQL, consistency vs availability)
      - Justify choices
      - Consider alternatives
      
💬 Example: "I chose eventual consistency here because the system is read-heavy and can tolerate slightly stale data."


**Step 8**✅ 8. Prepare for Follow-Up Questions
Interviewers often probe:
      - “What happens under failure?”
      - “How would you scale this for 10x traffic?”
      - “What are the bottlenecks?”
      - Stay calm and walk through it logically.
