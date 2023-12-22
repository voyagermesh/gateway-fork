package gatewayapi

import (
	"sigs.k8s.io/gateway-api/apis/v1alpha2"
)

func (t *Translator) ProcessBackendTLSPolicies(backendTlsPolicies []*v1alpha2.BackendTLSPolicy,
	gateways []*GatewayContext,
	routes []RouteContext,
	xdsIR XdsIRMap) []*v1alpha2.BackendTLSPolicy {
	res := []*v1alpha2.BackendTLSPolicy{}
	for _, poli := range backendTlsPolicies {
		policy := poli.DeepCopy()
		if policy.Status.Ancestors != nil {
			res = append(res, policy)
			//for k, status := range policy.Status.Ancestors {
			//	for j, cond := range status.Conditions {
			//		if cond.Status == v1.ConditionUnknown {
			//			policy.Status.Ancestors[k].Conditions[j].Status = v1.ConditionTrue
			//			fmt.Println("here we go ***************** ")
			//		}
			//	}
			//}
		}
	}

	return res
}
