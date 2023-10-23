package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/google/go-containerregistry/pkg/crane"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Targets []string `yaml:"targets"`
	Images  []struct {
		Name string `yaml:"name"`
		// Repo string   `yaml:"repo"`
		Tags []string `yaml:"tags"`
	} `yaml:"images"`
}

func main() {
	// Read the file
	data, err := ioutil.ReadFile("images.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	config := Config{}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	// fmt.Printf("%+v\n", config)

	for _, val := range config.Images {

		fmt.Println(val.Name)
		// fmt.Println(val.Repo)
		fmt.Println(val.Tags)
		fmt.Println("------")
	}

	// imagesToSync := Config{}

	// for each image.repo + "/" + image.name, get list of tags from target as existingTags
	for _, item := range config.Images {
		// existingTags, _ := getTags(image.Repo + "/" + image.Name)
		// imageName := fmt.Sprintf("%s/%s", item.Repo, item.Name)
		// get the existing tags in the target repo
		existingTags, _ := crane.ListTags(item.Name)
		fmt.Println(existingTags)
	}

	fmt.Println("testing...")
	getCacheRepoName(
		"ghcr.io/open-telemetry/opentelemetry-collector-releases/opentelemetry-collector-contrib",
		config.Targets[0],
	)

}

func getCacheRepoName(imageName string, targetRepo string) string {
	repoSplit := strings.Split(imageName, "/")
	repoSplitLen := len(repoSplit)
	var path string = targetRepo + "/" + repoSplit[repoSplitLen-1]
	fmt.Println(path)
	return path
}

// func getCacheRepoName(imageName string, imageRepo string, targetRepo string) {
// 	// remove the host name from image.Repo
// 	// concat target repo + updated image.Repo + image.Name
// 	var path string = targetRepo + "/"
// 	repoSplit := strings.Split(imageRepo, "/")
// 	for i, v := range repoSplit {
// 		if i > 0 {
// 			path = path + v + "/"
// 		}
// 	}
// 	path = path + imageName
// 	fmt.Println(path)
// }
