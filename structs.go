package main

type checkJson struct {
	ClubName string `json:"name"`
}

type Properties struct {
	UserTokens []string
	SandboxId  string
}

type Req struct {
	RelyingParty string
	TokenType    string
	P            Properties `json:"Properties"`
}

type tokenResponseStructure struct {
	IssueInstant  string `json:"IssueInstant"`
	NotAfter      string `json:"NotAfter"`
	Token         string `json:"Token"`
	DisplayClaims struct {
		Xui []struct {
			Uhs string `json:"uhs"`
			Agg string `json:"agg"`
		} `json:"xui"`
	} `json:"DisplayClaims"`
}
