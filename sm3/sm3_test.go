/*
Copyright Suzhou Tongji Fintech Research Institute 2017 All Rights Reserved.
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

package sm3

import (
	"strconv"
	"testing"

	tjfoc "github.com/tjfoc/gmsm/sm3"
)

func BenchmarkSm3(t *testing.B) {
	t.ReportAllocs()
	msg := "standardTS"
	var sm3 SM3
	for i := 0; i < t.N; i++ {
		// Sm3Sum(msg)
		sm3.Write([]byte(msg))
	}
}

func BenchmarkSm3_Tjfoc(t *testing.B) {
	t.ReportAllocs()
	msg := "standardTS"
	var sm3 tjfoc.SM3
	for i := 0; i < t.N; i++ {
		// Sm3Sum(msg)
		sm3.Write([]byte(msg))
	}
}

func BenchmarkSM3_TjfocCount(b *testing.B) {
	msg := "standardTS"
	var sm3 tjfoc.SM3
	for i := 0; i < 10; i++ {
		msg = msg + msg
		b.Run("len"+strconv.Itoa(len(msg)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sm3.Write([]byte(msg))
			}
		})
	}
}

func BenchmarkSM3Count(b *testing.B) {
	msg := "standardTS"
	var sm3 SM3
	for i := 0; i < 10; i++ {
		msg = msg + msg
		b.Run("len"+strconv.Itoa(len(msg)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sm3.Write([]byte(msg))
			}
		})
	}
}
