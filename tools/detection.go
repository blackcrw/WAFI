package tools

import (
	"fmt"
	"strings"

	"github.com/blackcrw/wafi/pkg/nettools"
)

type detection_params struct {
	url                    string
	tls_certificate_verify bool
	user_agent             string
}

func NewDetection() *detection_params { return &detection_params{} }

func (this *detection_params) SetURL(url string) { switch !strings.HasSuffix(url, "/") { case false: this.url = url + "/"; case true: this.url = url } }
func (this *detection_params) SetTlsCertificateVerify(crt bool) { this.tls_certificate_verify = crt }
func (this *detection_params) SetUserAgent(uat string) { this.user_agent = uat }

func (this *detection_params) RunnerAggressive() {/* The aggressive detection mode will use goroutines to force several requests to the target until it blocks it so that we can detect which WAF the target is using. */}

func (this *detection_params) RunnerPassive() {
	var wafs_functions_array = []func(string) bool{aesecure}

	var net = nettools.NewNETClient()
	net.SetURL(this.url)
	net.SetTlsCertificateVerify(true)
	net.Runner()
	
	for _, result := range wafs_functions_array {
		fmt.Println(result("x"))
	}
}