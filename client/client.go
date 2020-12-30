// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ExecuteRequestRequest struct {
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ExecuteRequestRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequestRequest) GoString() string {
	return s.String()
}

func (s *ExecuteRequestRequest) SetService(v string) *ExecuteRequestRequest {
	s.Service = &v
	return s
}

func (s *ExecuteRequestRequest) SetServiceParameters(v string) *ExecuteRequestRequest {
	s.ServiceParameters = &v
	return s
}

type ExecuteRequestResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ExecuteRequestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequestResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteRequestResponseBody) SetMessage(v string) *ExecuteRequestResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteRequestResponseBody) SetRequestId(v string) *ExecuteRequestResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExecuteRequestResponseBody) SetData(v string) *ExecuteRequestResponseBody {
	s.Data = &v
	return s
}

func (s *ExecuteRequestResponseBody) SetCode(v int32) *ExecuteRequestResponseBody {
	s.Code = &v
	return s
}

type ExecuteRequestResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ExecuteRequestResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteRequestResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteRequestResponse) GoString() string {
	return s.String()
}

func (s *ExecuteRequestResponse) SetHeaders(v map[string]*string) *ExecuteRequestResponse {
	s.Headers = v
	return s
}

func (s *ExecuteRequestResponse) SetBody(v *ExecuteRequestResponseBody) *ExecuteRequestResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-hangzhou": tea.String("saf.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("saf"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteRequestWithOptions(request *ExecuteRequestRequest, runtime *util.RuntimeOptions) (_result *ExecuteRequestResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ExecuteRequestResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ExecuteRequest"), tea.String("2017-03-31"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteRequest(request *ExecuteRequestRequest) (_result *ExecuteRequestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteRequestResponse{}
	_body, _err := client.ExecuteRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
