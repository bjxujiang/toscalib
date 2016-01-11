/*
Copyright 2015 - Olivier Wulveryck

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
package toscalib

// Input corresponds to  `yaml:"inputs,omitempty" json:"inputs,omitempty"`
type Input struct {
	Value            map[string][]string `yaml:"-,inline"`
	Type             string              `yaml:"type" json:"type"`
	Description      string              `yaml:"description,omitempty" json:"description,omitempty"` // Not required
	Constraints      Constraints         `yaml:"constraints,omitempty" json:"constraints,omitempty"`
	ValidSourceTypes interface{}         `yaml:"valid_source_types,omitempty" json:"valid_source_types,omitempty"`
	Occurrences      interface{}         `yaml:"occurrences,omitempty" json:"occurrences,omitempty"`
}
