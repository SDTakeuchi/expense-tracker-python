package infra

// postAlbums adds an album from JSON received in the request body.
func createItem(c *gin.Context) {
	// validate submitted data here
	newItem := Item{}
	now := time.Now()
	// newItem.User = c.User()
	newItem.CreatedAt = now
	newItem.UpdatedAt = now

	// Call BindJSON to bind the received JSON to newItem.
	err := c.BindJSON(&newItem)
	if err != nil {
		c.String(http.StatusBadRequest, "Request failed: "+err.Error())
	}

	db.NewRecord(newItem)
	db.Create(&newItem)
	if !db.NewRecord(newItem) {
		c.JSON(http.StatusOK, newItem)
	}
}

// getItems responds with the list of all Items as JSON.
func getItems(c *gin.Context) {
	items := []Item{}
	err := db.Find(&items).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
	} else {
		c.IndentedJSON(http.StatusOK, items)
	}
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getItemByID(c *gin.Context) {
	id := c.Param("id")
	item := Item{}

	err := db.Where("ID = ?", id).First(&item).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item not found"})
	} else {
		c.IndentedJSON(http.StatusOK, item)
	}
}

func updateAlbumByID(c *gin.Context) {
	album := Album{}
	id := c.Param("id")

	data := Album{}
	err := c.BindJSON(&data)
	if err != nil {
		c.String(http.StatusBadRequest, "Request failed"+err.Error())
	}

	db.Where("ID = ?", id).First(&album).Updates(&data)
}

func deleteAlbumByID(c *gin.Context) {
	album := Album{}
	id := c.Param("id")

	db.Where("ID = ?", id).Delete(&album)
}
