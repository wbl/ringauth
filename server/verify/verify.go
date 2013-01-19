package main
import "bufio"
import "common/libs/basepack"
import "crypto/sha256"
import "encoding/base64"
import "fmt"
import "github.com/agl/pond/bbssig"
import "log"
import "os"

func main(){
	var err error
	var challenge [] byte
	var response [] byte
	var group *bbssig.Group = new(bbssig.Group)
	var gfile *os.File
	gfile, err = os.Open(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}
	var ghandle = bufio.NewReader(gfile)
	{
		var gb = basepack.Unpack(ghandle)
		var success bool
		_, success= group.Unmarshal(gb)
		if ! success {
			log.Fatal("Unable to unmarshal public key")
		}
	}
	challenge, err = base64.StdEncoding.DecodeString(os.Args[1])
	if err != nil {
		log.Fatal("Invalid challenge")
	}
	response, err = base64.StdEncoding.DecodeString(os.Args[2])
	if err != nil {
		log.Fatal("Invalid response")
	}
	if(group.Verify(challenge, sha256.New(), response)){
		fmt.Printf("yes\n")
		os.Exit(0)
	} else {
		fmt.Printf("no\n")
		os.Exit(1)
	}
}