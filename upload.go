package main
import(
  "net/http"
  "os"
  "io"
  "encoding/json"
)

type File struct {
    Name string `json:"name"`
}

type Files []File

func upload_handler(w http.ResponseWriter, r *http.Request) {

  r.ParseMultipartForm(32 << 20)
  file, handler, err := r.FormFile("files")
  if err != nil {
    panic(err)
  }
  defer file.Close()
  f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
  if err != nil {
    panic(err)
  }
  defer f.Close()
  io.Copy(f, file)

  files := Files{
        File{Name: handler.Filename},
  }

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(files); err != nil {
        panic(err)
  }

}
