package main
import ("github.com/agl/pond/bbssig"
	"common/libs/basepack"
	"common/libs/objpack"
	"os"
	"log")

func main(){
	/* Generate parameters for a new ring and store in public and
	private. */
	var sk *bbssig.PrivateKey
	var randread, err = os.Open("/dev/random")
	if err != nil {
		log.Fatal(err)
	}

	sk,err = bbssig.GenerateGroup(randread)
	if err != nil {
		log.Fatal(err)
	}
	/* Need to output marshelled form of group ring */
	var pubhandle *os.File
	pubhandle, err = os.Create(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	objpack.PackGroup(pubhandle, sk.Group)
	pubhandle.Close()

	/*Now to save private key*/
	var priv []byte
	var privhandle *os.File
	privhandle, err = os.Create(os.Args[2])
	if err != nil {
		log.Fatal("err")
	}
	priv = sk.Marshal()
	/*TODO: pros and cons of encrypting this file*/
	/*We write the group first as it is needed information for
	 unmarshalling*/
	objpack.PackGroup(privhandle, sk.Group)
	basepack.Packout(privhandle, priv)
	privhandle.Close()
}