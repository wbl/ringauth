package objpack

import "common/libs/basepack"
import "github.com/agl/pond/bbssig"
import "io"
import "bufio"
import "log"

func PackGroup(out io.Writer, g * bbssig.Group){
	var pub []byte
	pub = g.Marshal()
	basepack.Packout(out, pub)
}


func UnPackPart(partreader *bufio.Reader, group *bbssig.Group, 
	part *bbssig.MemberKey){
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