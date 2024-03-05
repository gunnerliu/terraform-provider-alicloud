package alicloud

import (
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/gunnerliu/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

type OceanServiceTenantUser struct {
	client *connectivity.AliyunClient
}

func (s *OceanServiceTenantUser) DescribeOceanTenantUser(tenantId string, userName string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	action := "DescribeTenantUsers"
	conn, err := client.NewOceanbaseClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["TenantId"] = tenantId
	request["UserName"] = userName

	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(1*time.Minute, func() *resource.RetryError {
		response, err = conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-09-01"), StringPointer("AK"), query, request, &runtime)

		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		addDebug(action, response, request)
		return nil
	})

	if err != nil {
		if IsExpectedErrors(err, []string{"IllegalOperation.Resource", "UnknownError"}) {
			return object, WrapErrorf(Error(GetNotFoundMessage("TenantUser", tenantId+"-"+userName)), NotFoundMsg, response)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, tenantId+"-"+userName, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.TenantUsers[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, userName, "$.TenantUsers[*]", response)
	}

	if len(v.([]interface{})) == 0 {
		return make(map[string]interface{}), nil
	}

	return v.([]interface{})[0].(map[string]interface{}), nil
}
