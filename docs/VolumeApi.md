# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VolumeCreateInitProjectIdGet**](VolumeApi.md#VolumeCreateInitProjectIdGet) | **Get** /volume/create-init/{projectId} | Khởi tạo volume.
[**VolumeCreatePost**](VolumeApi.md#VolumeCreatePost) | **Post** /volume/create | Khởi tạo volume.
[**VolumeDeleteIdDelete**](VolumeApi.md#VolumeDeleteIdDelete) | **Delete** /volume/delete/{id} | Xóa volume.
[**VolumeResizePost**](VolumeApi.md#VolumeResizePost) | **Post** /volume/resize | Điều chỉnh máy ảo.

# **VolumeCreateInitProjectIdGet**
> ModelApiResponse VolumeCreateInitProjectIdGet(ctx, projectId, optional)
Khởi tạo volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int32**|  | 
 **optional** | ***VolumeApiVolumeCreateInitProjectIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeApiVolumeCreateInitProjectIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRootId** | **optional.Int32**| Id IAM User root | 
 **projectId** | **optional.Int32**| Project-Id | 
 **xProjectId** | **optional.Int32**| X-Project-Id | 

### Return type

[**ModelApiResponse**](ApiResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeCreatePost**
> ModelApiResponse VolumeCreatePost(ctx, optional)
Khởi tạo volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumeApiVolumeCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeApiVolumeCreatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VolumeCreateModel**](VolumeCreateModel.md)|  | 
 **userRootId** | **optional.**| Id IAM User root | 
 **projectId** | **optional.**| Project-Id | 
 **xProjectId** | **optional.**| X-Project-Id | 

### Return type

[**ModelApiResponse**](ApiResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeDeleteIdDelete**
> ModelApiResponse VolumeDeleteIdDelete(ctx, id, optional)
Xóa volume.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***VolumeApiVolumeDeleteIdDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeApiVolumeDeleteIdDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userRootId** | **optional.Int32**| Id IAM User root | 
 **projectId** | **optional.Int32**| Project-Id | 
 **xProjectId** | **optional.Int32**| X-Project-Id | 

### Return type

[**ModelApiResponse**](ApiResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeResizePost**
> ModelApiResponse VolumeResizePost(ctx, optional)
Điều chỉnh máy ảo.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumeApiVolumeResizePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumeApiVolumeResizePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VolumeResizeModel**](VolumeResizeModel.md)|  | 
 **userRootId** | **optional.**| Id IAM User root | 
 **projectId** | **optional.**| Project-Id | 
 **xProjectId** | **optional.**| X-Project-Id | 

### Return type

[**ModelApiResponse**](ApiResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json, text/json, application/_*+json
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

