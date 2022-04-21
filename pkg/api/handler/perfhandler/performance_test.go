package perfhandler

import (
	"encoding/json"
	"errors"
	"github.com/artback/queryperformance/mocks"
	"github.com/artback/queryperformance/pkg/queryperf"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestPerfHandler_GetQueryPerformance(t *testing.T) {
	t.Parallel()
	type args struct {
		request *http.Request
	}
	type Repository struct {
		options queryperf.Options
		result  queryperf.Result
		error   error
	}
	tests := []struct {
		name       string
		args       args
		want       queryperf.Result
		repository Repository
		wantErr    bool
	}{
		{
			name: "get request without query",
			args: args{
				httptest.NewRequest(http.MethodGet, "/", nil),
			},
			want: queryperf.Result{
				queryperf.Measure{Query: "select something"},
			},
			repository: Repository{
				options: queryperf.Options{},
				result: queryperf.Result{
					queryperf.Measure{Query: "select something"},
				}},
		},
		{
			name: "get request with query",
			args: args{
				httptest.NewRequest(http.MethodGet, "/?mincalls=2&statement=SELECT&statement=INSERT", nil),
			},
			want: queryperf.Result{
				queryperf.Measure{Query: "select something"},
			},
			repository: Repository{
				options: queryperf.Options{MinCalls: 2, Statements: []string{"SELECT", "INSERT"}},
				result: queryperf.Result{
					queryperf.Measure{Query: "select something"},
				}},
		},
		{
			name: "repository error",
			args: args{
				httptest.NewRequest(http.MethodGet, "/", nil),
			},
			wantErr:    true,
			repository: Repository{error: errors.New("something happened")},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockCtrl := gomock.NewController(t)
			repo := mocks.NewQueryPerfRepository(mockCtrl)
			repo.EXPECT().QueriesByMeanTime(gomock.Any(), tt.repository.options).Return(&tt.repository.result, tt.repository.error).Times(1)
			p := PerfHandler{
				Repository: repo,
			}
			app := fiber.New()
			app.Get("/", p.GetQueryPerformance)
			resp, err := app.Test(tt.args.request)
			if err != nil {
				t.Fatal("Fiber app.Test internal error")
			}
			var results queryperf.Result
			_ = json.NewDecoder(resp.Body).Decode(&results)
			if !reflect.DeepEqual(results, tt.want) {
				t.Errorf("handler returned wrong body: got %v want %v",
					results, tt.want)
			}
			mockCtrl.Finish()
		})
	}
}
