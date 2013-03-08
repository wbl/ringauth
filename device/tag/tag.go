package main

import "bufio"
import "common/libs/objpack"
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
	objpack.UnPackPart(partreader, group, part)
	var tagbytes = part.Tag()
	fmt.Printf("Tag is %s\n", base64.StdEncoding.EncodeToString(tagbytes))
}
