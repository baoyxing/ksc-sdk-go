// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package knad

import (
	"github.com/baoyxing/ksc-sdk-go/ksc"
	"github.com/baoyxing/ksc-sdk-go/ksc/kscquery"
	"github.com/baoyxing/ksc-sdk-go/ksc/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/corehandlers"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
)

// Knad provides the API operation methods for making requests to
// knad. See this package's package overview docs
// for details on the service.
//
// Knad methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type Knad struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "knad"      // Name of service.
	EndpointsID = ServiceName // ID to lookup a service endpoint with.
	ServiceID   = "knad"      // ServiceID is a unique identifer of a specific service.
)

// New creates a new instance of the Knad client with a session.
// If additional configuration is needed for the client instance use the optional
// ksc.Config parameter to add your extra config.
//
// Example:
//
//	// Create a Knad client from just a session.
//	svc := knad.New(mySession)
//
//	// Create a Knad client with additional configuration
//	svc := knad.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Knad {
	c := p.ClientConfig(EndpointsID, cfgs...)
	c.Endpoint = utils.Url(&utils.UrlInfo{
		UseSSL:      false,
		Locate:      true,
		UseInternal: false,
	}, utils.ServiceInfo{
		Service: EndpointsID,
		Region:  c.SigningRegion,
	})
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// extraNew create int can support ssl or region locate set
func ExtraNew(info *utils.UrlInfo, p client.ConfigProvider, cfgs ...*aws.Config) *Knad {
	c := p.ClientConfig(EndpointsID, cfgs...)
	c.Endpoint = utils.Url(info, utils.ServiceInfo{
		Service: EndpointsID,
		Region:  c.SigningRegion,
	})
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// SdkNew create int can support ssl or region locate set
func SdkNew(p client.ConfigProvider, cfgs *ksc.Config, info ...*utils.UrlInfo) *Knad {
	_info := utils.UrlInfo{}
	if len(info) > 0 && len(info) == 1 {
		if info[0].UseSSL {
			_info.UseSSL = info[0].UseSSL
		}
		if info[0].Locate {
			_info.Locate = info[0].Locate
		}
		if info[0].UseInternal {
			_info.UseInternal = info[0].UseInternal
		}
		if info[0].CustomerDomain != "" {
			_info.CustomerDomain = info[0].CustomerDomain
		}
		if info[0].CustomerDomainIgnoreService {
			_info.CustomerDomainIgnoreService = info[0].CustomerDomainIgnoreService
		}

	}
	return ExtraNew(&_info, p, &aws.Config{Region: cfgs.Region})
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *Knad {
	svc := &Knad{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2023-03-23",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	// remove aws user-agent
	svc.Handlers.Build.Remove(corehandlers.SDKVersionUserAgentHandler)
	// add ksc user-agent
	svc.Handlers.Build.PushBackNamed(ksc.SDKVersionUserAgentHandler)
	svc.Handlers.Build.PushBackNamed(kscquery.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(kscquery.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(kscquery.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(kscquery.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Knad operation and runs any
// custom request initialization.
func (c *Knad) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
