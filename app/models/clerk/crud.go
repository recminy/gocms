package clerk

import "cms/pkg/model"

func Find(uid string) (clerk Clerk, err error) {

	if err = model.DB.First(&clerk, 1).Error; err != nil {
		return clerk, err
	}
	return clerk, nil
}
