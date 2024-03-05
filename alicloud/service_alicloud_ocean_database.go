package alicloud

import (
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/gunnerliu/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

type OceanServiceDatabase struct {
	client *connectivity.AliyunClient
}

func (s *OceanServiceDatabase) DescribeOceanDatabase(tenantId string, instanceId string, databaseName string) (object map[string]interface{}, err error) {
	client := s.client
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	action := "DescribeDatabases"
	conn, err := client.NewOceanbaseClient()
	if err != nil {
		return object, WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["TenantId"] = tenantId
	request["InstanceId"] = instanceId
	request["DatabaseName"] = databaseName

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
			return object, WrapErrorf(Error(GetNotFoundMessage("Database", tenantId+"-"+databaseName)), NotFoundMsg, response)
		}
		return object, WrapErrorf(err, DefaultErrorMsg, tenantId+"-"+databaseName, action, AlibabaCloudSdkGoERROR)
	}

	v, err := jsonpath.Get("$.Databases[*]", response)
	if err != nil {
		return object, WrapErrorf(err, FailedGetAttributeMsg, tenantId+"-"+databaseName, "$.Databases[*]", response)
	}

	if len(v.([]interface{})) == 0 {
		return object, WrapErrorf(Error(GetNotFoundMessage("Database", tenantId+"-"+databaseName)), NotFoundMsg, response)
	}

	return v.([]interface{})[0].(map[string]interface{}), nil
}
