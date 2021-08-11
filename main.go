package main

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	cdn "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v1/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cdn/v1/region"
	"os"
	"strings"
)

func main() {
	ak := os.Getenv("ACCESS_KEY_ID")
	sk := os.Getenv("SECRET_ACCESS_KEY")
	defaultType := os.Getenv("TYPE")
	reg := os.Getenv("REGION")
	urls := os.Getenv("URLS")
	enterpriseProjectId := os.Getenv("ENTERPRISE_PROJECT_ID")
	if reg == "" {
		reg = "cn-north-1"
	}
	if defaultType == "" {
		defaultType = "file"
	}
	auth := global.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		Build()

	client := cdn.NewCdnClient(
		cdn.CdnClientBuilder().
			WithRegion(region.ValueOf(reg)).
			WithCredential(auth).
			Build())

	request := &model.CreateRefreshTasksRequest{}
	if enterpriseProjectId != "" {
		enterpriseProjectIdRequest := "1"
		request.EnterpriseProjectId = &enterpriseProjectIdRequest
	}
	typeRefreshTaskRefreshTaskRequestBody := model.GetRefreshTaskRequestBodyTypeEnum().DIRECTORY
	if defaultType != "directory" {
		typeRefreshTaskRefreshTaskRequestBody = model.GetRefreshTaskRequestBodyTypeEnum().FILE
	}
	refreshTaskbody := &model.RefreshTaskRequestBody{
		Type: &typeRefreshTaskRefreshTaskRequestBody,
		Urls: strings.Split(urls, "|"),
	}
	request.Body = &model.RefreshTaskRequest{
		RefreshTask: refreshTaskbody,
	}
	response, err := client.CreateRefreshTasks(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
		os.Exit(2)
	}
}
