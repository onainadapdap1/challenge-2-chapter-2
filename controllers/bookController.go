package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Description string `json:"desc"`
}

var BookDatas = []Book{}

func CreateBook(c *gin.Context) {
	// set objek book
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.ID = len(BookDatas) + 1
	BookDatas = append(BookDatas, newBook)

	c.JSON(http.StatusCreated, gin.H{
		"succes": true,
		"car": newBook,
	})
	
}

func UpdateBook(c *gin.Context) {
	bookID := c.Param("bookID")
	BookID, _ := strconv.Atoi(bookID)
	condition := false
	var updatedBook Book
	
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if BookID == book.ID {
			condition = true
			BookDatas[i] = updatedBook
			BookDatas[i].ID = BookID
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data not found",
			"error_message": fmt.Sprintf("book with id %v not found", BookID),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H {
		"succes": true,
		"message": fmt.Sprintf("book with id %v has been successfully updated", BookID),
	})
}

func GetAllCar(c *gin.Context) {
	var datas []Book
	condition := false
	if len(BookDatas) > 0 {
		condition = true
		datas = append(datas, BookDatas...)
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data not found",
			"error_message": "No found data at all",
		})
		return	
	}

	c.JSON(http.StatusOK, gin.H{
		"succes": true,
		"cars": datas,
	})
}

func GetBookByID(c *gin.Context) {
	bookID := c.Param("bookID")
	BookID, _ := strconv.Atoi(bookID)
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if BookID == book.ID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
			"error_message": fmt.Sprintf("car with id %d not found", BookID),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"succes": true,
		"car": bookData,
	})
}

func DeleteBook(c *gin.Context) {
	bookID := c.Param("bookID")
	BookID, _ :=strconv.Atoi(bookID)
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if BookID == book.ID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", BookID),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", BookID),
	})
}