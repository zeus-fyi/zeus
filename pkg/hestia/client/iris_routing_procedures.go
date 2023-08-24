package hestia_client

//func (h *Hestia) CreateIrisRoutingProcedure(ctx context.Context, rr hestia_req_types.IrisRoutingProcedureRequest) error {
//	h.PrintReqJson(rr)
//	resp, err := h.R().
//		SetBody(rr).
//		Post(hestia_endpoints.IrisCreateProcedurePath)
//	if err != nil || resp.StatusCode() >= 400 {
//		if err == nil {
//			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
//		}
//		log.Err(err).Msg("Hestia: CreateIrisRoutingProcedure")
//		return err
//	}
//	h.PrintRespJson(resp.Body())
//	return err
//}
