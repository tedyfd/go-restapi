package main
import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main(){
	fmt.Println("hello")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("Port not found")
	}
	fmt.Println("Port:", portString)

}