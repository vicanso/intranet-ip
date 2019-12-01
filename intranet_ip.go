// Copyright 2019 tree xie
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

package intranetip

import "net"

var (
	// intranetIPBlocks intranet ip blocks
	intranetIPBlocks []*net.IPNet
)

// https://stackoverflow.com/questions/43274579/golang-check-if-ip-address-is-in-a-network/43274687
func init() {
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	} {
		_, block, _ := net.ParseCIDR(cidr)
		intranetIPBlocks = append(intranetIPBlocks, block)
	}
}

// Is determine if the IP address is an intranet IP
func Is(ip net.IP) bool {
	for _, block := range intranetIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}
