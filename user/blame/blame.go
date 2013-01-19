package main
import "bufio"
import "common/libs/basepack"
import "encoding/base64"
import "fmt"
import "github.com/agl/pond/bbssig"
import "log"
import "os"

func main(){
	var err error
	var group = new(bbssig.Group)
	var priv = new(bbssig.PrivateKey)
	var privhandle *os.File
	privhandle, err =os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	var privreader *bufio.Reader
	privreader = bufio.NewReader(privhandle)
	{
		var gp [] byte
		var success bool
		gp = basepack.Unpack(privreader)
		_, success = group.Unmarshal(gp)
		if ! success {
			log.Fatal("Unable to read group data")
		}
		var pb [] byte
		pb = basepack.Unpack(privreader)
		_, success = priv.Unmarshal(group, pb)
		if ! success {
			log.Fatal("Unable to read private key")
		}
	}
	var sig [] byte
	sig, err = base64.StdEncoding.DecodeString(os.Args[1])
	if err != nil {
		log.Fatal("Signature is invalid")
	}
	var tag [] byte
	var success bool
	tag, success=priv.Open(sig)
	if ! success {
		log.Fatal("Signature unknown")
	} else {
		fmt.Printf("The tag is %s\n", base64.StdEncoding.EncodeToString(tag))
	}
}