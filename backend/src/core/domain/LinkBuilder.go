package domain

import "github.com/google/uuid"

type LinkBuilder struct {
	link *Link
}

func (builder *LinkBuilder) WithID(id uuid.UUID) *LinkBuilder {
	builder.link.id = id
	return builder
}

func (builder *LinkBuilder) WithContent(content string) *LinkBuilder {
	builder.link.content = content
	return builder
}

func (builder *LinkBuilder) WithCategory(category Category) *LinkBuilder {
	builder.link.categories = append(builder.link.categories, category)
	return builder
}

func (builder *LinkBuilder) Build() *Link {
	return builder.link
}

func NewLinkBuilder() *LinkBuilder {
	return &LinkBuilder{link: &Link{}}
}
