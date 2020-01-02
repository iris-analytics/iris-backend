package provider

import "github.com/sethgrid/pester"

// GetPesterHTTPClient returns an instance of Pester
func GetPesterHTTPClient() *pester.Client {
	httpClient := pester.New()
	httpClient.Concurrency = 1
	httpClient.MaxRetries = 5
	httpClient.Backoff = pester.LinearBackoff
	httpClient.KeepLog = true

	return httpClient
}
