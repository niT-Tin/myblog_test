package service

import "myblog/repositories"

type IItemService interface {
	SelectAllItem() repositories.Merge
	SelectItemByLimit(int) repositories.Merge
}

type ItemService struct {
	ItemRepo repositories.IItem
}

func NewItemService(itemRepo repositories.IItem) IItemService {
	return &ItemService{ItemRepo: itemRepo}
}

// SelectAllItem 查询所有显示项目数据
func (is *ItemService) SelectAllItem() repositories.Merge {
	all, err := is.ItemRepo.SelectAll()
	if err != nil {
		return NoErrorRepeat(err, "select all items error")
	}
	return repositories.Merge{
		IS: all,
	}
}

// SelectItemByLimit 根据限制查询显示项目数据
func (is *ItemService) SelectItemByLimit(limit int) repositories.Merge {
	byLimit, err := is.ItemRepo.SelectByLimit(limit)
	if err != nil {
		return NoErrorRepeat(err, "select item by limit error")
	}
	return repositories.Merge{
		IS: byLimit,
	}
}
