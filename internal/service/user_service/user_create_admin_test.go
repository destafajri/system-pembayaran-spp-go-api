package user_service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/destafajri/system-pembayaran-spp-go-api/internal/domain/model"
	mock_repository "github.com/destafajri/system-pembayaran-spp-go-api/mocks/repo"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_userServiceimpl_CreateAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := mock_repository.NewMockUserRepository(ctrl)
	fmt.Println(c)

	type args struct {
		request   *model.CreateAdminRequest
		timestamp time.Time
	}
	tests := []struct {
		name          string
		user          *userServiceimpl
		args          args
		beforeTest    func(userRepo *mock_repository.MockUserRepository)
		want          *model.CreateAdminResponse
		wantErr       bool
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "status 200",
			args: args{
				request: &model.CreateAdminRequest{
					Email:    "tes@mail.com",
					Username: "test",
					Password: "Tes123",
				},
				timestamp: time.Now(),
			},
			beforeTest: func(userRepo *mock_repository.MockUserRepository) {
				userRepo.EXPECT().
					CreateAdmin(
						gomock.Any(),
					).
					Return(
						&model.CreateAdminResponse{
							ID:       "185c52f0-429a-419b-83d9-0a335e54b293",
							Email:    "tes@mail.com",
							Username: "test",
						},
						nil,
					)
			},
			want: &model.CreateAdminResponse{
				ID:       "185c52f0-429a-419b-83d9-0a335e54b293",
				Email:    "tes@mail.com",
				Username: "test",
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mock_repository.NewMockUserRepository(ctrl)
			recorder := httptest.NewRecorder()

			w := &userServiceimpl{
				userRepository: mockUserRepo,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockUserRepo)
			}

			got, err := w.CreateAdmin(tt.args.request, tt.args.timestamp)
			if (err != nil) != tt.wantErr {
				t.Errorf("userServiceimpl.CreateAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userServiceimpl.CreateAdmin() = %v, want %v", got, tt.want)
			}

			tt.checkResponse(recorder)
			assert.Nil(t, err)
		})
	}
}
