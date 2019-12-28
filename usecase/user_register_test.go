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
		// TODO: Add iregal test cases.
		{name: "regal_test",
			fields: newFields(),
			args: args{
				model.User{
					UserID:   "regal-test-userID",
					Password: "regal-test-password",
					NickName: "regal-test-nick_nsme",
					Profile:  "regal-test-profile",
				},
			},
			want: model.User{
				UserID:     "regal-test-userID",
				NickName:   "regal-test-nick_nsme",
				Password:   "regal-test-password",
				Profile:    "regal-test-profile",
				InsertDate: time.Now(),
				UpdateDate: time.Now(),
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
			if !reflect.DeepEqual(got.UserID, tt.want.UserID) {
				t.Errorf("UeserUsecsse.Register() = %v, want %v", got.UserID, tt.want.UserID)
			}
			if !reflect.DeepEqual(got.Password, tt.want.Password) {
				t.Errorf("UeserUsecsse.Register() = %v, want %v", got.Password, tt.want.Password)
			}
			if !reflect.DeepEqual(got.NickName, tt.want.NickName) {
				t.Errorf("UeserUsecsse.Register() = %v, want %v", got.NickName, tt.want.NickName)
			}
			if !reflect.DeepEqual(got.Profile, tt.want.Profile) {
				t.Errorf("UeserUsecsse.Register() = %v, want %v", got.Profile, tt.want.Profile)
			}
			if !reflect.DeepEqual(got.ImageURL, tt.want.ImageURL) {
				t.Errorf("UeserUsecsse.Register() = %v, want %v", got.ImageURL, tt.want.ImageURL)
			}
			if got.InsertDate.IsZero() == true {
				t.Errorf("UeserUsecsse.Register() = %v is initial value", got.InsertDate)
			}
			if got.UpdateDate.IsZero() == true {
				t.Errorf("UeserUsecsse.Register() = %v is initial value", got.UpdateDate)
			}
		})
	}
}
