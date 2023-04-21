package photo

func FromDomainMapper(photo *Photo) *Photo {
	return &Photo{
		ID:        photo.ID,
		Title:     photo.Title,
		UserID:    photo.UserID,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UpdatedAt: photo.UpdatedAt,
	}
}

func ArrayToDomainMapper(photos *[]Photo) *[]Photo {
	booksDomain := make([]Photo, len(*photos))
	for i, photo := range *photos {
		booksDomain[i] = photo
	}

	return &booksDomain
}
