package main

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "embed"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRoutingHelloWorld(t *testing.T) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	request := httptest.NewRequest("GET", "/", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "hello world", string(bytes))
}

func TestCtx(t *testing.T) {
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "Guest")
		return c.SendString("hello " + name)
	})

	request := httptest.NewRequest("GET", "/hello?name=ana", nil)
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "hello ana", string(bytes))
}

func TestHttpRequest(t *testing.T) {
	app := fiber.New()
	app.Get("/request", func(c *fiber.Ctx) error {
		first := c.Get("firstname")
		last := c.Cookies("lastname")
		return c.SendString("hello " + first + last)
	})

	request := httptest.NewRequest("GET", "/request", nil)

	request.Header.Set("firstname", "ana")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: " jannatu"})

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "hello ana jannatu", string(bytes))
}

func TestRouteParameter(t *testing.T) {
	app := fiber.New()
	app.Get("/users/:userId/orders/:orderId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		orderId := c.Params("orderId")
		return c.SendString("Get Order " + orderId + " from " + userId)
	})

	request := httptest.NewRequest("GET", "/users/1/orders/3", nil)

	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Get Order 3 from 1", string(bytes))
}

func TestFormRequest(t *testing.T) {
	app := fiber.New()
	app.Post("/hello", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return c.SendString("hello " + name)
	})

	body := strings.NewReader("name=Ana")
	request := httptest.NewRequest("POST", "/hello", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "hello Ana", string(bytes))
}

// go:embed source/hehe.txt
var contohFile []byte

func TestFormUploadRequest2(t *testing.T) {
	app := fiber.New()
	app.Post("/upload", func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		err = c.SaveFile(file, "./target/"+file.Filename)
		if err != nil {
			return err
		}
		return c.SendString("Upload Success")
	})

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	file, err := writer.CreateFormFile("file", "nama.txt")
	assert.Nil(t, err)
	file.Write(contohFile)
	writer.Close()

	request := httptest.NewRequest("POST", "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Upload Success", string(bytes))
}

type LoginRequest struct {
	Username string `json:"username"`
	Pass     string `json:"pass"`
}

func TestRequestBody(t *testing.T) {
	app := fiber.New()
	app.Post("/login", func(c *fiber.Ctx) error {
		body := c.Body()

		request := new(LoginRequest)
		err := json.Unmarshal(body, request)
		if err != nil {
			return err
		}

		return c.SendString("hello " + request.Username)
	})

	body := strings.NewReader(`{"username":"Ana", "pass":"hehe"}`)
	request := httptest.NewRequest("POST", "/login", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "hello Ana", string(bytes))
}

type RegisterRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Pass     string `json:"pass" xml:"pass" form:"pass"`
	Name     string `json:"name" xml:"name" form:"name"`
}

func TestBodyParser(t *testing.T) {
	app := fiber.New()
	app.Post("/register", func(c *fiber.Ctx) error {
		request := new(RegisterRequest)
		err := c.BodyParser(request)
		if err != nil {
			return err
		}

		return c.SendString("Register Succces " + request.Username)
	})

	body := strings.NewReader(`{"username":"Ana", "pass":"hehe", "name":"anaa"}`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Succces Ana", string(bytes))
}

func TestBodyParserJSON(t *testing.T) {
	app := fiber.New()

	TestBodyParser(t)
	body := strings.NewReader(`{"username":"Ana", "pass":"hehe", "name":"anaa"}`)
	request := httptest.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := app.Test(request)
	assert.Nil(t, err)
	assert.Equal(t, 200, response.StatusCode)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Succces Ana", string(bytes))
}
