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

**Note** − Cassandra uses the Gossip Protocol in the background to allow the nodes to communicate with each other and detect any faulty nodes in the cluster.

# 5. Components of Cassandra:
  Key components of cassandra are as follows:-
  - **Node** -> It is the place where data is stored.
  - **Data center** -> It is a collection of related nodes.
  - **Cluster** -> A cluster is a component that contains one or more data centers.
  - **Commit log** -> The commit log is crash-recovery mechanism in cassandra, Every write opereation is written to the commit log.
  - **Mem-Table** -> A mem-table is a memory-resident data structure, After commit log, the data will be written to the mem-table. Sometimes, for a single-column family, there will be multiple mem-tables.
  - **SSTable** − It is a disk file to which the data is flushed from the mem-table when its contents reach a threshold value.
 - **Bloom filter** − These are nothing but quick, nondeterministic, algorithms for testing whether an element is a member of a set. It is a special kind of cache. Bloom filters are accessed after every query.


# 6. Cassandra Query Language:
Users can access Cassandra through its nodes using CQL(Cassandra Query Language). CQL treats the database (**Keyspace**) as a container of tables, **cqlsh**: a prompt to work with CQL.

# Write Operations: 
Every write activity of nodes is captured by the **commit logs** written in the nodes. Later the data will be captured and stored in the **mem-table**. Whenever the mem-table is full, data will be written into the **SStable** data file. All writes are automatically partitioned and replicated throughout the cluster. Cassandra periodically consolidates the SSTables, discarding unnecessary data.

# Read Operations
During read operations, Cassandra gets values from the mem-table and checks the bloom filter to find the appropriate SSTable that holds the required data.

