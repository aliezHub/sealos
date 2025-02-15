/*
Copyright 2023 cuisongliu@qq.com.

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

package main

import (
	"os"
	"testing"
)

func TestTerraform_Apply(t *testing.T) {
	tf := &Terraform{
		accessKey: os.Getenv("ALIYUN_AKID"),
		secretKey: os.Getenv("ALIYUN_AKSK"),
	}
	if err := tf.Apply("amd64"); err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("SUCCESS")
}

func TestTerraform_Destroy(t *testing.T) {
	tf := &Terraform{
		accessKey: os.Getenv("ALIYUN_AKID"),
		secretKey: os.Getenv("ALIYUN_AKSK"),
	}
	if err := tf.Destroy(); err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("SUCCESS")
}

func TestTerraform_Detail(t *testing.T) {
	tf := &Terraform{
		accessKey: os.Getenv("ALIYUN_AKID"),
		secretKey: os.Getenv("ALIYUN_AKSK"),
	}
	var err error
	var d *InfraDetail
	if d, err = tf.Detail(); err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("SUCCESS: %+v", d)
}
