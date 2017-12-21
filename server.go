package GO

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func servePage(w http.ResponseWriter, r* http.Request){
	fmt.Printf("Rendered")
}

func main(){
	router := mux.newRouter().StrictSlash(true)
	router.HandleServ("/",servePage)
	log.Fatal(http.ListenAndServ(":8080",router))
}