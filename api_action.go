package go-vnptcloud-client

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type ActionApiService service
/*
ActionApiService Lấy trạng thái đơn hàng
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ActionApiActionGetStatusGetOpts - Optional Parameters:
     * @param "TransactionId" (optional.String) - 
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id

*/

type ActionApiActionGetStatusGetOpts struct {
    TransactionId optional.String
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *ActionApiService) ActionGetStatusGet(ctx context.Context, localVarOptionals *ActionApiActionGetStatusGetOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/action/get-status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.TransactionId.IsSet() {
		localVarQueryParams.Add("transactionId", parameterToString(localVarOptionals.TransactionId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.UserRootId.IsSet() {
		localVarHeaderParams["User-Root-Id"] = parameterToString(localVarOptionals.UserRootId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.ProjectId.IsSet() {
		localVarHeaderParams["Project-Id"] = parameterToString(localVarOptionals.ProjectId.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XProjectId.IsSet() {
		localVarHeaderParams["X-Project-Id"] = parameterToString(localVarOptionals.XProjectId.Value(), "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
