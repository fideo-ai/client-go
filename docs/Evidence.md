# Evidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpTor** | Pointer to **bool** |  | [optional] 
**IpCountry** | Pointer to [**IPCountry**](IPCountry.md) |  | [optional] 
**CountryOfIp** | Pointer to **string** |  | [optional] 

## Methods

### NewEvidence

`func NewEvidence() *Evidence`

NewEvidence instantiates a new Evidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidenceWithDefaults

`func NewEvidenceWithDefaults() *Evidence`

NewEvidenceWithDefaults instantiates a new Evidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpTor

`func (o *Evidence) GetIpTor() bool`

GetIpTor returns the IpTor field if non-nil, zero value otherwise.

### GetIpTorOk

`func (o *Evidence) GetIpTorOk() (*bool, bool)`

GetIpTorOk returns a tuple with the IpTor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpTor

`func (o *Evidence) SetIpTor(v bool)`

SetIpTor sets IpTor field to given value.

### HasIpTor

`func (o *Evidence) HasIpTor() bool`

HasIpTor returns a boolean if a field has been set.

### GetIpCountry

`func (o *Evidence) GetIpCountry() IPCountry`

GetIpCountry returns the IpCountry field if non-nil, zero value otherwise.

### GetIpCountryOk

`func (o *Evidence) GetIpCountryOk() (*IPCountry, bool)`

GetIpCountryOk returns a tuple with the IpCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpCountry

`func (o *Evidence) SetIpCountry(v IPCountry)`

SetIpCountry sets IpCountry field to given value.

### HasIpCountry

`func (o *Evidence) HasIpCountry() bool`

HasIpCountry returns a boolean if a field has been set.

### GetCountryOfIp

`func (o *Evidence) GetCountryOfIp() string`

GetCountryOfIp returns the CountryOfIp field if non-nil, zero value otherwise.

### GetCountryOfIpOk

`func (o *Evidence) GetCountryOfIpOk() (*string, bool)`

GetCountryOfIpOk returns a tuple with the CountryOfIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfIp

`func (o *Evidence) SetCountryOfIp(v string)`

SetCountryOfIp sets CountryOfIp field to given value.

### HasCountryOfIp

`func (o *Evidence) HasCountryOfIp() bool`

HasCountryOfIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


