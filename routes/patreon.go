package routes

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

var httpClient = http.Client{}

const PatreonImageProxyRoute = "/patreon/"

func PatreonImageProxy(w http.ResponseWriter, r *http.Request) {
	patreonUrl := r.URL.Path[len(PatreonImageProxyRoute):]

	parsedURL, err := url.Parse(patreonUrl)
	if err != nil {
		log.Fatal(err)
		return
	}

	if parsedURL.Host != "c10.patreonusercontent.com" && parsedURL.Host != "c8.patreon.com" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotAcceptable)

		w.Write([]byte(`{"status": "Only whitelisted patreon urls are allowed."}`))
		return
	}

	// TODO: check if domains are allowed

	req, err := http.NewRequest(http.MethodGet, parsedURL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "oengus.io/patreon-fetcher")
	req.Header.Set("Origin", "https://oengus.io")

	res, httpErr := httpClient.Do(req)
	if httpErr != nil {
		log.Fatal(httpErr)
	}

	// defer calls are not executed until the function returns
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	w.Header().Set("Content-Type", res.Header.Get("Content-Type"))
	w.Header().Set("ETag", res.Header.Get("ETag"))
	//w.Header().Set("Cache-Control", res.Header.Get("Cache-Control"))
	w.Header().Set("Cache-Control", "public, max-age=31536000") // seems to just be this value

	w.Write(body)
}
