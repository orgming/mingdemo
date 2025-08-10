package demo

import (
	demoService "github.com/orgming/mingdemo/app/provider/demo"
	"github.com/orgming/mingdemo/framework/contract"
	"github.com/orgming/mingdemo/framework/gin"
)

type DemoApi struct {
	service *Service
}

func Register(r *gin.Engine) error {
	api := NewDemoApi()
	r.Bind(&demoService.DemoProvider{})

	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.POST("/demo/demo3", api.DemoPost)

	return nil
}

func NewDemoApi() *DemoApi {
	return &DemoApi{service: NewService()}
}

// Demo
// @Summary Demo
// @Description Demo
// @Tags Demo
// @Accept json
// @Produce json
// @Success 200 {object} string "{"code":200,"msg":"success","data":""}"
// @Router /demo [get]
func (api *DemoApi) Demo(c *gin.Context) {
	appService := c.MustMake(contract.AppKey).(contract.App)
	baseFolder := appService.BaseFolder()

	c.JSON(200, baseFolder)
}

func (api *DemoApi) Demo2(c *gin.Context) {
	demoProvider := c.MustMake(demoService.DemoKey).(demoService.IService)
	students := demoProvider.GetAllStudent()
	userDTOs := StudentsToUserDTOs(students)

	c.JSON(200, userDTOs)
}

func (api *DemoApi) DemoPost(c *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, nil)
}
