# cache-service

## Setup
1. Clone the Rep
2. Setup go
3. Setup Redis
4. Start Redis server <br> ```chmod +x initialize_redis.sh && ./initialize_redis.sh```
5. build server <br> ```go build -o ./build/rpcServer ./server/handlers/*.go```
6. build client <br> ```go build -o ./build/rpcClient ./client/*.go```

## Usage
### Console

### Programming Interface