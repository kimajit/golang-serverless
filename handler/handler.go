package handler

import (
	"context"
	"encoding/json"
	sec "lambdafunction/helper/aws/secrets"
	param "lambdafunction/helper/circleci/parameter"
	ccitrigger "lambdafunction/helper/circleci/request"
	lambdalogger "lambdafunction/logger"
	"lambdafunction/models"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

var (
	logger               = lambdalogger.CreateLogger()
	client models.Client = new(http.Client)
	// unmarshaler models.JSONUnMarshaler
	secretsId = "dev/AWS/Setup"
)

func Handler(ctx context.Context, event events.CloudWatchEvent) error {
	logger.Println("Getting secrets from aws secrets manager")
	secrets, err := sec.GetAWSsecrets(secretsId)
	if err != nil {
		return err
	}
	url := secrets["baseUrl"] + secrets["vcs"] + "/" + secrets["org"] + "/" + secrets["pipelineName"] + "/pipeline"
	circleCiToken := secrets["circleCIToken"]
	var ECRscanEvent events.ECRScanEventDetailType
	jsonErr := json.Unmarshal(event.Detail, &ECRscanEvent)
	if jsonErr != nil {
		return jsonErr
	}
	repoName := ECRscanEvent.RepositoryName
	imageTag := ECRscanEvent.ImageTags[0]

	params, paramsErr := param.GetParams(secrets, imageTag, repoName)
	if paramsErr != nil {
		return paramsErr
	}
	logger.Println(url)
	logger.Println(params)
	logger.Println(circleCiToken)

	logger.Println("creating post Request for CircleCi pipeline")
	respStat, reqErr := ccitrigger.PostRequestToCircleCI(url, params, circleCiToken, client)
	if reqErr != nil {
		return reqErr
	}
	log.Println("triggered", respStat)
	return nil
}
