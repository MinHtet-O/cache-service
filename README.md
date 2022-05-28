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

#### 1. Set Cache
```    
ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Second)
defer cancel()
_, err := c.Set(ctx, &pb.CacheSetInput{Key: key, Value: []byte(value)})
if err != nil {
		fmt.Println(err.Error())
		return
}
```

	
#### 2. Get Cache
#### 3. Set User Cache
#### 4. Get User Cache

### Programming Interface