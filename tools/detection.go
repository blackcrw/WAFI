package tools

import (
	"fmt"
	"log"
	"strings"

	"github.com/blackcrw/akumascan/pkg/nettools"
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

func (this *detection_params) RunnerAggressive() {}

func (this *detection_params) RunnerPassive() {
	var wafs_functions_array = []func(string) bool{aesecure}

	var net = nettools.NewNETClient()
	net.SetURL(this.url)
	net.SetTlsCertificateVerify(true)

	var response, err = net.Runner()

	if err != nil { log.Fatalln(err) }

	for _, result := range wafs_functions_array { fmt.Println(result(response.Raw)) }
}