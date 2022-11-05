# pull Image or run docker compose up
docker pull bitnami/zookeeper
docker pull bitnami/kafka




# command kafka that I learnt
- Check Kafka is available
```
kafka-topics --bootstrap-server=localhost:9092 --list
```
-  Create new topic
```
kafka-topics --bootstrap-server=localhost:9092 --topic=pekhello --create
```
- Create consumer
```
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=pekhello
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=pekhello --group=myconsumer
kafka-console-consumer --bootstrap-server=localhost:9092 --include="pekhello|testhello" --group=pek
```
- Create Producer
```
kafka-console-producer --bootstrap-server=localhost:9092 --topic=pekhello
kafka-console-producer --bootstrap-server=localhost:9092 --topic=testhello
kafka-console-producer --bootstrap-server=localhost:9092 --topic=OpenAccountEvent
kafka-console-producer --bootstrap-server=localhost:9092 --topic=DepositFundEvent
kafka-console-producer --bootstrap-server=localhost:9092 --topic=WithdrawFundEvent
kafka-console-producer --bootstrap-server=localhost:9092 --topic=CloseAccountEvent
```

# Json for test Call Producer
OpenAccountEvent
{"ID":"abc123","AccountHolder":"Pek","AccountType":1,"OpeningBalance": 1000}

DepositFundEvent
{"ID":"abc123","Amount":400}

WithdrawFundEvent
{"ID":"abc123","Amount":10}

CloseAccountEvent
{"ID":"abc123"}

# Postmanc Collection for test Producer and Consumer 
- Import go_kafka_poc/consumer/postman/ktb_kafka.postman_collection.json to Postman





