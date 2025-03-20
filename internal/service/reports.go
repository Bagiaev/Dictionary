package service

import (
	"database/sql"
	"dictionary/internal/reports"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ReportsService struct {
	db          *sql.DB
	logger      echo.Logger
	reportsRepo *reports.Repo
}

func (s *ReportsService) GetReportById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}
	report, err := s.reportsRepo.GetReportById(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}
	return c.JSON(http.StatusOK, Response{Object: report})
}

func (s *ReportsService) CreateReport(c echo.Context) error {
	var report reports.Report
	err := c.Bind(&report)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	err = s.reportsRepo.CreateReport(report.Title, report.Description)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

func (s *ReportsService) UpdateReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	_, err = s.reportsRepo.GetReportById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(s.NewError(InvalidParams))
		}
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	var report reports.Report
	err = c.Bind(&report)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	err = s.reportsRepo.UpdateReport(id, report.Title, report.Description)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}

func (s *ReportsService) DeleteReport(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InvalidParams))
	}

	_, err = s.reportsRepo.GetReportById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(s.NewError(InvalidParams))
		}
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	err = s.reportsRepo.DeleteReport(id)
	if err != nil {
		s.logger.Error(err)
		return c.JSON(s.NewError(InternalServerError))
	}

	return c.String(http.StatusOK, "OK")
}
