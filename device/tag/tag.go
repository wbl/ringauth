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
	var parthandle *os.File
	parthandle, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var partreader = bufio.NewReader(parthandle)
	var group *bbssig.Group = new(bbssig.Group)
	var part *bbssig.MemberKey = new(bbssig.MemberKey)
	{
		var gb = basepack.Unpack(partreader)
		var success bool
		_, success = group.Unmarshal(gb)
		if ! success {
			log.Fatal("Unable to unpack group")
		}
		var mb = basepack.Unpack(partreader)
		_, success = part.Unmarshal(group, mb)
		if ! success {
			log.Fatal("Unable to unpack part")
		}
	}
	var tagbytes = part.Tag()
	fmt.Printf("Tag is %s\n", base64.StdEncoding.EncodeToString(tagbytes))
}