package db

import (
	"go-server/pkg/model"

	"github.com/pkg/errors"
)

func (h *DBHandler) CreateBoard(board *model.Board) (*model.Board, error) {
	result := h.gDB.Create(board)

	return board, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) GetBoardList() ([]*model.Board, error) {
	boardList := []*model.Board{}
	result := h.gDB.Find(&boardList)

	return boardList, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) UpdateBoard(id uint, newBoard *model.Board) (*model.Board, error) {
	oldBoard := &model.Board{}
	result := h.gDB.Model(oldBoard).Where(id).Updates(newBoard)

	return oldBoard, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) DeleteBoardByID(id string) error {

	result := h.gDB.Delete(&model.Board{}, id)

	return errors.Wrap(result.Error, "db handler error")
}
