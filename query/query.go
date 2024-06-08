package query

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Aidunlin/go-srm/model"
	"github.com/a-h/templ"
)

// Builder object for constructing URL query strings.
type QueryBuilder struct {
	ctx    context.Context
	values url.Values
}

// Creates a new QueryBuilder.
func New(ctx context.Context) QueryBuilder {
	return QueryBuilder{
		ctx:    ctx,
		values: url.Values{},
	}
}

// Sets a key-value.
func (q QueryBuilder) With(key, value string) QueryBuilder {
	q.values.Set(key, value)
	return q
}

// Sets all the key-values from a string-to-string map.
func (q QueryBuilder) WithMap(params map[string]string) QueryBuilder {
	for key, value := range params {
		q.values.Set(key, value)
	}
	return q
}

// Sets all the key-values from a url value map.
func (q QueryBuilder) WithUrlValues(values url.Values) QueryBuilder {
	for key, value := range values {
		q.values.Set(key, value[0])
	}
	return q
}

// Gets and sets values from all the table parameters.
func (q QueryBuilder) WithTableParams() QueryBuilder {
	p := model.GetStudentTableParams(q.ctx)
	if len(p.Filter) > 0 {
		q.values.Set("filter", p.Filter)
	}
	q.values.Set("sortby", p.Sortby)
	q.values.Set("order", p.Order)
	q.values.Set("page", fmt.Sprint(p.Page))
	if len(p.Search) > 0 {
		q.values.Set("q", p.Search)
	}
	return q
}

// Gets and sets values from all the advanced search parameters.
func (q QueryBuilder) WithAdvancedSearch() QueryBuilder {
	p := model.GetAdvancedSearchForm(q.ctx).ToMap()
	return q.WithMap(p)
}

// Removes specified keys.
func (q QueryBuilder) Without(keys ...string) QueryBuilder {
	for _, key := range keys {
		q.values.Del(key)
	}
	return q
}

// Outputs a query string (with '?') to be used in templ.
func (q QueryBuilder) Build() templ.SafeURL {
	return templ.URL(fmt.Sprintf("?%v", q.values.Encode()))
}
