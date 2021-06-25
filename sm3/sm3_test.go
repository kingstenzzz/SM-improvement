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
	"fmt"
	"io/ioutil"
	"testing"
)

func ReadFile(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read error")
	}
	return content
}

func byteToString(b []byte) string {
	ret := ""
	for i := 0; i < len(b); i++ {
		ret += fmt.Sprintf("%02x", b[i])
	}
	fmt.Println("ret = ", ret)
	return ret
}

func BenchmarkSm3(t *testing.B) {
	t.ReportAllocs()
	msg := ReadFile("./1.jpg")
	var sm3 SM3
	for i := 0; i < t.N; i++ {
		// Sm3Sum(msg)
		sm3.Write(msg)
	}
}
