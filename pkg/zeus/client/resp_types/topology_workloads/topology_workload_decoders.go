package topology_workloads

import (
	"encoding/json"
	"errors"

	"github.com/rs/zerolog/log"
	v1 "k8s.io/api/apps/v1"
	v1core "k8s.io/api/core/v1"
	v1networking "k8s.io/api/networking/v1"
)

func (t *TopologyBaseInfraWorkload) DecodeBytes(jsonBytes []byte) error {
	metaType, err := t.IdWorkloadFromBytes(jsonBytes)
	if err != nil {
		log.Err(err).Msg("TopologyBaseInfraWorkload: DecodeBytes")
		return err
	}
	switch metaType.Kind {
	case "Deployment":
		t.Deployment = &v1.Deployment{}
		err = json.Unmarshal(jsonBytes, t.Deployment)
	case "StatefulSet":
		t.StatefulSet = &v1.StatefulSet{}
		err = json.Unmarshal(jsonBytes, t.StatefulSet)
	case "ConfigMap":
		t.ConfigMap = &v1core.ConfigMap{}
		err = json.Unmarshal(jsonBytes, t.ConfigMap)
	case "Service":
		t.Service = &v1core.Service{}
		err = json.Unmarshal(jsonBytes, t.Service)
	case "Ingress":
		t.Ingress = &v1networking.Ingress{}
		err = json.Unmarshal(jsonBytes, t.Ingress)
	default:
		err = errors.New("TopologyBaseInfraWorkload: DecodeBytes, no matching kind found")
		log.Err(err).Msg("TopologyBaseInfraWorkload: DecodeBytes, no matching kind found")
	}
	return err
}
