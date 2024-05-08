package watchimgs

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func WatchImgs(file []byte) *genai.GenerateContentResponse {
	ctx := context.Background()

	// Create a new gemini client
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Select model to use
	model := client.GenerativeModel("gemini-pro-vision")
	model.SetTemperature(1)

	// Prompt to the ai, for generate the response
	prompt := []genai.Part{
		genai.ImageData("png", file),
		genai.Text("Describe this img"),
	}

	// Response from the ai
	resp, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
