package main

import (
	"github.com/gunnerliu/terraform-provider-alicloud/alicloud"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: alicloud.Provider})
}
