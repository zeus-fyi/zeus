package iris_programmable_proxy_v1_beta

//
//type IrisV1Beta struct {
//	iris_programmable_proxy.Iris
//}
//
//func NewIrisV1BetaClient(irisBase iris_programmable_proxy.Iris) IrisV1Beta {
//	return IrisV1Beta{
//		irisBase,
//	}
//}

//func (i *IrisV1Beta) CreateProcedure(ctx context.Context, procedure IrisRoutingProcedure) error {
//	hc := hestia_client.NewHestia(i.GetHestiaRoute(), i.Token)
//	rr := hestia_req_types.IrisRoutingProcedureRequest{
//		ProcedureName: procedure.Name,
//	}
//	err := hc.CreateIrisRoutingProcedure(ctx, rr)
//	if err != nil {
//		return err
//	}
//	return nil
//}
