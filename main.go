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
  	r.Run()
}

func Create(c *gin.Context){
	client := redis.NewClient(&redis.Options{
				Addr:     "localhost:6379", // Default Redis address
				Password: "",              // No password by default
				DB:       0,               // Default Redis database
			})
	ctx := context.Background()
	errs := client.Set(ctx, "key", "value 1", 0).Err()
	if errs != nil {
		fmt.Println("Error setting key:", errs)
	} else {
		fmt.Println("Key set successfully")
	}
}