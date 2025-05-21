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

type InstanceApiService service
/*
InstanceApiService Thông tin tài nguyên liên quan đến máy ảo
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceBundleInformationsGetOpts - Optional Parameters:
     * @param "RegionId" (optional.Int32) - 
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return ModelApiResponse
*/

type InstanceApiInstanceBundleInformationsGetOpts struct {
    RegionId optional.Int32
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceBundleInformationsGet(ctx context.Context, localVarOptionals *InstanceApiInstanceBundleInformationsGetOpts) (ModelApiResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ModelApiResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/bundle-informations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.RegionId.IsSet() {
		localVarQueryParams.Add("regionId", parameterToString(localVarOptionals.RegionId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Khởi tạo máy ảo.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceCreatePostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of InstanceCreateModel) - 
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return ModelApiResponse
*/

type InstanceApiInstanceCreatePostOpts struct {
    Body optional.Interface
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceCreatePost(ctx context.Context, localVarOptionals *InstanceApiInstanceCreatePostOpts) (ModelApiResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ModelApiResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json", "application/_*+json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Xóa máy ảo
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceDeleteDeleteOpts - Optional Parameters:
     * @param "Id" (optional.Int32) - 
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return ModelApiResponse
*/

type InstanceApiInstanceDeleteDeleteOpts struct {
    Id optional.Int32
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceDeleteDelete(ctx context.Context, localVarOptionals *InstanceApiInstanceDeleteDeleteOpts) (ModelApiResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ModelApiResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/delete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Id.IsSet() {
		localVarQueryParams.Add("id", parameterToString(localVarOptionals.Id.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Lấy thông tin chi tiết của máy ảo
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceDetailGetOpts - Optional Parameters:
     * @param "InstanceId" (optional.Int32) - 
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return InstanceDetailInfo
*/

type InstanceApiInstanceDetailGetOpts struct {
    InstanceId optional.Int32
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceDetailGet(ctx context.Context, localVarOptionals *InstanceApiInstanceDetailGetOpts) (InstanceDetailInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue InstanceDetailInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/detail"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.InstanceId.IsSet() {
		localVarQueryParams.Add("instanceId", parameterToString(localVarOptionals.InstanceId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v InstanceDetailInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Thay đổi thông tin máy ảo (tên)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceEditPutOpts - Optional Parameters:
     * @param "Body" (optional.Interface of InstanceEditModel) - 
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return ModelApiResponse
*/

type InstanceApiInstanceEditPutOpts struct {
    Body optional.Interface
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceEditPut(ctx context.Context, localVarOptionals *InstanceApiInstanceEditPutOpts) (ModelApiResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ModelApiResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/edit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json", "application/_*+json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Fake
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceFakeGetOpts - Optional Parameters:
     * @param "TransactionId" (optional.String) - 
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id

*/

type InstanceApiInstanceFakeGetOpts struct {
    TransactionId optional.String
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceFakeGet(ctx context.Context, localVarOptionals *InstanceApiInstanceFakeGetOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/fake"

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
/*
InstanceApiService Lấy danh sách loại GPU
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceGpusGetOpts - Optional Parameters:
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return []InstanceGpuTypeInfo
*/

type InstanceApiInstanceGpusGetOpts struct {
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceGpusGet(ctx context.Context, localVarOptionals *InstanceApiInstanceGpusGetOpts) ([]InstanceGpuTypeInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []InstanceGpuTypeInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/gpus"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []InstanceGpuTypeInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Lấy danh sách hệ điều hành
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceImagesGetOpts - Optional Parameters:
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return []ImageInfo
*/

type InstanceApiInstanceImagesGetOpts struct {
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceImagesGet(ctx context.Context, localVarOptionals *InstanceApiInstanceImagesGetOpts) ([]ImageInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []ImageInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/images"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []ImageInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Lấy danh sách IP Public
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceIpPublicsGetOpts - Optional Parameters:
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return []IpPublicInfo
*/

type InstanceApiInstanceIpPublicsGetOpts struct {
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceIpPublicsGet(ctx context.Context, localVarOptionals *InstanceApiInstanceIpPublicsGetOpts) ([]IpPublicInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []IpPublicInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/ip-publics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []IpPublicInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Lấy danh sách máy ảo
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceListGetOpts - Optional Parameters:
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return []InstanceInfo
*/

type InstanceApiInstanceListGetOpts struct {
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceListGet(ctx context.Context, localVarOptionals *InstanceApiInstanceListGetOpts) ([]InstanceInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []InstanceInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []InstanceInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Thay đổi hệ điều hành máy ảo
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceRebuildPostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of InstanceRebuildModel) - 
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return ModelApiResponse
*/

type InstanceApiInstanceRebuildPostOpts struct {
    Body optional.Interface
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceRebuildPost(ctx context.Context, localVarOptionals *InstanceApiInstanceRebuildPostOpts) (ModelApiResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ModelApiResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/rebuild"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json", "application/_*+json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Điều chỉnh máy ảo.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceResizePostOpts - Optional Parameters:
     * @param "Body" (optional.Interface of InstanceResizeModel) - 
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return ModelApiResponse
*/

type InstanceApiInstanceResizePostOpts struct {
    Body optional.Interface
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceResizePost(ctx context.Context, localVarOptionals *InstanceApiInstanceResizePostOpts) (ModelApiResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ModelApiResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/resize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json", "application/_*+json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 201 {
			var v ModelApiResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Lấy danh sách Security Group
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceSecurityGroupsGetOpts - Optional Parameters:
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return []SecurityGroupInfo
*/

type InstanceApiInstanceSecurityGroupsGetOpts struct {
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceSecurityGroupsGet(ctx context.Context, localVarOptionals *InstanceApiInstanceSecurityGroupsGetOpts) ([]SecurityGroupInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []SecurityGroupInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/security-groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []SecurityGroupInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Lấy danh sách VLAN Network
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceSnapshotsGetOpts - Optional Parameters:
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return []SnapshotInfo
*/

type InstanceApiInstanceSnapshotsGetOpts struct {
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceSnapshotsGet(ctx context.Context, localVarOptionals *InstanceApiInstanceSnapshotsGetOpts) ([]SnapshotInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []SnapshotInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/snapshots"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []SnapshotInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Lấy danh sách Port VLAN
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceVlanPortsGetOpts - Optional Parameters:
     * @param "VlanId" (optional.String) - 
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return []PortInfo
*/

type InstanceApiInstanceVlanPortsGetOpts struct {
    VlanId optional.String
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceVlanPortsGet(ctx context.Context, localVarOptionals *InstanceApiInstanceVlanPortsGetOpts) ([]PortInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []PortInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/vlan-ports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.VlanId.IsSet() {
		localVarQueryParams.Add("vlanId", parameterToString(localVarOptionals.VlanId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []PortInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
InstanceApiService Lấy danh sách VLAN Network
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *InstanceApiInstanceVlansGetOpts - Optional Parameters:
     * @param "UserRootId" (optional.Int32) -  Id IAM User root
     * @param "ProjectId" (optional.Int32) -  Project-Id
     * @param "XProjectId" (optional.Int32) -  X-Project-Id
@return []VlanInfo
*/

type InstanceApiInstanceVlansGetOpts struct {
    UserRootId optional.Int32
    ProjectId optional.Int32
    XProjectId optional.Int32
}

func (a *InstanceApiService) InstanceVlansGet(ctx context.Context, localVarOptionals *InstanceApiInstanceVlansGetOpts) ([]VlanInfo, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue []VlanInfo
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/instance/vlans"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

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
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []VlanInfo
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
