## What is Distributed System?
   A **Distributed System** is a collection of independent computers (nodes) that work together to appear as a **single coherent system** to the user.

## Each Node:
    1. Has its own memory and processor.
    2. Communicate over a **network**.
    3. May fail independently.
    4. Shares a common goal/workload.

## Why Use Distributed Systems?
    | Need                        | Explanation                                              |
    | --------------------------- | -------------------------------------------------------- |
    | **Scalability**             | Add more machines to handle more data/work.              |
    | **Fault Tolerance**         | If one node fails, the system still works.               |
    | **Availability**            | The system continues to operate (partial failure OK).    |
    | **Performance**             | Tasks run in parallel = faster results.                  |
    | **Geographic Distribution** | Nodes can be close to users for low latency (e.g., CDN). |

## Core characteristics of Distributed Systems.
    | Property                        | Explanation                                                                    |
    | ------------------------------- | ------------------------------------------------------------------------------ |
    | **Transparency**                | Users don't know it's distributed (location, replication, failure are hidden). |
    | **Fault Tolerance**             | Can continue operating despite node failures.                                  |
    | **Concurrency**                 | Multiple components execute simultaneously.                                    |
    | **Scalability**                 | Can grow by adding more nodes.                                                 |
    | **Resource Sharing**            | Nodes share files, data, printers, etc.                                        |
    | **Consistency vs Availability** | Often tradeoffs (→ see CAP theorem).                                           |

## CAP Theorem(Brewer's Theorem)
    In a distributed system, **you can only guarantee 2 out of 3.**
    | CAP Element                 | Description                                            |
    | --------------------------- | ------------------------------------------------------ |
    | **C** – Consistency         | Every node sees the same data at the same time.        |
    | **A** – Availability        | Every request receives a (non-error) response.         |
    | **P** – Partition Tolerance | System continues to function even if nodes/links fail. |

## Example:
  - MongoDB (Configurable): Can be CP or AP depending on settings.
  - Cassandra: AP — prioritizes availability and partition tolerance.


## Components of Distributed systems:
    | Component               | Role                                                                     |
    | ----------------------- | ------------------------------------------------------------------------ |
    | **Nodes**               | Independent servers or machines.                                         |
    | **Network**             | Medium for communication (TCP/IP, HTTP, gRPC).                           |
    | **Coordinator/Leader**  | Optional. For coordinating tasks (e.g., master in master-slave systems). |
    | **Data Storage**        | Often replicated and/or sharded.                                         |
    | **Consensus Protocols** | Algorithms for consistency (e.g., Raft, Paxos, ZAB).                     |
    | **Middleware**          | Connects distributed apps and provides abstraction.                      |


## Communication in Distributed Systems:
    | Method                          | Description                                                          |
    | ------------------------------- | -------------------------------------------------------------------- |
    | **Remote Procedure Call (RPC)** | Call a function on another node (e.g., gRPC).                        |
    | **Message Passing**             | Send/receive messages (e.g., Kafka, RabbitMQ).                       |
    | **Shared State**                | Some systems (e.g., distributed databases) share state across nodes. |

## Types of Distributed Systems:
    | Type                               | Example                         | Purpose                        |
    | ---------------------------------- | ------------------------------- | ------------------------------ |
    | **Distributed Databases**          | Cassandra, MongoDB, CockroachDB | Data replication, partitioning |
    | **Distributed File Systems**       | HDFS, GFS, Ceph                 | Store large files across nodes |
    | **Cluster Computing**              | Apache Spark, Kubernetes        | Parallel computation           |
    | **Content Delivery Network (CDN)** | Cloudflare, Akamai              | Deliver content close to users |
    | **Blockchain Systems**             | Bitcoin, Ethereum               | Distributed ledger, consensus  |

## Common Problems in Distributed Systems:
    | Problem                   | Explanation                                         |
    | ------------------------- | --------------------------------------------------- |
    | **Network Partitions**    | Nodes can't talk to each other.                     |
    | **Clock Skew**            | Each node has its own clock (→ time is unreliable). |
    | **Message Loss/Delay**    | Packets can be dropped or delayed.                  |
    | **Leader Election**       | Choosing a coordinator in a fault-tolerant way.     |
    | **Split Brain**           | Two partitions think they are both the leader.      |
    | **Consistency Models**    | Strong, Eventual, Causal, etc.                      |
    | **Concurrency/Conflicts** | Multiple writes at once — needs resolution.         |

## Consistency Models 
    | Model                      | Description                                    | Example Use         |
    | -------------------------- | ---------------------------------------------- | ------------------- |
    | **Strong Consistency**     | Read always returns the latest write.          | Banking systems     |
    | **Eventual Consistency**   | Over time, all nodes converge.                 | Cassandra, DynamoDB |
    | **Causal Consistency**     | Preserves causality of writes.                 | Collaborative apps  |
    | **Read-Your-Writes**       | After writing, you can always read your value. | Sessions            |
    | **Monotonic Reads/Writes** | Read/write history progresses forward only.    | Versioned systems   |

## Fault Tolerance Mechanisms:
    | Mechanism           | Description                                   |
    | ------------------- | --------------------------------------------- |
    | **Replication**     | Store copies of data on multiple nodes.       |
    | **Heartbeats**      | Nodes ping each other to check liveness.      |
    | **Leader Election** | Protocols like Raft ensure only one leader.   |
    | **Retry Logic**     | Clients retry failed operations with backoff. |
    | **Quorums**         | Use majority votes to commit actions.         |


## Concensus Protocols: 
     This algorithm used in distributed systems to ensure that multiple computers(nodes) agree on a single version of truth,
     even if some nodes fail or act malicisouly. 
    | Protocol  | Use Case                                         | Notes |
    | --------- | ------------------------------------------------ | ----- |
    | **Paxos** | Theoretical base; difficult to implement         |       |
    | **Raft**  | Easier consensus (used in etcd, Consul)          |       |
    | **ZAB**   | Used in Apache ZooKeeper                         |       |
    | **PBFT**  | Practical Byzantine Fault Tolerance (blockchain) |       |


## Real-World Examples:
    | System            | Type                  | Distributed Characteristics       |
    | ----------------- | --------------------- | --------------------------------- |
    | **Apache Kafka**  | Event streaming       | Partitioned logs, replication     |
    | **MongoDB**       | Document DB           | Sharding, replica sets            |
    | **Cassandra**     | Wide-column DB        | Peer-to-peer, tunable consistency |
    | **Kubernetes**    | Cluster orchestration | Leader election, fault tolerance  |
    | **Hadoop**        | File system + compute | HDFS + MapReduce, data locality   |
    | **Redis Cluster** | Key-value store       | Partitioning + Replication        |

## 🧠 Interview-Level Summary:
## Q1. What makes distributed systesm hard?
    - Partial Failures
    - Non-determinism
    - Clock synchronization issues.
    - Concurrency
    - Trade-offs b/w CAP

## Q2. How do you ensure Consistency?
  - Through **replication protocols**, and **consistency models**.

## Q3. What's the trade-off b/w availability and consistency?
  - High availability systems like cassandra sacrifice strict consistency for uptime.

## Practice:
  - Real-world failure scenarios
  - Leader election algorithms (Raft, Paxos)
  - Design a distributed database
  - Interview question answers from top companies
