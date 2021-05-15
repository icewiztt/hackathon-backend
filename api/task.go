package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/thanhqt2002/hackathon/db/sqlc"
	"github.com/thanhqt2002/hackathon/token"
)

type CreateTaskRequest struct {
	Shortname     string    `json:"shortname" binding:"required"`
	Problemname   string    `json:"problemname" binding:"required"`
	Content       string    `json:"content" binding:"required"`
	Subtasks      int32     `json:"subtasks" binding:"required"`
	Answers       []string  `json:"answers" binding:"required"`
	SubtasksScore []float64 `json:"subtasks_score" binding:"required"`
}

func (server *Server) CreateTask(ctx *gin.Context) {
	var req CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTaskParams{
		Shortname:     req.Shortname,
		Problemname:   req.Problemname,
		Content:       req.Content,
		Subtasks:      req.Subtasks,
		Answers:       req.Answers,
		SubtasksScore: req.SubtasksScore,
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if authPayload.Usertype == 1 {
		err := errors.New("this user is not allowed to add new task")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	task, err := server.store.CreateTask(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, task)
}

type GetTaskRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type GetTaskResponse struct {
	Shortname     string    `json:"shortname" binding:"required"`
	Problemname   string    `json:"problemname" binding:"required"`
	Content       string    `json:"content" binding:"required"`
	Subtasks      int32     `json:"subtasks" binding:"required"`
	SubtasksScore []float64 `json:"subtasks_score" binding:"required"`
}

func (server *Server) GetTask(ctx *gin.Context) {
	var req GetTaskRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	task, err := server.store.GetTask(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, GetTaskResponse{
		Shortname:     task.Shortname,
		Problemname:   task.Problemname,
		Content:       task.Content,
		Subtasks:      task.Subtasks,
		SubtasksScore: task.SubtasksScore,
	})
}

type ListTasksRequest struct {
	PageSize int32 `form:"PageSize" binding:"required,min=1,max=30"`
	PageID   int32 `form:"PageID" binding:"required,min=1"`
}

func (server *Server) ListTasks(ctx *gin.Context) {
	var req ListTasksRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	task_list, err := server.store.ListTasks(ctx, db.ListTasksParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	task := make([]GetTaskResponse, len(task_list))
	for i := 0; i < len(task_list); i++ {
		task[i] = GetTaskResponse{
			Shortname:     task_list[i].Shortname,
			Problemname:   task_list[i].Problemname,
			Content:       task_list[i].Content,
			Subtasks:      task_list[i].Subtasks,
			SubtasksScore: task_list[i].SubtasksScore,
		}
	}
	ctx.JSON(http.StatusOK, task)
}
