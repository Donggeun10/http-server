package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	crashReportFlatbuffers "example.com/crashReportFlatbuffers"
	"example.com/greetings"
	"github.com/gorilla/mux"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("Gladys")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin", "Avicii"}
	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	router := mux.NewRouter()

	router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	}).Methods("GET")

	router.Path("/init").HeadersRegexp("Content-Type", "application/flatbuffers-(v3|v4)").HandlerFunc(readAppload).Methods("POST")

	http.ListenAndServe(":9084", router)
}

func readAppload(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err == nil {

		contentType := r.Header.Get("Content-Type")
		if contentType == "application/flatbuffers-v4" {
			AppLoadV4Log := crashReportFlatbuffers.GetRootAsAppLoadLogV4(body, 0)
			apploadV4 := crashReportFlatbuffers.AppLoadV4FromFB(AppLoadV4Log)
			// initMessage = crashReportFlatbuffers.InitMessageFromApploadObj(apploadV4.GameCode, apploadV4.Os, apploadV4, "v4")
			fmt.Println("GameCode:", apploadV4.GameCode, ", Os:", apploadV4.Os, ", OsVersion:", apploadV4.OsVersion)
		} else if contentType == "application/flatbuffers-v3" {
			AppLoadV3Log := crashReportFlatbuffers.GetRootAsAppLoadLogV3(body, 0)
			apploadV3 := crashReportFlatbuffers.AppLoadV3FromFB(AppLoadV3Log)
			initMessage := crashReportFlatbuffers.InitMessageFromApploadV3(apploadV3.GameCode, apploadV3.Os, apploadV3, "v3")
			fmt.Println("GameCode:", initMessage.GameCode, ", Os:", initMessage.Os, ", Log:", initMessage.Log, ", LogVersion:", initMessage.LogVersion)
		} else {
			fmt.Println("Unknown content type : ", contentType)
			fmt.Fprintf(w, "You've requested the book: %s on page %s\n", "title", "page")
		}

	} else {
		fmt.Println("Error reading body")
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", "title", "page")
	}

}
