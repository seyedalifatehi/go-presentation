package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// === Add CORS middleware ===
	// CORS (Cross-Origin Resource Sharing)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // or specify: "http://localhost:3000" for frontend
	}))

	// === Add custom middleware ===
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		log.Printf("[%s] %s - %v", c.Method(), c.Path(), duration)
		return err
	})

	// === Huma setup ===
	humaConfig := huma.DefaultConfig("Test API", "1.0.0")
	humaConfig.Servers = []*huma.Server{
		{URL: "http://localhost:8000"},
	}

	api := humafiber.New(app, humaConfig)

	// === Sample GET endpoint ===
	huma.Register(api, huma.Operation{
		OperationID: "greeting-message",
		Method:      http.MethodGet,
		Path:        "/sample-get",
		Summary:     "Greeting message",
		Description: "Returns a greeting message",
		Tags:        []string{"greeting"},
	}, func(_ context.Context, _ *struct{}) (*struct {
		Body struct {
			Message string `json:"message"`
		}
	}, error) {
		resp := &struct {
			Body struct {
				Message string `json:"message"`
			}
		}{}
		resp.Body.Message = "hello world"
		return resp, nil
	})

	// === Sample POST endpoint ===
	type SamplePostRequest struct {
		Body struct {
			Name string `json:"name"`
		}
	}

	type SamplePostResponse struct {
		Body struct {
			Greeting string `json:"greeting"`
		}
	}

	huma.Register(api, huma.Operation{
		OperationID: "sample-post",
		Method:      http.MethodPost,
		Path:        "/sample-post",
		Summary:     "Sample POST endpoint",
		Description: "Accepts a name and returns a greeting",
		Tags:        []string{"greeting"},
	}, func(_ context.Context, req *SamplePostRequest) (*SamplePostResponse, error) {
		resp := &SamplePostResponse{}
		resp.Body.Greeting = "Hello, " + req.Body.Name + "!"
		return resp, nil
	})

	// === Start the server ===
	log.Println("🚀 Server running on http://localhost:8000")
	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}
