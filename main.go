package herb

import (
	"flag"
)

var (
	host        = flag.String("host", "localhost", "remote host")
	username         = flag.String("username", "main", "username")
	privateKey   = flag.String("privateKey", "", "private key")
)

func main(){

	flag.Parse()
}
