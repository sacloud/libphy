// Copyright 2021-2025 The phy-api-go authors
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

// Bool returns a pointer to the given bool value.
func Bool(b bool) *bool { return &b }

// String returns a pointer to the given string value.
func String(s string) *string { return &s }

// Int returns a pointer to the given int value.
func Int(i int) *int { return &i }

// Int8 returns a pointer to the given int8 value.
func Int8(i int8) *int8 { return &i }

// Int16 returns a pointer to the given int16 value.
func Int16(i int16) *int16 { return &i }

// Int32 returns a pointer to the given int32 value.
func Int32(i int32) *int32 { return &i }

// Int64 returns a pointer to the given int64 value.
func Int64(i int64) *int64 { return &i }

// Uint returns a pointer to the given uint value.
func Uint(i uint) *uint { return &i }

// Uint8 returns a pointer to the given uint8 value.
func Uint8(i uint8) *uint8 { return &i }

// Uint16 returns a pointer to the given uint16 value.
func Uint16(i uint16) *uint16 { return &i }

// Uint32 returns a pointer to the given uint32 value.
func Uint32(i uint32) *uint32 { return &i }

// Uint64 returns a pointer to the given uint64 value.
func Uint64(i uint64) *uint64 { return &i }

// Float32 returns a pointer to the given float32 value.
func Float32(f float32) *float32 { return &f }

// Float64 returns a pointer to the given float64 value.
func Float64(f float64) *float64 { return &f }

// Byte returns a pointer to the given byte value.
func Byte(b byte) *byte { return &b }

// Rune returns a pointer to the given rune value.
func Rune(r rune) *rune { return &r }
