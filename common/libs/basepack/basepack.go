package basepack
import "encoding/base64"
import "io"
import "bufio"

func Packout(h io.Writer, b []byte){
	var encoded = base64.StdEncoding.EncodeToString(b)
	io.WriteString(h, encoded)
	io.WriteString(h,"\n")
}

func Unpack(h *bufio.Reader) [] byte {
	var enc,_ = h.ReadString('\n')
	enc = enc[0:len(enc)-1] //chop newline: the docs are a lie
	var out [] byte
	out, _ = base64.StdEncoding.DecodeString(enc)
	return out
}