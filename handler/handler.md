## handlerメモ

- validator(https://github.com/go-playground/validator)
    - Unmarshalする構造体に`validate`タグをつけて検証することができる
    - `*validator.Validate.Struct`
    - 今回は、`err := at.Validator.Struct(b)`で、HTTPリクエストから格納した`b`で必須の`title`があるかを検証している。

- `response.go`
    - レスポンスデータをJSONに変換してステータスコードと一緒に`http.ResponseWriter`インターフェースを満たす型の値に書き込む」という処理をまとめる
        - ちなみに、`ResponseWriter`のメソッド3つは以下の順番で使う
            - Header
                - http.Serverのインスタンスを取得し、ヘッダを設定する
            - WriteHeader
                - statusCodeを設定(200の場合は省略可能)
            - Write
                - レスポンスボディを設定

- ヘルパー関数
    - `testing`パッケージに用意
    - 共通のアサーションを定義することで、冗長の削除と、どの位置でエラーが出たかをわかりやすくする
    - `testutil/handler.go`に`t.Hepler()`と宣言することで用いている

- エンドポイント
    - add_task, list_task
    - `http.HandlerFunc`型を満たすように`ServeHTTP`メソッドを実装する
        - `mux.go`でルーティングできるようにする
    - `RespondJSON`でレスポンスを返す(エラーの時はもちろんエラーが入る)
