package pb

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
)

func TestImageBuild_json(t *testing.T) {
	var src []byte = []byte("test")
	var buildcontext []byte
	var senc string
	senc = base64.URLEncoding.EncodeToString(src)
	buildcontext = []byte(senc)
	fmt.Printf("%q\n", buildcontext)
	senc = base64.StdEncoding.EncodeToString(src)
	buildcontext = []byte(senc)
	fmt.Printf("%q\n", buildcontext)

	obj := DockerImageBuildReqResp{
		BuildContext: src,
	}
	dst, err := json.Marshal(obj)
	if err != nil {
		t.Log(err)
	} else {
		fmt.Printf("%q\n", dst)
	}
}
