// main パッケージであることの宣言
package main
// importするパッケージの宣言
import (
	// DB とのやりとり用
	"database/sql"
	"encoding/json"
	// HTTP client, server を実装する用
	"net/http"
	"os"
	
	// postgresql のクライアントツール
	// _ : packageの中身を利用せずに、init()関数が必要な場合に利用する
	_ "github.com/lib/pq"
)

type Todo struct {
	Id int `json: "id"`
	Task string `json: "task"`
}

type JsonResponse struct {
	Type string `json: "type"`
	Data []Todo `json: "data"`
	Message string `json: "message"`
}

// function main の宣言
func main() {
	// 環境変数で指定した port 番号を取得、代入
	port := os.Getenv("PORT")
	
	http.HandleFunc("/todos", GetTodos)

	// http.HandleFunc(path string, handler Handler)
	http.HandleFunc("/api/sample", getHelloWorld)
	// post で代入された 8000ポート で起動
	// "string" + "string" -> "stringstring" 
	// http.ListenAndServe(address string, handler Handler)
	http.ListenAndServe(":" + port, nil)
}

func dbConn() (db *sql.DB) {
	psqlInfo := "host=postgres user=postgres password=postgres dbname=mydb port=5432 sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	
	checkErr(err)

	return db
}

// json の生成・response の作成処理
func getHelloWorld(w http.ResponseWriter, _r *http.Request) {
	// map: TS の連想配列の型指定のようなもの
	// TS: obj: { [key: string]: string } = { name: "john" }
	// Golang: obj := map[string]string{ name: "john" }
	ping := map[string]string{"message": "Hello World!!!"}
	// json.Marshal でGoの値をJSONデータにエンコード
	res, _ := json.Marshal(ping)
	// HTTP1.x サーバーレスポンス形式 で json を書き込む
	w.Write(res)
}

func GetTodos (w http.ResponseWriter, r *http.Request) {
	// DB接続関数
	db := dbConn()
	// クエリを叩く・エラーをひろう
	rows, err := db.Query("SELECT * FROM todo;") 
	
	// エラーチェック関数
	checkErr(err)
	// todo, todos を宣言
	todo := Todo{}
	todos := []Todo{}
	
	// for文で todo を todos につっこむ
	for rows.Next() {
		var id int
		var task string
		// 順番注意: DB Column 順に設定しないと DB 側でエラーが起きる
		err = rows.Scan(&task, &id)

		checkErr(err)
		todo.Id = id
		todo.Task = task
		
		todos = append(todos, todo)
	}
	response := JsonResponse{Type: "success", Data: todos}
	json.NewEncoder(w).Encode(response)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}