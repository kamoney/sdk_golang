package private_kamoney

// func (s *privateRequests) CreateOrder(in kamoney_sdk_dtos.CreateOrderRequestParams) (out kamoney_sdk_dtos.CreateOrderRequestResponse, err error) {
// 	in.Nonce = fmt.Sprint(utility.GenNonce())

// 	req, err := s.r.RequestHandler("POST", ENDPOINT_ORDER, in)
// 	if err != nil {
// 		log.Panicln("CO 01: ", err.Error())
// 		return
// 	}

// 	client := &http.Client{}

// 	queryStr := s.gerQueryString(req.URL.Query(), map[string]string{
// 		"nonce":    in.Nonce,
// 		"password": in.Password,
// 	})

// 	req.URL.RawQuery = queryStr
// 	s.r.signRequest(req)

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Panicln("CO 02: ", err.Error())
// 		return
// 	}
// 	defer resp.Body.Close()

// 	bodyBytes, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Panicln("CO 03: ", err.Error())
// 		return
// 	}
// 	fmt.Println(string(bodyBytes))
// 	err = json.Unmarshal(bodyBytes, &out)
// 	if err != nil {
// 		log.Println("CO 04: ", err.Error())
// 		return
// 	}

// 	return
// }
