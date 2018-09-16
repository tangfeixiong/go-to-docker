package api

import (
	"fmt"
)

type NameValuePair struct {
	Name  string
	Value string
}

type UidRange struct {
	From *int
	To   *int
}

type VolumeMount struct {
	Source      string
	Destination string
	Keep        bool
}

type DockerImageBuildMetadata struct {
	FileServer   *FileServerDockerBuildConfig
	OpenshiftS2i *OpenshiftS2iBuildConfig
}

func (m *DockerImageBuildMetadata) IsIgnoreGitSubmodules() bool {
	if m != nil {
		if m.FileServer != nil {
			if m.FileServer.GitSource != nil {
				return !m.FileServer.GitSource.RecursiveSubmodules
			}
		}
	}
	return false
}

func (m *DockerImageBuildMetadata) GitRef() string {
	if m != nil {
		if m.FileServer != nil {
			if m.FileServer.GitSource != nil {
				return m.FileServer.GitSource.REF
			}
		}
	}
	return ""
}

func (m *DockerImageBuildMetadata) GitProjectDir() string {
	if m != nil {
		if m.FileServer != nil {
			if m.FileServer.GitSource != nil {
				return m.FileServer.GitSource.ProjectPath
			}
		}
	}
	return ""
}

func (m *DockerImageBuildMetadata) DockerBaseImage() string {
	if m != nil {
		if m.FileServer != nil {
			return m.FileServer.ProductionImageDirectiveFrom()
		}
	}
	return ""
}

func (m *DockerImageBuildMetadata) GitForceRemote() bool {
	if m != nil {
		if m.FileServer != nil {
			if m.FileServer.GitSource == nil {
				return m.FileServer.GitSource.ForceRemote
			}
		}
	}
	return false
}

type FileServerDockerBuildConfig struct {
	GitSource         *GitRepoConfig
	ArchiveSource     *string // a tar, zip, gz represents a project
	HostPathSource    *string
	NfsSource         *string
	GlusterSource     *string
	CephRadosSource   *string
	SwiftSource       *string
	S3Source          *string
	AzureSource       *string
	GoogleDriveSource *string
	AliOssSource      *string
	TencentCosSource  *string
	BaiduBosSource    *string

	GitStaging *GitRepoConfig // if empty, use defualt GofileserverBasedBuildConfig

	GofileserverBasedBuildConfig struct {
		dir  string
		addr string
	}
}

const (
	gofileserver_image_ref = "docker.io/tangfeixiong/gofileserver:latest"
)

func GofileserverBasedBuildConfig() *FileServerDockerBuildConfig {
	return &FileServerDockerBuildConfig{}
}

func (m *FileServerDockerBuildConfig) ProductionImageDirectiveFrom() string {
	if m.GitStaging != nil {
		return m.GitStaging.ProductionDockerImgRef
	}
	if m.GitSource != nil {
		if len(m.GitSource.ProductionDockerfilePath) != 0 {
			fmt.Println("Must checkout after clone/pull into local")
			return ""
		}
	}
	return gofileserver_image_ref
}

func (m *FileServerDockerBuildConfig) ExportCmdArgsAsOpenshiftS2iBuild() (dst []string) {
	dst = []string{}
	if m != nil {
		if m.GitSource != nil {
			if len(m.GitSource.URL) != 0 {
				dst = append(dst, m.GitSource.URL)
				if len(m.GitSource.ProjectBuildBasedDockerImgRef) != 0 {
					dst = append(dst, m.GitSource.ProjectBuildBasedDockerImgRef)
					if len(m.GitSource.ProductionDockerImgRef) != 0 {
						dst = append(dst, m.GitSource.ProductionDockerImgRef)
					}
				}
			}
		}
	}
	return
}

type GitRepoConfig struct {
	Credentials                   string
	URL                           string // https://github.com/... or or git://github.com/... or http://<your>/<git>/<server>/<repo>/...
	REF                           string // master, branch, tag, or commit
	RecursiveSubmodules           bool
	ForceRemote                   bool // Force git clone/pull regardless already done
	ProjectPath                   string
	ProjectBuildDockerfilePath    string
	ProjectBuildCommand           []string
	ProjectBuildArgs              []string
	ProjectBuildEnv               []NameValuePair
	ProjectBuildBasedDockerImgRef string
	ProductionDockerfilePath      string
	ProductionDockerImgRef        string
}

type OpenshiftS2iBuildConfig struct {
	Env              []NameValuePair
	AssembleUser     string
	TarExcludeRegexp string
	ImageScriptsUrl  string
	ScriptsUrl       string
	DisplayName      string
	Description      string
	AllowedUids      []UidRange
	AssembleMounts   []VolumeMount
	Volumes          []string
	CapDrop          []string
	RuntimeMounts    []VolumeMount
	NetworkMode      string
	KeepSymlinks     bool

	Quiet                   bool
	Incremental             bool
	RemovePreviousImage     bool
	ImagePullPolicy         string
	PreviousImagePullPolicy string
	RuntimeImagePullPolicy  string
	PreserveTempDir         bool
	DockerConfigPath        string
	TarDestination          string
}
