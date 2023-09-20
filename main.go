package main


import ( 
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()
	r.GET("/create-redis", Create)
	r.GET("/get-redis", Get)
	r.GET("/delete-redis", Delete)
  	r.Run()
}

func RedisConnection()(*redis.Client, context.Context){
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Default Redis address
		Password: "",              // No password by default
		DB:       0,               // Default Redis database
	})

	ctx := context.Background()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err) 
	}

	return client, ctx
}

func Create(c *gin.Context){
	client, ctx := RedisConnection()
	 
	errs := client.Set(ctx, "key", "value 22", 0).Err()
	if errs != nil {
		fmt.Println("Error setting key:", errs)
	} else {
		fmt.Println("Key set successfully")
	}
	c.JSON(200, "redis: data inserted successfully")
	return 
}

func Get(c *gin.Context){
	client, ctx := RedisConnection()
	value, err := client.Get(ctx, "key").Result()
	if err != nil {
		fmt.Println("Error getting key:", err)
	} else {
		fmt.Println("Value:", value)
	}
	c.JSON(200, "redis: get data 'key': "+ value)
	return 
}

func Delete(c *gin.Context){
	client, ctx := RedisConnection()
	err := client.Del(ctx, "key").Err()
	if err != nil {
		fmt.Println("Error deleting key:", err)
	} else {
		fmt.Println("Key deleted")
	}

	c.JSON(200, "redis: data deleted !!")
	return 
}