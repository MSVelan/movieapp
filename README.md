
# Table of Contents

1.  [Services implemented:](#orgbbb501b)
    1.  [Movie metadata service](#org59da5cd)
    2.  [Rating service](#org5d67754)
    3.  [Movie service](#org7c8bd09)
2.  [Implementations](#orgc84668e)
3.  [Setup](#orgf931096)
    1.  [Hashicorp/Consul Setup](#org82cdea9)
    2.  [Kafka Setup](#org4d077ba)
        1.  [Docker](#orgbd35618)
        2.  [Run locally:](#org39590ce)
    3.  [Run the services](#org5f14861)

This repository consists of three microservices: movie, metadata and rating. This is a practice project that I used to learn about microservices.


<a id="orgbbb501b"></a>

# Services implemented:


<a id="org59da5cd"></a>

## Movie metadata service

-   API: Get metadata for a movie and Put metadata for a movie
-   Database: Movie metadata database (in-memory implementation)
-   Interacts with services: None
-   Data model type: Movie metadata


<a id="org5d67754"></a>

## Rating service

-   API: Get the aggregated rating for a record and write a rating
-   Database: Rating database (in-memory implementation)
-   Interacts with service: None
-   Data model type: Rating


<a id="org7c8bd09"></a>

## Movie service

-   API: Get movie details, including aggregated ratings and movie metadata
-   Database: None
-   Interacts with service: Movie metadata and rating
-   Data model type: Movie details
    
    ![img](./img/services.png)


<a id="orgc84668e"></a>

# Implementations

-   service discovery with HashiCorp/Consul
-   synchronous communication using HTTP API endpoints as well as gRPC endpoints.
-   asynchronous communication using Apache Kafka
-   in-memory data stores for the microservices

For more details on how I implemented this, please check out [notes.md](./notes.md).


<a id="orgf931096"></a>

# Setup


<a id="org82cdea9"></a>

## Hashicorp/Consul Setup

    docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul hashicorp/consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0


<a id="org4d077ba"></a>

## Kafka Setup


<a id="orgbd35618"></a>

### Docker

Get the Docker image:

    docker pull apache/kafka:4.0.0

Start the Kafka Docker container:

    docker run -p 9092:9092 apache/kafka:4.0.0


<a id="org39590ce"></a>

### Run locally:

Generate a Cluster UUID

    KAFKA_CLUSTER_ID="$(bin/kafka-storage.sh random-uuid)"

Format Log Directories

    bin/kafka-storage.sh format --standalone -t $KAFKA_CLUSTER_ID -c config/server.properties

Start the Kafka Server

    bin/kafka-server-start.sh config/server.properties


<a id="org5f14861"></a>

## Run the services

    go run metadata/cmd/main.go
    go run rating/cmd/main.go
    go run movie/cmd/main.go

