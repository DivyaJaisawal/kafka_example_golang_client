package server1

/*
type GreetingRequest struct {
	GreetMessage string `json:"message"`
}

type GreetingResponse struct {
	Success bool                `json:"success"`
	Data    interface{}         `json:"data"`
	Errors  []map[string]string `json:"errors"`
}

func (c *GopayGatewayClient) ApproveAuthorisation(request *ApproveAuthorisationRequest) (*ApproveAuthorisationResponse, error) {
	logFields := util.BuildContext("GopayGatewayApproveAuthorisation")

	url := fmt.Sprintf("%s/v1/authorisations/%s?action=approve", config.GopayURL(), request.AuthorisationId)
	headers := map[string]string{
		"Content-Type":  "application/json",
		"pin":            request.Pin,
		"authorization": config.GopayAuthKey(),
	}

	body := &ApproveAuthorisationRequest {
		CustomerPaymentHandle: request.CustomerPaymentHandle,
	}

	requestByteArr, err := json.Marshal(body)
	if err != nil {
		util.Log.WithFields(logFields).Errorf("error marshalling request body %v", err)
		return nil, err
	}

	res, err := appcontext.GopayGatewayHystrixHttpClient().
		MakeHttpRequest("PATCH", url, headers, bytes.NewReader(requestByteArr))

	if err != nil {
		util.Log.WithFields(logFields).Errorf("error calling gopay gateway %v", err)
		return nil, err
	}

	defer res.Body.Close()
	var approveResponse ApproveAuthorisationResponse
	responseByteArr, err := ioutil.ReadAll(res.Body)

	if err != nil {
		util.Log.WithFields(logFields).Errorf("error reading response body %v", err)
		return nil, err
	}

	err = json.Unmarshal(responseByteArr, &approveResponse)
	if err != nil {
		util.Log.WithFields(logFields).Errorf("error un-marshalling response body %v", err)
		return nil, err
	}
	return &approveResponse, nil
}

*/