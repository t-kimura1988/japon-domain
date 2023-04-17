package group

type GroupDetail struct {
	GroupID               int64  `json:"group_id"`
	UserId                int64  `json:"create_user_id"`
	GroupJoinNum          int32  `json:"group_join_num"`
	GroupName             string `json:"group_name"`
	GroupDetail           string `json:"group_detail"`
	GroupCreateFamilyName string `json:"group_create_family_name"`
	GroupCreateGivenName  string `json:"group_create_given_name"`
}
