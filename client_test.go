package redash_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/redash-go"
)

func Test_NewClient_OK(t *testing.T) {
	assert := assert.New(t)
	_, err := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	assert.NoError(err)
}

func Test_NewClient_Err(t *testing.T) {
	assert := assert.New(t)
	_, err := redash.NewClient(":redash.example.com", testRedashAPIKey)
	assert.ErrorContains(err, `parse ":redash.example.com": missing protocol scheme`)
}

func Test_Get_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/1", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		assert.Equal("foo=bar", req.URL.Query().Encode())
		return httpmock.NewStringResponse(http.StatusOK, `{"zoo":"baz"}`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, close, err := client.Get(context.Background(), "api/queries/1", map[string]string{"foo": "bar"})
	defer close()
	assert.NoError(err)
	assert.Equal("200", res.Status)
	if res.Body == nil {
		assert.FailNow("res.Body is nil")
	}
	body, _ := io.ReadAll(res.Body)
	assert.Equal(`{"zoo":"baz"}`, string(body))
}

func Test_Get_Err(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodGet, "https://redash.example.com/api/queries/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusNotFound, ``), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, close, err := client.Get(context.Background(), "api/queries/1", map[string]string{"foo": "bar"})
	defer close()
	assert.ErrorContains(err, "GET api/queries/1 failed: HTTP status code not OK: 404")
}

func Test_Post_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/queries/1", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		if req.Body == nil {
			assert.FailNow("req.Body is nil")
		}
		body, _ := io.ReadAll(req.Body)
		assert.Equal(`{"foo":"bar"}`, string(body))
		return httpmock.NewStringResponse(http.StatusOK, `{"zoo":"baz"}`), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, close, err := client.Post(context.Background(), "api/queries/1", map[string]string{"foo": "bar"})
	defer close()
	assert.NoError(err)
	assert.Equal("200", res.Status)
	if res.Body == nil {
		assert.FailNow("res.Body is nil")
	}
	body, _ := io.ReadAll(res.Body)
	assert.Equal(`{"zoo":"baz"}`, string(body))
}

func Test_Post_Err(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodPost, "https://redash.example.com/api/queries/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusNotFound, ``), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, close, err := client.Post(context.Background(), "api/queries/1", map[string]string{"foo": "bar"})
	defer close()
	assert.ErrorContains(err, "POST api/queries/1 failed: HTTP status code not OK: 404")
}

func Test_Delete_OK(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/queries/1", func(req *http.Request) (*http.Response, error) {
		assert.Equal(
			http.Header(
				http.Header{
					"Authorization": []string{"Key " + testRedashAPIKey},
					"Content-Type":  []string{"application/json"},
					"User-Agent":    []string{"redash-go"},
				},
			),
			req.Header,
		)
		return httpmock.NewStringResponse(http.StatusOK, ``), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	res, close, err := client.Delete(context.Background(), "api/queries/1")
	defer close()
	assert.NoError(err)
	assert.Equal("200", res.Status)
}

func Test_Delete_Err(t *testing.T) {
	assert := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(http.MethodDelete, "https://redash.example.com/api/queries/1", func(req *http.Request) (*http.Response, error) {
		return httpmock.NewStringResponse(http.StatusNotFound, ``), nil
	})

	client, _ := redash.NewClient("https://redash.example.com", testRedashAPIKey)
	_, close, err := client.Delete(context.Background(), "api/queries/1")
	defer close()
	assert.ErrorContains(err, "DELETE api/queries/1 failed: HTTP status code not OK: 404")
}
