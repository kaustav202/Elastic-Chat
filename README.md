# Elastic-Chat

A distributed messaging system that can scale at ease, maintaining ultra-low latency and high throughput with many concurrent users. Highly robust architecture that can withstand large sudden fluctuations in traffic.

### Salient Features

- Supports notifications, group chats, file sharing, last seen and online status.
- Distributed architecture enables easy horizontal scaling, both across servers and virtually with multiple instances. New services provisioned dynamically to handle increased requests resulting in ultra low latency.
- Back-end services written using GO, enabling high concurrent processing for handling increased workload. Database used is Cassandra nosql db, which suits flexible data type and is distributed with easy scale-out achieving high throughput. The backend services are configured using layer 7 load balancer - nginx.
- Redis Pub/Sub model is used to route messages to the connected node.  Each instance is subscribed to a particular channel in Redis and gets notified on receiving messages.
- The frontend for application is build using React.js and served through Nginx container. All the services are containarized using Docker and configured using Docker-Compose.
- Terraform is used for provisioning infrastructure on AWS, CI/CD pipeline is implemented using github actions along with AWS codedeploy.


<br>

# Architecture

![Untitled Diagram-Page-1 drawio](https://github.com/kaustav202/Elastic-Chat/assets/89788120/d2137ebd-f285-4742-b67a-f3a9fb878a98)

-  The Websocket protocol is chosen for real-time full duplex communication for messaging between two users or in groups. Simple REST APIs handle all other functionalities such as history, auth.
-  Async read/writes to the database utilizing Kafka message queues handle heavy write load to the database without affecting performance ( decoupled ). A consumer Go instance writes the messages in batches.
-  Redis cache is used to store websocket connection information and what node the target user is residing on, to route the message to that node without needing DB lookup. Redis is also used to retrieve groups and users corresponding to those, which is used to send group messages.
-  Service discovery like envoy proxy ( alternatively service mesh ) is used for inter service communication.
