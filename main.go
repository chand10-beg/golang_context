package main

import (
	"context"
	"fmt"
	"time"

	"github.com/icrowley/fake"
)

func main() {

	ctx := context.Background()
	ctx2 := context.WithValue(ctx, "main", "2")
	conte := fake.Password(3, 3, true, false, false)
	// fmt.Println(RandStringBytes(3))
	fmt.Println(conte)
	ctx3, ctx4 := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	ctx4()
	select {
	case <-ctx3.Done():
		fmt.Println("ct done")
	case <-time.Tick(2 * time.Second):
		fmt.Println("10 sec")
	}
	fmt.Println(ctx2.Value("main"))
	fmt.Println(ctx.Value("main"))
}

// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex
// var c, tQ int

// func main() {

// 	file, err := os.Open("ques.csv")
// 	// newtimer := time.NewTimer(time.Second)
// 	// stoptimer := newtimer.Stop()
// 	// Checks for the error

// 	if err != nil {
// 		log.Fatal("Error while reading the file", err)
// 	}

// 	// Closes the file
// 	defer file.Close()
// 	reader := csv.NewReader(file)

// 	records, err := reader.ReadAll()

// 	if err != nil {
// 		fmt.Println("Error reading records")
// 	}
// 	stop := time.After(5 * time.Second)
// 	// for _, eachrecord := range records {
// 	// 	if len(eachrecord) > 0 {
// 	// 		tQ+=1
// 	// 	}
// 	// }
// 	for _, eachrecord := range records {
// 		fmt.Printf("Hello %v,  please answer\n", eachrecord[0])
// 		cha := make(chan int)
// 		go func() {
// 			var ans int
// 			fmt.Scanf("%s\n", &ans)
// 			cha <- ans
// 		}()
// 		select {
// 		case <-stop:
// 			fmt.Println("EXIT: 5 seconds")
// 			return
// 		case ans := <-cha:

// 			if len(eachrecord) < 1 {
// 				fmt.Printf("No more questions!!!!!")
// 				break
// 			}
// 			num, err := strconv.Atoi(eachrecord[1])
// 			if err != nil {
// 				fmt.Printf("error %v", err)
// 			}

// 			if ans == num {
// 				fmt.Println("answer is correct!")
// 				c += 1
// 				tQ += 1
// 			} else {
// 				fmt.Println("you're wrong!")
// 				tQ += 1
// 			}
// 		}
// 		fmt.Println(tQ, c)

// 	}

// }
