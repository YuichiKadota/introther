package usecase

import (
	"reflect"
	"testing"
	"time"

	"github.com/YuichiKadota/introther/domain/model"
	repository "github.com/YuichiKadota/introther/domain/repository/user"
)

type fields struct {
	userRepo repository.UserProfileRepo
}

type testRepo struct{}

func newFields() fields {
	fields := fields{userRepo: &testRepo{}}
	return fields
}

func (r *testRepo) Get(userID string) (*model.User, error) {
	var user model.User
	return &user, nil
}

func (r *testRepo) Insert(user *model.User) (*model.User, error) {

	return user, nil
}

func (r *testRepo) Update(user *model.User) (*model.User, error) {

	return user, nil
}

func (r *testRepo) Delete(user *model.User) (bool, error) {

	return true, nil
}

func TestUeserUsecsse_Register(t *testing.T) {

	type args struct {
		in0 model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "normal_test",
			fields: newFields(),
			args: args{
				model.User{
					UserID:     "normal_test_userID",
					NickName:   "normal_test_nick_nsme",
					Profile:    "normal_test_profile",
					InsertDate: time.Date(2019, 12, 31, 23, 59, 59, 0, time.Local),
					UpdateDate: time.Date(2019, 12, 31, 23, 59, 59, 0, time.Local),
				},
			},
			want: model.User{
				UserID:     "normal_test_userID",
				NickName:   "normal_test_nick_nsme",
				Profile:    "normal_test_profile",
				InsertDate: time.Date(2019, 12, 31, 23, 59, 59, 0, time.Local),
				UpdateDate: time.Date(2019, 12, 31, 23, 59, 59, 0, time.Local),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			u := &UeserUsecsse{
				userRepo: tt.fields.userRepo,
			}

			got, err := u.Register(&tt.args.in0)
			if (err != nil) != tt.wantErr {
				t.Errorf("UeserUsecsse.Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UeserUsecsse.Register() = %v, want %v", got, tt.want)
			}
		})
	}
}
