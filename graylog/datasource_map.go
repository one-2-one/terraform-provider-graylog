package graylog

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/one-2-one/terraform-provider-graylog/graylog/datasource/dashboard"
	"github.com/one-2-one/terraform-provider-graylog/graylog/datasource/sidecar"
	"github.com/one-2-one/terraform-provider-graylog/graylog/datasource/stream"
	"github.com/one-2-one/terraform-provider-graylog/graylog/datasource/system/indices/indexset"
	"github.com/one-2-one/terraform-provider-graylog/graylog/datasource/user"
)

var dataSourcesMap = map[string]*schema.Resource{
	"graylog_dashboard": dashboard.DataSource(),
	"graylog_index_set": indexset.DataSource(),
	"graylog_sidecar":   sidecar.DataSource(),
	"graylog_stream":    stream.DataSource(),
	"graylog_user":      user.DataSource(),
}
