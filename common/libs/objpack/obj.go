package objpack

import "common/libs/basepack"
import "github.com/agl/pond/bbssig"
import "io"

func PackGroup(out io.Writer, g * bbssig.Group){
	var pub []byte
	pub = g.Marshal()
	basepack.Packout(out, pub)
}