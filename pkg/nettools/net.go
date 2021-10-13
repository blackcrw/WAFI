/* https://github.com/blackcrw/akumascan/projects/1#card-70646744 */

package nettools

import (
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
)

type net_params struct{
	url                    string
	method                 string
	tls_certificate_verify bool
	user_agent             string
	redirect               func(req *http.Request, via []*http.Request) error
}

type Response struct {
	RawIo    io.Reader
	Raw      string
	URL      string
	Response *http.Response
}

func NewNETClient() *net_params { return &net_params{ method: "GET", user_agent: "akumascan - Advanced Web Application Firewall Scanner", tls_certificate_verify: true } }

func (this *net_params) SetURL(url string) { this.url = url }
func (this *net_params) SetTlsCertificateVerify(crt bool) { this.tls_certificate_verify = crt }
func (this *net_params) SetUserAgent(uat string) { this.user_agent = uat }

func (this *net_params) Runner() (*Response, error) {
	var client = &http.Client{
		CheckRedirect: this.redirect,
		Transport: &http.Transport{
			DisableKeepAlives: false,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: this.tls_certificate_verify,
			},
		},
	}

	var request, err = http.NewRequest(this.method, this.url, nil)

	if err != nil { return nil, err }

	request.Header.Set("User-Agent", this.user_agent)

	response, err := client.Do(request)

	if err != nil { return nil, err }

	raw, err := ioutil.ReadAll(response.Body)

	if err != nil { return nil, err }

	return &Response{ Raw: string(raw), URL: this.url, Response: response }, nil
}