/* this or self: https://github.com/blackcrw/akumascan/projects/1#card-70646744 */

package tools

import (
	"log"
	"strings"
	"sync"

	"github.com/blackcrw/akumascan/pkg/nettools"
	"github.com/blackcrw/akumascan/pkg/printer"
)

var wg sync.WaitGroup

type detection_params struct {
	url                    string
	tls_certificate_verify bool
	user_agent             string
	wafs_functions_array   []func(chan [3]string, chan bool, *nettools.Response)
}

func NewDetection() *detection_params {
	return &detection_params{
		wafs_functions_array: []func(chan [3]string, chan bool, *nettools.Response) {aesecure, wordfence},
	}
}

func (this *detection_params) SetURL(url string) { switch !strings.HasSuffix(url, "/") { case false: this.url = url + "/"; case true: this.url = url } }
func (this *detection_params) SetTlsCertificateVerify(crt bool) { this.tls_certificate_verify = crt }
func (this *detection_params) SetUserAgent(uat string) { this.user_agent = uat }

func (this *detection_params) RunnerAggressive() {}

func (this *detection_params) RunnerPassive() {
	var net = nettools.NewNETClient()
	net.SetURL(this.url)
	net.SetTlsCertificateVerify(true)

	var response, err = net.Runner()

	if err != nil { log.Fatalln(err) }

	for _, detection := range this.wafs_functions_array { 
		var channel = make(chan [3]string)
		var quit = make(chan bool)

		wg.Add(1)
		go detection(channel, quit, response)

		while: for {
			select{
			case x := <-channel:
				if x[1] != "" {
					printer.PrintString(x[0], x[1], x[2])
				} else if x[1] != "false" && x[1] != "true" {
					printer.PrintString(x[0], x[2])
				} else if x[1] == "true" {
					var response = printer.ScanQ("WAF found and you were blocked! Do you wish to continue ?!")
					if response != "y" || response != "Y" { break while }
				}
			case <-quit:
				wg.Done()
				break while
			}
		}

		wg.Wait()
	}
}