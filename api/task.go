package api

import (
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
	Official      bool      `json:"official" binding:"required"`
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
		Official:      req.Official,
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	authorized_user, err := server.store.GetUser(ctx, authPayload.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if authorized_user.Usertype == 1 {
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
	ID            int32     `json:"id" binding:"required"`
	Shortname     string    `json:"shortname" binding:"required"`
	Problemname   string    `json:"problemname" binding:"required"`
	Content       string    `json:"content" binding:"required"`
	Subtasks      int32     `json:"subtasks" binding:"required"`
	SubtasksScore []float64 `json:"subtasks_score" binding:"required"`
}

func (server *Server) ListTasks(ctx *gin.Context) {
	task_list, err := server.store.ListTasks(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	task := make([]GetTaskResponse, len(task_list))
	for i := 0; i < len(task_list); i++ {
		task[i] = GetTaskResponse{
			ID:            task_list[i].ID,
			Shortname:     task_list[i].Shortname,
			Problemname:   task_list[i].Problemname,
			Content:       task_list[i].Content,
			Subtasks:      task_list[i].Subtasks,
			SubtasksScore: task_list[i].SubtasksScore,
		}
	}
	ctx.JSON(http.StatusOK, task)
}

func (server *Server) ListTasksAdmin(ctx *gin.Context) {
	task_list, err := server.store.ListTasksAdmin(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	authorized_user, err := server.store.GetUser(ctx, authPayload.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if authorized_user.Usertype == 1 {
		err := errors.New("this user is not allowed to edit tasks")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, task_list)
}
