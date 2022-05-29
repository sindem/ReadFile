package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/readfile", fileProcessingHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func invokeWordCount(text string, w http.ResponseWriter) {
	fmt.Println("Inside invokeWordCount ")
	client := &http.Client{}
	serviceUrl := "http://localhost:8000/wordcounts"

	filedata := url.Values{}
	filedata.Set("textstring", text)
	encodedData := filedata.Encode()
	fmt.Println(encodedData)

	req, err := http.NewRequest("POST", serviceUrl, strings.NewReader(encodedData))
	if err != nil {
		fmt.Println("Error encountered ", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error encountered making client request ", err)
	} else {
		var wordcount []WordCount
		data, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(data, &wordcount)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(wordcount)
	}
}

func fileProcessingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Sorry,  GET request processing is not supported for this endpoint")

	case "POST":
		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(10 << 20)

		file, _, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		invokeWordCount(string(fileBytes), w)

	default:
		fmt.Fprintf(w, "Sorry, only POST methods is supported.")

	}
}

type WordCount struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func (wc WordCount) String() string {
	return fmt.Sprintf("{\n word:	%s"+",\n count:"+"	%d \n}", wc.Word, wc.Count)
}
