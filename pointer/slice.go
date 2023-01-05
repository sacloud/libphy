// Copyright 2021-2023 The phy-api-go authors
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

package pointer

// StringSlice returns a pointer to the given tags value
func StringSlice(v []string) *[]string { return &v }

// IntSlice returns a pointer to the given tags value
func IntSlice(v []int) *[]int { return &v }

// Int64Slice returns a pointer to the given tags value
func Int64Slice(v []int64) *[]int64 { return &v }

// UintSlice returns a pointer to the given tags value
func UintSlice(v []uint) *[]uint { return &v }

// Uint64Slice returns a pointer to the given tags value
func Uint64Slice(v []uint64) *[]uint64 { return &v }

// ByteSlice returns a pointer to the given tags value
func ByteSlice(v []byte) *[]byte { return &v }
