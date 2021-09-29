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

// func (CompanyController CompanyController) Login(c echo.Context) error {
// 	fmt.Println("Login")
// 	CompanyLogin := requests.CompanyLogin{}
// 	c.Bind(&CompanyLogin)

// 	ctx := c.Request().Context()
// 	Company, error := CompanyController.CompanyUseCase.Login(ctx, CompanyLogin.Email, CompanyLogin.Password)

// 	if error != nil {
// 		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
// 	}

// 	return controllers.NewSuccessResponse(c, responses.FromDomain(Company))
// }

func (CompanyController CompanyController) GetCompanyById(c echo.Context) error {
	fmt.Println("GetById")

	CompanyId, err := strconv.Atoi(c.Param("CompanyId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	Company, err := CompanyController.CompanyUseCase.GetCompanyById(ctx, CompanyId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(Company))
}

func (CompanyController CompanyController) GetAllCompany(c echo.Context) error {
	fmt.Println("GetAllCompany")

	ctx := c.Request().Context()
	Company, err := CompanyController.CompanyUseCase.GetAllCompany(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.ListFromDomainAll(Company))
}

func (CompanyController CompanyController) UpdateCompany(c echo.Context) error {
	fmt.Println("UpdateCompany")

	CompanyUpdate := requests.CompanyUpdate{}
	c.Bind(&CompanyUpdate)

	ctx := c.Request().Context()
	Company, err := CompanyController.CompanyUseCase.UpdateCompany(ctx, CompanyUpdate.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(Company))
}

func (CompanyController CompanyController) DeleteCompany(c echo.Context) error {
	fmt.Println("DeleteCompany")

	CompanyId, err := strconv.Atoi(c.Param("CompanyId"))
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	ctx := c.Request().Context()
	Company, err := CompanyController.CompanyUseCase.DeleteCompany(ctx, CompanyId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(Company))
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

func (CompanyController CompanyController) HardDeleteAllCompanies(c echo.Context) error {
	fmt.Println("HardDeleteAllcompanies")

	ctx := c.Request().Context()
	err := CompanyController.CompanyUseCase.HardDeleteAllCompanies(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, responses.FromDomain(companies.Domain{}))
}
