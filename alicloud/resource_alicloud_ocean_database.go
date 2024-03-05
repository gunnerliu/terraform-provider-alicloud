package alicloud

import (
	"fmt"
	"log"
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/gunnerliu/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAliCloudOceanDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudOceanDatabaseCreate,
		Read:   resourceAliCloudOceanDatabaseRead,
		Update: resourceAliCloudOceanBaseDatabaseUpdate,
		Delete: resourceAliCloudOceanDatabaseDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(80 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"database_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"encoding": {
				Type:     schema.TypeString,
				Required: true,
			},
			"collation": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudOceanDatabaseCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateDatabase"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	conn, err := client.NewOceanbaseClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	// 参数赋值
	request["DatabaseName"] = d.Get("database_name")
	request["TenantId"] = d.Get("tenant_id")
	request["Encoding"] = d.Get("encoding")
	request["InstanceId"] = d.Get("instance_id")
	if v, ok := d.GetOk("collation"); ok {
		request["Collation"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	// 发送请求
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ocean_data_base", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.DatabaseName", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudOceanBaseDatabaseUpdate(d, meta)
}

func resourceAliCloudOceanDatabaseRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	oceanServiceDatabase := OceanServiceDatabase{client}

	objectRaw, err := oceanServiceDatabase.DescribeOceanDatabase(d.Get("tenant_id").(string), d.Get("instance_id").(string), d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ocean_database DescribeDatabase Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("user_type", objectRaw["Status"])
	d.Set("description", objectRaw["Description"])
	d.Set("encoding", objectRaw["Encoding"])
	d.Set("db_type", objectRaw["DbType"])
	d.Set("database_name", objectRaw["DatabaseName"])
	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("tenant_id", objectRaw["TenantId"])
	d.Set("required_size", objectRaw["RequiredSize"])
	d.Set("data_size", objectRaw["DataSize"])
	d.Set("collation", objectRaw["Collation"])
	d.Set("instance_id", objectRaw["InstanceId"])
	d.Set("tenant_name", objectRaw["TenantName"])

	return nil
}

func resourceAliCloudOceanBaseDatabaseUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceAliCloudOceanDatabaseRead(d, meta)
}

func resourceAliCloudOceanDatabaseDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteDatabases"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	conn, err := client.NewOceanbaseClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["TenantId"] = d.Get("tenant_id")
	request["InstanceId"] = d.Get("instance_id")
	request["DatabaseNames"] = fmt.Sprintf("[\"%s\"]", d.Id())

	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 5*time.Second)
	err = resource.Retry(d.Timeout(schema.TimeoutDelete), func() *resource.RetryError {
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
		if IsExpectedErrors(err, []string{"UnknownError", "IllegalOperation.Resource"}) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
	}

	return nil
}
