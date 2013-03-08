package main
import "bufio"
import "common/libs/basepack"
import "common/libs/objpack"
import "crypto/sha256"
import "encoding/base64"
import "github.com/agl/pond/bbssig"
import "log"
import "os"

func main(){
	var challenge = os.Args[1]
	var err error
	var parthandle *os.File
	parthandle, err = os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	var partreader = bufio.NewReader(parthandle)
	var randreader *os.File
	randreader, err = os.Open("/dev/random")
	if err != nil {
		log.Fatal(err)
	}
	var group *bbssig.Group = new(bbssig.Group)
	var mem *bbssig.MemberKey = new(bbssig.MemberKey)
	var hash = sha256.New()
	var cb [] byte
	var response [] byte
	objpack.UnPackPart(partreader, group, mem)
	cb, err= base64.StdEncoding.DecodeString(challenge)
	if err != nil {
		log.Fatal("Invalid Challenge")
	}
	response, err=mem.Sign(randreader, cb, hash)
	if err != nil {
		log.Fatal(err)
	}
	basepack.Packout(os.Stdout, response)
}