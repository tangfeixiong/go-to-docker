package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Reads all .json files in the current folder
// and encodes them as strings literals in textfiles.go
func main() {
	fs, _ := ioutil.ReadDir(".")
	out, _ := os.Create("swagger.pb.go")
	out.Write([]byte("package pb \n\nconst (\n"))
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".json") {
			name := strings.TrimPrefix(f.Name(), "service.")
			out.Write([]byte(strings.TrimSuffix(name, ".json") + " = `"))
			f, _ := os.Open(f.Name())
			origs := new(bytes.Buffer)
			io.Copy(origs, f)
			dests := new(bytes.Buffer)
			for {
				b, err := origs.ReadByte()
				if err == io.EOF {
					break
				}
				if b != byte('`') {
					dests.WriteByte(b)
				} else {
					dests.WriteString("0x60")
				}
			}
			io.Copy(out, dests)
			//io.Copy(out, f)
			out.Write([]byte("`\n"))
		}
	}
	out.Write([]byte(")\n"))
}
