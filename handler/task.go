package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"TodoTask/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) CreateTask(c echo.Context) (err error) {
	u := &model.User{
		ID: bson.ObjectIdHex(userIDFromToken(c)),
	}
	p := &model.Task{
		ID:   bson.NewObjectId(),
		UserId: u.ID.Hex(),
	}
	if err = c.Bind(p); err != nil {
		return
	}

	// Validation
	if p.TaskName == "" || p.Description == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid  fields"}
	}

	// Find user from database
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("TodoTask").C("users").FindId(u.ID).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
		}
		return
	}

	// Save post in database
	if err = db.DB("TodoTask").C("tasks").Insert(p); err != nil {
		return
	}
	return c.JSON(http.StatusCreated, p)
}

func (h *Handler) FetchTasks(c echo.Context) (err error) {
	userID := userIDFromToken(c)
	// Defaults
	
	// Retrieve posts from database
	tasks := []*model.Task{}
	db := h.DB.Clone()
	if err = db.DB("TodoTask").C("tasks").
		Find(bson.M{"UserId": userID}).
		All(&tasks); err != nil {
		return
	}
	defer db.Close()

	return c.JSON(http.StatusOK, tasks)
}

func (h *Handler) UpdateTask(c echo.Context) (err error) {
	userID := userIDFromToken(c)
	id := c.Param("id")
	p := &model.Task{
	}
	if err = c.Bind(p); err != nil {
		return
	}
	p.UserId = userID
	// Add a follower to user

	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("TodoTask").C("tasks").
		UpdateId(bson.ObjectIdHex(id), p); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
			//return c.JSON(http.StatusOK, p)
		}
	}
	return c.JSON(http.StatusOK, p)
}

func (h *Handler) CompleteTask(c echo.Context) (err error) {
	id := c.Param("id")
	db := h.DB.Clone()
	defer db.Close()

	if err = db.DB("TodoTask").C("tasks").
		UpdateId(bson.ObjectIdHex(id), bson.M{"$set" : bson.M{"Status" : 1}}); err != nil {
		if err == mgo.ErrNotFound {
			return echo.ErrNotFound
			//return c.JSON(http.StatusOK, p)
		}
	}
	return c.JSON(http.StatusOK, "Task completed sucessfully")
}
