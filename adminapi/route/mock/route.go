package mock

import (
	"github.com/ServiceComb/go-chassis/core/config/model"
	"github.com/stretchr/testify/mock"
)

//RouterMock ..
type RouterMock struct {
	mock.Mock
}

//Init ..
func (m *RouterMock) Init() error {
	return nil
}

//SetRouteRule ..
func (m *RouterMock) SetRouteRule(map[string][]*model.RouteRule) {}

//FetchRouteRule ..
func (m *RouterMock) FetchRouteRule() map[string][]*model.RouteRule {
	return nil
}

//FetchRouteRuleByServiceName ..
func (m *RouterMock) FetchRouteRuleByServiceName(s string) []*model.RouteRule {
	args := m.Called(s)
	rules, ok := args.Get(0).([]*model.RouteRule)
	if !ok {
		return nil
	}
	return rules
}
