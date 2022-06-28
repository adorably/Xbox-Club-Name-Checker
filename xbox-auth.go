package main

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func authorizeTokens(token string) {
	defer WaitGroup.Done()
	per := Properties{
		UserTokens: []string{token},
		SandboxId:  "RETAIL",
	}

	group := Req{
		RelyingParty: "http://xboxlive.com",
		TokenType:    "JWT",
		P:            per,
	}

	b, _ := json.Marshal(group)

	req := fasthttp.AcquireRequest()
	req.SetBody(b)
	req.Header.SetMethod("POST")
	req.SetRequestURI("https://xsts.auth.xboxlive.com/xsts/authorize")
	res := fasthttp.AcquireResponse()
	
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	if err := fasthttp.Do(req, res); err != nil {
		return
	}
	

	body := res.Body()

	if res.StatusCode() == 200 {
		var tokenResponse tokenResponseStructure

		err := json.Unmarshal(body, &tokenResponse)
		if err == nil {
			authorizedTokens = append(authorizedTokens, "XBL3.0 x="+tokenResponse.DisplayClaims.Xui[0].Uhs+";"+tokenResponse.Token+"")
		}
	}
}
