package test

import (
	"delivery_tracking_api/consumer1/mocks"
	"delivery_tracking_api/consumer1/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetOrder(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().GetOrder(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Order{OrderId: c.Param("id"), OrderDate: "15-08-2024 15:04:05", OrderTotal: 234.55, Status: model.Status{State: "placed", StatusDate: "15-08-2024 15:04:05"}, Address: model.Address{City: "bangalore", State: "karnataka", Pincode: 560037}})
	})

	router := gin.Default()
	router.GET("/consumer/order/:id", mockOrderController.GetOrder)

	request := httptest.NewRequest(http.MethodGet, "/consumer/order/1001", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	respOrder := `{"OrderId": "1001", "OrderDate": "15-08-2024 15:04:05", "OrderTotal": 234.55, "Status": {"State": "placed", "StatusDate": "15-08-2024 15:04:05"}, "Address": {"City": "bangalore", "State": "karnataka", "Pincode": 560037}}`
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, respOrder, recorder.Body.String())
}

func TestGetOrder2(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().GetOrder(gomock.Any()).DoAndReturn(func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Order{OrderId: c.Param("id"), OrderDate: "15-08-2024 15:04:05", OrderTotal: 234.55, Status: model.Status{State: "placed", StatusDate: "15-08-2024 15:04:05"}, Address: model.Address{City: "bangalore", State: "karnataka", Pincode: 560037}})
	})

	router := gin.Default()
	router.GET("/consumer/order/:id", mockOrderController.GetOrder)

	request := httptest.NewRequest(http.MethodGet, "/consumer/order/1002", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	respOrder := `{"OrderId": "1001", "OrderDate": "15-08-2024 15:04:05", "OrderTotal": 234.55, "Status": {"State": "placed", "StatusDate": "15-08-2024 15:04:05"}, "Address": {"City": "bangalore", "State": "karnataka", "Pincode": 560037}}`
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, respOrder, recorder.Body.String())
}
func TestGetStatus(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().GetStatus(gomock.Any()).DoAndReturn(func(c *gin.Context){
		c.HTML(http.StatusOK,"placed.html", gin.H{
			"id" : c.Param("id"),
			"orderDate" : "15-08-2024 15:04:05",
			"status" : "placed",
			"statusDate" : "15-08-2024 15:04:05",
			"total" : 234.55,
			"city" : "bangalore",
			"state" : "karnataka",
			"pincode": 560037,
		})
	})

	router := gin.Default()
	router.LoadHTMLGlob("C:/Users/SARADHM/Downloads/practice/delivery-tracking-demo/Delivery_Tracking_API/consumer1/templates/*")
	router.GET("/consumer/status/:id", mockOrderController.GetStatus)

	request := httptest.NewRequest(http.MethodGet, "/consumer/status/1001", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), `placed`)
}

func TestGetStatus2(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().GetStatus(gomock.Any()).DoAndReturn(func(c *gin.Context){
		c.HTML(http.StatusOK,"delivered.html", gin.H{
			"id" : c.Param("id"),
			"orderDate" : "15-08-2024 15:04:05",
			"status" : "delivered",
			"statusDate" : "15-08-2024 15:04:05",
			"total" : 234.55,
			"city" : "bangalore",
			"state" : "karnataka",
			"pincode": 560037,
		})
	})

	router := gin.Default()
	router.LoadHTMLGlob("C:/Users/SARADHM/Downloads/practice/delivery-tracking-demo/Delivery_Tracking_API/consumer1/templates/*")
	router.GET("/consumer/status/:id", mockOrderController.GetStatus)

	request := httptest.NewRequest(http.MethodGet, "/consumer/status/1001", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), `delivered`)
}

func TestGetStatus3(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockOrderController := mocks.NewMockOrderControllerInterface(mockCtrl)

	mockOrderController.EXPECT().GetStatus(gomock.Any()).DoAndReturn(func(c *gin.Context){
		c.HTML(http.StatusOK,"shipped.html", gin.H{
			"id" : c.Param("id"),
			"orderDate" : "15-08-2024 15:04:05",
			"status" : "shipped",
			"statusDate" : "15-08-2024 15:04:05",
			"total" : 234.55,
			"city" : "bangalore",
			"state" : "karnataka",
			"pincode": 560037,
		})
	})

	router := gin.Default()
	router.LoadHTMLGlob("C:/Users/SARADHM/Downloads/practice/delivery-tracking-demo/Delivery_Tracking_API/consumer1/templates/*")
	router.GET("/consumer/status/:id", mockOrderController.GetStatus)

	request := httptest.NewRequest(http.MethodGet, "/consumer/status/1001", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), `delivered`)
}
