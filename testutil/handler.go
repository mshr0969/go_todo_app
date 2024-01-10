package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func AssertJson(t *testing.T, want, got []byte) {
	t.Helper()

	var jw, jg any
	// Unmarshalは、JSONを第二引数の構造体に入れる
	if err := json.Unmarshal(want, &jw); err != nil {
		t.Fatalf("cannnot unmarshal want %q: %v", want, err)
	}
	if err := json.Unmarshal(got, &jg); err != nil {
		t.Fatalf("cannnot unmarshal got %q: %v", got, err)
	}
	// want, gotそれぞれUnmarshalしたものを比較
	if diff := cmp.Diff(jg, jw); diff != "" {
		t.Errorf("got differs: (-got +want)\n%s", diff)
	}
}

// w.Result()で返されたレスポンスを検証する
// body: LoadFileで読みこんだ期待するデータ(want)
func AssertResponse(t *testing.T, got *http.Response, status int, body []byte) {
	t.Helper()
	t.Cleanup(func() { _ = got.Body.Close()})
	gb, err := io.ReadAll(got.Body)
	if err != nil {
		t.Fatal(err)
	}
	if got.StatusCode != status {
		t.Fatalf("want status %d, but got %d, body: %q", status, got.StatusCode, gb)
	}

	if len(gb) == 0 && len(body) == 0{
		return
	}
	AssertJson(t, body, gb)
}

// ゴールデンテストで使用。ファイルを読み込んで返す
func LoadFile(t *testing.T, path string) []byte {
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("cannnot read from %q: %v", path, err)
	}
	return bt
}
