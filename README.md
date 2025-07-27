
# Table of Contents

1.  [Services implemented:](#org2c420e0)
    1.  [Movie metadata service](#org054d86e)
    2.  [Rating service](#org9670ce7)
    3.  [Movie service](#org404af45)
2.  [Other Details](#org52a8ed9)
3.  [Setup](#org353a65c)
    1.  [Hashicorp/Consul Setup](#org70c7fe0)
    2.  [Kafka Setup](#org85b71cb)
        1.  [Docker](#orgb9fb8e9)
        2.  [Run locally:](#org64e36c7)
    3.  [Run the services](#org19a4b6c)

This repository consists of three microservices: movie, metadata and rating. This is a practice project that I used to learn about microservices.


<a id="org2c420e0"></a>

# Services implemented:


<a id="org054d86e"></a>

## Movie metadata service

-   API: Get metadata for a movie
-   Database: Movie metadata database
-   Interacts with services: None
-   Data model type: Movie metadata


<a id="org9670ce7"></a>

## Rating service

-   API: Get the aggregated rating for a record and write a rating
-   Database: Rating database
-   Interacts with service: None
-   Data model type: Rating


<a id="org404af45"></a>

## Movie service

-   API: Get movie details, including aggregated ratings and movie metadata
-   Database: None
-   Interacts with service: Movie metadata and rating
-   Data model type: Movie details
    
    ![img](./img/services.png)


<a id="org52a8ed9"></a>

# Other Details

Implemented service discovery with HashiCorp/Consul, implemented synchronous communication using HTTP API endpoints as well as gRPC endpoints.
For more details on how I implemented this, please check out [notes.md](./notes.md).


<a id="org353a65c"></a>

# Setup


<a id="org70c7fe0"></a>

## Hashicorp/Consul Setup

    docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul hashicorp/consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0


<a id="org85b71cb"></a>

## Kafka Setup


<a id="orgb9fb8e9"></a>

### Docker

-   Get the Docker image:

    docker pull apache/kafka:4.0.0

-   Start the Kafka Docker container:

    docker run -p 9092:9092 apache/kafka:4.0.0


<a id="org64e36c7"></a>

### Run locally:

-   Generate a Cluster UUID

    KAFKA_CLUSTER_ID="$(bin/kafka-storage.sh random-uuid)"

-   Format Log Directories

    bin/kafka-storage.sh format --standalone -t $KAFKA_CLUSTER_ID -c config/server.properties

-   Start the Kafka Server

    bin/kafka-server-start.sh config/server.properties


<a id="org19a4b6c"></a>

## Run the services

    go run metadata/cmd/main.go
    go run rating/cmd/main.go
    go run movie/cmd/main.go

