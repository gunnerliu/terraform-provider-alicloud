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

func resourceAliCloudOceanTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceAliCloudOceanTenantCreate,
		Read:   resourceAliCloudOceanTenantRead,
		Update: resourceAliCloudOceanBaseTenantUpdate,
		Delete: resourceAliCloudOceanTenantDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(80 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"tenant_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tenant_mode": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: StringInSlice([]string{"Oracle", "MySQL"}, false),
			},
			"charset": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_vpc_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_v_switch_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"primary_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"unit_num": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"read_only_zone_list": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"log_disk": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceAliCloudOceanTenantCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)

	action := "CreateTenant"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	conn, err := client.NewOceanbaseClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})

	// 参数赋值
	request["TenantName"] = d.Get("tenant_name")
	request["Cpu"] = d.Get("cpu")
	request["Memory"] = d.Get("memory")
	request["TimeZone"] = d.Get("time_zone")
	request["TenantMode"] = d.Get("tenant_mode")
	request["Charset"] = d.Get("charset")
	request["InstanceId"] = d.Get("instance_id")
	request["UserVpcId"] = d.Get("user_vpc_id")
	request["UserVSwitchId"] = d.Get("user_v_switch_id")
	request["PrimaryZone"] = d.Get("primary_zone")
	if v, ok := d.GetOk("description"); ok {
		request["Description"] = v
	}
	if v, ok := d.GetOk("unit_num"); ok {
		request["UnitNum"] = v
	}
	if v, ok := d.GetOk("read_only_zone_list"); ok {
		request["ReadOnlyZoneList"] = v
	}
	if v, ok := d.GetOk("log_disk"); ok {
		request["LogDisk"] = v
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
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ocean_tenant", action, AlibabaCloudSdkGoERROR)
	}

	id, _ := jsonpath.Get("$.TenantId", response)
	d.SetId(fmt.Sprint(id))

	return resourceAliCloudOceanBaseTenantUpdate(d, meta)
}

func resourceAliCloudOceanTenantRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	oceanTenantService := OceanBaseServiceTenant{client}

	objectRaw, err := oceanTenantService.DescribeOceanTenant(d.Get("instance_id").(string), d.Id())
	if err != nil {
		if !d.IsNewResource() && NotFoundError(err) {
			log.Printf("[DEBUG] Resource alicloud_ocean_tenant DescribeOceanBaseInstance Failed!!! %s", err)
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	d.Set("create_time", objectRaw["CreateTime"])
	d.Set("tenant_id", objectRaw["TenantId"])
	d.Set("tenant_name", objectRaw["TenantName"])
	d.Set("tenant_mode", objectRaw["TenantMode"])
	d.Set("user_vpc_id", objectRaw["VpcId"])
	d.Set("status", objectRaw["Status"])
	d.Set("primary_zone", objectRaw["PrimaryZone"])
	d.Set("description", objectRaw["Description"])
	d.Set("charset", objectRaw["Charset"])
	d.Set("time_zone", objectRaw["TimeZone"])

	return nil
}

func resourceAliCloudOceanBaseTenantUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	var request map[string]interface{}
	var response map[string]interface{}
	var query map[string]interface{}
	update := false
	d.Partial(true)
	action := "ModifyTenantResource"
	conn, err := client.NewOceanbaseClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	query = make(map[string]interface{})
	request["TenantId"] = d.Id()
	if !d.IsNewResource() {
		if d.HasChange("cpu") {
			update = true
			request["Cpu"] = d.Get("cpu")
			d.SetPartial("cpu")
		}
		if d.HasChange("memory") {
			update = true
			request["Memory"] = d.Get("memory")
			d.SetPartial("memory")
		}
		if d.HasChange("instance_id") {
			update = true
			request["Memory"] = d.Get("instance_id")
			d.SetPartial("instance_id")
		}
		if d.HasChange("log_disk") {
			if v, ok := d.GetOk("log_disk"); ok {
				update = true
				request["LogDisk"] = v
				d.SetPartial("log_disk")
			}
		}
		if d.HasChange("read_only_zone_list") {
			if v, ok := d.GetOk("read_only_zone_list"); ok {
				update = true
				request["ReadOnlyZoneList"] = v
				d.SetPartial("read_only_zone_list")
			}
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
	return resourceAliCloudOceanTenantRead(d, meta)
}

func resourceAliCloudOceanTenantDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*connectivity.AliyunClient)
	action := "DeleteTenants"
	var request map[string]interface{}
	var response map[string]interface{}
	query := make(map[string]interface{})
	conn, err := client.NewOceanbaseClient()
	if err != nil {
		return WrapError(err)
	}
	request = make(map[string]interface{})
	request["TenantIds"] = fmt.Sprintf("[\"%s\"]", d.Id())
	request["InstanceId"] = d.Get("instance_id")

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
