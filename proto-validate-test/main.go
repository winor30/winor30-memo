package main

import (
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	article "github.com/winor30/winor30-memo/proto-validate-test/gen/proto/article/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	req := &article.CreateArticlesRequest{}
	req.Inputs = append(
		req.Inputs,
		&article.CreateArticleInput{
			Id:     &wrapperspb.StringValue{Value: "id1"},
			Title:  "title1",
			Author: "author1",
		},
		&article.CreateArticleInput{
			Id:     &wrapperspb.StringValue{Value: "id2"},
			Title:  "title2",
			Author: "author2",
		},
		&article.CreateArticleInput{
			Title:  "title3",
			Author: "author3",
		},
		&article.CreateArticleInput{
			Title:  "title4",
			Author: "author4",
		},
	)

	if err := protovalidate.Validate(req); err != nil {
		panic(err)
	}

	fmt.Printf("CreateArticlesRequest:{%s}\n", req.String())
}
