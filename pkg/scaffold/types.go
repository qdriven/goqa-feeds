package scaffold

type ProjectSetting struct {
	ProjectType        string            `mapstructure: projectType`
	ProjectTemplateUrl string            `mapstructure: projectTemplateUrl`
	Addl               map[string]string `mapstructure: addl`
}
