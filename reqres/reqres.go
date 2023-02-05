package reqres

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func Get[T any](ctx context.Context, url string) (T, error) {
	var m T

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return m, err
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return m, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return m, err
	}

	data, err := ParseJson[T](body)
	if err != nil {
		return m, err
	}
	return data, nil
}

func ParseJson[T any](b []byte) (T, error) {
	var m T
	if err := json.Unmarshal(b, &m); err != nil {
		return m, err
	}
	return m, nil
}

func toJson(T any) ([]byte, error) {
	byt, err := json.Marshal(T)
	if err != nil {
		return nil, err
	}
	return byt, nil
}

func Post[T any](ctx context.Context, url string, payload any) (T, error) {
	var m T

	b, err := toJson(payload)
	if err != nil {
		return m, err
	}

	byteReader := bytes.NewReader(b)
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, url, byteReader)
	if err != nil {
		return m, err
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return m, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return m, err
	}

	data, err := ParseJson[T](body)
	if err != nil {
		return m, err
	}
	return data, nil
}
