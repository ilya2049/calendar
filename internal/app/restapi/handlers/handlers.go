package handlers

func New(noteHandler *NoteHandler) *Handlers {
	return &Handlers{
		Note: noteHandler,
	}
}

type Handlers struct {
	Note *NoteHandler
}
