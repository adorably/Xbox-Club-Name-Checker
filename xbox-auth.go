package main

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func authorizeTokens(token string) {

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

	if err := fasthttp.Do(req, res); err != nil {
		// I don't want to go to school mommy
	}
	fasthttp.ReleaseRequest(req)

	body := res.Body()

	if res.StatusCode() == 200 {
		var tokenResponse tokenResponseStructure

		err := json.Unmarshal(body, &tokenResponse)
		if err == nil {
			authorizedTokens = append(authorizedTokens, "XBL3.0 x="+tokenResponse.DisplayClaims.Xui[0].Uhs+";"+tokenResponse.Token+"")
		}
	}

	fasthttp.ReleaseResponse(res)
	WaitGroup.Done()
}
