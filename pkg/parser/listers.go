// Copyright 2019 Altinity Ltd and/or its affiliates. All rights reserved.
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

package parser

import (
	chiv1 "github.com/altinity/clickhouse-operator/pkg/apis/clickhouse.altinity.com/v1"
	"gopkg.in/yaml.v2"
)

func ListPodHostnames(chi *chiv1.ClickHouseInstallation) []string {
	// Number of full deployment ids - which is number of Stateful Sets and number of Pods
	fullDeploymentIDsNum := len(chi.Status.FullDeploymentIDs)

	names := make([]string, fullDeploymentIDsNum)

	for i, fullDeploymentID := range chi.Status.FullDeploymentIDs {
		names[i] = CreatePodHostname(fullDeploymentID)
	}

	return names
}

func ListPodFQDNs(chi *chiv1.ClickHouseInstallation) []string {
	// Number of full deployment ids - which is number of Stateful Sets and number of Pods
	fullDeploymentIDsNum := len(chi.Status.FullDeploymentIDs)

	names := make([]string, fullDeploymentIDsNum)

	for i, fullDeploymentID := range chi.Status.FullDeploymentIDs {
		names[i] = CreatePodFQDN(chi.Namespace, fullDeploymentID)
	}

	return names
}

func ListStatefulSetNames(chi *chiv1.ClickHouseInstallation) []string {
	// Number of full deployment ids - which is number of Stateful Sets and number of Pods
	fullDeploymentIDsNum := len(chi.Status.FullDeploymentIDs)

	names := make([]string, fullDeploymentIDsNum)

	for i, fullDeploymentID := range chi.Status.FullDeploymentIDs {
		names[i] = CreateStatefulSetName(fullDeploymentID)
	}

	return names
}

func Yaml(chi *chiv1.ClickHouseInstallation) string {
	if data, err := yaml.Marshal(chi); err != nil {
		return ""
	} else {
		return string(data)
	}
}
