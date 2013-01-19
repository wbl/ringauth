package main
import "bufio"
import "common/libs/basepack"
import "github.com/agl/pond/bbssig"
import "fmt"
import "log"
import "os"


func main(){
	var err error
	var ofhandle *os.File
	ofhandle, err = os.Create(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var privhandle *os.File
	privhandle, err = os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	var privreader *bufio.Reader
	privreader = bufio.NewReader(privhandle)
	var group *bbssig.Group = new(bbssig.Group)
	var privkey *bbssig.PrivateKey = new(bbssig.PrivateKey)
	var part *bbssig.MemberKey = new(bbssig.MemberKey)
	var randreader *os.File
	randreader, err = os.Open("/dev/random")
	if err != nil {
		log.Fatal(err)
	}
	var gb [] byte
	{
		gb = basepack.Unpack(privreader)
		var success bool
		_, success=group.Unmarshal(gb)
		if ! success {
			log.Fatal("Unable to unpack group data")
		}
		var pb = basepack.Unpack(privreader)
		_, success = privkey.Unmarshal(group, pb)
		if ! success {
			log.Fatal("Unable to unpack private key")
		}
	}
	part, err = privkey.NewMember(randreader)
	if err != nil {
		log.Fatal(err)
	}
	{
		var partbyte = part.Marshal()
		basepack.Packout(ofhandle,gb)
		basepack.Packout(ofhandle, partbyte)
		fmt.Printf("Tag for part is ")
		basepack.Packout(os.Stdout, part.Tag())
	}
	ofhandle.Close()
	privhandle.Close()
	randreader.Close()
}