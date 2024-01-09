## handlerメモ

- validator(https://github.com/go-playground/validator)
    - Unmarshalする構造体に`validate`タグをつけて検証することができる
    - `*validator.Validate.Struct`
    - 今回は、`err := at.Validator.Struct(b)`で、HTTPリクエストから格納した`b`で必須の`title`があるかを検証している。
