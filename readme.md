## Instructions
### Clone the repository using git clone --- 
### Run locally docker compose up &&  go run . to start up the server and connect to the database

```go
func doWork(d time.Duration, s string, myChan chan string) {
	fmt.Println("doing work......")
	time.Sleep(d)
	fmt.Println("work is done", s)
	myChan <- s
	wg.Done()
}

var wg *sync.WaitGroup

func main() {
	start := time.Now()
	resChannel := make(chan string)
	wg = &sync.WaitGroup{}

	wg.Add(3)
	go doWork(time.Second*2, "Hello 1", resChannel)
	go doWork(time.Second*4, "Hello 2", resChannel)
	go doWork(time.Second*4, "Hello 3", resChannel)

	strings := make(map[int]string)

	go func() {
		iteration := 1
		for res := range resChannel {
			strings[iteration] = res
			iteration++
		}
		fmt.Println("it took", time.Since(start))
	}()

	wg.Wait()
	close(resChannel)
	time.Sleep(time.Second)
	fmt.Println(strings)
	// server := gin.Default()
	// database.ConnectDatabase()
	// routes.RegisterRoutes(server)
	// server.Run(":8080")
}


```

#### For reference this is how we handle concurrent data in Golang