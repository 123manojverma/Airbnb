package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseUrl string, pathPreix string) http.HandlerFunc {
	target, err := url.Parse(targetBaseUrl)

	if err != nil {
		fmt.Println("Error parsing target URL:", err)
		return nil
	}

	// proxy := httputil.NewSingleHostReverseProxy(target)

	// originalDirector := proxy.Director

	// proxy.Director = func(r *http.Request) {
	// 	originalDirector(r)

	// 	r.Host = target.Host

	// 	fmt.Println("Proxying request to:", targetBaseUrl+r.URL.Path)
	// 	fmt.Println("Original request path:", r.URL.Path)
	// 	fmt.Println("Path prefix:", pathPreix)

	// 	r.URL.Path = strings.TrimPrefix(r.URL.Path, pathPreix)

	// 	fmt.Println("Modified request path:", r.URL.Path)

	// 	if userId, ok := r.Context().Value("userId").(string); ok {
	// 		r.Header.Set("X-User-ID", userId)
	// 		fmt.Println("userId",userId)
	// 	}
	// }

	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {

			// Forward request to target service
			pr.SetURL(target)

			// Remove gateway prefix
			pr.Out.URL.Path = strings.TrimPrefix(
				pr.In.URL.Path,
				pathPreix,
			)

			// Forward authenticated user ID
			if userID, ok := pr.In.Context().Value("userId").(string); ok {
				pr.Out.Header.Set("X-User-ID", userID)
				fmt.Println("Got userId",userID)
			}else{
				fmt.Println("userId is missing")
			}
		},
	}

	return proxy.ServeHTTP
}
