/* https://github.com/blackcrw/akumascan/projects/1#card-70649811 */

package tools

import (
	"github.com/blackcrw/akumascan/pkg/nettools"
)

func aesecure(channel chan [3]string, quit chan bool, response *nettools.Response) {
	channel<-[3]string{"AESecure", "", "teste"}

	quit<-true
}