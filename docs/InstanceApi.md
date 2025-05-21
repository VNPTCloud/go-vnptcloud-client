# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InstanceBundleInformationsGet**](InstanceApi.md#InstanceBundleInformationsGet) | **Get** /instance/bundle-informations | Thông tin tài nguyên liên quan đến máy ảo
[**InstanceCreatePost**](InstanceApi.md#InstanceCreatePost) | **Post** /instance/create | Khởi tạo máy ảo.
[**InstanceDeleteDelete**](InstanceApi.md#InstanceDeleteDelete) | **Delete** /instance/delete | Xóa máy ảo
[**InstanceDetailGet**](InstanceApi.md#InstanceDetailGet) | **Get** /instance/detail | Lấy thông tin chi tiết của máy ảo
[**InstanceEditPut**](InstanceApi.md#InstanceEditPut) | **Put** /instance/edit | Thay đổi thông tin máy ảo (tên)
[**InstanceFakeGet**](InstanceApi.md#InstanceFakeGet) | **Get** /instance/fake | Fake
[**InstanceGpusGet**](InstanceApi.md#InstanceGpusGet) | **Get** /instance/gpus | Lấy danh sách loại GPU
[**InstanceImagesGet**](InstanceApi.md#InstanceImagesGet) | **Get** /instance/images | Lấy danh sách hệ điều hành
[**InstanceIpPublicsGet**](InstanceApi.md#InstanceIpPublicsGet) | **Get** /instance/ip-publics | Lấy danh sách IP Public
[**InstanceListGet**](InstanceApi.md#InstanceListGet) | **Get** /instance/list | Lấy danh sách máy ảo
[**InstanceRebuildPost**](InstanceApi.md#InstanceRebuildPost) | **Post** /instance/rebuild | Thay đổi hệ điều hành máy ảo
[**InstanceResizePost**](InstanceApi.md#InstanceResizePost) | **Post** /instance/resize | Điều chỉnh máy ảo.
[**InstanceSecurityGroupsGet**](InstanceApi.md#InstanceSecurityGroupsGet) | **Get** /instance/security-groups | Lấy danh sách Security Group
[**InstanceSnapshotsGet**](InstanceApi.md#InstanceSnapshotsGet) | **Get** /instance/snapshots | Lấy danh sách VLAN Network
[**InstanceVlanPortsGet**](InstanceApi.md#InstanceVlanPortsGet) | **Get** /instance/vlan-ports | Lấy danh sách Port VLAN
[**InstanceVlansGet**](InstanceApi.md#InstanceVlansGet) | **Get** /instance/vlans | Lấy danh sách VLAN Network

# **InstanceBundleInformationsGet**
> ModelApiResponse InstanceBundleInformationsGet(ctx, optional)
Thông tin tài nguyên liên quan đến máy ảo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceBundleInformationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceBundleInformationsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **regionId** | **optional.Int32**|  | 
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

# **InstanceCreatePost**
> ModelApiResponse InstanceCreatePost(ctx, optional)
Khởi tạo máy ảo.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceCreatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceCreatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InstanceCreateModel**](InstanceCreateModel.md)|  | 
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

# **InstanceDeleteDelete**
> ModelApiResponse InstanceDeleteDelete(ctx, optional)
Xóa máy ảo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceDeleteDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceDeleteDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int32**|  | 
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

# **InstanceDetailGet**
> InstanceDetailInfo InstanceDetailGet(ctx, optional)
Lấy thông tin chi tiết của máy ảo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceDetailGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceDetailGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceId** | **optional.Int32**|  | 
 **userRootId** | **optional.Int32**| Id IAM User root | 
 **projectId** | **optional.Int32**| Project-Id | 
 **xProjectId** | **optional.Int32**| X-Project-Id | 

### Return type

[**InstanceDetailInfo**](InstanceDetailInfo.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceEditPut**
> ModelApiResponse InstanceEditPut(ctx, optional)
Thay đổi thông tin máy ảo (tên)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceEditPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceEditPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InstanceEditModel**](InstanceEditModel.md)|  | 
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

# **InstanceFakeGet**
> InstanceFakeGet(ctx, optional)
Fake

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceFakeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceFakeGetOpts struct
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

# **InstanceGpusGet**
> []InstanceGpuTypeInfo InstanceGpusGet(ctx, optional)
Lấy danh sách loại GPU

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceGpusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceGpusGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRootId** | **optional.Int32**| Id IAM User root | 
 **projectId** | **optional.Int32**| Project-Id | 
 **xProjectId** | **optional.Int32**| X-Project-Id | 

### Return type

[**[]InstanceGpuTypeInfo**](InstanceGpuTypeInfo.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceImagesGet**
> []ImageInfo InstanceImagesGet(ctx, optional)
Lấy danh sách hệ điều hành

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceImagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceImagesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRootId** | **optional.Int32**| Id IAM User root | 
 **projectId** | **optional.Int32**| Project-Id | 
 **xProjectId** | **optional.Int32**| X-Project-Id | 

### Return type

[**[]ImageInfo**](ImageInfo.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceIpPublicsGet**
> []IpPublicInfo InstanceIpPublicsGet(ctx, optional)
Lấy danh sách IP Public

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceIpPublicsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceIpPublicsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRootId** | **optional.Int32**| Id IAM User root | 
 **projectId** | **optional.Int32**| Project-Id | 
 **xProjectId** | **optional.Int32**| X-Project-Id | 

### Return type

[**[]IpPublicInfo**](IPPublicInfo.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceListGet**
> []InstanceInfo InstanceListGet(ctx, optional)
Lấy danh sách máy ảo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRootId** | **optional.Int32**| Id IAM User root | 
 **projectId** | **optional.Int32**| Project-Id | 
 **xProjectId** | **optional.Int32**| X-Project-Id | 

### Return type

[**[]InstanceInfo**](InstanceInfo.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceRebuildPost**
> ModelApiResponse InstanceRebuildPost(ctx, optional)
Thay đổi hệ điều hành máy ảo

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceRebuildPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceRebuildPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InstanceRebuildModel**](InstanceRebuildModel.md)|  | 
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

# **InstanceResizePost**
> ModelApiResponse InstanceResizePost(ctx, optional)
Điều chỉnh máy ảo.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceResizePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceResizePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InstanceResizeModel**](InstanceResizeModel.md)|  | 
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

# **InstanceSecurityGroupsGet**
> []SecurityGroupInfo InstanceSecurityGroupsGet(ctx, optional)
Lấy danh sách Security Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceSecurityGroupsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceSecurityGroupsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRootId** | **optional.Int32**| Id IAM User root | 
 **projectId** | **optional.Int32**| Project-Id | 
 **xProjectId** | **optional.Int32**| X-Project-Id | 

### Return type

[**[]SecurityGroupInfo**](SecurityGroupInfo.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceSnapshotsGet**
> []SnapshotInfo InstanceSnapshotsGet(ctx, optional)
Lấy danh sách VLAN Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceSnapshotsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceSnapshotsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRootId** | **optional.Int32**| Id IAM User root | 
 **projectId** | **optional.Int32**| Project-Id | 
 **xProjectId** | **optional.Int32**| X-Project-Id | 

### Return type

[**[]SnapshotInfo**](SnapshotInfo.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceVlanPortsGet**
> []PortInfo InstanceVlanPortsGet(ctx, optional)
Lấy danh sách Port VLAN

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceVlanPortsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceVlanPortsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanId** | **optional.String**|  | 
 **userRootId** | **optional.Int32**| Id IAM User root | 
 **projectId** | **optional.Int32**| Project-Id | 
 **xProjectId** | **optional.Int32**| X-Project-Id | 

### Return type

[**[]PortInfo**](PortInfo.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstanceVlansGet**
> []VlanInfo InstanceVlansGet(ctx, optional)
Lấy danh sách VLAN Network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***InstanceApiInstanceVlansGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InstanceApiInstanceVlansGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userRootId** | **optional.Int32**| Id IAM User root | 
 **projectId** | **optional.Int32**| Project-Id | 
 **xProjectId** | **optional.Int32**| X-Project-Id | 

### Return type

[**[]VlanInfo**](VlanInfo.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

