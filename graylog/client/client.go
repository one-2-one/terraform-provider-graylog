package client

import (
	"net/http"

	"github.com/suzuki-shunsuke/go-httpclient/httpclient"

	authzSharesEntities "github.com/one-2-one/terraform-provider-graylog/graylog/client/authz/shares/entities"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/dashboard"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/dashboard/position"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/dashboard/widget"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/event/definition"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/event/notification"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/role"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/sidecar"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/sidecar/collector"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/sidecar/configuration"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/stream"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/stream/alarmcallback"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/stream/alert/condition"
	streamOutput "github.com/one-2-one/terraform-provider-graylog/graylog/client/stream/output"
	streamRule "github.com/one-2-one/terraform-provider-graylog/graylog/client/stream/rule"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/system/grok"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/system/indices/indexset"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/system/input"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/system/input/extractor"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/system/input/staticfield"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/system/ldap/setting"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/system/output"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/system/pipeline/connection"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/system/pipeline/pipeline"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/system/pipeline/rule"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/user"
	userOutput "github.com/one-2-one/terraform-provider-graylog/graylog/client/user/output"
	userToken "github.com/one-2-one/terraform-provider-graylog/graylog/client/user/token"
	"github.com/one-2-one/terraform-provider-graylog/graylog/client/view"
	"github.com/one-2-one/terraform-provider-graylog/graylog/config"
)

type Client struct {
	APIVersion              string
	AuthzSharesEntities     authzSharesEntities.Client
	AlarmCallback           alarmcallback.Client
	AlertCondition          condition.Client
	Collector               collector.Client
	Dashboard               dashboard.Client
	DashboardWidget         widget.Client
	DashboardWidgetPosition position.Client
	EventDefinition         definition.Client
	EventNotification       notification.Client
	Extractor               extractor.Client
	Grok                    grok.Client
	IndexSet                indexset.Client
	Input                   input.Client
	InputStaticField        staticfield.Client
	LDAPSetting             setting.Client
	Output                  output.Client
	Pipeline                pipeline.Client
	PipelineConnection      connection.Client
	PipelineRule            rule.Client
	Role                    role.Client
	Sidecar                 sidecar.Client
	SidecarConfiguration    configuration.Client
	Stream                  stream.Client
	StreamOutput            streamOutput.Client
	StreamRule              streamRule.Client
	View                    view.Client
	User                    user.Client
	UserToken               userToken.Client
	UserOutput              userOutput.Client
}

func New(m interface{}) (Client, error) {
	cfg := m.(config.Config)

	httpClient := httpclient.New(cfg.Endpoint)
	xRequestedBy := cfg.XRequestedBy
	if xRequestedBy == "" {
		xRequestedBy = "terraform-provider-graylog"
	}
	httpClient.SetRequest = func(req *http.Request) error {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Requested-By", xRequestedBy)
		req.SetBasicAuth(cfg.AuthName, cfg.AuthPassword)
		return nil
	}

	return Client{
		APIVersion: cfg.APIVersion,
		AlarmCallback: alarmcallback.Client{
			Client: httpClient,
		},
		AlertCondition: condition.Client{
			Client: httpClient,
		},
		Collector: collector.Client{
			Client: httpClient,
		},
		Dashboard: dashboard.Client{
			Client: httpClient,
		},
		DashboardWidget: widget.Client{
			Client: httpClient,
		},
		DashboardWidgetPosition: position.Client{
			Client: httpClient,
		},
		EventDefinition: definition.Client{
			Client: httpClient,
		},
		EventNotification: notification.Client{
			Client: httpClient,
		},
		Extractor: extractor.Client{
			Client: httpClient,
		},
		Grok: grok.Client{
			Client: httpClient,
		},
		IndexSet: indexset.Client{
			Client: httpClient,
		},
		Input: input.Client{
			Client: httpClient,
		},
		InputStaticField: staticfield.Client{
			Client: httpClient,
		},
		LDAPSetting: setting.Client{
			Client: httpClient,
		},
		Output: output.Client{
			Client: httpClient,
		},
		Pipeline: pipeline.Client{
			Client: httpClient,
		},
		PipelineConnection: connection.Client{
			Client: httpClient,
		},
		PipelineRule: rule.Client{
			Client: httpClient,
		},
		Role: role.Client{
			Client: httpClient,
		},
		Sidecar: sidecar.Client{
			Client: httpClient,
		},
		SidecarConfiguration: configuration.Client{
			Client: httpClient,
		},
		Stream: stream.Client{
			Client: httpClient,
		},
		StreamOutput: streamOutput.Client{
			Client: httpClient,
		},
		StreamRule: streamRule.Client{
			Client: httpClient,
		},
		View: view.Client{
			Client: httpClient,
		},
		User: user.Client{
			Client: httpClient,
		},
		UserToken: userToken.Client{
			Client: httpClient,
		},
		UserOutput: userOutput.Client{
			Client: httpClient,
		},
		AuthzSharesEntities: authzSharesEntities.Client{
			Client: httpClient,
		},
	}, nil
}
