package group

import (
	"fmt"
	//"errors"
	"github.com/t-kimura1988/japon-domain/src/domain"
	"github.com/t-kimura1988/japon-domain/src/model/group"
	//"gorm.io/gorm"
	"github.com/t-kimura1988/japon-domain/src/exceptions"
	"net/http"
)

type GroupDetailInput struct {
	GroupId int64
}

func GroupDetailDomain(input GroupDetailInput) (*group.GroupDetail, error) {

	db := domain.DB

	group := group.GroupDetail{}
	groupJoinCntSubQ := db.Debug().Table(
		"post_groups",
	).Select(
		"count(1)",
	).Where(
		"post_groups.group_id = ?",
		input.GroupId,
	)
	db.Debug().Table(
		"groups",
	).Select(
		"groups.id as group_id, groups.group_name, groups.user_id, "+
			"users.family_name group_create_family_name, users.given_name group_create_given_name, "+
			"(?) as group_join_num",
		groupJoinCntSubQ,
	).Joins(
		"left join users on users.id = groups.user_id",
	).Find(&group, "groups.id = ?", input.GroupId)

	if group.GroupID == 0 {
		return nil, &exceptions.DataNotFound{
			Message: fmt.Sprintf("グループ詳細データは存在しません。[group_id=%d]", input.GroupId),
			Code:    http.StatusNotFound,
		}
	}

	return &group, nil
}
