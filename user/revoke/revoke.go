package main

import "bufio"
import "common/libs/basepack"
import "encoding/base64"
import "github.com/agl/pond/bbssig"
import "fmt"
import "log"
import "os"

func main() {
	var err error
	var group = new(bbssig.Group)
	var sk = new(bbssig.PrivateKey)
	var part = new(bbssig.MemberKey)
	var rev *bbssig.Revocation
	var revbyte []byte
	var partfile *os.File
	partfile, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var privfile *os.File
	privfile, err = os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	var partreader = bufio.NewReader(partfile)
	var privreader = bufio.NewReader(privfile)
	/* Now we have a detail, namely need to skip some stuff from
	the part file before we can use it. But first load the
	private key*/
	{
		var gb = basepack.Unpack(privreader)
		var success bool
		_, success = group.Unmarshal(gb)
		if !success {
			log.Fatal("Could not load key")
		}
		var sb = basepack.Unpack(privreader)
		_, success = sk.Unmarshal(group, sb)
		if !success {
			log.Fatal("Could not load private key")
		}
	}
	{
		var success bool
		basepack.Unpack(partreader) /*Skip first entry*/
		var memb = basepack.Unpack(partreader)
		_, success = part.Unmarshal(group, memb)
		if !success {
			log.Fatal("Part could not be upacked")
		}
	}
	rev = sk.GenerateRevocation(part)
	revbyte = rev.Marshal()
	fmt.Printf("The revokation string is %s\n",
		base64.StdEncoding.EncodeToString(revbyte))
}
