package companies

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"main.go/business/companies"
	"main.go/controllers"
	"main.go/controllers/companies/requests"
	"main.go/controllers/companies/responses"
)

type CompanyController struct {
	CompanyUseCase companies.Usecase
}

func NewCompanyController(CompanyUseCase companies.Usecase) *CompanyController {
	return &CompanyController{
		CompanyUseCase: CompanyUseCase,
	}
}

func (CompanyController CompanyController) GetCompanyById(c echo.Context) error {
	fmt.Println("GetById")

	companyId, err := strconv.Atoi(c.Param("companyId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	company, err := CompanyController.CompanyUseCase.GetCompanyById(ctx, companyId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(company))
}

func (CompanyController CompanyController) GetAllCompany(c echo.Context) error {
	fmt.Println("GetAllCompany")

	ctx := c.Request().Context()
	company, err := CompanyController.CompanyUseCase.GetAllCompany(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ListFromDomainAll(company))
}

func (CompanyController CompanyController) UpdateCompany(c echo.Context) error {
	fmt.Println("UpdateCompany")

	companyUpdate := requests.CompanyUpdate{}
	c.Bind(&companyUpdate)

	ctx := c.Request().Context()
	company, err := CompanyController.CompanyUseCase.UpdateCompany(ctx, companyUpdate.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(company))
}

func (CompanyController CompanyController) DeleteCompany(c echo.Context) error {
	fmt.Println("DeleteCompany")

	companyId, err := strconv.Atoi(c.Param("companyId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	company, err := CompanyController.CompanyUseCase.DeleteCompany(ctx, companyId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(company))
}

func (CompanyController CompanyController) RegisterCompany(c echo.Context) error {
	fmt.Println("RegisterCompany")
	companyRegister := requests.CompanyCreate{}
	c.Bind(&companyRegister)

	ctx := c.Request().Context()
	company, error := CompanyController.CompanyUseCase.RegisterCompany(ctx, companyRegister.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(company))
}
