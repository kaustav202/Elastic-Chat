# Elastic-Chat

A distributed messaging system that can scale at ease, maintaining ultra-low latency and high throughput with many concurrent users. Highly robust architecture that can withstand large sudden fluctuations in traffic.

### Salient Features

- Supports notifications, group chats, file sharing, last seen and online status.
- Distributed architecture enables easy horizontal scaling, both across servers and virtually with multiple instances. New services provisioned dynamically to handle increased requests resulting in ultra low latency.
- Back-end services written using GO, enabling high concurrent processing for handling increased workload. Database used is Cassandra nosql db, which suits flexible data type and is distributed with easy scale-out achieving high throughput. The backend services are configured using layer 7 load balancer - nginx.
- Redis Pub/Sub model is used to route messages to the connected node.  Each instance is subscribed to a particular channel in Redis and gets notified on receiving messages.
- The frontend for application is build using React.js and served through Nginx container. All the services are containarized using Docker and configured using Docker-Compose.
- Terraform is used for provisioning infrastructure on AWS, CI/CD pipeline is implemented using github actions along with AWS codedeploy.
