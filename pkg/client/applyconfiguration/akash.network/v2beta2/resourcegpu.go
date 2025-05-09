/*
Copyright The Akash Network Authors.

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

package v2beta2

import (
	v1beta3 "github.com/akash-network/akash-api/go/node/types/v1beta3"
)

// ResourceGPUApplyConfiguration represents a declarative configuration of the ResourceGPU type for use
// with apply.
type ResourceGPUApplyConfiguration struct {
	Units      *uint32             `json:"units,omitempty"`
	Attributes *v1beta3.Attributes `json:"attributes,omitempty"`
}

// ResourceGPUApplyConfiguration constructs a declarative configuration of the ResourceGPU type for use with
// apply.
func ResourceGPU() *ResourceGPUApplyConfiguration {
	return &ResourceGPUApplyConfiguration{}
}

// WithUnits sets the Units field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Units field is set to the value of the last call.
func (b *ResourceGPUApplyConfiguration) WithUnits(value uint32) *ResourceGPUApplyConfiguration {
	b.Units = &value
	return b
}

// WithAttributes sets the Attributes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Attributes field is set to the value of the last call.
func (b *ResourceGPUApplyConfiguration) WithAttributes(value v1beta3.Attributes) *ResourceGPUApplyConfiguration {
	b.Attributes = &value
	return b
}
