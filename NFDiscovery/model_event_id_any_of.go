/*
NRF NFDiscovery Service

NRF NFDiscovery Service. © 2021, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// EventIdAnyOf the model 'EventIdAnyOf'
type EventIdAnyOf string

// List of EventId_anyOf
const (
	LOAD_LEVEL_INFORMATION EventIdAnyOf = "LOAD_LEVEL_INFORMATION"
	NETWORK_PERFORMANCE EventIdAnyOf = "NETWORK_PERFORMANCE"
	NF_LOAD EventIdAnyOf = "NF_LOAD"
	SERVICE_EXPERIENCE EventIdAnyOf = "SERVICE_EXPERIENCE"
	UE_MOBILITY EventIdAnyOf = "UE_MOBILITY"
	UE_COMMUNICATION EventIdAnyOf = "UE_COMMUNICATION"
	QOS_SUSTAINABILITY EventIdAnyOf = "QOS_SUSTAINABILITY"
	ABNORMAL_BEHAVIOUR EventIdAnyOf = "ABNORMAL_BEHAVIOUR"
	USER_DATA_CONGESTION EventIdAnyOf = "USER_DATA_CONGESTION"
	NSI_LOAD_LEVEL EventIdAnyOf = "NSI_LOAD_LEVEL"
)

var allowedEventIdAnyOfEnumValues = []EventIdAnyOf{
	"LOAD_LEVEL_INFORMATION",
	"NETWORK_PERFORMANCE",
	"NF_LOAD",
	"SERVICE_EXPERIENCE",
	"UE_MOBILITY",
	"UE_COMMUNICATION",
	"QOS_SUSTAINABILITY",
	"ABNORMAL_BEHAVIOUR",
	"USER_DATA_CONGESTION",
	"NSI_LOAD_LEVEL",
}

func (v *EventIdAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventIdAnyOf(value)
	for _, existing := range allowedEventIdAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventIdAnyOf", value)
}

// NewEventIdAnyOfFromValue returns a pointer to a valid EventIdAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventIdAnyOfFromValue(v string) (*EventIdAnyOf, error) {
	ev := EventIdAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventIdAnyOf: valid values are %v", v, allowedEventIdAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventIdAnyOf) IsValid() bool {
	for _, existing := range allowedEventIdAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventId_anyOf value
func (v EventIdAnyOf) Ptr() *EventIdAnyOf {
	return &v
}

type NullableEventIdAnyOf struct {
	value *EventIdAnyOf
	isSet bool
}

func (v NullableEventIdAnyOf) Get() *EventIdAnyOf {
	return v.value
}

func (v *NullableEventIdAnyOf) Set(val *EventIdAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventIdAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventIdAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventIdAnyOf(val *EventIdAnyOf) *NullableEventIdAnyOf {
	return &NullableEventIdAnyOf{value: val, isSet: true}
}

func (v NullableEventIdAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventIdAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

