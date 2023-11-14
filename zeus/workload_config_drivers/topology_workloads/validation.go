package topology_workloads

import "fmt"

func (t *TopologyBaseInfraWorkload) ValidateWorkloads() error {
	if t.Service != nil {
		seen := make(map[string]bool)
		for _, port := range t.Service.Spec.Ports {
			if seen[port.Name] {
				return fmt.Errorf("%s duplicate port name %s", t.Service.Name, port.Name)
			}
		}
	}
	if t.StatefulSet != nil && t.Deployment != nil {
		return fmt.Errorf("both sts and dep cannot be set in one workload")
	}
	if t.StatefulSet != nil {
		seen := make(map[string]bool)
		for _, cont := range t.StatefulSet.Spec.Template.Spec.Containers {
			if seen[cont.Name] {
				return fmt.Errorf("sts %s duplicate container name %s", t.StatefulSet.Name, cont.Name)
			}
		}
	}
	if t.Deployment != nil {
		seen := make(map[string]bool)
		for _, cont := range t.Deployment.Spec.Template.Spec.Containers {
			if seen[cont.Name] {
				return fmt.Errorf("dep %s duplicate container name %s", t.Deployment.Name, cont.Name)
			}
		}
	}
	return nil
}
