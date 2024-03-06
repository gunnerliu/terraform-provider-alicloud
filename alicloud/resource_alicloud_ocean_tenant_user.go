// Package alicloud. This file is generated automatically. Please do not modify it manually, thank you!
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

func resourceAliCloudOceanTenantUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudOceanTenantUserCreate,
		Read:   resourceAliCloudOceanTenantUserRead,
		Update: resourceAliCloudOceanBaseTenantUserUpdate,
		Delete: resourceAliCloudOceanTenantUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(80 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringInSlice([]string{"Normal", "Admin", "ReadonlyAccount"}, false),
			},
			"user_password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"roles": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudOceanTenantUserCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateTenantUser"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	conn, err := client.NewOceanbaseClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	// 参数赋值
	request["InstanceId"] = d.Get("instance_id")
	request["TenantId"] = d.Get("tenant_id")
	request["UserName"] = d.Get("user_name")
	request["UserType"] = d.Get("user_type")
	request["UserPassword"] = d.Get("user_password")
	if v, ok := d.GetOk("roles"); ok {
		request["Roles"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("encryption_type"); ok {
		request["EncryptionType"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ocean_tenant_user", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.TenantUser.UserName", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudOceanBaseTenantUserUpdate(d, meta)
}

func resourceAliCloudOceanTenantUserRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	oceanServiceTenantUser := OceanServiceTenantUser{client}

	objectRaw, err := oceanServiceTenantUser.DescribeOceanTenantUser(d.Get("tenant_id").(string), d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ocean_tenant_user DescribeTenantUser Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("user_type", objectRaw["UserType"])
	d.Set("description", objectRaw["Description"])
	d.Set("user_status", objectRaw["UserStatus"])
	d.Set("user_name", objectRaw["UserName"])
	d.Set("tenant_id", objectRaw["TenantId"])
	d.Set("instance_id", objectRaw["InstanceId"])

	return nil
}

func resourceAliCloudOceanBaseTenantUserUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)
	action := "ModifyTenantUserPassword"
	conn, err := client.NewOceanbaseClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["UserName"] = d.Id()
	request["TenantId"] = d.Get("tenant_id")
	request["InstanceId"] = d.Get("instance_id")
	if !d.IsNewResource() {
		if d.HasChange("user_password") {
			update = true
			request["UserPassword"] = d.Get("user_password")
			d.SetPartial("user_password")
		}

		if update {
			runtime := util.RuntimeOptions{}
			runtime.SetAutoretry(true)
			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
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
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}
		}
	}

	update = false
	action = "ModifyTenantUserRoles"
	conn, err = client.NewOceanbaseClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["UserName"] = d.Id()
	request["TenantId"] = d.Get("tenant_id")
	request["InstanceId"] = d.Get("instance_id")

	if !d.IsNewResource() {
		if d.HasChange("roles") {
			update = true
			request["UserRole"] = d.Get("roles")
			d.SetPartial("roles")
		}

		if update {
			runtime := util.RuntimeOptions{}
			runtime.SetAutoretry(true)
			wait := incrementalWait(3*time.Second, 5*time.Second)
			err = resource.Retry(d.Timeout(schema.TimeoutUpdate), func() *resource.RetryError {
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
				return WrapErrorf(err, DefaultErrorMsg, d.Id(), action, AlibabaCloudSdkGoERROR)
			}
		}
	}

	d.Partial(false)
	return resourceAliCloudOceanTenantUserRead(d, meta)
}

func resourceAliCloudOceanTenantUserDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	action := "DeleteTenantUsers"
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
	request["Users"] = fmt.Sprintf("[\"%s\"]", d.Id())

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
