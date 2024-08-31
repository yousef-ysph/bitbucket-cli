package bitbucketapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"slices"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Password string `json:"password"`
	User     string `json:"user"`
	Token    string `json:"token"`
}

func GetConfig() (Config, error) {
	homeDirectory, err := os.UserHomeDir()
	configuration := Config{}
	if err != nil {
		log.Fatal(err)
		return configuration, err
	}

	file, fileErr := os.Open(homeDirectory + "/.bitbucketcmd.json")
	defer file.Close()
	if fileErr != nil {

		fmt.Println("error:", fileErr)
		return configuration, fileErr
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
		return configuration, err
	}
	return configuration, nil
}

type Pipelines struct {
	PullRequests map[string]any `yaml:"pull-requests"`
	Custom       map[string]any `yaml:"custom"`
	Tags         map[string]any `yaml:"tags"`
	Branches     map[string]any `yaml:"branches"`
}
type BitbucketFileConfig struct {
	Image      any            `yaml:"image"`
	Defenition map[string]any `yaml:"definitions"`
	Pipelines  Pipelines      `yaml:"pipelines"`
}

func GetPipelineNames(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

	file, fileErr := os.Open("bitbucket-pipelines.yml")
	var configFile BitbucketFileConfig

	pipelines := []string{}
	defer file.Close()
	if fileErr != nil {

		return pipelines, cobra.ShellCompDirectiveError

	}

	yamlDecoder := yaml.NewDecoder(file)
	err := yamlDecoder.Decode(&configFile)
	if err != nil {
		return pipelines, cobra.ShellCompDirectiveError

	}
	for key, _ := range configFile.Pipelines.Custom {
		pipelines = slices.Insert(pipelines, len(pipelines), "custom:"+key)
		pipelines = slices.Insert(pipelines, len(pipelines), key)

	}
	for key, _ := range configFile.Pipelines.PullRequests {
		pipelines = slices.Insert(pipelines, len(pipelines), "pull-requests:"+key)
	}
	for key, _ := range configFile.Pipelines.Branches {
		pipelines = slices.Insert(pipelines, len(pipelines), "branches:"+key)
	}
	for key, _ := range configFile.Pipelines.Tags {
		pipelines = slices.Insert(pipelines, len(pipelines), "tags:"+key)
	}
	return pipelines, cobra.ShellCompDirectiveDefault

}

func HttpRequestWithBitbucketAuth(method string, path string, data any, contentType string) (http.Response, error) {
	config, configErr := GetConfig()
	if configErr != nil {
		log.Fatal(configErr)
		return http.Response{}, configErr
	}
	if config.Token == "" && (config.Password == "" || config.User == "") {
		log.Fatal("config file is not correct")
		return http.Response{}, nil
	}
	authorizationToken := ""
	if config.Token != "" {
		authorizationToken = "Bearer " + config.Token
	} else {
		authorizationToken = "Basic " + base64.StdEncoding.EncodeToString([]byte(config.User+":"+config.Password))
	}
	var postBody []byte
	if reflect.TypeOf(data).String() == "[]uint8" {
		postBody = data.([]byte)
	} else {
		var jsonErr error = nil
		postBody, jsonErr = json.Marshal(data)
		if jsonErr != nil {
			log.Fatal(jsonErr)
			return http.Response{}, jsonErr
		}

	}
	req, err := http.NewRequest(method, "https://api.bitbucket.org/2.0/repositories/"+path, bytes.NewBuffer(postBody))
	if err != nil {
		log.Fatal(err)
		return http.Response{}, err
	}
	req.Header.Set("Authorization", authorizationToken)
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("Accept", contentType)
	}
	client := &http.Client{}
	response, responseErr := client.Do(req)
	return *response, responseErr
}
func HttpRequestWithBitbucketAuthJson(method string, path string, data any) (http.Response, error) {
	return HttpRequestWithBitbucketAuth(method, path, data, "application/json")
}

type FetchToChanResponse struct {
	Resp http.Response
	Err  error
}

func FetchToChannel(method string, url string, data any, contentType string, c chan FetchToChanResponse) {
	var fetchToChannel FetchToChanResponse
	fetchToChannel.Resp, fetchToChannel.Err = HttpRequestWithBitbucketAuthJson("GET", url, map[string]string{})
	c <- fetchToChannel
}
func FetchToChannelJson(method string, url string, data any, c chan FetchToChanResponse) {

	FetchToChannel(method, url, data, "application/json", c)
}
