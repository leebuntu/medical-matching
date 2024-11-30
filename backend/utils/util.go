package utils

import (
	"medical-matching/constants"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SortMapByValueAndGetKeys(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})
	return keys
}

func ParseGEOQuery(ctx *gin.Context) (float64, float64, float64, error) {
	longitude, err := strconv.ParseFloat(ctx.Query("longitude"), 64)
	if err != nil {
		return 0, 0, 0, err
	}
	latitude, err := strconv.ParseFloat(ctx.Query("latitude"), 64)
	if err != nil {
		return 0, 0, 0, err
	}
	radius, err := strconv.ParseFloat(ctx.Query("radius"), 64)
	if err != nil {
		return 0, 0, 0, err
	}

	return longitude, latitude, radius, nil
}

func CheckBindData(c *gin.Context, bindData interface{}) (int, error) {
	userID := c.GetInt("userID")

	if bindData != nil {
		if err := c.ShouldBindJSON(bindData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": constants.BadRequest})
			return 0, err
		}
	}

	return userID, nil
}
