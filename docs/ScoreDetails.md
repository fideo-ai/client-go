# ScoreDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scorer** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **float64** |  | [optional] 
**Evidence** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Weight** | Pointer to **float64** |  | [optional] 

## Methods

### NewScoreDetails

`func NewScoreDetails() *ScoreDetails`

NewScoreDetails instantiates a new ScoreDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreDetailsWithDefaults

`func NewScoreDetailsWithDefaults() *ScoreDetails`

NewScoreDetailsWithDefaults instantiates a new ScoreDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScorer

`func (o *ScoreDetails) GetScorer() string`

GetScorer returns the Scorer field if non-nil, zero value otherwise.

### GetScorerOk

`func (o *ScoreDetails) GetScorerOk() (*string, bool)`

GetScorerOk returns a tuple with the Scorer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScorer

`func (o *ScoreDetails) SetScorer(v string)`

SetScorer sets Scorer field to given value.

### HasScorer

`func (o *ScoreDetails) HasScorer() bool`

HasScorer returns a boolean if a field has been set.

### GetScore

`func (o *ScoreDetails) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ScoreDetails) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ScoreDetails) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *ScoreDetails) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetEvidence

`func (o *ScoreDetails) GetEvidence() map[string]map[string]interface{}`

GetEvidence returns the Evidence field if non-nil, zero value otherwise.

### GetEvidenceOk

`func (o *ScoreDetails) GetEvidenceOk() (*map[string]map[string]interface{}, bool)`

GetEvidenceOk returns a tuple with the Evidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvidence

`func (o *ScoreDetails) SetEvidence(v map[string]map[string]interface{})`

SetEvidence sets Evidence field to given value.

### HasEvidence

`func (o *ScoreDetails) HasEvidence() bool`

HasEvidence returns a boolean if a field has been set.

### GetWeight

`func (o *ScoreDetails) GetWeight() float64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ScoreDetails) GetWeightOk() (*float64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ScoreDetails) SetWeight(v float64)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ScoreDetails) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


