// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type CommonResponse struct {
	Message     string `json:"message"`
	Code        int16  `json:"code"`
	Description string `json:"description"`
}

type RegisterBaseInfo struct {
	UserName string `json:"userName"`
	PassWord string `json:"pw"`
	Role     int    `json:"role"`
}

type RegisterInfo struct {
	UserName   string `json:"userName"`
	SchoolNum  string `json:"sNo"`
	SchoolName string `json:"sName"`
	PetName    string `json:"petName,omitempty"`
	PhoneNo    string `json:"phoneNo,omitempty"`
	PassWord   string `json:"password"`
	Birthday   string `json:"birthday,omitempty"`
	Gender     int    `json:"gender,options=-1|0|1,default=-1"`
	Grade      int    `json:"grade,omitempty"`
	Role       int    `json:"role,default=0"`
}

type LoginInfo struct {
	SchoolNum string `json:"sNo"`
	PassWord  string `json:"password"`
}

type LoginResponseInfo struct {
	SchoolNo     string `json:"sNo"`
	PetName      string `json:"petName"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type LoginResponse struct {
	CommonResponse
	Data LoginResponseInfo `json:"data"`
}
