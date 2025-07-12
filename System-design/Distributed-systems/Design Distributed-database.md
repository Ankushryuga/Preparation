## Problem Statement:
Design a distributed database system that can store massive volumes of structured or semi-structured data 
across multiple machines, with **high availability**, **patition-tolerance**, and **horizontal scalability**.

## High Level Architecture:
                        +------------------------+
                        |   Client Applications  |
                        +-----------+------------+
                                    |
                           +--------v--------+
                           |   Query Router  |
                           +--------+--------+
                                    |
               +--------------------+--------------------+
               |                    |                    |
            +--v--+            +----v----+           +---v---+
            |Node1|            | Node2   |           | Node3 |
            +--+--+            +----+----+           +---+---+
               |                   |                    |
            +--v--+            +---v----+           +---v---+
            |Data |            | Data   |           | Data  |
            | +   |            | +      |           | +     |
            |Meta |            | Meta   |           | Meta  |
            +-----+            +--------+           +-------+


## Key Components:
      | Component                      | Purpose                                                                       |
      | ------------------------------ | ----------------------------------------------------------------------------- |
      | **Client**                     | Application or SDK that sends queries (CRUD, analytics, etc.).                |
      | **Query Router / Coordinator** | Forwards requests to the appropriate node(s) based on partition key or query. |
      | **Storage Nodes**              | Store data partitions. Each node may replicate others.                        |
      | **Metadata Service**           | Keeps track of schema, partitions, leader nodes, etc.                         |
      | **Consensus Engine**           | Manages leader election, replication, cluster membership (Raft or Paxos).     |


## Data Partitioning(Sharding)
  - Goal Split large dataset across nodes.
  - Why? One machine cannot hold or process all data.
  - How?
    - 1. Hash-based partitioning
         - sharId = hash(partitionKey) % num_shards
    - 2. Range-based partitioning
         - Example: store data with keys A-M on node1, N-Z on node2.
  - Partitioning keys:
    - Choose a good partition key to avoid hotspots(e.g., userId, orderId).

## Replication Strategy:
Need for **high availability** and **fault tolerance**.
    - Replication Factor (RF): Number of copies of data(e.g., RF=3).
    - strategies:
      - Leader-follower (Master-Slave): Writes go to a leader, then sync to replicas.
      - Multi-leader: writes can go to any node -> conflicts possible.
      - Leaderless(Dynamo style): No central coordinator; consistency via quorum.

# Example (Quorum):
  - Write to W=2 of 3 nodes.
  - Read from R=2 of 3 nodes.
  - If R + W > RF ⇒ consistency possible.


# Consistency models:
    => Choose based on use case(CAP tradeoff):
    
    | Model    | Description                    | Use Case                  |
    | -------- | ------------------------------ | ------------------------- |
    | Strong   | Read returns latest write      | Banking, critical systems |
    | Eventual | Data converges over time       | Social media feed         |
    | Tunable  | Choose consistency per request | Cassandra style           |

    # Use Vector Clocks or Timestamps to detect conflicts.

## Fault Tolerance:
    - Node Failure: Survive if 1 or more nodes crash.
    - Network Partition: Survice 
