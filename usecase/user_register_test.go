package usecase

import (
	"reflect"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"

	"errors"

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
	user.UserID = "duplicate_userID"
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
	regalTests := []struct {
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
					Profile:  "テストプロフィール",
				},
			},
			want: model.User{
				UserID:     "regal-test-userID",
				NickName:   "regal-test-nick_nsme",
				Password:   "regal-test-password",
				Profile:    "テストプロフィール",
				InsertDate: time.Now(),
				UpdateDate: time.Now(),
			},
		},
	}

	iregalTests := []struct {
		name    string
		fields  fields
		args    args
		want    error
		wantErr bool
	}{
		// TODO: Add iregal test cases.
		{name: "iregal_test_kana",
			fields: newFields(),
			args: args{
				model.User{
					UserID:   "テストユーザー（カナ）",
					Password: "テストパスワード（カナ）",
					NickName: "テストニックネーム（カナ）",
					Profile:  "亜gkdjfdじゃdbfbbbbcbbbdbbjfdvgjgdgggfhheyuuwjjjsnnncnbfbjjjjwmmdnhjjwiiwiwkoallk,sjjmmxmmmejjeiiikw,,,,,s,xkmmmekjjhenqksoeoke,m；亜djfkjdfぁsdjふぁls；djfsfvwfvweviwevcgewfcgweycgecgewycgweych；dFjlskdfjs；dFjlksdFjsl亜dvhんsjvbんjdhgfんrjfmbんvjmjrにdjkfhmんvjkjdんゔぃkhjんりfdkvhnりkdhvnF邪dhfんvjhrんづjhvんうrjdmfhんvjmdfhんvjmへrんfdjkhvんうrjfdmhvんrjdfhvん",
				},
			},
			want: errors.New("入力項目が不適切です。 NickName: must be in a valid format; Password: must be in a valid format; Profile: the length must be no more than 255; UserID: must be in a valid format."),
		},
		{name: "iregal_test_duplicate_userID",
			fields: newFields(),
			args: args{
				model.User{
					UserID:   "duplicate_userID",
					Password: "duplicate_password",
					NickName: "duplicate_nick_name",
					Profile:  "duplicate_profile",
				},
			},
			want: errors.New("そのユーザーIDは既に登録されています。"),
		},
	}

	for _, tt := range regalTests {
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

			err = bcrypt.CompareHashAndPassword([]byte(got.Password), []byte(tt.want.Password))
			if err != nil {
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

	for _, tt := range iregalTests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UeserUsecsse{
				userRepo: tt.fields.userRepo,
			}

			_, err := u.Register(&tt.args.in0)

			if err.Error() != tt.want.Error() {
				t.Errorf("UeserUsecsse.Register() = %v, want %v", err, tt.want)
			}
		})
	}
}
