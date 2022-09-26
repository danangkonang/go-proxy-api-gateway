# go-proxy-api-gateway

Simple proxy api gateway write with golang


```text
     +--------+        +--------------+        +----------+
     |        |        |              | -----> |          |
     |        |        |              |        | Server 1 |
     |        | -----> |              | <----- |          |
     |        |        |   Go Proxy   |        +----------+
     | Client |        |  Api Gateway |
     |        |        |              |        +----------+
     |        | <----- |              | -----> |          |
     |        |        |              |        | Server 2 |
     |        |        |              | <----- |          |
     +--------+        +--------------+        +----------+
```
