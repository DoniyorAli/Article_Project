package handlers

import "UacademyGo/Article/storage/inmemory"

type Handler struct {
	IM inmemory.InMemory
}