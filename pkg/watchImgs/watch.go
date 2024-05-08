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
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Select model to use
	model := client.GenerativeModel("gemini-pro-vision")
	model.SetTemperature(1)

	// wd, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println("Error al obtener el directorio de trabajo actual:", err)
	// 	return
	// }

	// Select the image to use
	// img, err := os.ReadFile(fmt.Sprintf("%s/%s", wd, "/imgs/landscape.png"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	prompt := []genai.Part{
		genai.ImageData("png", file),
		genai.Text("Describe this img"),
	}

	resp, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}
