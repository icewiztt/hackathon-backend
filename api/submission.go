package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/thanhqt2002/hackathon/db/sqlc"
)

type CreateSubmissionRequest struct {
	FromUserID        int32    `json:"FromUserID" binding:"required"`
	ToTaskID          int32    `json:"ToTaskID" binding:"required"`
	TaskSubtasks      int32    `json:"TaskSubtasks" binding:"required"`
	SubmissionAnswers []string `json:"SubmissionAnswers" binding:"required"`
}

func CheckContains(v []string, s string) bool {
	for _, u := range v {
		if strings.TrimSpace(u) == s {
			return true
		}
	}
	return false
}

func (server *Server) CalculateScore(ctx *gin.Context, req CreateSubmissionRequest, task db.Task) (result []bool, total float64, err error) {

	result = make([]bool, task.Subtasks)
	total = 0
	for i := 0; i < int(task.Subtasks); i++ {
		result[i] = CheckContains(strings.Split(task.Answers[i], "|"), strings.TrimSpace(req.SubmissionAnswers[i]))
		if result[i] {
			total += task.SubtasksScore[i]
		}
	}
	return
}

func (server *Server) CreateSubmission(ctx *gin.Context) {
	var req CreateSubmissionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if len(req.SubmissionAnswers) != int(req.TaskSubtasks) {
		err := fmt.Errorf("expected slice of length %v but recieved slice of length %v", req.TaskSubtasks, len(req.SubmissionAnswers))
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	task, err := server.store.GetTask(ctx, req.ToTaskID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique violation":
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}
		}
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if task.Subtasks != req.TaskSubtasks {
		err = fmt.Errorf("task length mismatched")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	SubmissionResults, SubmissionScore, err := server.CalculateScore(ctx, req, task)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateSubmissionParams{
		FromUserID:        req.FromUserID,
		ToTaskID:          req.ToTaskID,
		TaskSubtasks:      req.TaskSubtasks,
		SubmissionAnswers: req.SubmissionAnswers,
		SubmissionScore:   SubmissionScore,
		SubmissionResults: SubmissionResults,
	}

	submission, err := server.store.CreateSubmission(ctx, arg)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, submission)
}

/*
type GetTaskRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
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
	ctx.JSON(http.StatusOK, task)
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

	task, err := server.store.ListTasks(ctx, db.ListTasksParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, task)
}
*/
