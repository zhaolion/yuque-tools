package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func client(fakeServerURL string) *Client {
	var authToken string
	if v := os.Getenv("YUQUE_TEST_AUTH_TOKEN"); v != "" {
		authToken = v
	} else {
		authToken = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	c := NewClient(authToken)
	c.SetBaseURL(fakeServerURL)
	return c
}

func GoldenEndpointServer(
	t *testing.T,
	method, path string,
	statusCode int, golden string,
	handles ...func(w http.ResponseWriter, r *http.Request),
) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			_ = r.Body.Close()
		}()

		require.Equal(t, r.Method, method)
		require.Equal(t, r.URL.Path, path)

		for _, handle := range handles {
			handle(w, r)
		}
		w.WriteHeader(statusCode)
		responseFile := FindFile(golden)
		responseBody, err := ioutil.ReadFile(responseFile)
		if err != nil {
			panic(err)
		}
		_, _ = w.Write([]byte(responseBody))
	}))
}

func RedirectEndpointServer(
	t *testing.T,
	targetMethod, targetPath string, targetStatusCode int,
	targetHandle func(w http.ResponseWriter, r *http.Request),
	redirectMethod, redirectPath string, redirectCode int,
	locationQuery func() string,
) (redirectServer, targetServer *httptest.Server) {
	// 模拟重定向最终的服务器
	targetServer = EndpointServer(t, targetMethod, targetPath, targetStatusCode, targetHandle)

	redirectURL := targetServer.URL
	// 模拟重定向
	redirectServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			_ = r.Body.Close()
		}()
		require.Equal(t, r.Method, redirectMethod)
		require.Equal(t, r.URL.Path, redirectPath)
		w.Header().Set("Location", fmt.Sprintf("%s%s?%s", redirectURL, targetPath, locationQuery()))
		w.WriteHeader(redirectCode)
	}))

	return redirectServer, redirectServer
}

func FindFolder(folder string) string {
	lookupPaths := []string{""}
	nowPathPrefix := ""
	for i := 0; i < 5; i++ {
		lookupPaths = append(lookupPaths, filepath.Join(nowPathPrefix, folder))
		nowPathPrefix = filepath.Join("../", nowPathPrefix)
	}
	for _, p := range lookupPaths {
		exists, err := DirExists(p)
		if err != nil {
			panic(err)
		}
		if exists {
			return p
		}
	}
	panic(fmt.Errorf("can not find folder: %s", folder))
}

func EndpointServer(
	t *testing.T,
	method, path string,
	statusCode int,
	handles ...func(w http.ResponseWriter, r *http.Request),
) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			_ = r.Body.Close()
		}()

		require.Equal(t, r.Method, method)
		require.Equal(t, r.URL.Path, path)

		for _, handle := range handles {
			handle(w, r)
		}

		w.WriteHeader(statusCode)
	}))
}

func FindFile(p string) string {
	dir, file := filepath.Split(p)
	return filepath.Join(FindFolder(dir), file)
}

func DirExists(dir string) (bool, error) {
	f, err := os.Stat(dir)
	if err == nil {
		if f.IsDir() {
			return true, nil
		}
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
