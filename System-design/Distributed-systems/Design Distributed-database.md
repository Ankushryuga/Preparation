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
