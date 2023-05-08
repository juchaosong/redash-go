// Code generated from query.go using tools/withoutctx.go; DO NOT EDIT.

package redash

import (
	"context"
	"io"
)

func (client *ClientWithoutContext) ListQueries(input *ListQueriesInput) (*QueryPage, error) {
	return client.withCtx.ListQueries(context.Background(), input)
}

func (client *ClientWithoutContext) GetQuery(id int) (*Query, error) {
	return client.withCtx.GetQuery(context.Background(), id)
}

func (client *ClientWithoutContext) CreateFavoriteQuery(id int) error {
	return client.withCtx.CreateFavoriteQuery(context.Background(), id)
}

func (client *ClientWithoutContext) CreateQuery(input *CreateQueryInput) (*Query, error) {
	return client.withCtx.CreateQuery(context.Background(), input)
}

func (client *ClientWithoutContext) ForkQuery(id int) (*Query, error) {
	return client.withCtx.ForkQuery(context.Background(), id)
}

func (client *ClientWithoutContext) UpdateQuery(id int, input *UpdateQueryInput) (*Query, error) {
	return client.withCtx.UpdateQuery(context.Background(), id, input)
}

func (client *ClientWithoutContext) ArchiveQuery(id int) error {
	return client.withCtx.ArchiveQuery(context.Background(), id)
}

func (client *ClientWithoutContext) GetQueryResultsJSON(id int, out io.Writer) error {
	return client.withCtx.GetQueryResultsJSON(context.Background(), id, out)
}

func (client *ClientWithoutContext) GetQueryResultsCSV(id int, out io.Writer) error {
	return client.withCtx.GetQueryResultsCSV(context.Background(), id, out)
}

func (client *ClientWithoutContext) GetQueryResults(id int, ext string, out io.Writer) error {
	return client.withCtx.GetQueryResults(context.Background(), id, ext, out)
}

func (client *ClientWithoutContext) ExecQueryJSON(id int, out io.Writer) (*JobResponse, error) {
	return client.withCtx.ExecQueryJSON(context.Background(), id, out)
}

func (client *ClientWithoutContext) GetQueryTags() (*QueryTags, error) {
	return client.withCtx.GetQueryTags(context.Background())
}

func (client *ClientWithoutContext) RefreshQuery(id int) (*JobResponse, error) {
	return client.withCtx.RefreshQuery(context.Background(), id)
}

func (client *ClientWithoutContext) SearchQueries(input *SearchQueriesInput) (*QueryPage, error) {
	return client.withCtx.SearchQueries(context.Background(), input)
}

func (client *ClientWithoutContext) ListMyQueries(input *ListMyQueriesInput) (*QueryPage, error) {
	return client.withCtx.ListMyQueries(context.Background(), input)
}

func (client *ClientWithoutContext) ListFavoriteQueries(input *ListFavoriteQueriesInput) (*QueryPage, error) {
	return client.withCtx.ListFavoriteQueries(context.Background(), input)
}

func (client *ClientWithoutContext) FormatQuery(query string) (*FormatQueryOutput, error) {
	return client.withCtx.FormatQuery(context.Background(), query)
}

func (client *ClientWithoutContext) ListRecentQueries() ([]Query, error) {
	return client.withCtx.ListRecentQueries(context.Background())
}
