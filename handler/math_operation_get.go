package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gocalc/handler/operations"
	"math"
	"net/http"
	"strconv"
)

type response struct {
	Success bool    `json:"Success"`
	ErrCode int     `json:"ErrCode"`
	Value   float64 `json:"Value"`
}

const (
	ErrCodeNone = iota
	ErrCodeExactlyTwoArgumentsRequired
	ErrCodeInappropriateParams
	ErrCodeFailedToParse
	ErrCodeInfiniteValue
	ErrCodeNaNValue
)

func respond(c *gin.Context, status int, errCode int, value float64) {
	r := response{
		Success: errCode == ErrCodeNone,
		ErrCode: errCode,
		Value:   value,
	}
	responseBytes, _ := json.Marshal(r)
	c.String(status, string(responseBytes))
}

func MathOperationGet(op operations.Operation) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()

		if len(params) != 2 {
			respond(c, http.StatusBadRequest, ErrCodeExactlyTwoArgumentsRequired, 0)
			return
		}

		aValues, aValuesOk := params["a"]
		bValues, bValuesOk := params["b"]
		if !aValuesOk || !bValuesOk {
			respond(c, http.StatusBadRequest, ErrCodeInappropriateParams, 0)
			return
		}
		if len(aValues) != 1 && len(bValues) != 1 {
			respond(c, http.StatusBadRequest, ErrCodeInappropriateParams, 0)
			return
		}

		aString := aValues[0]
		bString := bValues[0]
		a, errA := strconv.ParseFloat(aString, 64)
		b, errB := strconv.ParseFloat(bString, 64)
		if errA != nil && errB != nil {
			respond(c, http.StatusBadRequest, ErrCodeFailedToParse, 0)
			return
		}

		value := op.Perform(a, b)
		if math.IsInf(value, 0) {
			respond(c, http.StatusBadRequest, ErrCodeInfiniteValue, 0)
			return
		}
		if math.IsNaN(value) {
			respond(c, http.StatusBadRequest, ErrCodeNaNValue, 0)
			return
		}

		respond(c, http.StatusOK, ErrCodeNone, value)
	}
}
