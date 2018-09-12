package pb

import (
	"bytes"
	"errors"
	"fmt"

	dockertypes "github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"
	networktypes "github.com/docker/docker/api/types/network"
	"github.com/golang/glog"

	mobypb "github.com/tangfeixiong/go-to-docker/pb/moby"
	containerpb "github.com/tangfeixiong/go-to-docker/pb/moby/container"
	filterspb "github.com/tangfeixiong/go-to-docker/pb/moby/filters"
	networkpb "github.com/tangfeixiong/go-to-docker/pb/moby/network"
)

var (
	errRequestRequired             = errors.New("request required")
	errImgRefRequired              = errors.New("image ref required")
	errImgKeyRequired              = errors.New("image ID or name required")
	errImgKeyTypeIllegal           = errors.New("illegal key, either 1(ID) or 2(name) permitted")
	errImgBuildContextNotSpecified = errors.New("image build context not specified")
)

func (m *DockerContainerRunReqResp) DeepCopyChecked() (*DockerContainerRunReqResp, error) {
	if m == nil {
		return new(DockerContainerRunReqResp), errRequestRequired
	}
	obj := &DockerContainerRunReqResp{
		Config:                     (*containerpb.Config)(nil),
		HostConfig:                 (*containerpb.HostConfig)(nil),
		NetworkingConfig:           (*networkpb.NetworkingConfig)(nil),
		Name:                       "",
		ImagePullOptions:           (*mobypb.ImagePullOptions)(nil),
		ContainerCreateCreatedBody: (*containerpb.ContainerCreateCreatedBody)(nil),
	}
	checkedErrors := []error{}
	var errorList []error
	obj.Config, errorList = m.Config.DeepCopyChecked()
	checkedErrors = append([]error{}, errorList...)
	errorList = errorList[:0]
	obj.HostConfig, errorList = m.HostConfig.DeepCopyChecked()
	checkedErrors = append(checkedErrors, errorList...)
	errorList = errorList[:0]
	obj.NetworkingConfig, errorList = m.NetworkingConfig.DeepCopyChecked()
	checkedErrors = append(checkedErrors, errorList...)
	errorList = errorList[:0]
	obj.Name = obj.Name
	if len(obj.Name) == 0 {
		glog.Warningln("container name not specified")
	}
	obj.ImagePullOptions, errorList = m.ImagePullOptions.DeepCopyChecked()
	checkedErrors = append(checkedErrors, errorList...)

	msgBuf := bytes.Buffer{}
	for _, err := range checkedErrors {
		switch {
		case err == containerpb.ErrNilPointer:
			glog.Warningln(err.Error())
		case err == containerpb.ErrConfigIsNil:
			glog.Errorf("config required: %v", err)
			msgBuf.WriteString(err.Error())
			msgBuf.WriteString("; ")
		case err == containerpb.ErrImageNotSpecified:
			glog.Errorf("config required: %v", err)
			msgBuf.WriteString(err.Error())
			msgBuf.WriteString("; ")
		case err == containerpb.ErrHostConfigIsNil:
			glog.Warningln(err.Error())
		case err == mobypb.ErrImagePullOptionsIsNil:
			glog.Warningln(err.Error())
		default:
			glog.Warning("Unknown error: %v", err)
			msgBuf.WriteString(err.Error())
		}
	}
	if msgBuf.Len() != 0 {
		msgBuf.Truncate(msgBuf.Len() - 2)
		return obj, fmt.Errorf("args checked errors: %s", msgBuf.String())
	}
	return obj, nil
}

func (m *DockerContainerRunReqResp) ExportAsDockerApiTypeContainerCreateConfig() dockertypes.ContainerCreateConfig {
	tgt := dockertypes.ContainerCreateConfig{
		Name:             "",
		Config:           (*containertypes.Config)(nil),
		HostConfig:       (*containertypes.HostConfig)(nil),
		NetworkingConfig: (*networktypes.NetworkingConfig)(nil),
		AdjustCPUShares:  false,
	}
	if m != nil {
		tgt.Name = m.Name
		tgt.Config = m.Config.ExportIntoDockerApiType()
		tgt.HostConfig = m.HostConfig.ExportIntoDockerApiType()
		tgt.NetworkingConfig = m.NetworkingConfig.ExportIntoDockerApiType()
		tgt.AdjustCPUShares = true
	}
	return tgt
}

func (m *DockerContainerListReqResp) DeepCopyCheckedArgs() (*DockerContainerListReqResp, error) {
	if m == nil {
		return new(DockerContainerListReqResp), errRequestRequired
	}
	obj := &DockerContainerListReqResp{
		ContainerListOptions: (*mobypb.ContainerListOptions)(nil),
		Containers:           make([]*mobypb.Container, 0),
	}
	errorList := []error{}
	obj.ContainerListOptions, errorList = m.ContainerListOptions.DeepCopyChecked()
	var err error = nil
	if len(errorList) != 0 {
		return obj, errorList[0]
	}
	return obj, err
}

func (m *DockerContainerInspectReqResp) DeepCopyCheckedArgs() (*DockerContainerInspectReqResp, error) {
	if m == nil {
		return new(DockerContainerInspectReqResp), errRequestRequired
	}
	obj := &DockerContainerInspectReqResp{
		Id:            m.Id,
		Name:          m.Name,
		KeyType:       m.KeyType,
		ContainerJson: (*mobypb.ContainerJSON)(nil),
	}
	errorList := []error{}
	switch {
	case m.KeyType == DockerContainerInspectReqResp_ID:
		if len(m.Id) == 0 {
			errorList = append(errorList, fmt.Errorf("id required"))
		}
	case m.KeyType == DockerContainerInspectReqResp_NAME:
		if len(m.Name) == 0 {
			errorList = append(errorList, fmt.Errorf("name required"))
		}
	default:
		errorList = append(errorList, fmt.Errorf("key type must be 1(ID) or 2(NAME)"))
	}

	if len(errorList) != 0 {
		return obj, fmt.Errorf("args checked errors: %v", errorList[0])
	}
	return obj, nil
}

func (m *DockerContainerRemoveReqResp) DeepCopyCheckedArgs() (*DockerContainerRemoveReqResp, error) {
	if m == nil {
		return new(DockerContainerRemoveReqResp), fmt.Errorf("request required")
	}
	obj := &DockerContainerRemoveReqResp{
		Id:                     m.Id,
		Name:                   m.Name,
		KeyType:                m.KeyType,
		ContainerRemoveOptions: (*mobypb.ContainerRemoveOptions)(nil),
	}
	errorList := []error{}
	switch {
	case m.KeyType == DockerContainerRemoveReqResp_ID:
		if len(m.Id) == 0 {
			errorList = append(errorList, fmt.Errorf("id required"))
		}
	case m.KeyType == DockerContainerRemoveReqResp_NAME:
		if len(m.Name) == 0 {
			errorList = append(errorList, fmt.Errorf("name required"))
		}
	default:
		errorList = append(errorList, fmt.Errorf("key type must ID(1) or NAME(2)"))
	}

	if m.ContainerRemoveOptions != nil {
		obj.ContainerRemoveOptions = &mobypb.ContainerRemoveOptions{
			RemoveVolumes: m.ContainerRemoveOptions.RemoveVolumes,
			RemoveLinks:   m.ContainerRemoveOptions.RemoveLinks,
			Force:         m.ContainerRemoveOptions.Force,
		}
	}

	if len(errorList) != 0 {
		return obj, fmt.Errorf("args checked errors: %v", errorList[0])
	}
	return obj, nil
}

func (m *DockerContainerPruneReqResp) DeepCopyCheckedArgs() (*DockerContainerPruneReqResp, error) {
	obj := new(DockerContainerPruneReqResp)
	if m != nil {
		obj.Filters = m.Filters.DeepCopyChecked()
	}
	return obj, nil
}

func (m *DockerImageInspectReqResp) DeepCopyCheckedArgs() (*DockerImageInspectReqResp, error) {
	if m == nil {
		return new(DockerImageInspectReqResp), errRequestRequired
	}

	errorList := []error{}
	tgt := &DockerImageInspectReqResp{
		Id:           m.Id,
		Ref:          m.Ref,
		KeyType:      m.KeyType,
		ImageInspect: (*mobypb.ImageInspect)(nil),
	}

	switch {
	case m.KeyType == DockerImageInspectReqResp_ID:
		if len(m.Id) == 0 {
			errorList = append(errorList, errors.New("image id required"))
		}
	case m.KeyType == DockerImageInspectReqResp_REF:
		if len(m.Ref) == 0 {
			errorList = append(errorList, errors.New("image name required"))
		}
	default:
		errorList = append(errorList, errors.New("unknown key type, either 1(id) or 2(ref) required"))
	}

	msgBuf := bytes.Buffer{}
	for _, err := range errorList {
		msgBuf.WriteString(err.Error())
		msgBuf.WriteString("; ")
	}
	if msgBuf.Len() != 0 {
		msgBuf.Truncate(msgBuf.Len() - 2)
		return tgt, fmt.Errorf("args checked errors: %s", msgBuf.String())
	}
	return tgt, nil
}

func (m *DockerImageListReqResp) DeepCopyCheckedArgs() (*DockerImageListReqResp, error) {
	if m == nil {
		return new(DockerImageListReqResp), errRequestRequired
	}

	errorList := []error{}
	tgt := &DockerImageListReqResp{
		ImageListOptions: new(mobypb.ImageListOptions),
		ImageSummaries:   make([]*mobypb.ImageSummary, 0),
	}

	var errs []error
	tgt.ImageListOptions, errs = m.ImageListOptions.DeepCopyChecked()
	errorList = append(errorList, errs...)

	msgBuf := bytes.Buffer{}
	for _, err := range errorList {
		switch {
		case err == mobypb.ErrNilPointer:
			glog.Warningln("Image list options not specified")
		default:
			glog.Warningf("Unknown error: %v", err)
			msgBuf.WriteString(err.Error())
			msgBuf.WriteString("; ")
		}
	}
	if msgBuf.Len() != 0 {
		msgBuf.Truncate(msgBuf.Len() - 2)
		return tgt, fmt.Errorf("args checked errors: %s", msgBuf.String())
	}
	return tgt, nil
}

func (m *DockerImageRemoveReqResp) DeepCopyCheckedArgs() (*DockerImageRemoveReqResp, error) {
	if m == nil {
		return new(DockerImageRemoveReqResp), errRequestRequired
	}

	errorList := []error{}
	tgt := &DockerImageRemoveReqResp{
		Id:                       m.Id,
		Ref:                      m.Ref,
		KeyType:                  m.KeyType,
		ImageRemoveOptions:       new(mobypb.ImageRemoveOptions),
		ImageDeleteResponseItems: make([]*mobypb.ImageDeleteResponseItem, 0),
	}

	switch {
	case m.KeyType == DockerImageRemoveReqResp_ID:
		if len(m.Id) == 0 {
			errorList = append(errorList, errors.New("image id required"))
		}
	case m.KeyType == DockerImageRemoveReqResp_REF:
		if len(m.Ref) == 0 {
			errorList = append(errorList, errors.New("image ref required"))
		}
	default:
		errorList = append(errorList, errors.New("unknown type key, either 1(id) or 2(ref) required"))
	}

	var errs []error
	tgt.ImageRemoveOptions, errs = m.ImageRemoveOptions.DeepCopyChecked()
	errorList = append(errorList, errs...)

	msgBuf := bytes.Buffer{}
	for _, err := range errorList {
		switch {
		case err == mobypb.ErrNilPointer:
			glog.Warningln("Image remove options not specified")
		default:
			glog.Warningf("Unknown error: %v", err)
			msgBuf.WriteString(err.Error())
			msgBuf.WriteString("; ")
		}
	}
	if msgBuf.Len() != 0 {
		return tgt, fmt.Errorf("args checked errors: %s", msgBuf.String())
	}
	return tgt, nil
}

func (m *DockerImagePruneReqResp) DeepCopyCheckedArgs() (*DockerImagePruneReqResp, error) {
	if m == nil {
		return new(DockerImagePruneReqResp), errRequestRequired
	}

	tgt := &DockerImagePruneReqResp{
		Filters:           new(filterspb.Args),
		ImagesPruneReport: (*mobypb.ImagesPruneReport)(nil),
	}

	tgt.Filters = m.Filters.DeepCopyChecked()

	return tgt, nil
}

func (m *DockerImagePullReqResp) DeepCopyCheckedArgs() (*DockerImagePullReqResp, error) {
	if m == nil {
		return new(DockerImagePullReqResp), errRequestRequired
	}

	errorList := []error{}
	tgt := &DockerImagePullReqResp{
		RefStr:           m.RefStr,
		ImagePullOptions: new(mobypb.ImagePullOptions),
		RespBody:         make([]byte, 0),
	}

	if len(m.RefStr) == 0 {
		errorList = append(errorList, errImgRefRequired)
	}

	var errs []error
	tgt.ImagePullOptions, errs = m.ImagePullOptions.DeepCopyChecked()
	errorList = append(errorList, errs...)

	msgBuf := bytes.Buffer{}
	for _, err := range errorList {
		switch {
		case err == mobypb.ErrNilPointer:
			glog.Warningln("Image pull options not specified")
		default:
			glog.Warningf("Unknown error: %v", err)
			msgBuf.WriteString(err.Error())
			msgBuf.WriteString("; ")
		}
	}
	if msgBuf.Len() != 0 {
		msgBuf.Truncate(msgBuf.Len() - 2)
		return tgt, fmt.Errorf("args checked errors: %v", msgBuf.String())
	}
	return tgt, nil
}

func (m *DockerImagePushReqResp) DeepCopyCheckedArgs() (*DockerImagePushReqResp, error) {
	if m == nil {
		return new(DockerImagePushReqResp), errRequestRequired
	}

	errorList := []error{}
	tgt := &DockerImagePushReqResp{
		RefStr:           m.RefStr,
		ImagePushOptions: new(mobypb.ImagePushOptions),
		RespBody:         make([]byte, 0),
	}

	if len(m.RefStr) == 0 {
		errorList = append(errorList, errImgRefRequired)
	}

	var errs []error
	tgt.ImagePushOptions, errs = m.ImagePushOptions.DeepCopyChecked()
	errorList = append(errorList, errs...)

	msgBuf := bytes.Buffer{}
	for _, err := range errorList {
		switch {
		case err == mobypb.ErrNilPointer:
			glog.Warningln("Image push options not specified")
		default:
			glog.Warningf("Unknown error: %v", err)
			msgBuf.WriteString(err.Error())
			msgBuf.WriteString("; ")
		}
	}
	if msgBuf.Len() != 0 {
		msgBuf.Truncate(msgBuf.Len() - 2)
		return tgt, fmt.Errorf("args checked errors: %v", msgBuf.String())
	}
	return tgt, nil
}

func (m *DockerImageBuildReqResp) DeepCopyCheckedArgs() (*DockerImageBuildReqResp, error) {
	if m == nil {
		return new(DockerImageBuildReqResp), errRequestRequired
	}

	errorList := []error{}
	tgt := &DockerImageBuildReqResp{
		BuildContext:       ([]byte)(nil),
		ImageBuildOptions:  (*mobypb.ImageBuildOptions)(nil),
		ImageBuildResponse: (*mobypb.ImageBuildResponse)(nil),
	}

	if len(m.BuildContext) == 0 {
		errorList = append(errorList, errImgBuildContextNotSpecified)
	} else {
		tgt.BuildContext = make([]byte, len(m.BuildContext))
		for _, v := range m.BuildContext {
			tgt.BuildContext = append(tgt.BuildContext, v)
		}
	}

	var errs []error
	tgt.ImageBuildOptions, errs = m.ImageBuildOptions.DeepCopyChecked()
	errorList = append(errorList, errs...)

	msgBuf := bytes.Buffer{}
	for _, err := range errorList {
		switch {
		case err == errImgBuildContextNotSpecified:
			if tgt.ImageBuildOptions != nil && len(tgt.ImageBuildOptions.Dockerfile) == 0 {
				glog.Errorln(err.Error())
				errorList = append(errorList, err)
			} else {
				glog.Warningln(err.Error())
			}
		case err == mobypb.ErrImageBuildOptionsIsNil:
			if len(tgt.BuildContext) == 0 {
				glog.Errorln(err.Error())
				errorList = append(errorList, err)
			} else {
				glog.Warningln(err.Error())
			}
		case err == mobypb.ErrImageBuildDockefileNotSpecified:
			if tgt.ImageBuildOptions == nil && len(tgt.BuildContext) == 0 {
				glog.Errorln(err.Error())
				errorList = append(errorList, err)
			} else {
				glog.Warningln(err.Error())
			}
		case err == mobypb.ErrImageBuildOptCtxNotSpecified:
			break
		case err == mobypb.ErrNilPointer:
			glog.Errorf("Image build options not specified")
			msgBuf.WriteString(err.Error())
			msgBuf.WriteString("; ")
		default:
			glog.Warningf("Unknown error: %v", err)
			msgBuf.WriteString(err.Error())
			msgBuf.WriteString("; ")
		}
	}
	if msgBuf.Len() != 0 {
		msgBuf.Truncate(msgBuf.Len() - 2)
		return tgt, fmt.Errorf("args checked errors: %v", msgBuf.String())
	}
	return tgt, nil
}

func (req *DockerNetworkCreateReqResp) DeepCopyCheckedArgs() (*DockerNetworkCreateReqResp, error) {
	resp := new(DockerNetworkCreateReqResp)
	if req == nil {
		return resp, errRequestRequired
	}

	if len(req.Name) == 0 {
		glog.Warningln("Docker network name not specified")
	}
	resp.Name = req.Name

	var errorList []error
	resp.NetworkCreate, errorList = req.NetworkCreate.DeepCopyChecked()

	msgBuf := bytes.Buffer{}
	for _, err := range errorList {
		switch {
		case err == networkpb.ErrNilPointer:
			glog.Warningf("found nil pointer: %v", err)
		case err == mobypb.ErrNetworkCreateIsNil:
			glog.Warningln("Network creating args not specified")
		case err == networkpb.ErrIpamIsNil:
			glog.Warningln("Networking IPAM arguments not specified")
		case err == networkpb.ErrIPAMConfigListIsEmpty:
			glog.Warningln("Networking arguments of IPAM subnet or gateway not specified")
		default:
			glog.Warningf("Unknown error: %v", err)
			msgBuf.WriteString(err.Error())
			msgBuf.WriteString("; ")
		}
	}
	if msgBuf.Len() != 0 {
		msgBuf.Truncate(msgBuf.Len() - 2)
		return resp, errors.New(msgBuf.String())
	}
	return resp, nil
}

func (m *DockerNetworkInspectReqResp) DeepCopyCheckedArgs() (*DockerNetworkInspectReqResp, error) {
	obj := new(DockerNetworkInspectReqResp)
	if m == nil {
		return obj, fmt.Errorf("Docker network inspect request required")
	}
	errorList := []error{}
	obj = &DockerNetworkInspectReqResp{
		Id:      m.Id,
		Name:    m.Name,
		KeyType: m.KeyType,
	}

	switch {
	case m.KeyType == DockerNetworkInspectReqResp_ID:
		if len(m.Id) == 0 {
			errorList = append(errorList, errors.New("network id required"))
		}
	case m.KeyType == DockerNetworkInspectReqResp_NAME:
		if len(m.Name) == 0 {
			errorList = append(errorList, errors.New("network name required"))
		}
	default:
		errorList = append(errorList, errors.New("unknown key type, either 1 (ID) or 2(NAME) required"))
	}

	obj.NetworkInspectOptions = m.NetworkInspectOptions.DeepCopyChecked()

	if len(errorList) != 0 {
		return obj, fmt.Errorf("args checked errors: %v", errorList[0])
	}
	return obj, nil
}

func (m *DockerNetworkListReqResp) DeepCopyCheckedArgs() (*DockerNetworkListReqResp, error) {
	obj := new(DockerNetworkListReqResp)
	if obj != nil {
		obj.NetworkListOptions = m.NetworkListOptions.DeepCopyChecked()
		obj.NetworkResources = make([]*mobypb.NetworkResource, 0)
	}
	return obj, nil
}

func (m *DockerNetworkRemoveReqResp) DeepCopyCheckedArgs() (*DockerNetworkRemoveReqResp, error) {
	obj := new(DockerNetworkRemoveReqResp)
	if m == nil {
		return obj, fmt.Errorf("Docker network remove request required")
	}
	errorList := []error{}
	obj = &DockerNetworkRemoveReqResp{
		Id:      m.Id,
		Name:    m.Name,
		KeyType: m.KeyType,
	}

	switch {
	case m.KeyType == DockerNetworkRemoveReqResp_ID:
		if len(m.Id) == 0 {
			errorList = append(errorList, errors.New("network id required"))
		}
	case m.KeyType != DockerNetworkRemoveReqResp_NAME:
		if len(m.Name) == 0 {
			errorList = append(errorList, errors.New("network name required"))
		}
	default:
		errorList = append(errorList, errors.New("unknown key type, either 1(ID) or 2(NAME) is required"))
	}
	if len(errorList) != 0 {
		return obj, fmt.Errorf("args checked errors: %v", errorList[0])
	}
	return obj, nil
}

func (m *DockerNetworkPruneReqResp) DeepCopyCheckedArgs() (*DockerNetworkPruneReqResp, error) {
	obj := new(DockerNetworkPruneReqResp)
	if m != nil {
		obj.Filters = m.Filters.DeepCopyChecked()
		obj.NetworksPruneReport = (*mobypb.NetworksPruneReport)(nil)
	}
	return obj, nil
}
