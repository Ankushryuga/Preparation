# 1. Cassandra Introduction:
  Its a highly scalable, high-performance distributed database designed to handle large amounts of data across many commodity servers, providing high availability with no single point of failure.
  Its a NoSQL database.

# 2. NoSQL Database:
  It provides a mechanism to store and retrieve data other than the tabular relations used in relational databases. These databases are schema-free, support easy replication.
  **Primary objective of a NoSQL database:**
  - simplicitly of design
  - horizontal scaling
  - finer control over availability

# 3. Cassandra Architecture:
  Cassandra has peer-to-peer distributed system across its nodes, and data is distributed among all the nodes in a cluster.
    - All the nodes in a cluster play the same role, each node is independent and at the same time interconnected to other nodes.
    - Each node in a cluster can accept read and write requests, regardless of where the data is actually located in the cluster.
    - When a node goes down, read/write requests can be served from other nodes in the network.


# 4. Cassandra Data Replication:
  In Cassandra, one or more of the nodes in a cluster act as replicas for a given piece of data. If it is detected that some of the nodes responded with an out-of-date value, Cassandra will return the most recent value to the client. After returning the most recent value, Cassandra performs a read repair in the background to update the stale values.
  The following figure shows a schematic view of how Cassandra uses data replication among the nodes in a cluster to ensure no single point of failure.
  <img width="390" height="409" alt="image" src="https://github.com/user-attachments/assets/10ddf669-d956-4f40-a2f1-c496cc56ec4f" />
