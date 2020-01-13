package infra

import (
	"time"

	"github.com/YuichiKadota/introther/domain/model"
	repository "github.com/YuichiKadota/introther/domain/repository/post"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDBRepoImpl - DynamoDB処理の実装メソッドをもつ構造体
type DynamoDBRepoImpl struct {
	dynamoDB *dynamodb.DynamoDB
}

// NewDynamoDBRepoImpl - DynamoDB処理の実装を返す
func NewDynamoDBRepoImpl() (repository.PostRepo, error) {

	ddb := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	dynamoDBRepoImpl := &DynamoDBRepoImpl{
		dynamoDB: ddb,
	}

	return dynamoDBRepoImpl, nil
}

// Get - 仮定義
func (r *DynamoDBRepoImpl) Get(postID string) (*model.Post, error) {
	var post model.Post
	params := &dynamodb.GetItemInput{
		TableName: aws.String("post"), // テーブル名

		Key: map[string]*dynamodb.AttributeValue{
			"post_id": { // キー名
				S: aws.String(postID), // 持ってくるキーの値
			},
		},
		AttributesToGet: []*string{
			aws.String("post_user_id"),
			aws.String("posted_user_id"),
			aws.String("text"),
			//TODO 複数画像がある時にどうするか
			aws.String("image_url"),
			aws.String("insert_date"),
			aws.String("update_date"), // 欲しいデータの名前
		},
		ConsistentRead: aws.Bool(true), // 常に最新を取得するかどうか

		//返ってくるデータの種類
		ReturnConsumedCapacity: aws.String("NONE"),
	}

	resp, err := r.dynamoDB.GetItem(params)

	if err != nil {
		return &post, err
	}

	insertDate, err := time.Parse("2006/01/02 15:04:05", *resp.Item["insert_date"].S)
	if err != nil {
		return &post, err
	}

	updateDate, err := time.Parse("2006/01/02 15:04:05", *resp.Item["update_date"].S)
	if err != nil {
		return &post, err
	}

	//resp.Item[項目名].型 でデータへのポインタを取得
	post = model.Post{
		PostID:       postID,
		PostUserID:   *resp.Item["post_user_id"].S,
		PostedUserID: *resp.Item["posted_user_id"].S,
		Text:         *resp.Item["text"].S,
		ImageURL:     *resp.Item["image_url"].S,
		InsertDate:   insertDate,
		UpdateDate:   updateDate,
	}

	return &post, nil
}

// Insert - 投稿内容の登録を行う実装
func (r *DynamoDBRepoImpl) Insert(post *model.Post) (*model.Post, error) {

	var err error

	param := &dynamodb.UpdateItemInput{
		TableName: aws.String("post"), // テーブル名を指定

		Key: map[string]*dynamodb.AttributeValue{
			"post_id": {
				S: aws.String(post.PostID), // キー名を指定
			},
		},

		ExpressionAttributeNames: map[string]*string{
			"#post_user_id":   aws.String("post_user_id"),
			"#posted_user_id": aws.String("posted_user_id"),
			"#text":           aws.String("text"), // 項目名をプレースホルダに入れる
			"#image_url":      aws.String("image_url"),
			"#insert_date":    aws.String("insert_date"),
			"#update_date":    aws.String("update_date"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":post_user_id_value": {
				S: aws.String(post.PostUserID), // 値をプレースホルダに入れる
			},
			":posted_user_id_value": {
				S: aws.String(post.PostedUserID),
			},
			":text_value": {
				S: aws.String(post.Text), // 値をプレースホルダに入れる
			},
			":image_url_value": {
				S: aws.String(post.ImageURL), // 値をプレースホルダに入れる
			},
			":insert_date_value": {
				S: aws.String(post.InsertDate.Format("2006/01/02 15:04:05")), // 値をプレースホルダに入れる
			},
			":update_date_value": {
				S: aws.String(post.UpdateDate.Format("2006/01/02 15:04:05")), // 値をプレースホルダに入れる
			},
		},

		UpdateExpression: aws.String(
			"set #post_user_id = :post_user_id_value, " +
				"#posted_user_id = :posted_user_id_value, " +
				"#text = :text_value, " +
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
		return post, err
	}

	return post, nil
}

// Update - 仮定義
func (r *DynamoDBRepoImpl) Update(post *model.Post) (*model.Post, error) {
	return post, nil
}

// Delete - 仮定義
func (r *DynamoDBRepoImpl) Delete(post *model.Post) (bool, error) {
	return true, nil
}
