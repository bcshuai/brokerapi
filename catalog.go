package brokerapi

type IMetadataProvider interface {
	ProvideMetadata() IMetadata
}
type IMetadata interface {
	metadata()
}

type Service struct {
	ID              string                  `json:"id"`
	Name            string                  `json:"name"`
	Description     string                  `json:"description"`
	Bindable        bool                    `json:"bindable"`
	Tags            []string                `json:"tags,omitempty"`
	PlanUpdatable   bool                    `json:"plan_updateable"`
	Plans           []IMetadataProvider     `json:"plans"`   //ServicePlan
	Requires        []RequiredPermission    `json:"requires,omitempty"`
	Metadata        IMetadata               `json:"metadata,omitempty"`   //ServicePlanMetadata
	DashboardClient *ServiceDashboardClient `json:"dashboard_client,omitempty"`
}
func (service Service) ProvideMetadata() IMetadata{
	return service.Metadata
}

type ServiceDashboardClient struct {
	ID          string `json:"id"`
	Secret      string `json:"secret"`
	RedirectURI string `json:"redirect_uri"`
}

type ServicePlan struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Free        *bool                `json:"free,omitempty"`
	Metadata    IMetadata `json:"metadata,omitempty"`    //ServicePlanMetadata
}
func (plan ServicePlan) ProvideMetadata() IMetadata {
	return plan.Metadata
}

type ServicePlanMetadata struct {
	DisplayName string        `json:"displayName,omitempty"`
	Bullets     []string      `json:"bullets,omitempty"`
	Costs       []ServiceCost `json:"costs,omitempty"`
}
func (planMetaData ServicePlanMetadata) metadata() {

}

type ServiceCost struct {
	Amount map[string]float64 `json:"amount"`
	Unit   string             `json:"unit"`
}

type ServiceMetadata struct {
	DisplayName         string `json:"displayName,omitempty"`
	ImageUrl            string `json:"imageUrl,omitempty"`
	LongDescription     string `json:"longDescription,omitempty"`
	ProviderDisplayName string `json:"providerDisplayName,omitempty"`
	DocumentationUrl    string `json:"documentationUrl,omitempty"`
	SupportUrl          string `json:"supportUrl,omitempty"`
}
func (serviceMetadata ServiceMetadata) metadata() {}

func FreeValue(v bool) *bool {
	return &v
}

type RequiredPermission string

const (
	PermissionRouteForwarding = RequiredPermission("route_forwarding")
	PermissionSyslogDrain     = RequiredPermission("syslog_drain")
)
