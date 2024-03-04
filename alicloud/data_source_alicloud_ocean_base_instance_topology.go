package alicloud

import (
	"time"

	"github.com/PaesslerAG/jsonpath"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/aliyun/terraform-provider-alicloud/alicloud/connectivity"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceAlicloudOceanBaseInstanceTopology() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlicloudOceanBaseInstanceTopologyRead,
		Schema: map[string]*schema.Schema{
			"instance_topology": {
				Optional: true,
				ForceNew: true,
				Type:     schema.TypeMap,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tenants": {
							Computed: true,
							Type:     schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tenant_id": {
										Computed: true,
										Type:     schema.TypeString,
									},
									"tenant_name": {
										Computed: true,
										Type:     schema.TypeString,
									},
									"tenant_cpu": {
										Computed: true,
										Type:     schema.TypeFloat,
									},
									"tenant_memory": {
										Computed: true,
										Type:     schema.TypeFloat,
									},
									"tenant_mode": {
										Computed: true,
										Type:     schema.TypeString,
									},
									"tenant_zones": {
										Computed: true,
										Type:     schema.TypeList,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"tenant_zone_role": {
													Computed: true,
													Type:     schema.TypeString,
												},
												"is_primary_tenant_zone": {
													Computed: true,
													Type:     schema.TypeString,
												},
												"tenant_zone_id": {
													Computed: true,
													Type:     schema.TypeString,
												},
												"units": {
													Computed: true,
													Type:     schema.TypeList,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"unit_id": {
																Computed: true,
																Type:     schema.TypeString,
															},
															"unit_status": {
																Computed: true,
																Type:     schema.TypeString,
															},
															"node_id": {
																Computed: true,
																Type:     schema.TypeString,
															},
															"unit_cpu": {
																Computed: true,
																Type:     schema.TypeFloat,
															},
															"unit_memory": {
																Computed: true,
																Type:     schema.TypeFloat,
															},
															"enable_migrate_unit": {
																Computed: true,
																Type:     schema.TypeBool,
															},
															"manual_migrate": {
																Computed: true,
																Type:     schema.TypeBool,
															},
															"enable_cancel_migrateUnit": {
																Computed: true,
																Type:     schema.TypeBool,
															},
															"unit_data_size": {
																Computed: true,
																Type:     schema.TypeInt,
															},
														},
													},
												},
											},
										},
									},
									"tenant_status": {
										Computed: true,
										Type:     schema.TypeString,
									},
									"tenant_deploy_type": {
										Computed: true,
										Type:     schema.TypeString,
									},
									"tenant_unit_num": {
										Computed: true,
										Type:     schema.TypeInt,
									},
									"primary_zone_deploy_type": {
										Computed: true,
										Type:     schema.TypeString,
									},
								},
							},
						},
						"zones": {
							Type: schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceAlicloudOceanBaseInstanceTopologyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	request := map[string]interface{}{
		"RegionId": client.RegionId,
	}

	if v, ok := d.GetOk("instance_id"); ok {
		request["InstanceId"] = v
	}

	conn, err := client.NewOceanbaseClient()
	if err != nil {
		return WrapError(err)
	}
	var instanceTopology map[string]interface{}
	var response map[string]interface{}

	action := "DescribeInstanceTopology"
	runtime := util.RuntimeOptions{}
	runtime.SetAutoretry(true)
	wait := incrementalWait(3*time.Second, 3*time.Second)
	err = resource.Retry(5*time.Minute, func() *resource.RetryError {
		resp, err := conn.DoRequest(StringPointer(action), nil, StringPointer("POST"), StringPointer("2019-09-01"), StringPointer("AK"), nil, request, &runtime)
		if err != nil {
			if NeedRetry(err) {
				wait()
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		response = resp
		addDebug(action, response, request)
		return nil
	})
	if err != nil {
		return WrapErrorf(err, DataDefaultErrorMsg, "alicloud_ocean_base_instance_topology", action, AlibabaCloudSdkGoERROR)
	}
	resp, err := jsonpath.Get("$.InstanceTopology", response)
	if err != nil {
		return WrapErrorf(err, FailedGetAttributeMsg, action, "$.InstanceTopology", response)
	}
	result, _ := resp.(map[string]interface{})
	instanceTopology = result

	mapping := map[string]interface{}{}
	instanceTopologyMapping := map[string]interface{}{}
	tenantMappingsMaps := make([]map[string]interface{}, 0)
	if tenantMappingsRaw, ok := instanceTopology["Tenants"]; ok {
		for _, value0 := range tenantMappingsRaw.([]interface{}) {
			tenantMappings := value0.(map[string]interface{})
			tenantMappingsMap := map[string]interface{}{
				"tenant_id":                tenantMappings["TenantId"],
				"tenant_name":              tenantMappings["TenantName"],
				"tenant_cpu":               tenantMappings["TenantCpu"],
				"tenant_memory":            tenantMappings["TenantMemory"],
				"tenant_mode":              tenantMappings["TenantMode"],
				"tenant_deploy_type":       tenantMappings["TenantDeployType"],
				"primary_zone_deploy_type": tenantMappings["PrimaryZoneDeployType"],
			}
			if v, ok := tenantMappings["TenantZones"]; ok && len(v.([]interface{})) > 0 {
				tenantZoneMappingsMaps := make([]map[string]interface{}, 0)
				for _, tenantZoneRow := range v.(map[string]interface{}) {
					tenantZoneMapping := tenantZoneRow.(map[string]interface{})
					tenantZoneMappingMap := map[string]interface{}{
						"tenant_zone_role": tenantZoneMapping["TenantZoneRole"],
						"tenant_zone_id":   tenantZoneMapping["TenantZoneId"],
					}
					if v0, ok := tenantZoneMapping["Units"]; ok && len(v0.([]interface{})) > 0 {
						unitMappingsMaps := make([]map[string]interface{}, 0)
						for _, unitRow := range v0.(map[string]interface{}) {
							unitMapping := unitRow.(map[string]interface{})
							unitMappingMap := map[string]interface{}{
								"unit_id":                    unitMapping["UnitId"],
								"unit_status":                unitMapping["UnitStatus"],
								"node_id":                    unitMapping["NodeId"],
								"unit_cpu":                   unitMapping["UnitCpu"],
								"unit_memory":                unitMapping["UnitMemory"],
								"enable_migrate_unit":        unitMapping["EnableMigrateUnit"],
								"manual_migrate":             unitMapping["ManualMigrate"],
								"enable_cancel_migrate_unit": unitMapping["EnableCancelMigrateUnit"],
								"unit_data_size":             unitMapping["UnitDataSize"],
							}
							unitMappingsMaps = append(unitMappingsMaps, unitMappingMap)
						}
						tenantZoneMappingMap["units"] = unitMappingsMaps
					}
					tenantZoneMappingsMaps = append(tenantZoneMappingsMaps, tenantZoneMappingMap)
				}
				tenantMappingsMap["tenant_zones"] = tenantZoneMappingsMaps
			}
			tenantMappingsMaps = append(tenantMappingsMaps, tenantMappingsMap)
		}
	}
	instanceTopologyMapping["tenants"] = tenantMappingsMaps
	mapping["instance_topology"] = instanceTopologyMapping

	if err := d.Set("instance_topology", mapping); err != nil {
		return WrapError(err)
	}

	if output, ok := d.GetOk("output_file"); ok && output.(string) != "" {
		writeToFile(output.(string), mapping)
	}
	return nil
}
