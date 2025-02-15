// Copyright © 2023 sealos.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package infra

import (
	v1 "github.com/labring/sealos/controllers/infra/api/v1"
	"github.com/labring/sealos/controllers/infra/common"
)

// https://help.aliyun.com/document_detail/100410.html?spm=a2c4g.440317.0.0.22e04747yC0m0u#section-vat-mi8-821
const (
	AliyunArm64UbuntuImage = "ubuntu_20_04_arm64_20G_alibase_20230104.vhd"
	AliyunAmd64UbuntuImage = "ubuntu_20_04_x64_20G_alibase_20230208.vhd"

	AliyunAmd64CentosImage = "centos_7_9_x64_20G_alibase_20230109.vhd"

	AliyunArm64Flavor = "ecs.g6r.large"
	AliyunAmd64Flavor = "ecs.c7.large"
)

func GetPublicIP(hosts []v1.Hosts) []string {
	return getIPSByType(hosts, common.IPTypePublic)
}

func GetPrivateIP(hosts []v1.Hosts) []string {
	return getIPSByType(hosts, common.IPTypePrivate)
}

func getIPSByType(hosts []v1.Hosts, ipType string) (ips []string) {
	masterIps := make(map[string]struct{})
	var master0Ip string
	for _, h := range hosts {
		for _, m := range h.Metadata {
			for _, ip := range m.IP {
				if ip.IPType != ipType {
					continue
				}
				masterIps[ip.IPValue] = struct{}{}
				if _, ok := m.Labels[common.MasterO]; ok {
					master0Ip = ip.IPValue
				}
				break
			}
		}
	}
	if master0Ip != "" {
		ips = append(ips, master0Ip)
	}

	for ip := range masterIps {
		if ip != master0Ip {
			ips = append(ips, ip)
		}
	}
	return
}
