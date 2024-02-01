package adf

import (
	"fmt"
	"io"
	"os"

	"github.com/gofrontier-com/pony-express/pkg/core/adf"
	"gopkg.in/yaml.v2"
)

func LoadConfig(configFilePath string) (*adf.AppADFConfig, error) {
	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file's contents
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	cfg := adf.AppADFConfig{}

	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &cfg, nil
}

func hasWildcard(strArr []string) bool {
	for _, s := range strArr {
		if s == "*" {
			return true
		}
	}
	return false
}

func pipelineExists(pipelineName string, pipelines []*map[string]interface{}) bool {
	for _, pipeline := range pipelines {
		if (*pipeline)["name"] == pipelineName {
			return true
		}
	}
	return false
}

func VerifyPipelinesToDeploy(pipelineCfg []string, pipelines []*map[string]interface{}) (bool, error) {
	if hasWildcard(pipelineCfg) {
		return true, nil
	}

	for _, p := range pipelineCfg {
		if !pipelineExists(p, pipelines) {
			return false, fmt.Errorf("Pipeline '%s' is requested for deployment but does not exist", p)
		}
	}

	return true, nil
}
