package handlers

import (
	"bytes"
	"contentSquare/src/internal/models"
	"contentSquare/src/mocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockAggregateHandler struct {
	router *gin.RouterGroup
	DbRepo *mocks.MockDBRepository
}

func TestCountEventsWithFilter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// · Mocks · //
	response := models.CountResponse{
		Count: "1",
	}
	jsonResponse, _ := json.Marshal(response)

	// · Tests · //
	type want struct {
		code     int
		response string
		err      error
	}

	tests := []struct {
		name    string
		filters models.Filters
		url     string
		want    want
		result  string
		mocks   func(mAH mockAggregateHandler)
	}{
		{
			name: "Should count with filters succesfully",
			url:  "/count?date_from=2022-01-01T00:00:00&date_to=2022-01-05T23:59:59&event=MOUSE_UP",
			want: want{
				code:     http.StatusOK,
				response: string(jsonResponse),
				err:      nil,
			},
			mocks: func(mAH mockAggregateHandler) {
				mAH.DbRepo.EXPECT().CountEvents(gomock.Any(), models.Filters{
					Event:    "MOUSE_UP",
					DateFrom: "2022-01-01T00:00:00",
					DateTo:   "2022-01-05T23:59:59",
				}).Return(int64(1), nil)
			},
		},
		{
			name: "Should return error - Failed to query DB",
			url:  "/count?date_from=2022-01-01T00:00:00&date_to=2022-01-05T23:59:59&event=MOUSE_UP",
			want: want{
				code: http.StatusInternalServerError,
				response: `{
					"message": "Error getting count values"
				}`,
				err: errors.New("Error getting count values"),
			},
			mocks: func(mAH mockAggregateHandler) {
				mAH.DbRepo.EXPECT().CountEvents(gomock.Any(), models.Filters{
					Event:    "MOUSE_UP",
					DateFrom: "2022-01-01T00:00:00",
					DateTo:   "2022-01-05T23:59:59",
				}).Return(int64(0), errors.New("Error getting count values"))
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// Prepare
			w := httptest.NewRecorder()
			r := gin.Default()
			app := r.Group("/")

			mAH := mockAggregateHandler{
				router: app,
				DbRepo: mocks.NewMockDBRepository(gomock.NewController(t)),
			}

			tt.mocks(mAH)
			NewAggregateHandler(mAH.router, mAH.DbRepo)

			req, err := http.NewRequest("GET", tt.url, bytes.NewBufferString(""))
			require.NoError(t, err)
			r.ServeHTTP(w, req)
			assert.JSONEq(t, tt.want.response, w.Body.String())
			assert.Equal(t, tt.want.code, w.Code)
		})

	}

}

func TestCountDistinctUsersWithFilter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// · Mocks · //
	response := models.CountDistinctUsersResponse{
		CountDistinctUsers: "1",
	}
	jsonResponse, _ := json.Marshal(response)

	// · Tests · //
	type want struct {
		code     int
		response string
		err      error
	}

	tests := []struct {
		name    string
		filters models.Filters
		url     string
		want    want
		result  string
		mocks   func(mAH mockAggregateHandler)
	}{
		{
			name: "Should count different users with filters succesfully",
			url:  "/count_distinct_users?date_from=2022-01-01T00:00:00&date_to=2022-01-05T23:59:59&event=MOUSE_UP",
			want: want{
				code:     http.StatusOK,
				response: string(jsonResponse),
				err:      nil,
			},
			mocks: func(mAH mockAggregateHandler) {
				mAH.DbRepo.EXPECT().CountDistinctUsers(gomock.Any(), models.Filters{
					Event:    "MOUSE_UP",
					DateFrom: "2022-01-01T00:00:00",
					DateTo:   "2022-01-05T23:59:59",
				}).Return(int64(1), nil)
			},
		},
		{
			name: "Should return error - Failed to query DB",
			url:  "/count_distinct_users?date_from=2022-01-01T00:00:00&date_to=2022-01-05T23:59:59&event=MOUSE_UP",
			want: want{
				code: http.StatusInternalServerError,
				response: `{
					"message": "Error counting distinct users"
				}`,
				err: errors.New("Error counting distinct users"),
			},
			mocks: func(mAH mockAggregateHandler) {
				mAH.DbRepo.EXPECT().CountDistinctUsers(gomock.Any(), models.Filters{
					Event:    "MOUSE_UP",
					DateFrom: "2022-01-01T00:00:00",
					DateTo:   "2022-01-05T23:59:59",
				}).Return(int64(0), errors.New("Error counting distinct users"))
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// Prepare
			w := httptest.NewRecorder()
			r := gin.Default()
			app := r.Group("/")

			mAH := mockAggregateHandler{
				router: app,
				DbRepo: mocks.NewMockDBRepository(gomock.NewController(t)),
			}

			tt.mocks(mAH)
			NewAggregateHandler(mAH.router, mAH.DbRepo)

			req, err := http.NewRequest("GET", tt.url, bytes.NewBufferString(""))
			require.NoError(t, err)
			r.ServeHTTP(w, req)
			assert.JSONEq(t, tt.want.response, w.Body.String())
			assert.Equal(t, tt.want.code, w.Code)
		})

	}

}

func TestExistsWithFilter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// · Mocks · //
	response := models.ExistsResponse{
		Exists: "true",
	}
	jsonResponse, _ := json.Marshal(response)

	// · Tests · //
	type want struct {
		code     int
		response string
		err      error
	}

	tests := []struct {
		name    string
		filters models.Filters
		url     string
		want    want
		result  string
		mocks   func(mAH mockAggregateHandler)
	}{
		{
			name: "Should check if exists succesfully",
			url:  "/exists?event=CLICK&user_id=8864487207",
			want: want{
				code:     http.StatusOK,
				response: string(jsonResponse),
				err:      nil,
			},
			mocks: func(mAH mockAggregateHandler) {
				mAH.DbRepo.EXPECT().Exists(gomock.Any(), models.Filters{
					Event:  "CLICK",
					UserId: "8864487207",
				}).Return(true, nil)
			},
		},
		{
			name: "Should return error - Failed to query DB",
			url:  "/exists?event=CLICK&user_id=8864487207",
			want: want{
				code: http.StatusInternalServerError,
				response: `{
					"message": "Error getting exists query"
				}`,
				err: errors.New("Error getting exists query"),
			},
			mocks: func(mAH mockAggregateHandler) {
				mAH.DbRepo.EXPECT().Exists(gomock.Any(), models.Filters{
					Event:  "CLICK",
					UserId: "8864487207",
				}).Return(false, errors.New("Error getting exists query"))
			},
		},
	}

	// · Runner · //
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			// Prepare
			w := httptest.NewRecorder()
			r := gin.Default()
			app := r.Group("/")

			mAH := mockAggregateHandler{
				router: app,
				DbRepo: mocks.NewMockDBRepository(gomock.NewController(t)),
			}

			tt.mocks(mAH)
			NewAggregateHandler(mAH.router, mAH.DbRepo)

			req, err := http.NewRequest("GET", tt.url, bytes.NewBufferString(""))
			require.NoError(t, err)
			r.ServeHTTP(w, req)
			assert.JSONEq(t, tt.want.response, w.Body.String())
			assert.Equal(t, tt.want.code, w.Code)
		})

	}

}
