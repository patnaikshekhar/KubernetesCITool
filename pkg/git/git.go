package git

import (
	"fmt"

	pb "github.com/patnaikshekhar/kubernetescitool/interface"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"gopkg.in/yaml.v2"
)

// GetBuildDefinitionFromRepo gets a build definition from a repository
func GetBuildDefinitionFromRepo(githubHookData *GitHubPush) (*pb.BuildRequest, error) {
	var result *pb.BuildRequest

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: githubHookData.Repository.CloneURL,
	})

	if err != nil {
		return nil, err
	}

	ref, err := r.Head()
	if err != nil {
		return nil, err
	}

	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	files, err := commit.Files()
	if err != nil {
		return nil, err
	}

	err = files.ForEach(func(f *object.File) error {
		if f.Name == "build.yaml" {
			contents, err := f.Contents()
			if err != nil {
				return err
			}
			err = yaml.UnmarshalStrict([]byte(contents), &result)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, fmt.Errorf("No build.xml detected")
	}

	return result, nil
}
