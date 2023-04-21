package sosmed

func FromDomainMapper(photo *SocialMedia) *SocialMedia {
	return &SocialMedia{
		ID:             photo.ID,
		UserID:         photo.UserID,
		SocialMediaUrl: photo.SocialMediaUrl,
		UpdatedAt:      photo.UpdatedAt,
	}
}

func ArrayToDomainMapper(photos *[]SocialMedia) *[]SocialMedia {
	booksDomain := make([]SocialMedia, len(*photos))
	for i, photo := range *photos {
		booksDomain[i] = photo
	}

	return &booksDomain
}
