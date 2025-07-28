
# Table of Contents

1.  [Services implemented:](#org70f6d56)
    1.  [Movie metadata service](#orgd33f4a8)
    2.  [Rating service](#org72bd92c)
    3.  [Movie service](#orgb0543b4)
2.  [Implementations](#orgf8f60c4)
3.  [Setup](#org0bcf71a)
    1.  [Hashicorp/Consul Setup](#orgac065ef)
    2.  [Kafka Setup](#orgf813f66)
        1.  [Docker](#org5f358aa)
        2.  [Run locally:](#orgd7a9d54)
    3.  [MySQL setup:](#org368adae)
        1.  [Docker](#org2fdb513)
        2.  [Local setup](#orge881fb6)
    4.  [Run the services](#org0b03544)

This repository consists of three microservices: movie, metadata and rating. This is a practice project that I used to learn about microservices.


<a id="org70f6d56"></a>

# Services implemented:


<a id="orgd33f4a8"></a>

## Movie metadata service

-   API: Get metadata for a movie and Put metadata for a movie
-   Database: Movie metadata database (in-memory MySQL implementation)
-   Interacts with services: None
-   Data model type: Movie metadata


<a id="org72bd92c"></a>

## Rating service

-   API: Get the aggregated rating for a record and write a rating
-   Database: Rating database (in-memory and MySQL implementation)
-   Interacts with service: None
-   Data model type: Rating


<a id="orgb0543b4"></a>

## Movie service

-   API: Get movie details, including aggregated ratings and movie metadata
-   Database: None
-   Interacts with service: Movie metadata and rating
-   Data model type: Movie details
    
    ![img](./img/services.png)


<a id="orgf8f60c4"></a>

# Implementations

-   service discovery with HashiCorp/Consul
-   synchronous communication using HTTP API endpoints as well as gRPC endpoints.
-   asynchronous communication using Apache Kafka
-   in-memory data stores and MySQL for the microservices

For more details on how I implemented this, please check out [notes.md](./notes.md).


<a id="org0bcf71a"></a>

# Setup


<a id="orgac065ef"></a>

## Hashicorp/Consul Setup

    docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul hashicorp/consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0


<a id="orgf813f66"></a>

## Kafka Setup


<a id="org5f358aa"></a>

### Docker

Get the Docker image:

    docker pull apache/kafka:4.0.0

Start the Kafka Docker container:

    docker run -p 9092:9092 apache/kafka:4.0.0


<a id="orgd7a9d54"></a>

### Run locally:

Generate a Cluster UUID

    KAFKA_CLUSTER_ID="$(bin/kafka-storage.sh random-uuid)"

Format Log Directories

    bin/kafka-storage.sh format --standalone -t $KAFKA_CLUSTER_ID -c config/server.properties

Start the Kafka Server

    bin/kafka-server-start.sh config/server.properties


<a id="org368adae"></a>

## MySQL setup:


<a id="org2fdb513"></a>

### Docker

    docker run --name movieapp_db -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=movieapp -p 3306:3306 -d mysql:latest

    docker exec -i movieapp_db mysql movieapp -h localhost -P 3306 --protocol=tcp -uroot -ppassword < schema/schema.sql

Create .env file in pkg/mysql/ directory and add password:

    PASSWORD=password


<a id="orge881fb6"></a>

### Local setup

Create movieapp database making sure that mysql is running in port 3306 (default port).

Create schema using:

    mysql -u root -p movieapp < schema/schema.sql

Create `.env` file in `pkg/mysql/` directory and add password:

    PASSWORD=<your-root-password>


<a id="org0b03544"></a>

## Run the services

    go run metadata/cmd/main.go
    go run rating/cmd/main.go
    go run movie/cmd/main.go

