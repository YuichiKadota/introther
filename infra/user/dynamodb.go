package infra

import (
	"os"

	"github.com/YuichiKadota/introther/domain/model"
	repository "github.com/YuichiKadota/introther/domain/repository/user"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDBRepoImpl - DynamoDB処理の実装メソッドをもつ構造体
type DynamoDBRepoImpl struct {
	dynamoDB *dynamodb.DynamoDB
}

// NewDynamoDBRepoImpl - DynamoDB処理の実装を返す
func NewDynamoDBRepoImpl() (repository.UserProfileRepo, error) {

	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion(os.Getenv("REGION")))

	dynamoDBRepoImpl := &DynamoDBRepoImpl{
		dynamoDB: ddb,
	}

	return dynamoDBRepoImpl, nil
}

// Get - 仮定義
func (r *DynamoDBRepoImpl) Get(userID string) (*model.User, error) {
	var user model.User
	return &user, nil
}

// Insert - ユーザー登録を行う実装
func (r *DynamoDBRepoImpl) Insert(user *model.User) (*model.User, error) {

	var err error

	param := &dynamodb.UpdateItemInput{
		TableName: aws.String(os.Getenv("DYNAMODB_NAME")), // テーブル名を指定

		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(user.UserID), // キー名を指定
			},
		},

		ExpressionAttributeNames: map[string]*string{
			"#password":    aws.String("password"),
			"#nick_name":   aws.String("nick_name"),
			"#profile":     aws.String("profile"), // 項目名をプレースホルダに入れる
			"#image_url":   aws.String("image_url"),
			"#insert_date": aws.String("insert_date"),
			"#update_date": aws.String("update_date"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":nick_name_value": {
				S: aws.String(user.NickName), // 値をプレースホルダに入れる
			},
			":password_value": {
				S: aws.String(user.Password),
			},
			":profile_value": {
				S: aws.String(user.Profile), // 値をプレースホルダに入れる
			},
			":image_url_value": {
				S: aws.String(user.ImageURL), // 値をプレースホルダに入れる
			},
			":insert_date_value": {
				S: aws.String(user.InsertDate.Format("2006/01/02 15:04:05")), // 値をプレースホルダに入れる
			},
			":update_date_value": {
				S: aws.String(user.UpdateDate.Format("2006/01/02 15:04:05")), // 値をプレースホルダに入れる
			},
		},

		UpdateExpression: aws.String(
			"set #password = :password_value, " +
				"#nick_name = :nick_name_value, " +
				"#profile = :profile_value, " +
				"#image_url = :image_url_value, " +
				"#insert_date = :insert_date_value, " +
				"#update_date = :update_date_value ",
		), //プレースホルダを利用して更新の式を書く

		//あとは返してくる情報の種類を指定する
		ReturnConsumedCapacity:      aws.String("NONE"), //(デフォルト値) 何も返さない
		ReturnItemCollectionMetrics: aws.String("NONE"), //(デフォルト値) 何も返さない
		ReturnValues:                aws.String("NONE"), //(デフォルト値) 何も返さない
	}

	_, err = r.dynamoDB.UpdateItem(param) //実行

	if err != nil {
		return user, err
	}

	return user, nil
}

// Update - 仮定義
func (r *DynamoDBRepoImpl) Update(user *model.User) (*model.User, error) {
	return user, nil
}

// Delete - 仮定義
func (r *DynamoDBRepoImpl) Delete(user *model.User) (bool, error) {
	return true, nil
}
