package main

import (
	"math/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handleKill(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n", r.Header)

	_, err := io.Copy(os.Stdout, r.Body)
	if err != nil {
		log.Println(err)
		return
	}

  c, err := inClusterLogin()
  if err != nil {
    panic(err.Error())
  }

  switch rand.Intn(killOptions) {
  case 0:
    killRandomPod(c)
  case 1:
    //killRandomDeployment(c)
  case 2:
    //killRandomStatefulSet(c)
  case 3:
    //updateRandomDeployement(c)
  case 4:
    //updateRandomStatefulSet(c)
  case 5:
    //killRandomWorker(c)
  }
}

func main() {
	log.Println("server started")
	http.HandleFunc("/kill", handleKill)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
