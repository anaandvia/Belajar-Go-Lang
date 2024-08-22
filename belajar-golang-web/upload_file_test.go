package belajargolangweb

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "uploadform.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// ambil file
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	// create destinasi file
	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)

	if err != nil {
		panic(err)
	}
	// taruh file di destinasi
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload_success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})

}
func TestUploadFormServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// go::embed resources/download.jpeg
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Ana")
	file, _ := writer.CreateFormFile("file", "contoh.jpeg")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8082/upload", body)

	request.Header.Set("Content-type", writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	response := recorder.Result()
	bodyResponse, _ := io.ReadAll(response.Body)
	fmt.Println(string(bodyResponse))
}
