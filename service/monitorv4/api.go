// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package monitorv4

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opAddAlarmReceives = "AddAlarmReceives"

// AddAlarmReceivesRequest generates a "ksc/request.Request" representing the
// client's request for the AddAlarmReceives operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See AddAlarmReceives for more information on using the AddAlarmReceives
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the AddAlarmReceivesRequest method.
//    req, resp := client.AddAlarmReceivesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/AddAlarmReceives
func (c *Monitorv4) AddAlarmReceivesRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddAlarmReceives,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AddAlarmReceives API operation for monitorv4.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for monitorv4's
// API operation AddAlarmReceives for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/AddAlarmReceives
func (c *Monitorv4) AddAlarmReceives(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddAlarmReceivesRequest(input)
	return out, req.Send()
}

// AddAlarmReceivesWithContext is the same as AddAlarmReceives with the addition of
// the ability to pass a context and additional request options.
//
// See AddAlarmReceives for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Monitorv4) AddAlarmReceivesWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddAlarmReceivesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteAlarmReceives = "DeleteAlarmReceives"

// DeleteAlarmReceivesRequest generates a "ksc/request.Request" representing the
// client's request for the DeleteAlarmReceives operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DeleteAlarmReceives for more information on using the DeleteAlarmReceives
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DeleteAlarmReceivesRequest method.
//    req, resp := client.DeleteAlarmReceivesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/DeleteAlarmReceives
func (c *Monitorv4) DeleteAlarmReceivesRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteAlarmReceives,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteAlarmReceives API operation for monitorv4.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for monitorv4's
// API operation DeleteAlarmReceives for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/DeleteAlarmReceives
func (c *Monitorv4) DeleteAlarmReceives(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteAlarmReceivesRequest(input)
	return out, req.Send()
}

// DeleteAlarmReceivesWithContext is the same as DeleteAlarmReceives with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteAlarmReceives for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Monitorv4) DeleteAlarmReceivesWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteAlarmReceivesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeAlarmPolicy = "DescribeAlarmPolicy"

// DescribeAlarmPolicyRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeAlarmPolicy operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeAlarmPolicy for more information on using the DescribeAlarmPolicy
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeAlarmPolicyRequest method.
//    req, resp := client.DescribeAlarmPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/DescribeAlarmPolicy
func (c *Monitorv4) DescribeAlarmPolicyRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeAlarmPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeAlarmPolicy API operation for monitorv4.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for monitorv4's
// API operation DescribeAlarmPolicy for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/DescribeAlarmPolicy
func (c *Monitorv4) DescribeAlarmPolicy(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeAlarmPolicyRequest(input)
	return out, req.Send()
}

// DescribeAlarmPolicyWithContext is the same as DescribeAlarmPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAlarmPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Monitorv4) DescribeAlarmPolicyWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeAlarmPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeAlarmReceives = "DescribeAlarmReceives"

// DescribeAlarmReceivesRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeAlarmReceives operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeAlarmReceives for more information on using the DescribeAlarmReceives
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeAlarmReceivesRequest method.
//    req, resp := client.DescribeAlarmReceivesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/DescribeAlarmReceives
func (c *Monitorv4) DescribeAlarmReceivesRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeAlarmReceives,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeAlarmReceives API operation for monitorv4.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for monitorv4's
// API operation DescribeAlarmReceives for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/DescribeAlarmReceives
func (c *Monitorv4) DescribeAlarmReceives(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeAlarmReceivesRequest(input)
	return out, req.Send()
}

// DescribeAlarmReceivesWithContext is the same as DescribeAlarmReceives with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAlarmReceives for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Monitorv4) DescribeAlarmReceivesWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeAlarmReceivesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribePolicyObject = "DescribePolicyObject"

// DescribePolicyObjectRequest generates a "ksc/request.Request" representing the
// client's request for the DescribePolicyObject operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribePolicyObject for more information on using the DescribePolicyObject
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribePolicyObjectRequest method.
//    req, resp := client.DescribePolicyObjectRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/DescribePolicyObject
func (c *Monitorv4) DescribePolicyObjectRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribePolicyObject,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribePolicyObject API operation for monitorv4.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for monitorv4's
// API operation DescribePolicyObject for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/DescribePolicyObject
func (c *Monitorv4) DescribePolicyObject(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribePolicyObjectRequest(input)
	return out, req.Send()
}

// DescribePolicyObjectWithContext is the same as DescribePolicyObject with the addition of
// the ability to pass a context and additional request options.
//
// See DescribePolicyObject for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Monitorv4) DescribePolicyObjectWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribePolicyObjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetAlertUser = "GetAlertUser"

// GetAlertUserRequest generates a "ksc/request.Request" representing the
// client's request for the GetAlertUser operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetAlertUser for more information on using the GetAlertUser
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetAlertUserRequest method.
//    req, resp := client.GetAlertUserRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/GetAlertUser
func (c *Monitorv4) GetAlertUserRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetAlertUser,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// GetAlertUser API operation for monitorv4.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for monitorv4's
// API operation GetAlertUser for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/GetAlertUser
func (c *Monitorv4) GetAlertUser(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetAlertUserRequest(input)
	return out, req.Send()
}

// GetAlertUserWithContext is the same as GetAlertUser with the addition of
// the ability to pass a context and additional request options.
//
// See GetAlertUser for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Monitorv4) GetAlertUserWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetAlertUserRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetUserGroup = "GetUserGroup"

// GetUserGroupRequest generates a "ksc/request.Request" representing the
// client's request for the GetUserGroup operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetUserGroup for more information on using the GetUserGroup
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetUserGroupRequest method.
//    req, resp := client.GetUserGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/GetUserGroup
func (c *Monitorv4) GetUserGroupRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetUserGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// GetUserGroup API operation for monitorv4.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for monitorv4's
// API operation GetUserGroup for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/GetUserGroup
func (c *Monitorv4) GetUserGroup(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetUserGroupRequest(input)
	return out, req.Send()
}

// GetUserGroupWithContext is the same as GetUserGroup with the addition of
// the ability to pass a context and additional request options.
//
// See GetUserGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Monitorv4) GetUserGroupWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetUserGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAlarmPolicy = "ListAlarmPolicy"

// ListAlarmPolicyRequest generates a "ksc/request.Request" representing the
// client's request for the ListAlarmPolicy operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListAlarmPolicy for more information on using the ListAlarmPolicy
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListAlarmPolicyRequest method.
//    req, resp := client.ListAlarmPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/ListAlarmPolicy
func (c *Monitorv4) ListAlarmPolicyRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAlarmPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ListAlarmPolicy API operation for monitorv4.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for monitorv4's
// API operation ListAlarmPolicy for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/ListAlarmPolicy
func (c *Monitorv4) ListAlarmPolicy(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAlarmPolicyRequest(input)
	return out, req.Send()
}

// ListAlarmPolicyWithContext is the same as ListAlarmPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See ListAlarmPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Monitorv4) ListAlarmPolicyWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAlarmPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateAlertUserStatus = "UpdateAlertUserStatus"

// UpdateAlertUserStatusRequest generates a "ksc/request.Request" representing the
// client's request for the UpdateAlertUserStatus operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See UpdateAlertUserStatus for more information on using the UpdateAlertUserStatus
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the UpdateAlertUserStatusRequest method.
//    req, resp := client.UpdateAlertUserStatusRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/UpdateAlertUserStatus
func (c *Monitorv4) UpdateAlertUserStatusRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateAlertUserStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateAlertUserStatus API operation for monitorv4.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for monitorv4's
// API operation UpdateAlertUserStatus for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/monitor-2021-01-01/UpdateAlertUserStatus
func (c *Monitorv4) UpdateAlertUserStatus(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateAlertUserStatusRequest(input)
	return out, req.Send()
}

// UpdateAlertUserStatusWithContext is the same as UpdateAlertUserStatus with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateAlertUserStatus for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Monitorv4) UpdateAlertUserStatusWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateAlertUserStatusRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}
