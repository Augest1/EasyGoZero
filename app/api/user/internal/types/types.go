// Code generated by goctl. DO NOT EDIT.
package types

type ReqLogin struct {
	IdentityKind int64  `json:"identityKind"`
	Identifier   string `json:"identifier"`
	Credential   string `json:"credential"`
}

type ResLogin struct {
	Token string `json:"token"`
}

type ResUserInfo struct {
}
