// main パッケージであることの宣言
package main
// importするパッケージの宣言
import (
	"encoding/json"
	// HTTP client, server を実装する用
	"net/http"
	"os"
)

// function main の宣言
func main() {
	// 環境変数で指定した port 番号を取得、代入
	port := os.Getenv("PORT")

	// http.HandleFunc(path string, handler Handler)
	http.HandleFunc("/api/sample", getHelloWorld)
	// post で代入された 8000ポート で起動
	// "string" + "string" -> "stringstring" 
	// http.ListenAndServe(address string, handler Handler)
	http.ListenAndServe(":" + port, nil)
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