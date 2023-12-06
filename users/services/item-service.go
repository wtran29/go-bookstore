package services

var (
	ItemsService itemsRepository = &itemsService{}
)

type itemsService struct{}

type itemsRepository interface {
	GetItem()
	SetItem()
}

func (i *itemsService) GetItem() {

}

func (i *itemsService) SetItem() {

}
