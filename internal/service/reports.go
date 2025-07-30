package service

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// функция GetReport получаем report по id
//
//localhost:8000/api/report/:id
func (s *Service) GetReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.reportsRepo
	report, err := repo.GetReportById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.JSON(http.StatusOK, Response{Object: report})
}

// CreateReport добавляет в базу report (created_at добавляется, updated_at остается пустым)
// localhost:8000/api/reports
func (s *Service) CreateReport(c echo.Context) error {
	var reportOne Report
	err := c.Bind(&reportOne)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	err = repo.RCreateReport(reportOne.Title, reportOne.Description)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.String(http.StatusOK, "Ok")
}

// UpdateReport для обновления report (добавляется update_at)
//
//localhost:8000/api/report/:id
func (s *Service) UpdateReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	var reportOne Report

	if err := c.Bind(&reportOne); err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	repo := s.reportsRepo
	err = repo.RUpdateReport(id, reportOne.Title, reportOne.Description)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

// DeleteReport удаляет report по id
//
//localhost:8000/api/report/:id
func (s *Service) DeleteReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	repo := s.reportsRepo
	err = repo.DeleteReportById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}
