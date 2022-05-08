/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConditionApplyConfiguration represents an declarative configuration of the Condition type for use
// with apply.
type ConditionApplyConfiguration struct {
	Type               *string             `json:"type,omitempty"`
	Status             *v1.ConditionStatus `json:"status,omitempty"`
	ObservedGeneration *int64              `json:"observedGeneration,omitempty"`
	LastTransitionTime *v1.Time            `json:"lastTransitionTime,omitempty"`
	Reason             *string             `json:"reason,omitempty"`
	Message            *string             `json:"message,omitempty"`
}

// ConditionApplyConfiguration constructs an declarative configuration of the Condition type for use with
// apply.
func Condition() *ConditionApplyConfiguration {
	return &ConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithType(value string) *ConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *ConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithObservedGeneration(value int64) *ConditionApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithLastTransitionTime(value v1.Time) *ConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithReason(value string) *ConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *ConditionApplyConfiguration) WithMessage(value string) *ConditionApplyConfiguration {
	b.Message = &value
	return b
}
