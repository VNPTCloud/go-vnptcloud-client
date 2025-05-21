# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionGetStatusGet**](ActionApi.md#ActionGetStatusGet) | **Get** /action/get-status | Lấy trạng thái đơn hàng

# **ActionGetStatusGet**
> ActionGetStatusGet(ctx, optional)
Lấy trạng thái đơn hàng

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ActionApiActionGetStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ActionApiActionGetStatusGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionId** | **optional.String**|  | 
 **userRootId** | **optional.Int32**| Id IAM User root | 
 **projectId** | **optional.Int32**| Project-Id | 
 **xProjectId** | **optional.Int32**| X-Project-Id | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

