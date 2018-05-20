package api

import (
	"net/http"
	"net/url"
	"strconv"
)

// GetFalsePositivesFeed gets newly discovered files which are considered possible false positives. An infected scan result is considered to be false positive if 2 or less engines detected the file as being infected. The feed is updated on a daily basis and contains files that are detected in the previous day. This feed contains data about all engines.
func (api *API) GetFalsePositivesFeed(engine string, page int) string {
	fpURL := URL + "feed/false-positives/"
	if engine != "" {
		fpURL += engine
	}
	url, _ := url.Parse(fpURL)
	q := url.Query()
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	url.RawQuery = q.Encode()
	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}

// GetInfectedHashesFeed gets newly discovered malicious hashes. The feed is updated on a daily basis and contains files that are detected as being malicious in the previous day by at least 3 engines.
func (api *API) GetInfectedHashesFeed(fmtType string, page int) string {
	url, _ := url.Parse(URL + "feed/infected")
	q := url.Query()
	switch fmtType {
	case "bro":
		q.Set("type", "bro")
	case "csv":
		q.Set("type", "csv")
	case "json":
		q.Set("type", "json")
	}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	url.RawQuery = q.Encode()
	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}

// GetHashesFeed gets newly discovered hashes
func (api *API) GetHashesFeed(page int) string {
	url, _ := url.Parse(URL + "feed/hashes")
	q := url.Query()
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	url.RawQuery = q.Encode()
	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}

// GetHashDownloadLink Retrieve the download link for a specific file. Any of the md5, sha1 and sha256 hashes can be used for downloading the file. This endpoint must be called for each file.
func (api *API) GetHashDownloadLink(hash string) string {
	req, _ := http.NewRequest("GET", URL+"file/"+hash+"/download", nil)
	req.Header.Add("Authorization", "apikey "+api.Token)
	return FmtResponse(api.Client.Do(req))
}
