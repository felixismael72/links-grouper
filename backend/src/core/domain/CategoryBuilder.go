package domain

import "github.com/google/uuid"

type CategoryBuilder struct {
	category *Category
}

func (builder *CategoryBuilder) WithID(id uuid.UUID) *CategoryBuilder {
	builder.category.id = id
	return builder
}

func (builder *CategoryBuilder) WithName(name string) *CategoryBuilder {
	builder.category.name = name
	return builder
}

func (builder *CategoryBuilder) Build() *Category {
	return builder.category
}

func NewCategoryBuilder() *CategoryBuilder {
	return &CategoryBuilder{category: &Category{}}
}
