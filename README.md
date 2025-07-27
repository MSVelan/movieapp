
# Table of Contents

1.  [Services implemented:](#orgad719e3)
    1.  [Movie metadata service](#org6d456d9)
    2.  [Rating service](#org7edbb22)
    3.  [Movie service](#orgcffc0a8)
2.  [Other Details](#org19319e7)

This repository consists of three microservices: movie, metadata and rating. This is a practice project that I used to learn about microservices.


<a id="orgad719e3"></a>

# Services implemented:


<a id="org6d456d9"></a>

## Movie metadata service

-   API: Get metadata for a movie
-   Database: Movie metadata database
-   Interacts with services: None
-   Data model type: Movie metadata


<a id="org7edbb22"></a>

## Rating service

-   API: Get the aggregated rating for a record and write a rating
-   Database: Rating database
-   Interacts with service: None
-   Data model type: Rating


<a id="orgcffc0a8"></a>

## Movie service

-   API: Get movie details, including aggregated ratings and movie metadata
-   Database: None
-   Interacts with service: Movie metadata and rating
-   Data model type: Movie details
    
    ![img](./img/services.png)


<a id="org19319e7"></a>

# Other Details

Implemented service discovery with HashiCorp/Consul, implemented synchronous communication using HTTP API endpoints as well as gRPC endpoints.
For more details on how I implemented this, please check out [notes.md](./notes.md).
