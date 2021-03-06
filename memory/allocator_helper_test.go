// Encoding: utf-8
// Cloud Foundry Java Buildpack
// Copyright (c) 2015 the original author or authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Helper file for testing Allocator type. Only compiled when testing.
package memory

// Expose internal bucket strings for testing
func GetBuckets(alloc Allocator) []string {
	a, _ := alloc.(*allocator)
	bs := []string{}
	for _, b := range a.buckets {
		bs = append(bs, b.String())
	}
	return bs
}
