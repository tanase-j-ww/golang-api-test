package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Memo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var memoList []Memo

func createMemo(c *gin.Context) {
	var newMemo Memo
	if err := c.ShouldBindJSON(&newMemo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMemo.ID = len(memoList) + 1
	memoList = append(memoList, newMemo)

	c.JSON(http.StatusCreated, newMemo)
}

func getAllMemos(c *gin.Context) {
	c.JSON(http.StatusOK, memoList)
}

func getMemoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, memo := range memoList {
		if memo.ID == id {
			c.JSON(http.StatusOK, memo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Memo not found"})
}

func updateMemo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedMemo Memo
	if err := c.ShouldBindJSON(&updatedMemo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, memo := range memoList {
		if memo.ID == id {
			memoList[i] = updatedMemo
			c.JSON(http.StatusOK, gin.H{"message": "Memo updated successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Memo not found"})
}

func deleteMemo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, memo := range memoList {
		if memo.ID == id {
			memoList = append(memoList[:i], memoList[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Memo deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Memo not found"})
}
