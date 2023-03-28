/*
Copyright 2022.

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

package vectoragent

import (
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	ApiPort             = 8686
	MetricsExporterPort = 9598
)

func (ctrl *Controller) createVectorAgentService() *corev1.Service {
	labels := ctrl.labelsForVectorAgent()

	ports := []corev1.ServicePort{
		{
			Name:       "prom-exporter",
			Protocol:   corev1.Protocol("TCP"),
			Port:       MetricsExporterPort,
			TargetPort: intstr.FromInt(MetricsExporterPort),
		},
	}

	if ctrl.Vector.Spec.Agent.Api.Enabled {
		ports = append(ports, corev1.ServicePort{
			Name:       "api",
			Protocol:   corev1.Protocol("TCP"),
			Port:       ApiPort,
			TargetPort: intstr.FromInt(ApiPort),
		})
	}

	service := &corev1.Service{
		ObjectMeta: ctrl.objectMetaVectorAgent(labels),
		Spec: corev1.ServiceSpec{
			Ports:    ports,
			Selector: labels,
		},
	}

	return service
}
