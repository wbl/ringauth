package main
import "encoding/base64"
import "time"
import "crypto/sha256"
import "os"
import "fmt"
import "io"
//usage: challenge servicename
func main(){
	var chash []byte
	var sha = sha256.New()
	var t = time.Now().UnixNano()
	var challenge string
	io.WriteString(sha, fmt.Sprintf("%s:%d:", os.Args[1], t))
	chash=sha.Sum(chash)
	challenge=base64.StdEncoding.EncodeToString(chash)
	fmt.Printf("%s\n", challenge)
	os.Exit(0)
}
