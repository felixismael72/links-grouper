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

func (builder *LinkBuilder) WithCategories(categories []Category) *LinkBuilder {
	builder.link.categories = categories
	return builder
}

func (builder *LinkBuilder) Build() *Link {
	return builder.link
}
