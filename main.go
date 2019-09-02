package main


import (
        "fmt"
        "log"
        "net/http"
        "text/template"
)


type Person struct {
        Id string
        Name string
}

//HTML テンプレートを読み込む関数
func readerTemplate(w http.ResponseWriter, r *http.Request) {

        person :=Person{Id: "1", Name: "golangtest"}

        //docker上にコピーしたtemplateファイルを指定する（ファイルは後ほど作成する）
        parseedTemplate, err := template.ParseFiles("/go/src/github.com/gouser/templates/template.html")
        if err != nil {
                log.Fatalf("Not found file: %V", err)
        }
        // テンプレートにPerson の値を渡す
        err = parsedTemplate.Execute(w, person)
        if err != nil {
                log.Printf("Error occurred while executing the template or writting its output: %v", err)
                return
        }
}

func main() {
        fmt.Println("API Server Start ...")
        //ハンドラー登録
        http.HandleFunc("/", readerTemplate)
        // localhost:8080で動くようにする
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                log.Fatal("error starting http server: ", err)
                return
        }
}
