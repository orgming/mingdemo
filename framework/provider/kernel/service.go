package kernel

import (
	"github.com/orgming/ming/framework/gin"
	"net/http"
)

type MingKernelService struct {
	engine *gin.Engine
}

// NewMingKernelService 初始化 web 引擎服务实例
func NewMingKernelService(params ...any) (any, error) {
	return &MingKernelService{engine: params[0].(*gin.Engine)}, nil
}

func (s *MingKernelService) HttpEngine() http.Handler {
	return s.engine
}
