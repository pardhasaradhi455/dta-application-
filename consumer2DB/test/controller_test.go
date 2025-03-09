package test

import (
	"delivery_tracking_api/consumer2DB/mocks"
	"delivery_tracking_api/consumer2DB/model"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInsertOrder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//create a mock order controller
	mockOrderController := mocks.NewMockControllerInterface(mockCtrl)

	mockOrderController.EXPECT().InsertOrder(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		var order model.Order
		c.ShouldBindJSON(&order)
		c.JSON(http.StatusAccepted, "Record Inserted Successfully")
	})

	//set up a router
	router := gin.Default()
	router.POST("/db/insert", mockOrderController.InsertOrder)

	orderJson := `{"OrderId": "1001", "OrderDate": "15-08-2024 15:04:05", "OrderTotal": 234.55, "Status": {"State": "placed", "StatusDate": "15-08-2024 15:04:05"}, "Address": {"City": "bangalore", "State": "karnataka", "Pincode": 560037}}`

	//create a POST request to router
	request := httptest.NewRequest(http.MethodPost, "/db/insert", nil)
	request.Body = io.NopCloser(strings.NewReader(orderJson))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	//assert the response usong response recorder
	assert.Equal(t, http.StatusAccepted, recorder.Code)
	assert.JSONEq(t, `"Record Inserted Successfully"`, recorder.Body.String())
}

func TestFetchOrderByKey(t *testing.T){
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockControllerInterface(mockCtrl)

	mockOrderController.EXPECT().FetchOrderByKey(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Order{OrderId: c.Param("id"), OrderDate: "15-08-2024 15:04:05", OrderTotal: 234.55, Status: model.Status{State: "placed", StatusDate: "15-08-2024 15:04:05"}, Address: model.Address{City: "bangalore", State: "karnataka", Pincode: 560037}})
	})

	router := gin.Default()
	router.GET("/db/fetch/:id", mockOrderController.FetchOrderByKey)

	request := httptest.NewRequest(http.MethodGet, "/db/fetch/1001", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	respOrder := `{"OrderId": "1001", "OrderDate": "15-08-2024 15:04:05", "OrderTotal": 234.55, "Status": {"State": "placed", "StatusDate": "15-08-2024 15:04:05"}, "Address": {"City": "bangalore", "State": "karnataka", "Pincode": 560037}}`
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, respOrder, recorder.Body.String())
}

func TestFetchAllOrders(t *testing.T){
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockControllerInterface(mockCtrl)

	mockOrderController.EXPECT().FetchAllOrders(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, []model.Order{
			{OrderId: "1001", OrderDate: "15-08-2024 15:04:05", OrderTotal: 234.55, Status: model.Status{State: "placed", StatusDate: "15-08-2024 15:04:05"}, Address: model.Address{City: "bangalore", State: "karnataka", Pincode: 560037}},
		})
	})

	router := gin.Default()
	router.GET("/db/fetch", mockOrderController.FetchAllOrders)

	request := httptest.NewRequest(http.MethodGet, "/db/fetch", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	expected := `[{"OrderId": "1001", "OrderDate": "15-08-2024 15:04:05", "OrderTotal": 234.55, "Status": {"State": "placed", "StatusDate": "15-08-2024 15:04:05"}, "Address": {"City": "bangalore", "State": "karnataka", "Pincode": 560037}}]`
	assert.JSONEq(t, expected, recorder.Body.String())
}