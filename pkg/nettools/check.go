package nettools

import (
	"net"
	"net/url"

	recover_error "github.com/blackcrw/akumascan/pkg/recover"
)

// IsURL :: This function will be used for URL validation
func URLValidate(URL string) bool {
	defer recover_error.NetTools_URL()

	var uri, err = url.ParseRequestURI(URL)

	if err != nil { panic(err) }

	switch uri.Scheme { case "http":; case "https":; default: panic("URL Validade: Invalid scheme") }

	return true
}

func GetHost(URL string) (string, error) {
	var uri, err = url.ParseRequestURI(URL)

	if err != nil { return "", err }

	_, err = net.LookupHost(uri.Host)

	if err != nil { return "", err }

	return uri.Host, nil
}
