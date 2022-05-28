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
```    go
var ctx = context.Background()
_, err := c.Set(ctx, &pb.CacheSetInput{Key: "hello", Value: []byte("123")})
if err != nil {
	panic(err)
}
```
#### 2. Get Cache
```    go
var ctx = context.Background()
resp, err := c.Get(ctx, &pb.CacheGetInput{Key: "hello"})
if err != nil {
	panic(err)
}
fmt.Println(resp)
```
#### 3. Set User Cache
```    go
_, err = c.SetUser(ctx, &pb.User{
	Name:     "john",
	Class:    "V,
	Rollnum:  22,
	Metadata: []byte("smart student"),
})
if err != nil {
	panic(err)
}
```
#### 4. Get User Cache
```    go

user, err := c.GetUser(ctx, &pb.GetUserInput{
		Name:    name,
		Class:   class,
		Rollnum: rollInt,
})
if err != nil {
	panic(err)
}
fmt.Println(user)
```

### Programming Interface