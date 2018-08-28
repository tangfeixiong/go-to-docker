package pb

import (
	"fmt"

	"github.com/golang/glog"

	mobypb "github.com/tangfeixiong/go-to-docker/pb/moby"
	networkpb "github.com/tangfeixiong/go-to-docker/pb/moby/network"
)

func (req *DockerImagePullReqResp) CopyWithRequestVaildation() (*DockerImagePullReqResp, error) {
	resp := new(DockerImagePullReqResp)
	if req == nil {
		glog.Errorln("Request required")
		return resp, fmt.Errorf("Pull request required")
	}

	if len(req.RefStr) == 0 {
		glog.Errorln("Image ref required")
		return resp, fmt.Errorf("Image ref reqired")
	}
	resp.RefStr = req.RefStr

	var errs []error
	resp.ImagePullOptions, errs = req.ImagePullOptions.CopyWithValidation()
	for _, err := range errs {
		switch {
		case err == mobypb.ErrNilPointer:
			glog.Warningln("Image pull options not specified")
		default:
			glog.Warningf("Unknown error: %v", err)
			return resp, fmt.Errorf("Unknown error: %v", err)
		}
	}
	return resp, nil
}

func (req *DockerNetworkCreateReqResp) CopyWithRequestValidation() (*DockerNetworkCreateReqResp, error) {
	resp := new(DockerNetworkCreateReqResp)
	if req == nil {
		glog.Errorln("Network request required")
		return resp, fmt.Errorf("Network request required")
	}

	if len(req.Name) == 0 {
		glog.Warningln("Network name not specified")
	}
	resp.Name = req.Name

	var errs []error
	resp.NetworkCreate, errs = req.NetworkCreate.CopyWithValidation()
	for _, err := range errs {
		switch {
		case err == mobypb.ErrNetworkCreateIsNil:
			glog.Errorln("Network creating args required")
			return resp, fmt.Errorf("Network creating args required")
		case err == networkpb.ErrIpamIsNil:
			glog.Warningln("Network creating argument of IPAM not specified")
		case err == networkpb.ErrIPAMConfigIsEmpty:
			glog.Warningln("Network creating argument of IPAM subnet or gateway not specified")
		default:
			glog.Warningf("Unknown error: %v", err)
			return resp, fmt.Errorf("Unkown error: %s", err.Error())
		}
	}

	return resp, nil
}
