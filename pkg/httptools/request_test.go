package httptools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRequest(t *testing.T) {
	t.Run("test response http code 200 and empty body", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			p, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("request body: %s\n", p)

			w.WriteHeader(http.StatusOK)
		}))

		defer srv.Close()

		// 发送请求
		payload := map[string]any{
			"name": "test",
			"age":  18,
		}

		body, _ := json.Marshal(payload)
		b := bytes.NewReader(body)

		req, err := http.NewRequest(http.MethodPost, srv.URL, b)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		r, err := request(req, 10*time.Second)
		if err != nil {
			t.Fatal(err)
		}
		if r.StatusCode != http.StatusOK {
			t.Fatal("expected status code 200, got", r.StatusCode)
		}

		if *r.Body != nil {
			t.Fatal("expected *r.body nil, got", r.Body)
		}
	})

	t.Run("test response code 204 and empty body", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			p, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("request body: %s\n", p)
			w.WriteHeader(http.StatusNoContent)
		}))

		defer srv.Close()

		// 发送请求
		payload := map[string]any{
			"name": "test",
			"age":  18,
		}

		body, _ := json.Marshal(payload)
		b := bytes.NewReader(body)

		req, err := http.NewRequest(http.MethodPost, srv.URL, b)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		r, err := request(req, 10*time.Second)
		if err != nil {
			t.Fatal(err)
		}

		if r.StatusCode != http.StatusNoContent {
			t.Fatal("expected status code 204, got", r.StatusCode)
		}

		if *r.Body != nil {
			t.Fatal("expected *r.body nil, got", r.Body)
		}
	})

	t.Run("test response code 200 and text body", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			p, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("request body: %s\n", p)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("success"))
		}))
		defer srv.Close()

		// 发送请求
		payload := map[string]any{
			"name": "test",
			"age":  18,
		}

		body, _ := json.Marshal(payload)
		b := bytes.NewReader(body)

		req, err := http.NewRequest(http.MethodPost, srv.URL, b)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		r, err := request(req, 10*time.Second)
		if err != nil {
			t.Fatal(err)
		}

		if r.StatusCode != http.StatusOK {
			t.Fatal("expected status code 200, got", r.StatusCode)
		}

		if *r.Body == nil {
			t.Fatal("expected *r.body is not nil, got", r.Body)
		}

		j, _ := json.Marshal(r)
		t.Logf("response: %+v\n", string(j))

	})

}

func TestJsonPost(t *testing.T) {
	t.Run("test response http code 200 and empty body", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			p, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatal(err)
			}

			w.WriteHeader(http.StatusOK)
		}))

		defer srv.Close()

		payload := map[string]any{
			"name": "test",
			"age":  18,
		}

		// 发送请求
		req := Request{
			Method:  "POST",
			Url:     srv.URL,
			Input:   payload,
			Timeout: 10 * time.Second,
		}

		res, err := JsonPost(&req)
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != http.StatusOK {
			t.Fatal("expected status code 200, got", res.StatusCode)
		}

		if res.Body == nil || *res.Body != nil {
			t.Fatal("expected res.body nil, got", res.Body)
		}

	})
}
