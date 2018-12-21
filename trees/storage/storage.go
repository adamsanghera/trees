package storage

import "github.com/adamsanghera/trees/trees/treespb"

type Storage interface {
	// Commands

	// Queries
	Search(*treespb.SearchRequest) (*treespb.SearchResponse, error)
}
