package inmemory

import (
	"article/models"
	"errors"
	"strings"
	"time"
)

// AddAuthor ...
func (im Inmemory) AddAuthor(id string, entity models.CreateAuthorModel) error {
	var author models.Author
	author.ID = id
	author.Firstname = entity.Firstname
	author.Lastname = entity.Lastname
	author.CreatedAt = time.Now()

	im.Db.InMemoryAuthorData = append(im.Db.InMemoryAuthorData, author)
	return nil
}

// GetAuthorByID ...
func (im Inmemory) GetAuthorByID(id string) (models.Author, error) {
	var result models.Author
	for _, v := range im.Db.InMemoryAuthorData {
		if v.ID == id && v.DeletedAt == nil {
			result = v
			return result, nil
		}
	}
	return result, errors.New("author not found")
}

// GetAuthorList ...
func (im Inmemory) GetAuthorList(offset, limit int, search string) (resp []models.Author, err error) {
	var off, counter int
	for _, v := range im.Db.InMemoryAuthorData {
		if v.DeletedAt == nil &&
			strings.Contains(strings.ToLower(v.Firstname+v.Lastname), strings.ToLower(search)) {
			if offset <= off {
				counter++
				resp = append(resp, v)
			}
			if counter >= limit {
				break
			}
			off++
		}
	}
	return resp, err
}

// UpdateAuthor ...
func (im Inmemory) UpdateAuthor(req models.UpdateAuthorModel) error {
	for i, v := range im.Db.InMemoryAuthorData {
		if v.ID == req.ID {
			if v.DeletedAt != nil {
				return errors.New("Author already deleted!")
			}
			t := time.Now()
			im.Db.InMemoryAuthorData[i].UpdatedAt = &t
			if req.Firstname != "" {
				im.Db.InMemoryAuthorData[i].Firstname = req.Firstname
			}
			if req.Lastname != "" {
				im.Db.InMemoryAuthorData[i].Lastname = req.Lastname
			}
		}
	}
	return nil
}

// DeleteAuthor ...
func (im Inmemory) DeleteAuthor(id string) error {
	for i, v := range im.Db.InMemoryAuthorData {
		if v.ID == id && v.DeletedAt == nil {
			t := time.Now()
			im.Db.InMemoryAuthorData[i].DeletedAt = &t
			return nil
		}
	}
	return errors.New("Author had been deleted already!")
}
