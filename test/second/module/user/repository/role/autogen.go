package role

import dto "metis/test/second/model/dto"

type iAutoGen interface {
	SelectByIDs(ids ...int64) []*entity.Role
	SelectByID(id int64) dto.Role
}
