package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

type testGetAPIKeyCase map[string]struct {
	given http.Header
	want  struct {
		key string
		ok  bool
	}
}

func TestGetAPIKey(t *testing.T) {
	cases := testGetAPIKeyCase{
		"normal": {
			given: http.Header{
				"Authorization": []string{"ApiKey mykeyaight"}},
			want: struct {
				key string
				ok  bool
			}{
				key: "mykeyaight", ok: true}},
		"missing prefix": {
			given: http.Header{
				"Authorization": []string{"mykeyaight"}},
			want: struct {
				key string
				ok  bool
			}{
				key: "", ok: false}},
		"wrong prefix": {
			given: http.Header{
				"Authorization": []string{"bearer: mykeyaight"}},
			want: struct {
				key string
				ok  bool
			}{
				key: "", ok: false}},
	}

	for name, c := range cases {
		key, err := auth.GetAPIKey(c.given)
		if key != c.want.key {
			t.Fatalf(
				"[%s] key incorrect:\n want:%s\nget:%s",
				name, c.want.key, key)
		} else if c.want.ok && err != nil {
			t.Fatalf(
				"[%s] should not error:\nget:%s",
				name, err)
		} else if !c.want.ok && err == nil {
			t.Fatalf(
				"[%s] should error:\n get:%s",
				name, err)
		}
	}
}
