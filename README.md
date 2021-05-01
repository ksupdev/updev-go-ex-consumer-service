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
1. implement context.go to be interface of Context
2. implement context_consumer.go :ซึ่งเป็น implementation ของ context.go ``ถ้าเราจำได้ service ที่เป็นก็จะมี context_http.go`` เหมือนกัน
3. implement microservice.go to be ......
    - ซึ่งเป็นส่วนที่ใช้ในการกำหนด function การทำงานของ microservice ที่เป็นรูปแบบ Consumer
    > Consumer : จะเหมาะกับการทำงานเบื้องหลัง ที่ต้องใช้เวลาการทำงานมากและจะ Response ผลลัพธ์เมื่อทำงานเสร็จแล้ว ``จะไม่ส่งผลลัพท์หาผู้เรียกโดยตรง`` อาจจะส่งเป็น Email,Message,หรือ mark flag ใน database เป้นต้น โดย Consumer จะทำงานเป็นรูปแบบ Async โดยปกติจะทำงานกับพวก Queue หรือก็คือไปอ่านข้อมูลจาก Database หรือ Message Queue (kafka)

    - implement ``newKafkaConsumer`` สำหรับใช้ในการสร้าง object ของ Kafka consumer
        ```powershell
            // install kafka lib
            % go get github.com/confluentinc/confluent-kafka-go/kafka
            go: found github.com/confluentinc/confluent-kafka-go/kafka in github.com/confluentinc/confluent-kafka-go v1.6.1
        ```
    - implement func ``Consume`` : จะเป็น public func ซึ่งใช้ในการ connect ไปที่ Kafka server โดยเราจะมีการกำหนดในส่วนของ kafka topic,group ด้วย แต่ func Consume จะไม่ได้ connect ไปที่ kafka แต่จะไปเรียกใช้ private func consumeSingle และจะเรียกใช้โดยใช้ ``Gorouetins`` ซึ่งจะทำให้เป็นการทำงานแบบ concurrently กับ การทำงานอื่นๆ

    - implement func ``consumeSingle`` : จะเป็น private func ซึ่งจะช่วยในการจัดการการเชื่อมต่อกับ Kafka ซึ่งก็คือจะมีการ Connect ไปยัง kafka server และ subscribe ไปที่ topic และ group ที่เราสนใจ และถ้ามีข้อมูลเกิดขึ้นจาก producer กับ topic ที่เรา subscribe ไว้ก็จะมีการ callback ไปที่ func ที่เราได้ทำการส่งเข้ามาผ่าน ServiceHandleFunc เพื่อทำงานต่อไป สำหรับระยะเวลาในการ subscribe เราสามารถกำหนด ``c.ReadMessage(readTimeout)`` และกำหนดให้ readTimeout = -1 เพื่อไม่ให้มี timeout หรือก็คืออ่านไปเลื่อยๆๆนั้นเอง

        ```golang
            msg, err := c.ReadMessage(readTimeout)
            if err != nil {
                kafkaErr, ok := err.(kafka.Error)
                if ok {
                    if kafkaErr.Code() == kafka.ErrTimedOut {
                        if readTimeout == -1 {
                            // No timeout just continue to read message again
                            continue
                        }
                    }
                }
                ms.Log("Consumer", err.Error())
                return
            }

            // Execute Handler
            h(NewConsumerContext(ms, string(msg.Value)))
        ```
    - implement func `Start` : จะทำการสร้ง chanel เพื่อใช้ในการรบรับคำสั่งที่จะถูกส่งมาจาก terminal และจะมีการ รอดูผลลัพท์จาก ``exitChannel`` ที่จะถูกส่งมาจาก goroutine เพื่อจะบอกว่าหยุดการทำงานเมื่อไร 
    ```golang
        [main.go]
        prod := NewProducer(servers, ms)
        go func() {
            for i := 0; i < 10; i++ {
                prod.SendMessage(topic, "", map[string]interface{}{"message_id": i})
                time.Sleep(time.Second)
            }

            // Exit program
            ms.Stop()
        }()
        ------
        [microservice.go]
        func (ms *Microservice) Stop() {
            if ms.exitChannel == nil {
                return
            }
            ms.exitChannel <- true
        }
    ```
    - implement func `Stop` : จะถูกเรียกเพื่อเป็นการกำหนดค่าให้ Channel

    
    