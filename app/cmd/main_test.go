package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/steinfletcher/apitest"
)

func Test_main(t *testing.T) {
	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Get("/").
		Expect(t).
		Body("Hello, World!").
		Status(http.StatusOK).
		End()
}

func Test_GetShops(t *testing.T) {

	expected := `[{"DeletedAt":null,"ID":1,"reference":"PRU001","name":"Prueba","time_open":"2023-05-17T19:19:45.371+02:00","time_close":"2023-05-17T19:19:45.371+02:00","days_open":"Lunes,Martes,Miércoles","CreatedAt":"2023-05-17T19:19:45.372+02:00","UpdatedAt":"2023-05-17T19:19:45.372+02:00"},{"DeletedAt":null,"ID":2,"reference":"Prueba 2","name":"La Pelu de Rober","time_open":"2023-06-18T09:00:00+02:00","time_close":"2023-06-18T20:00:00+02:00","days_open":"Lunes,Martes,Miércoles","CreatedAt":"2023-06-28T14:12:13.844+02:00","UpdatedAt":"2023-06-28T14:12:13.844+02:00"},{"DeletedAt":null,"ID":3,"reference":"Prueba 2","name":"La Pelu de Rober","time_open":"2023-06-18T09:00:00+02:00","time_close":"2023-06-18T20:00:00+02:00","days_open":"Lunes,Martes,Miércoles","CreatedAt":"2023-06-28T17:35:48.132+02:00","UpdatedAt":"2023-06-28T17:35:48.132+02:00"}]`

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Get("/api/v1/shop").
		Expect(t).
		Body(expected).
		Status(http.StatusOK).
		End()
}

func Test_GetShopByID(t *testing.T) {

	expected := `{"DeletedAt":null,"ID":1,"reference":"PRU001","name":"Prueba","time_open":"2023-05-17T19:19:45.371+02:00","time_close":"2023-05-17T19:19:45.371+02:00","days_open":"Lunes,Martes,Miércoles","CreatedAt":"2023-05-17T19:19:45.372+02:00","UpdatedAt":"2023-05-17T19:19:45.372+02:00"}`

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(newApp())).
		Get("/api/v1/shop/1").
		Expect(t).
		Body(expected).
		Status(http.StatusOK).
		End()
}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		// copy body
		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}
