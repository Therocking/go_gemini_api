package imgai

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	watchimgs "githup.com/Therocking/go_gemini/pkg/watchImgs"
)

type ErrMsg struct {
	Msg string
}

func ImgAi(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)

		err, _ := json.Marshal(ErrMsg{Msg: "No such file"})
		w.Write(err)
		return
	}
	defer file.Close()

	fileByte, _ := io.ReadAll(file)
	resp := watchimgs.WatchImgs(fileByte)

	json.NewEncoder(w).Encode(resp)
}
