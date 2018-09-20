package api

import (
	// "fmt"

	dockertypes "github.com/docker/docker/api/types"
)

type SimpleDockerBuildContext struct {
	DockerfileContent         string `json:"dockerfile_content,omitempty"`
	*SimpleSourceRepoMetadata `json:",omitempty"`
}

type SimpleSourceRepoMetadata struct {
	GitReopDescription  *string `json:"git_repo_description,omitempty"`
	FilePathDescription *string `json:"file_path_description,omitempty"`
}

type GitRepoConfig struct {
	Credentials         []string `json:"credentials,omitempty"`
	URL                 string   `json:"url"`           // https://github.com/... or or git://github.com/... or http://<your>/<git>/<server>/<repo>/...
	REF                 string   `json:"ref,omitempty"` // master, branch, tag, or commit
	RecursiveSubmodules bool     `json:"recursive_submodules,omitempty"`
	AllBranches         bool     `json"all_branchs,omitempty"`
	AllTags             bool     `json:"-"`
	ForceRemote         bool     `json:"-"` // Force git clone regardless already done
}

type FilePathConfig struct {
	Directory string `json:"directory"`
}

type HttpSourceConfig struct {
	Credentials []string `json:"credentials,omitempty"`
	URL         string   `json:"url"`
}

type ArchiveRepoConfig struct {
	FilePathSource *FilePathConfig   `json:"file_path_source,omitempty"`
	HttpSource     *HttpSourceConfig `json:"http_source,omitempty"`
	WebdavSource   *HttpSourceConfig `json:"webdav_source,omitemtpy""`
	RestfulSource  *HttpSourceConfig `json:"restful_source,omitempty"`
}

type ImageBuildConfig struct {
	GitSource         *GitRepoConfig     `json:"git_source,omitempty"`
	RestfulSource     *HttpSourceConfig  `json:"-"`
	WebdavSource      *HttpSourceConfig  `json:"-"`
	ArchiveSource     *ArchiveRepoConfig `json:"arcive_source,omitempty"`
	FilePathSource    *FilePathConfig    `json:"file_path_source,omitempty"`
	CephRadosSource   *string            `json:"-"`
	SwiftSource       *string            `json:"-"`
	S3Source          *string            `json:"-"`
	AzureSource       *string            `json:"-"`
	GoogleDriveSource *string            `json:"-"`
	AliOssSource      *string            `json:"-"`
	TencentCosSource  *string            `json:"-"`
	BaiduBosSource    *string            `json:"-"`

	ProjectMetadata *ProjectMetadata         `json:"project_metadata,omitempty"`
	DockerSrcBldCfg *DockerSourceBuildConfig `json:"docker_source_build_config,omitempty"`
	DockerImgBldCfg *DockerImageBuildConfig  `json:"docker_image_build_config",omitempty`
	// VirtualMachineImage
	// OpenshiftS2i *OpenshiftS2iBuildConfig

	User string `json:"user"`
}

func (m *ImageBuildConfig) GitRecursiveSubmodules() bool {
	if m != nil {
		if m.GitSource != nil {
			return m.GitSource.RecursiveSubmodules
		}
	}
	return false
}

func (m *ImageBuildConfig) GitRef() string {
	if m != nil {
		if m.GitSource != nil {
			return m.GitSource.REF
		}
	}
	return ""
}

func (m *ImageBuildConfig) GitForceRemote() bool {
	if m != nil {
		if m.GitSource == nil {
			return m.GitSource.ForceRemote
		}
	}
	return false
}

//func (m *ImageBuildConfig) ProjectDir() string {
//	if m != nil {
//		if m.ProjectMetadata != nil {
//			return m.ProjectMetadata.HomeDir
//		}
//	}
//	return ""
//}

//func (m *ImageBuildConfig) ImgRef() string {
//	if m != nil {
//		if m.DockerImageBuildConfig != nil {
//			return m.DockerImageBuildConfig.ImageBuildConfig.ImgRef
//		}
//	}
//	return ""
//}

type NameValuePair struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ProjectMetadata struct {
	Name                  string          `json:"name,omitempty"`
	SourceBuildCmd        []string        `json:"source_build_cmd,omitempty"`
	SourceBuildArgs       []string        `json:"source_build_args,omitempty"`
	SourceBuildEnv        []NameValuePair `json:"source_build_env,omitempty"`
	SourceBuildWorkingDir string          `json:"source_build_working_dir,omitempty"` // relative to source context dir

	DockerfilePath string `json:"dockerfile_path,omitempty"` // Project have already create Dockerfile
}

type DockerSourceBuildConfig struct {
	ContainerCreateConfig *dockertypes.ContainerCreateConfig `json:"container_create_config"` // to build destination image, otherwise use MicroServiceBldCfg
	ImgRef                string                             `json:"img_ref"`                 // override

}

type DockerImageBuildConfig struct {
	FromImgRef string `json:"from_img_ref,omitempty"` // override From directive into Dockerfile

	ImageBuildOptions *dockertypes.ImageBuildOptions `json:"image_pull_options,omitempty"`
	ImgRef            string                         `json:"img_ref,omitempty"` // override into dockertypes.ImageBuildOptions

	BuildContextDir  string `json:"build_context_dir,omitempty"`
	IncludeBldCtxDir bool   `json:"include_build_context_dir,omitempty"`

	MicroServiceBldCfg *MicroServiceDockerBuildConfig `json:"micro_service_build_config,omitempty"` // Dedicated build
}

type MicroServiceDockerBuildConfig struct {
	FileServerBldCfg *FileServerDockerBuildConfig `json:"file_server_build_config,omitempty"`
	PhpWebBldCfg     *PhpWebDockerBuildConfig     `json:"php_web_build_config,omitempty"`
}

type FileServerDockerBuildConfig struct {
	BasedGofileserverBldCfg struct {
		dir  string // default is /tmp/download
		addr string // default is :48080
	} `json:"based_gofileserver_build_config"`
}

type PhpWebDockerBuildConfig struct {
	BasedAlpinePhpBuiltinBldCfg struct {
		dir  string
		addr string
	}
}

func BasedGofileserverDockerBuildConfig(imgref string, ctxdir string, includectxdir bool) *DockerImageBuildConfig {
	return &DockerImageBuildConfig{
		ImgRef:           imgref,
		BuildContextDir:  ctxdir,
		IncludeBldCtxDir: includectxdir,
		MicroServiceBldCfg: &MicroServiceDockerBuildConfig{
			FileServerBldCfg: &FileServerDockerBuildConfig{},
		},
	}
}

func DockerImageBuildConfigWithGitRepo(project *ProjectMetadata, dockerBldCfg *DockerImageBuildConfig, gitSource *GitRepoConfig) *ImageBuildConfig {
	return &ImageBuildConfig{
		GitSource:       gitSource,
		ProjectMetadata: project,
		DockerImgBldCfg: dockerBldCfg,
	}
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

	GitSource             *GitRepoConfig
	SourceBuildFromImgRef string
	TargetImgRef          string
}

func (m *OpenshiftS2iBuildConfig) ExportCmdArgsAsOpenshiftS2iBuild() (dst []string) {
	dst = []string{}
	if m != nil {
		if m.GitSource != nil {
			if len(m.GitSource.URL) != 0 {
				dst = append(dst, m.GitSource.URL)
				if len(m.SourceBuildFromImgRef) != 0 {
					dst = append(dst, m.SourceBuildFromImgRef)
					if len(m.TargetImgRef) != 0 {
						dst = append(dst, m.TargetImgRef)
					}
				}
			}
		}
	}
	return
}
