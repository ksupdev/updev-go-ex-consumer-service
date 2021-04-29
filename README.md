# updev-go-ex-consumer-service
Demo go service connect kafka

## setup project
1. clone project ``git clone https://github.com/ksupdev/updev-go-ex-consumer-service.git``
2. go mod init ``go mod init github.com/ksupdev/updev-go-ex-consumer-service``
```powershell
% go mod init github.com/ksupdev/updev-go-ex-consumer-service
go: creating new go.mod: module github.com/ksupdev/updev-go-ex-consumer-service
% go run main.go
Hi%   
```

## Labs Noted
### setup KAFKA
1. create docker compose file for create zookeeper, kafka 
2. Run docker (zookeeper, kafka)
    - run docker compose
    ```powershell
    docker-compose up -d
    ```
    - Check container
    ```powershell
        % docker ps
        CONTAINER ID   IMAGE                             COMMAND                  CREATED          STATUS                          PORTS                                        NAMES
        f0ecd9ec0831   3dsinteractive/kafka:2.0-custom   "/app-entrypoint.sh …"   43 seconds ago   Up 40 seconds                   9092/tcp, 0.0.0.0:9094->9094/tcp             updev-go-ex-consumer-service_kafka_1
        bc585e42ba0f   3dsinteractive/zookeeper:3.0      "/app-entrypoint.sh …"   43 seconds ago   Up 42 seconds                   2888/tcp, 0.0.0.0:2181->2181/tcp, 3888/tcp   updev-go-ex-consumer-service_zookeeper_1
    ```
### setup go project