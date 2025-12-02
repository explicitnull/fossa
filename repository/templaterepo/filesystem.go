package templaterepo

import (
	"context"
	"errors"
	"fmt"
	"os"

	"fossa/service/asset"
	"fossa/service/template"
)

var ErrGenericTemplateNotFound = fmt.Errorf("generic template not found")

type FileSystem struct {
	templatesPath string
}

func NewFileSystem(templatesPath string) *FileSystem {
	return &FileSystem{
		templatesPath: templatesPath,
	}
}

func (s *FileSystem) FetchTemplatesByJobType(ctx context.Context, jobType string) ([]template.Template, error) {

	results := []template.Template{}

	for step := range asset.Steps {
		jobFilename := s.templatesPath + "/" + jobType + "/" + step + ".j2"
		// fmt.Printf("@@@@@ filename: %s\n", jobFilename)

		templateExists, err := s.isFileExist(jobFilename)
		if err != nil {
			return nil, err
		}

		// var templateExists bool

		// _, err := os.Stat(filename)
		// if !errors.Is(err, fs.ErrNotExist)
		// 	 = true
		// }

		var fileBytes []byte

		if templateExists {
			fileBytes, err = os.ReadFile(jobFilename)
			if err != nil {
				return nil, err
			}
		} else {
			// use generic template as fallback
			genericFilename := s.templatesPath + "/generic/" + step + ".j2"

			fileBytes, err = os.ReadFile(genericFilename)
			if err != nil {
				return nil, err
			}
		}

		t := template.Template{
			JobType:             jobType,
			Step:                step,
			Content:             string(fileBytes),
			GenericTemplateUsed: !templateExists,
		}

		results = append(results, t)

		// fmt.Printf("+++++ loaded template for step %s with len %d\n", step, len(fileBytes))

	}

	// mock data
	// t := template.Template{
	// 	ID:      "1",
	// 	JobType: "install_optics_to_connect_two_cenic_devices",
	// 	Step:    "installation",
	// 	Content: "example-content:\n A device: {{ device_a_port }}",
	// }

	// results = append(results, t)

	return results, nil
}

func (s *FileSystem) isFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return true, err // return true to avoid overwrite
}
