package services

import (
	"fmt"
	"go_restapi/domain/item/models"
	"go_restapi/domain/item/repositories"
	"go_restapi/helpers"

	"gorm.io/gorm"
)

type itemService struct {
	itemRepo repositories.ItemRepository
}

// Create implements ItemService.
func (service *itemService) Create(item models.Item) helpers.Response {
	var response helpers.Response

	err := service.itemRepo.Create(item)

	if err != nil {
		response.Status = 500
		response.Message = "Failed to create item"
	} else {
		response.Status = 200
		response.Message = "Success to create item"
	}

	return response
}

// Delete implements ItemService.
func (service *itemService) Delete(idItem int) helpers.Response {
	var response helpers.Response

	err := service.itemRepo.Delete(idItem)

	if err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to delete item : ", idItem)
	} else {
		response.Status = 200
		response.Message = "Success to delete item"
	}

	return response
}

// GetAll implements ItemService.
func (service *itemService) GetAll() helpers.Response {
	var response helpers.Response

	data, err := service.itemRepo.GetAll()

	if err != nil {
		response.Status = 500
		response.Message = "Failed to get all items"
	} else {
		response.Status = 200
		response.Message = "Success to get all items"
		response.Data = data
	}

	return response
}

// GetById implements ItemService.
func (service *itemService) GetById(idItem int) helpers.Response {
	var response helpers.Response

	data, err := service.itemRepo.GetById(idItem)

	if err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to get item: ", idItem)
	} else {
		response.Status = 200
		response.Message = "Success to get all items"
		response.Data = data
	}

	return response
}

// Update implements ItemService.
func (service *itemService) Update(idItem int, item models.Item) helpers.Response {
	var response helpers.Response

	err := service.itemRepo.Update(idItem, item)

	if err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to update item: ", idItem)
	} else {
		response.Status = 200
		response.Message = "Success to update item"
	}

	return response
}

type ItemService interface {
	Create(item models.Item) helpers.Response
	Update(idItem int, item models.Item) helpers.Response
	Delete(idItem int) helpers.Response
	GetById(idItem int) helpers.Response
	GetAll() helpers.Response
}

func NewItemService(db *gorm.DB) ItemService {
	return &itemService{itemRepo: repositories.NewItemRepository(db)}
}
