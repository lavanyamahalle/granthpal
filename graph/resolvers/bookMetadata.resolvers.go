package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"dario.cat/mergo"
	"github.com/PICT-LibraryAutomation/granthpal/graph"
	"github.com/PICT-LibraryAutomation/granthpal/remote/models"
	"github.com/PICT-LibraryAutomation/granthpal/utils"
)

// Authors is the resolver for the authors field.
func (r *bookMetadataResolver) Authors(ctx context.Context, obj *graph.BookMetadata) ([]*graph.Author, error) {
	var as []models.Author
	if err := r.Remote.First(&models.BookMetadata{ID: obj.ID}).Association("Authors").Find(&as); err != nil {
		return nil, err
	}

	return utils.Map(as, func(a models.Author) *graph.Author {
		return a.ToGraphQL()
	}), nil
}

// Publication is the resolver for the publication field.
func (r *bookMetadataResolver) Publication(ctx context.Context, obj *graph.BookMetadata) (*graph.Publication, error) {
	var p models.Publication
	if err := r.Remote.First(&p, "id = ?", obj.PublicationID).Error; err != nil {
		return nil, err
	}

	return p.ToGraphQL(), nil
}

// Books is the resolver for the books field.
func (r *bookMetadataResolver) Books(ctx context.Context, obj *graph.BookMetadata) ([]*graph.Book, error) {
	var bs []models.Book
	if err := r.Remote.First(&models.BookMetadata{ID: obj.ID}).Association("Books").Find(&bs); err != nil {
		return nil, err
	}

	return utils.Map(bs, func(b models.Book) *graph.Book {
		return b.ToGraphQL()
	}), nil
}

// CreateBookMeta is the resolver for the createBookMeta field.
func (r *mutationResolver) CreateBookMeta(ctx context.Context, inp graph.CreateBookMetadataInp) (*graph.BookMetadata, error) {
	meta := models.BookMetadata{
		ID:            inp.ID,
		Title:         inp.Title,
		Isbn:          inp.Isbn,
		PublicationID: inp.PublicationID,
		Authors: utils.Map(inp.AuthorIDs, func(id string) models.Author {
			return models.Author{ID: id}
		}),
	}

	if err := r.Remote.Save(&meta).Error; err != nil {
		return nil, err
	}

	return meta.ToGraphQL(), nil
}

// DeleteBookMeta is the resolver for the deleteBookMeta field.
func (r *mutationResolver) DeleteBookMeta(ctx context.Context, id string) (*string, error) {
	if err := r.Remote.Delete(&models.BookMetadata{ID: id}).Error; err != nil {
		return nil, err
	}

	return &id, nil
}

// UpdateBookMeta is the resolver for the updateBookMeta field.
func (r *mutationResolver) UpdateBookMeta(ctx context.Context, inp graph.UpdateBookMetadataInp) (*graph.BookMetadata, error) {
	var b models.BookMetadata
	if err := r.Remote.First(&b, "id = ?", inp.ID).Error; err != nil {
		return nil, err
	}

	if err := mergo.Merge(&b, inp, mergo.WithOverride); err != nil {
		return nil, err
	}

	if err := r.Remote.Save(&b).Error; err != nil {
		return nil, err
	}

	return b.ToGraphQL(), nil
}

// ChangeBookMetaAuthors is the resolver for the changeBookMetaAuthors field.
func (r *mutationResolver) ChangeBookMetaAuthors(ctx context.Context, id string, authors []string) (*graph.BookMetadata, error) {
	var b models.BookMetadata
	if err := r.Remote.First(&b, "id = ?", id).Error; err != nil {
		return nil, err
	}

	b.Authors = utils.Map(authors, func(id string) models.Author {
		return models.Author{ID: id}
	})

	if err := r.Remote.Save(&b).Error; err != nil {
		return nil, err
	}

	return b.ToGraphQL(), nil
}

// BookMeta is the resolver for the bookMeta field.
func (r *queryResolver) BookMeta(ctx context.Context, id string) (*graph.BookMetadata, error) {
	var b models.BookMetadata
	if err := r.Remote.First(&b, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return b.ToGraphQL(), nil
}

// BookMetas is the resolver for the bookMetas field.
func (r *queryResolver) BookMetas(ctx context.Context) ([]*graph.BookMetadata, error) {
	var bs []models.BookMetadata
	if err := r.Remote.Find(&bs).Error; err != nil {
		return nil, err
	}

	return utils.Map(bs, func(b models.BookMetadata) *graph.BookMetadata {
		return b.ToGraphQL()
	}), nil
}

// BookMetadata returns graph.BookMetadataResolver implementation.
func (r *Resolver) BookMetadata() graph.BookMetadataResolver { return &bookMetadataResolver{r} }

type bookMetadataResolver struct{ *Resolver }
