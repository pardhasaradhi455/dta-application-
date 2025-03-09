package test

import (
	"delivery_tracking_api/producer/mocks"
	"delivery_tracking_api/producer/model"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPlaceOrder(t *testing.T) {
	//set a mockcontrol
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//create a mock order controller
	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().PlaceOrder(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		var order model.Order
		c.ShouldBindJSON(&order)
		c.JSON(http.StatusAccepted, "order successfully added with id : "+order.OrderId)
	})

	//set up a router
	router := gin.Default()
	router.POST("/producer/place", mockOrderController.PlaceOrder)

	orderJson := `{"OrderId": "1001", "OrderDate": "15-08-2024 15:04:05", "OrderTotal": 234.55, "Status": {"State": "placed", "StatusDate": "15-08-2024 15:04:05"}, "Address": {"City": "bangalore", "State": "karnataka", "Pincode": 560037}}`

	//create a POST request to router
	request := httptest.NewRequest(http.MethodPost, "/producer/place", nil)
	request.Body = io.NopCloser(strings.NewReader(orderJson))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	//assert the response usong response recorder
	assert.Equal(t, http.StatusAccepted, recorder.Code)
	assert.JSONEq(t, `"order successfully added with id : 1001"`, recorder.Body.String())
}

func TestPlaceOrder2(t *testing.T) {
	//set a mockcontrol
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//create a mock order controller
	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().PlaceOrder(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		var order model.Order
		c.ShouldBindJSON(&order)
		c.JSON(http.StatusAccepted, "order successfully added with id : "+order.OrderId)
	})

	//set up a router
	router := gin.Default()
	router.POST("/producer/place", mockOrderController.PlaceOrder)

	orderJson := `{"OrderId": "1002", "OrderDate": "15-08-2024 15:04:05", "OrderTotal": 234.55, "Status": {"State": "placed", "StatusDate": "15-08-2024 15:04:05"}, "Address": {"City": "bangalore", "State": "karnataka", "Pincode": 560037}}`

	//create a POST request to router
	request := httptest.NewRequest(http.MethodPost, "/producer/place", nil)
	request.Body = io.NopCloser(strings.NewReader(orderJson))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	//assert the response usong response recorder
	assert.Equal(t, http.StatusAccepted, recorder.Code)
	assert.JSONEq(t, `"order successfully added with id : 1001"`, recorder.Body.String())
}

func TestGetPendingOrders(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().GetPendingOrders(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, []model.Order{
			{OrderId: "1001", OrderDate: "15-08-2024 15:04:05", OrderTotal: 234.55, Status: model.Status{State: "placed", StatusDate: "15-08-2024 15:04:05"}, Address: model.Address{City: "bangalore", State: "karnataka", Pincode: 560037}},
		})
	})

	router := gin.Default()
	router.GET("/producer/pending", mockOrderController.GetPendingOrders)

	request := httptest.NewRequest(http.MethodGet, "/producer/pending", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expected := `[{"OrderId": "1001", "OrderDate": "15-08-2024 15:04:05", "OrderTotal": 234.55, "Status": {"State": "placed", "StatusDate": "15-08-2024 15:04:05"}, "Address": {"City": "bangalore", "State": "karnataka", "Pincode": 560037}}]`
	assert.JSONEq(t, expected, recorder.Body.String())
}

func TestGetPendingOrders2(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().GetPendingOrders(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, []model.Order{
			{OrderId: "1001", OrderDate: "15-08-2024 15:04:05", OrderTotal: 234.55, Status: model.Status{State: "delivered", StatusDate: "15-08-2024 15:04:05"}, Address: model.Address{City: "bangalore", State: "karnataka", Pincode: 560037}},
		})
	})

	router := gin.Default()
	router.GET("/producer/pending", mockOrderController.GetPendingOrders)

	request := httptest.NewRequest(http.MethodGet, "/producer/pending", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expected := `[{"OrderId": "1001", "OrderDate": "15-08-2024 15:04:05", "OrderTotal": 234.55, "Status": {"State": "placed", "StatusDate": "15-08-2024 15:04:05"}, "Address": {"City": "bangalore", "State": "karnataka", "Pincode": 560037}}]`
	assert.JSONEq(t, expected, recorder.Body.String())
}

func TestGetDeliveredOrders(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().GetDeliveredOrders(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, []model.Order{
			{OrderId: "1001", OrderDate: "15-08-2024 15:04:05", OrderTotal: 234.55, Status: model.Status{State: "delivered", StatusDate: "15-08-2024 15:04:05"}, Address: model.Address{City: "bangalore", State: "karnataka", Pincode: 560037}},
		})
	})

	router := gin.Default()
	router.GET("/producer/delivered", mockOrderController.GetDeliveredOrders)

	request := httptest.NewRequest(http.MethodGet, "/producer/delivered", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expected := `[{"OrderId": "1001", "OrderDate": "15-08-2024 15:04:05", "OrderTotal": 234.55, "Status": {"State": "delivered", "StatusDate": "15-08-2024 15:04:05"}, "Address": {"City": "bangalore", "State": "karnataka", "Pincode": 560037}}]`
	assert.JSONEq(t, expected, recorder.Body.String())
}

func TestGetDeliveredOrders2(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().GetDeliveredOrders(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, []model.Order{
			{OrderId: "1001", OrderDate: "15-08-2024 15:04:05", OrderTotal: 234.55, Status: model.Status{State: "shipped", StatusDate: "15-08-2024 15:04:05"}, Address: model.Address{City: "bangalore", State: "karnataka", Pincode: 560037}},
		})
	})

	router := gin.Default()
	router.GET("/producer/delivered", mockOrderController.GetDeliveredOrders)

	request := httptest.NewRequest(http.MethodGet, "/producer/delivered", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expected := `[{"OrderId": "1001", "OrderDate": "15-08-2024 15:04:05", "OrderTotal": 234.55, "Status": {"State": "delivered", "StatusDate": "15-08-2024 15:04:05"}, "Address": {"City": "bangalore", "State": "karnataka", "Pincode": 560037}}]`
	assert.JSONEq(t, expected, recorder.Body.String())
}

func TestChangeState(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().ChangeState(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusAccepted, "order state changed successfully with id : "+c.Query("id"))
	})

	router := gin.Default()
	router.GET("/producer/changeState", mockOrderController.ChangeState)

	request := httptest.NewRequest(http.MethodGet, "/producer/changeState?id=1001&state=out-for-delivery", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusAccepted, recorder.Code)
	assert.JSONEq(t, `"order state changed successfully with id : 1001"`, recorder.Body.String())
}

func TestChangeState2(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().ChangeState(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusAccepted, "order state changed successfully with id : "+c.Query("id"))
	})

	router := gin.Default()
	router.GET("/producer/changeState", mockOrderController.ChangeState)

	request := httptest.NewRequest(http.MethodGet, "/producer/changeState?id=1002&state=out-for-delivery", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusAccepted, recorder.Code)
	assert.JSONEq(t, `"order state changed successfully with id : 1001"`, recorder.Body.String())
}

func TestChangeAddress(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().ChangeAddress(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusAccepted, "order address changed successfully with id : "+c.Param("id"))
	})

	router := gin.Default()
	router.POST("/producer/changeAddress/:id", mockOrderController.ChangeAddress)

	addressJson := `{"City": "bangalore", "State": "karnataka", "Pincode": 560037}`

	request := httptest.NewRequest(http.MethodPost, "/producer/changeAddress/1001", nil)
	request.Body = io.NopCloser(strings.NewReader(addressJson))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusAccepted, recorder.Code)
	assert.JSONEq(t, `"order address changed successfully with id : 1001"`, recorder.Body.String())
}

func TestChangeAddress2(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().ChangeAddress(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusAccepted, "order address changed successfully with id : "+c.Param("id"))
	})

	router := gin.Default()
	router.POST("/producer/changeAddress/:id", mockOrderController.ChangeAddress)

	addressJson := `{"City": "bangalore", "State": "karnataka", "Pincode": 560037}`

	request := httptest.NewRequest(http.MethodPost, "/producer/changeAddress/1002", nil)
	request.Body = io.NopCloser(strings.NewReader(addressJson))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusAccepted, recorder.Code)
	assert.JSONEq(t, `"order address changed successfully with id : 1001"`, recorder.Body.String())
}
