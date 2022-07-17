package model

//TokenDetails ...
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

//AccessDetails ...
type AccessDetails struct {
	AccessUUID string
	UserName   string
}

//Token ...
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
