/*
 * Generated file models/multidomain/v4/content/content_model.go.
 *
 * Product version: 4.4.1-beta-1
 *
 * Part of the Nutanix Multidomain Versioned APIs
 *
 * (c) 2026 Nutanix Inc.  All rights reserved
 *
 */

/*
  Multi domain content.
*/
package content

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	import1 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/common/v1/response"
	import3 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/multidomain/v4/error"
	import2 "github.com/nutanix/ntnx-api-golang-clients/multidomain-go-client/v4/models/prism/v4/config"
	"time"
)

type ContentRepository struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Create time of a Content Repository.
	*/
	CreateTime *time.Time `json:"createTime,omitempty"`
	/*
	  The user-defined description of a Content Repository.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Last update time of a Content Repository.
	*/
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The user-defined name of a Content Repository.
	*/
	Name *string `json:"name"`
	/*
	  External identifier of the owner of a Content Repository.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  External identifier of the project to which the Content Repository belongs.
	*/
	ProjectExtId *string `json:"projectExtId"`

	Publisher *ContentRepositoryPublisherInfo `json:"publisher,omitempty"`

	State *ContentRepositoryState `json:"state,omitempty"`
	/*
	  Subscription list of Content Repository at the source PC.
	*/
	Subscribers []ContentRepositorySubscriberInfo `json:"subscribers,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`
	/*
	  Total number of items in a Content Repository.
	*/
	TotalItemCount *int64 `json:"totalItemCount,omitempty"`
}

func (p *ContentRepository) MarshalJSON() ([]byte, error) {
	type ContentRepositoryProxy ContentRepository

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ContentRepositoryProxy
		Name         *string `json:"name,omitempty"`
		ProjectExtId *string `json:"projectExtId,omitempty"`
	}{
		ContentRepositoryProxy: (*ContentRepositoryProxy)(p),
		Name:                   p.Name,
		ProjectExtId:           p.ProjectExtId,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ContentRepository) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ContentRepository
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewContentRepository()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CreateTime != nil {
		p.CreateTime = known.CreateTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.LastUpdateTime != nil {
		p.LastUpdateTime = known.LastUpdateTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.ProjectExtId != nil {
		p.ProjectExtId = known.ProjectExtId
	}
	if known.Publisher != nil {
		p.Publisher = known.Publisher
	}
	if known.State != nil {
		p.State = known.State
	}
	if known.Subscribers != nil {
		p.Subscribers = known.Subscribers
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.TotalItemCount != nil {
		p.TotalItemCount = known.TotalItemCount
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "createTime")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "lastUpdateTime")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "ownerExtId")
	delete(allFields, "projectExtId")
	delete(allFields, "publisher")
	delete(allFields, "state")
	delete(allFields, "subscribers")
	delete(allFields, "tenantId")
	delete(allFields, "totalItemCount")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewContentRepository() *ContentRepository {
	p := new(ContentRepository)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.ContentRepository"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

type ContentRepositoryItem struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Create time of a Repository item.
	*/
	CreateTime *time.Time `json:"createTime,omitempty"`
	/*
	  The user-defined description of a Repository item.
	*/
	Description *string `json:"description,omitempty"`
	/*
	  A globally unique identifier of an instance that is suitable for external consumption.
	*/
	ExtId *string `json:"extId,omitempty"`
	/*
	  Last update time of a Repository item.
	*/
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	/*
	  A HATEOAS style link for the response.  Each link contains a user-friendly name identifying the link and an address for retrieving the particular resource.
	*/
	Links []import1.ApiLink `json:"links,omitempty"`
	/*
	  The user-defined name of a Repository item.
	*/
	Name *string `json:"name"`
	/*
	  External identifier of the owner of a Repository item.
	*/
	OwnerExtId *string `json:"ownerExtId,omitempty"`
	/*
	  The size in bytes of a Repository item file.
	*/
	SizeBytes *int64 `json:"sizeBytes,omitempty"`

	State *ContentRepositoryItemState `json:"state,omitempty"`
	/*
	  A globally unique identifier that represents the tenant that owns this entity. The system automatically assigns it, and it and is immutable from an API consumer perspective (some use cases may cause this ID to change - For instance, a use case may require the transfer of ownership of the entity, but these cases are handled automatically on the server).
	*/
	TenantId *string `json:"tenantId,omitempty"`

	Type *ContentRepositoryItemType `json:"type"`
}

func (p *ContentRepositoryItem) MarshalJSON() ([]byte, error) {
	type ContentRepositoryItemProxy ContentRepositoryItem

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ContentRepositoryItemProxy
		Name *string                    `json:"name,omitempty"`
		Type *ContentRepositoryItemType `json:"type,omitempty"`
	}{
		ContentRepositoryItemProxy: (*ContentRepositoryItemProxy)(p),
		Name:                       p.Name,
		Type:                       p.Type,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ContentRepositoryItem) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ContentRepositoryItem
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewContentRepositoryItem()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.CreateTime != nil {
		p.CreateTime = known.CreateTime
	}
	if known.Description != nil {
		p.Description = known.Description
	}
	if known.ExtId != nil {
		p.ExtId = known.ExtId
	}
	if known.LastUpdateTime != nil {
		p.LastUpdateTime = known.LastUpdateTime
	}
	if known.Links != nil {
		p.Links = known.Links
	}
	if known.Name != nil {
		p.Name = known.Name
	}
	if known.OwnerExtId != nil {
		p.OwnerExtId = known.OwnerExtId
	}
	if known.SizeBytes != nil {
		p.SizeBytes = known.SizeBytes
	}
	if known.State != nil {
		p.State = known.State
	}
	if known.TenantId != nil {
		p.TenantId = known.TenantId
	}
	if known.Type != nil {
		p.Type = known.Type
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "createTime")
	delete(allFields, "description")
	delete(allFields, "extId")
	delete(allFields, "lastUpdateTime")
	delete(allFields, "links")
	delete(allFields, "name")
	delete(allFields, "ownerExtId")
	delete(allFields, "sizeBytes")
	delete(allFields, "state")
	delete(allFields, "tenantId")
	delete(allFields, "type")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewContentRepositoryItem() *ContentRepositoryItem {
	p := new(ContentRepositoryItem)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.ContentRepositoryItem"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Various states of a Content Repository Item.
*/
type ContentRepositoryItemState int

const (
	CONTENTREPOSITORYITEMSTATE_UNKNOWN    ContentRepositoryItemState = 0
	CONTENTREPOSITORYITEMSTATE_REDACTED   ContentRepositoryItemState = 1
	CONTENTREPOSITORYITEMSTATE_PENDING    ContentRepositoryItemState = 2
	CONTENTREPOSITORYITEMSTATE_READY      ContentRepositoryItemState = 3
	CONTENTREPOSITORYITEMSTATE_FAILED     ContentRepositoryItemState = 4
	CONTENTREPOSITORYITEMSTATE_DELETING   ContentRepositoryItemState = 5
	CONTENTREPOSITORYITEMSTATE_CONFLICTED ContentRepositoryItemState = 6
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ContentRepositoryItemState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PENDING",
		"READY",
		"FAILED",
		"DELETING",
		"CONFLICTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ContentRepositoryItemState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PENDING",
		"READY",
		"FAILED",
		"DELETING",
		"CONFLICTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ContentRepositoryItemState) index(name string) ContentRepositoryItemState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"PENDING",
		"READY",
		"FAILED",
		"DELETING",
		"CONFLICTED",
	}
	for idx := range names {
		if names[idx] == name {
			return ContentRepositoryItemState(idx)
		}
	}
	return CONTENTREPOSITORYITEMSTATE_UNKNOWN
}

func (e *ContentRepositoryItemState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ContentRepositoryItemState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ContentRepositoryItemState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ContentRepositoryItemState) Ref() *ContentRepositoryItemState {
	return &e
}

/*
Type of Repository item.
*/
type ContentRepositoryItemType int

const (
	CONTENTREPOSITORYITEMTYPE_UNKNOWN  ContentRepositoryItemType = 0
	CONTENTREPOSITORYITEMTYPE_REDACTED ContentRepositoryItemType = 1
	CONTENTREPOSITORYITEMTYPE_IMAGE    ContentRepositoryItemType = 2
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ContentRepositoryItemType) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IMAGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ContentRepositoryItemType) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IMAGE",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ContentRepositoryItemType) index(name string) ContentRepositoryItemType {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"IMAGE",
	}
	for idx := range names {
		if names[idx] == name {
			return ContentRepositoryItemType(idx)
		}
	}
	return CONTENTREPOSITORYITEMTYPE_UNKNOWN
}

func (e *ContentRepositoryItemType) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ContentRepositoryItemType:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ContentRepositoryItemType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ContentRepositoryItemType) Ref() *ContentRepositoryItemType {
	return &e
}

/*
Detail of a publisher Content Repository.
*/
type ContentRepositoryPublisherInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Last metadata sync time of the Content Repository as part of the subscription.
	*/
	LastSyncTime *time.Time `json:"lastSyncTime,omitempty"`
	/*
	  The Domain Manager or Prism Central hosting the publisher Content Repository.
	*/
	RemoteDomainManagerExtId *string `json:"remoteDomainManagerExtId"`
}

func (p *ContentRepositoryPublisherInfo) MarshalJSON() ([]byte, error) {
	type ContentRepositoryPublisherInfoProxy ContentRepositoryPublisherInfo

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ContentRepositoryPublisherInfoProxy
		RemoteDomainManagerExtId *string `json:"remoteDomainManagerExtId,omitempty"`
	}{
		ContentRepositoryPublisherInfoProxy: (*ContentRepositoryPublisherInfoProxy)(p),
		RemoteDomainManagerExtId:            p.RemoteDomainManagerExtId,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ContentRepositoryPublisherInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ContentRepositoryPublisherInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewContentRepositoryPublisherInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.LastSyncTime != nil {
		p.LastSyncTime = known.LastSyncTime
	}
	if known.RemoteDomainManagerExtId != nil {
		p.RemoteDomainManagerExtId = known.RemoteDomainManagerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "lastSyncTime")
	delete(allFields, "remoteDomainManagerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewContentRepositoryPublisherInfo() *ContentRepositoryPublisherInfo {
	p := new(ContentRepositoryPublisherInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.ContentRepositoryPublisherInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Various states of a Content Repository.
*/
type ContentRepositoryState int

const (
	CONTENTREPOSITORYSTATE_UNKNOWN     ContentRepositoryState = 0
	CONTENTREPOSITORYSTATE_REDACTED    ContentRepositoryState = 1
	CONTENTREPOSITORYSTATE_SUBSCRIBING ContentRepositoryState = 2
	CONTENTREPOSITORYSTATE_SUBSCRIBED  ContentRepositoryState = 3
	CONTENTREPOSITORYSTATE_READY       ContentRepositoryState = 4
	CONTENTREPOSITORYSTATE_STALE       ContentRepositoryState = 5
	CONTENTREPOSITORYSTATE_DELETING    ContentRepositoryState = 6
	CONTENTREPOSITORYSTATE_CONFLICTED  ContentRepositoryState = 7
)

// Returns the name of the enum given an ordinal number
//
// Deprecated: Please use GetName instead of name
func (e *ContentRepositoryState) name(index int) string {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUBSCRIBING",
		"SUBSCRIBED",
		"READY",
		"STALE",
		"DELETING",
		"CONFLICTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the name of the enum
func (e ContentRepositoryState) GetName() string {
	index := int(e)
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUBSCRIBING",
		"SUBSCRIBED",
		"READY",
		"STALE",
		"DELETING",
		"CONFLICTED",
	}
	if index < 0 || index >= len(names) {
		return "$UNKNOWN"
	}
	return names[index]
}

// Returns the enum type given a string value
func (e *ContentRepositoryState) index(name string) ContentRepositoryState {
	names := [...]string{
		"$UNKNOWN",
		"$REDACTED",
		"SUBSCRIBING",
		"SUBSCRIBED",
		"READY",
		"STALE",
		"DELETING",
		"CONFLICTED",
	}
	for idx := range names {
		if names[idx] == name {
			return ContentRepositoryState(idx)
		}
	}
	return CONTENTREPOSITORYSTATE_UNKNOWN
}

func (e *ContentRepositoryState) UnmarshalJSON(b []byte) error {
	var enumStr string
	if err := json.Unmarshal(b, &enumStr); err != nil {
		return errors.New(fmt.Sprintf("Unable to unmarshal for ContentRepositoryState:%s", err))
	}
	*e = e.index(enumStr)
	return nil
}

func (e *ContentRepositoryState) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(e.name(int(*e)))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

func (e ContentRepositoryState) Ref() *ContentRepositoryState {
	return &e
}

/*
Detail of a subscriber Content Repository.
*/
type ContentRepositorySubscriberInfo struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  Last metadata sync time of the Content Repository as part of the subscription.
	*/
	LastSyncTime *time.Time `json:"lastSyncTime,omitempty"`
	/*
	  The Domain Manager or Prism Central hosting the subscriber Content Repository.
	*/
	RemoteDomainManagerExtId *string `json:"remoteDomainManagerExtId"`

	SubscriptionState *ContentRepositoryState `json:"subscriptionState,omitempty"`
}

func (p *ContentRepositorySubscriberInfo) MarshalJSON() ([]byte, error) {
	type ContentRepositorySubscriberInfoProxy ContentRepositorySubscriberInfo

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ContentRepositorySubscriberInfoProxy
		RemoteDomainManagerExtId *string `json:"remoteDomainManagerExtId,omitempty"`
	}{
		ContentRepositorySubscriberInfoProxy: (*ContentRepositorySubscriberInfoProxy)(p),
		RemoteDomainManagerExtId:             p.RemoteDomainManagerExtId,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ContentRepositorySubscriberInfo) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ContentRepositorySubscriberInfo
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewContentRepositorySubscriberInfo()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.LastSyncTime != nil {
		p.LastSyncTime = known.LastSyncTime
	}
	if known.RemoteDomainManagerExtId != nil {
		p.RemoteDomainManagerExtId = known.RemoteDomainManagerExtId
	}
	if known.SubscriptionState != nil {
		p.SubscriptionState = known.SubscriptionState
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "lastSyncTime")
	delete(allFields, "remoteDomainManagerExtId")
	delete(allFields, "subscriptionState")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewContentRepositorySubscriberInfo() *ContentRepositorySubscriberInfo {
	p := new(ContentRepositorySubscriberInfo)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.ContentRepositorySubscriberInfo"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
Subscription parameters to setup a target Content Repository.
Credentials are not needed as the Domains are expected to be paired.
*/
type ContentRepositorySubscriptionSpec struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*
	  The external identifier of the project to which the Content Repository subscription belongs.
	The project name supplied by the subscriber must match the project name associated with the Content Repository.
	*/
	ProjectExtId *string `json:"projectExtId"`
	/*
	  The external identifier of a source Domain Manager or Prism Central.
	*/
	SourceDomainManagerExtId *string `json:"sourceDomainManagerExtId"`
}

func (p *ContentRepositorySubscriptionSpec) MarshalJSON() ([]byte, error) {
	type ContentRepositorySubscriptionSpecProxy ContentRepositorySubscriptionSpec

	// Step 1: Marshal known fields via proxy to enforce required fields
	baseStruct := struct {
		*ContentRepositorySubscriptionSpecProxy
		ProjectExtId             *string `json:"projectExtId,omitempty"`
		SourceDomainManagerExtId *string `json:"sourceDomainManagerExtId,omitempty"`
	}{
		ContentRepositorySubscriptionSpecProxy: (*ContentRepositorySubscriptionSpecProxy)(p),
		ProjectExtId:                           p.ProjectExtId,
		SourceDomainManagerExtId:               p.SourceDomainManagerExtId,
	}

	known, err := json.Marshal(baseStruct)
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ContentRepositorySubscriptionSpec) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ContentRepositorySubscriptionSpec
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewContentRepositorySubscriptionSpec()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.ProjectExtId != nil {
		p.ProjectExtId = known.ProjectExtId
	}
	if known.SourceDomainManagerExtId != nil {
		p.SourceDomainManagerExtId = known.SourceDomainManagerExtId
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "projectExtId")
	delete(allFields, "sourceDomainManagerExtId")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewContentRepositorySubscriptionSpec() *ContentRepositorySubscriptionSpec {
	p := new(ContentRepositorySubscriptionSpec)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.ContentRepositorySubscriptionSpec"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/content/content-repositories Post operation
*/
type CreateContentRepositoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfCreateContentRepositoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *CreateContentRepositoryApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias CreateContentRepositoryApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *CreateContentRepositoryApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias CreateContentRepositoryApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewCreateContentRepositoryApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewCreateContentRepositoryApiResponse() *CreateContentRepositoryApiResponse {
	p := new(CreateContentRepositoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.CreateContentRepositoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *CreateContentRepositoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *CreateContentRepositoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfCreateContentRepositoryApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/content/content-repositories/{extId} Delete operation
*/
type DeleteContentRepositoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfDeleteContentRepositoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *DeleteContentRepositoryApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias DeleteContentRepositoryApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *DeleteContentRepositoryApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias DeleteContentRepositoryApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewDeleteContentRepositoryApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewDeleteContentRepositoryApiResponse() *DeleteContentRepositoryApiResponse {
	p := new(DeleteContentRepositoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.DeleteContentRepositoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *DeleteContentRepositoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *DeleteContentRepositoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfDeleteContentRepositoryApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/content/content-repositories/{extId} Get operation
*/
type GetContentRepositoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfGetContentRepositoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *GetContentRepositoryApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias GetContentRepositoryApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *GetContentRepositoryApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias GetContentRepositoryApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewGetContentRepositoryApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewGetContentRepositoryApiResponse() *GetContentRepositoryApiResponse {
	p := new(GetContentRepositoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.GetContentRepositoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *GetContentRepositoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *GetContentRepositoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfGetContentRepositoryApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/content/content-repositories Get operation
*/
type ListContentRepositoriesApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListContentRepositoriesApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListContentRepositoriesApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListContentRepositoriesApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ListContentRepositoriesApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListContentRepositoriesApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListContentRepositoriesApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListContentRepositoriesApiResponse() *ListContentRepositoriesApiResponse {
	p := new(ListContentRepositoriesApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.ListContentRepositoriesApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListContentRepositoriesApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListContentRepositoriesApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListContentRepositoriesApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/content/content-repositories/{extId}/items Get operation
*/
type ListContentRepositoryItemsApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfListContentRepositoryItemsApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *ListContentRepositoryItemsApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias ListContentRepositoryItemsApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *ListContentRepositoryItemsApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias ListContentRepositoryItemsApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewListContentRepositoryItemsApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewListContentRepositoryItemsApiResponse() *ListContentRepositoryItemsApiResponse {
	p := new(ListContentRepositoryItemsApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.ListContentRepositoryItemsApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *ListContentRepositoryItemsApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *ListContentRepositoryItemsApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfListContentRepositoryItemsApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/content/content-repositories/{extId}/$actions/subscribe Post operation
*/
type SubscribeContentRepositoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfSubscribeContentRepositoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *SubscribeContentRepositoryApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SubscribeContentRepositoryApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *SubscribeContentRepositoryApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SubscribeContentRepositoryApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSubscribeContentRepositoryApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSubscribeContentRepositoryApiResponse() *SubscribeContentRepositoryApiResponse {
	p := new(SubscribeContentRepositoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.SubscribeContentRepositoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *SubscribeContentRepositoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *SubscribeContentRepositoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfSubscribeContentRepositoryApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/content/content-repositories/{extId}/$actions/sync Post operation
*/
type SyncContentRepositoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfSyncContentRepositoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *SyncContentRepositoryApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias SyncContentRepositoryApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *SyncContentRepositoryApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias SyncContentRepositoryApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewSyncContentRepositoryApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewSyncContentRepositoryApiResponse() *SyncContentRepositoryApiResponse {
	p := new(SyncContentRepositoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.SyncContentRepositoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *SyncContentRepositoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *SyncContentRepositoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfSyncContentRepositoryApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

/*
REST response for all response codes in API path /multidomain/v4.4.b1/content/content-repositories/{extId} Put operation
*/
type UpdateContentRepositoryApiResponse struct {
	ObjectType_ *string `json:"$objectType,omitempty"`

	Reserved_ map[string]interface{} `json:"$reserved,omitempty"`

	UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
	/*

	 */
	DataItemDiscriminator_ *string `json:"$dataItemDiscriminator,omitempty"`

	Data *OneOfUpdateContentRepositoryApiResponseData `json:"data,omitempty"`

	Metadata *import1.ApiResponseMetadata `json:"metadata,omitempty"`
}

func (p *UpdateContentRepositoryApiResponse) MarshalJSON() ([]byte, error) {
	// Create Alias to avoid infinite recursion
	type Alias UpdateContentRepositoryApiResponse

	// Step 1: Marshal the known fields
	known, err := json.Marshal(Alias(*p))
	if err != nil {
		return nil, err
	}

	// Step 2: Convert known to map for merging
	var knownMap map[string]interface{}
	if err := json.Unmarshal(known, &knownMap); err != nil {
		return nil, err
	}
	delete(knownMap, "$unknownFields")

	// Step 3: Merge unknown fields
	for k, v := range p.UnknownFields_ {
		knownMap[k] = v
	}

	// Step 4: Marshal final merged map
	return json.Marshal(knownMap)
}

func (p *UpdateContentRepositoryApiResponse) UnmarshalJSON(b []byte) error {
	// Step 1: Unmarshal into a generic map to capture all fields
	var allFields map[string]interface{}
	if err := json.Unmarshal(b, &allFields); err != nil {
		return err
	}

	// Step 2: Unmarshal into a temporary struct with known fields
	type Alias UpdateContentRepositoryApiResponse
	known := &Alias{}
	if err := json.Unmarshal(b, known); err != nil {
		return err
	}

	// Step 3: Assign known fields
	*p = *NewUpdateContentRepositoryApiResponse()

	if known.ObjectType_ != nil {
		p.ObjectType_ = known.ObjectType_
	}
	if known.Reserved_ != nil {
		p.Reserved_ = known.Reserved_
	}
	if known.UnknownFields_ != nil {
		p.UnknownFields_ = known.UnknownFields_
	}
	if known.DataItemDiscriminator_ != nil {
		p.DataItemDiscriminator_ = known.DataItemDiscriminator_
	}
	if known.Data != nil {
		p.Data = known.Data
	}
	if known.Metadata != nil {
		p.Metadata = known.Metadata
	}

	// Step 4: Remove known JSON fields from allFields map
	delete(allFields, "$objectType")
	delete(allFields, "$reserved")
	delete(allFields, "$unknownFields")
	delete(allFields, "$dataItemDiscriminator")
	delete(allFields, "data")
	delete(allFields, "metadata")

	// Step 5: Assign remaining fields to UnknownFields_
	for key, value := range allFields {
		p.UnknownFields_[key] = value
	}

	return nil
}

func NewUpdateContentRepositoryApiResponse() *UpdateContentRepositoryApiResponse {
	p := new(UpdateContentRepositoryApiResponse)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "multidomain.v4.content.UpdateContentRepositoryApiResponse"
	p.Reserved_ = map[string]interface{}{"$fv": "v4.r4.b1"}
	p.UnknownFields_ = map[string]interface{}{}

	return p
}

func (p *UpdateContentRepositoryApiResponse) GetData() interface{} {
	if nil == p.Data {
		return nil
	}
	return p.Data.GetValue()
}

func (p *UpdateContentRepositoryApiResponse) SetData(v interface{}) error {
	if nil == p.Data {
		p.Data = NewOneOfUpdateContentRepositoryApiResponseData()
	}
	e := p.Data.SetValue(v)
	if nil == e {
		if nil == p.DataItemDiscriminator_ {
			p.DataItemDiscriminator_ = new(string)
		}
		*p.DataItemDiscriminator_ = *p.Data.Discriminator
	}
	return e
}

type OneOfListContentRepositoryItemsApiResponseData struct {
	Discriminator *string                 `json:"-"`
	ObjectType_   *string                 `json:"-"`
	oneOfType400  *import3.ErrorResponse  `json:"-"`
	oneOfType2001 []ContentRepositoryItem `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListContentRepositoryItemsApiResponseData() *OneOfListContentRepositoryItemsApiResponseData {
	p := new(OneOfListContentRepositoryItemsApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListContentRepositoryItemsApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListContentRepositoryItemsApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []ContentRepositoryItem:
		p.oneOfType2001 = v.([]ContentRepositoryItem)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<multidomain.v4.content.ContentRepositoryItem>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<multidomain.v4.content.ContentRepositoryItem>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListContentRepositoryItemsApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<multidomain.v4.content.ContentRepositoryItem>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListContentRepositoryItemsApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import3.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import3.ErrorResponse)
							}
							*p.oneOfType400 = *vOneOfType400
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType400.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType400.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<multidomain.v4.content.ContentRepositoryItem>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new([]ContentRepositoryItem)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType2001 == nil || len(*vOneOfType2001) == 0 || ((*vOneOfType2001)[0].ObjectType_ != nil && "multidomain.v4.content.ContentRepositoryItem" == *((*vOneOfType2001)[0].ObjectType_)) {
							p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<multidomain.v4.content.ContentRepositoryItem>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<multidomain.v4.content.ContentRepositoryItem>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new([]ContentRepositoryItem)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || (vOneOfType2001 != nil && (*vOneOfType2001)[0].ObjectType_ != nil && "multidomain.v4.content.ContentRepositoryItem" == *((*vOneOfType2001)[0].ObjectType_)) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<multidomain.v4.content.ContentRepositoryItem>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<multidomain.v4.content.ContentRepositoryItem>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListContentRepositoryItemsApiResponseData"))
}

func (p *OneOfListContentRepositoryItemsApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<multidomain.v4.content.ContentRepositoryItem>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListContentRepositoryItemsApiResponseData")
}

type OneOfCreateContentRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfCreateContentRepositoryApiResponseData() *OneOfCreateContentRepositoryApiResponseData {
	p := new(OneOfCreateContentRepositoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfCreateContentRepositoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfCreateContentRepositoryApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfCreateContentRepositoryApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfCreateContentRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import3.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import3.ErrorResponse)
							}
							*p.oneOfType400 = *vOneOfType400
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType400.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType400.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import2.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import2.TaskReference)
							}
							*p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2001.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfCreateContentRepositoryApiResponseData"))
}

func (p *OneOfCreateContentRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfCreateContentRepositoryApiResponseData")
}

type OneOfListContentRepositoriesApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 []ContentRepository    `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfListContentRepositoriesApiResponseData() *OneOfListContentRepositoriesApiResponseData {
	p := new(OneOfListContentRepositoriesApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfListContentRepositoriesApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfListContentRepositoriesApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case []ContentRepository:
		p.oneOfType2001 = v.([]ContentRepository)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = "List<multidomain.v4.content.ContentRepository>"
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = "List<multidomain.v4.content.ContentRepository>"
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfListContentRepositoriesApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if "List<multidomain.v4.content.ContentRepository>" == *p.Discriminator {
		return p.oneOfType2001
	}
	return nil
}

func (p *OneOfListContentRepositoriesApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import3.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import3.ErrorResponse)
							}
							*p.oneOfType400 = *vOneOfType400
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType400.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType400.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["List<multidomain.v4.content.ContentRepository>"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new([]ContentRepository)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For arrays, verify the array item ObjectType matches
						if vOneOfType2001 == nil || len(*vOneOfType2001) == 0 || ((*vOneOfType2001)[0].ObjectType_ != nil && "multidomain.v4.content.ContentRepository" == *((*vOneOfType2001)[0].ObjectType_)) {
							p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = "List<multidomain.v4.content.ContentRepository>"
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = "List<multidomain.v4.content.ContentRepository>"
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new([]ContentRepository)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if len(*vOneOfType2001) == 0 || (vOneOfType2001 != nil && (*vOneOfType2001)[0].ObjectType_ != nil && "multidomain.v4.content.ContentRepository" == *((*vOneOfType2001)[0].ObjectType_)) {
			p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = "List<multidomain.v4.content.ContentRepository>"
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = "List<multidomain.v4.content.ContentRepository>"
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfListContentRepositoriesApiResponseData"))
}

func (p *OneOfListContentRepositoriesApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if "List<multidomain.v4.content.ContentRepository>" == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfListContentRepositoriesApiResponseData")
}

type OneOfSyncContentRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfSyncContentRepositoryApiResponseData() *OneOfSyncContentRepositoryApiResponseData {
	p := new(OneOfSyncContentRepositoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSyncContentRepositoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSyncContentRepositoryApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfSyncContentRepositoryApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfSyncContentRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import3.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import3.ErrorResponse)
							}
							*p.oneOfType400 = *vOneOfType400
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType400.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType400.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import2.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import2.TaskReference)
							}
							*p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2001.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSyncContentRepositoryApiResponseData"))
}

func (p *OneOfSyncContentRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfSyncContentRepositoryApiResponseData")
}

type OneOfGetContentRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *ContentRepository     `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfGetContentRepositoryApiResponseData() *OneOfGetContentRepositoryApiResponseData {
	p := new(OneOfGetContentRepositoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfGetContentRepositoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfGetContentRepositoryApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case ContentRepository:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(ContentRepository)
		}
		*p.oneOfType2001 = v.(ContentRepository)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfGetContentRepositoryApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfGetContentRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import3.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import3.ErrorResponse)
							}
							*p.oneOfType400 = *vOneOfType400
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType400.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType400.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(ContentRepository)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "multidomain.v4.content.ContentRepository" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(ContentRepository)
							}
							*p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2001.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(ContentRepository)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "multidomain.v4.content.ContentRepository" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(ContentRepository)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfGetContentRepositoryApiResponseData"))
}

func (p *OneOfGetContentRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfGetContentRepositoryApiResponseData")
}

type OneOfDeleteContentRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfDeleteContentRepositoryApiResponseData() *OneOfDeleteContentRepositoryApiResponseData {
	p := new(OneOfDeleteContentRepositoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfDeleteContentRepositoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfDeleteContentRepositoryApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfDeleteContentRepositoryApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfDeleteContentRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import3.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import3.ErrorResponse)
							}
							*p.oneOfType400 = *vOneOfType400
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType400.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType400.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import2.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import2.TaskReference)
							}
							*p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2001.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfDeleteContentRepositoryApiResponseData"))
}

func (p *OneOfDeleteContentRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfDeleteContentRepositoryApiResponseData")
}

type OneOfSubscribeContentRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfSubscribeContentRepositoryApiResponseData() *OneOfSubscribeContentRepositoryApiResponseData {
	p := new(OneOfSubscribeContentRepositoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfSubscribeContentRepositoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfSubscribeContentRepositoryApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfSubscribeContentRepositoryApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfSubscribeContentRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import3.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import3.ErrorResponse)
							}
							*p.oneOfType400 = *vOneOfType400
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType400.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType400.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import2.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import2.TaskReference)
							}
							*p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2001.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfSubscribeContentRepositoryApiResponseData"))
}

func (p *OneOfSubscribeContentRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfSubscribeContentRepositoryApiResponseData")
}

type OneOfUpdateContentRepositoryApiResponseData struct {
	Discriminator *string                `json:"-"`
	ObjectType_   *string                `json:"-"`
	oneOfType400  *import3.ErrorResponse `json:"-"`
	oneOfType2001 *import2.TaskReference `json:"-"`
	// Holds data with unknown oneOf types
	UnknownValue_ interface{} `json:"-"`
}

func NewOneOfUpdateContentRepositoryApiResponseData() *OneOfUpdateContentRepositoryApiResponseData {
	p := new(OneOfUpdateContentRepositoryApiResponseData)
	p.Discriminator = new(string)
	p.ObjectType_ = new(string)
	return p
}

func (p *OneOfUpdateContentRepositoryApiResponseData) SetValue(v interface{}) error {
	if nil == p {
		return errors.New(fmt.Sprintf("OneOfUpdateContentRepositoryApiResponseData is nil"))
	}
	switch v.(type) {
	case import3.ErrorResponse:
		if nil == p.oneOfType400 {
			p.oneOfType400 = new(import3.ErrorResponse)
		}
		*p.oneOfType400 = v.(import3.ErrorResponse)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType400.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType400.ObjectType_
	case import2.TaskReference:
		if nil == p.oneOfType2001 {
			p.oneOfType2001 = new(import2.TaskReference)
		}
		*p.oneOfType2001 = v.(import2.TaskReference)
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		*p.Discriminator = *p.oneOfType2001.ObjectType_
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.oneOfType2001.ObjectType_
	default:
		return errors.New(fmt.Sprintf("%T(%v) is not expected type", v, v))
	}
	return nil
}

func (p *OneOfUpdateContentRepositoryApiResponseData) GetValue() interface{} {
	if p.UnknownValue_ != nil {
		return p.UnknownValue_
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return *p.oneOfType400
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return *p.oneOfType2001
	}
	return nil
}

func (p *OneOfUpdateContentRepositoryApiResponseData) UnmarshalJSON(b []byte) error {
	p.UnknownValue_ = nil
	// Try to handle nested structure like {"": {"value": {...}}}
	// This recursively unwraps {"field": {"value": {...}}} patterns for nested oneOf fields
	var rawMap map[string]interface{}
	if err := json.Unmarshal(b, &rawMap); err == nil {
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType400 := new(import3.ErrorResponse)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType400)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
							if nil == p.oneOfType400 {
								p.oneOfType400 = new(import3.ErrorResponse)
							}
							*p.oneOfType400 = *vOneOfType400
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType400.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType400.ObjectType_
							return nil
						}
					}
				}
			}
		}
		// Check if this field name exists in the map (handles nested structure)
		if nestedMap, ok := rawMap["ObjectType_"].(map[string]interface{}); ok {
			// Check for "value" wrapper
			if valueData, ok := nestedMap["value"]; ok {
				valueJSON, marshalErr := json.Marshal(valueData)
				if marshalErr == nil {
					vOneOfType2001 := new(import2.TaskReference)
					var unmarshalErr error
					// Unmarshal - if vField has oneOf fields, their UnmarshalJSON will handle nested patterns recursively
					unmarshalErr = json.Unmarshal(valueJSON, vOneOfType2001)
					if unmarshalErr == nil {
						// For struct items, verify the ObjectType matches
						if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
							if nil == p.oneOfType2001 {
								p.oneOfType2001 = new(import2.TaskReference)
							}
							*p.oneOfType2001 = *vOneOfType2001
							if nil == p.Discriminator {
								p.Discriminator = new(string)
							}
							*p.Discriminator = *p.oneOfType2001.ObjectType_
							if nil == p.ObjectType_ {
								p.ObjectType_ = new(string)
							}
							*p.ObjectType_ = *p.oneOfType2001.ObjectType_
							return nil
						}
					}
				}
			}
		}
	}

	// Fallback: try direct unmarshalling (for non-nested structures)
	vOneOfType400 := new(import3.ErrorResponse)
	if err := json.Unmarshal(b, vOneOfType400); err == nil {
		if vOneOfType400.ObjectType_ != nil && "multidomain.v4.error.ErrorResponse" == *vOneOfType400.ObjectType_ {
			if nil == p.oneOfType400 {
				p.oneOfType400 = new(import3.ErrorResponse)
			}
			*p.oneOfType400 = *vOneOfType400
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType400.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType400.ObjectType_
			return nil
		}
	}
	vOneOfType2001 := new(import2.TaskReference)
	if err := json.Unmarshal(b, vOneOfType2001); err == nil {
		if vOneOfType2001.ObjectType_ != nil && "prism.v4.config.TaskReference" == *vOneOfType2001.ObjectType_ {
			if nil == p.oneOfType2001 {
				p.oneOfType2001 = new(import2.TaskReference)
			}
			*p.oneOfType2001 = *vOneOfType2001
			if nil == p.Discriminator {
				p.Discriminator = new(string)
			}
			*p.Discriminator = *p.oneOfType2001.ObjectType_
			if nil == p.ObjectType_ {
				p.ObjectType_ = new(string)
			}
			*p.ObjectType_ = *p.oneOfType2001.ObjectType_
			return nil
		}
	}
	// Store raw when no known variant matched
	var unknownRaw map[string]interface{}
	if err := json.Unmarshal(b, &unknownRaw); err == nil {
		p.UnknownValue_ = unknownRaw
		if nil == p.Discriminator {
			p.Discriminator = new(string)
		}
		if ot, ok := unknownRaw["$objectType"].(string); ok && ot != "" {
			*p.Discriminator = ot
		} else {
			*p.Discriminator = "UNKNOWN"
		}
		if nil == p.ObjectType_ {
			p.ObjectType_ = new(string)
		}
		*p.ObjectType_ = *p.Discriminator
		return nil
	}
	return errors.New(fmt.Sprintf("Unable to unmarshal for OneOfUpdateContentRepositoryApiResponseData"))
}

func (p *OneOfUpdateContentRepositoryApiResponseData) MarshalJSON() ([]byte, error) {
	if p.UnknownValue_ != nil {
		return json.Marshal(p.UnknownValue_)
	}
	if p.oneOfType400 != nil && *p.oneOfType400.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType400)
	}
	if p.oneOfType2001 != nil && *p.oneOfType2001.ObjectType_ == *p.Discriminator {
		return json.Marshal(p.oneOfType2001)
	}
	return nil, errors.New("No value to marshal for OneOfUpdateContentRepositoryApiResponseData")
}

type FileDetail struct {
	Path        *string `json:"-"`
	ObjectType_ *string `json:"-"`
}

func NewFileDetail() *FileDetail {
	p := new(FileDetail)
	p.ObjectType_ = new(string)
	*p.ObjectType_ = "FileDetail"

	return p
}
