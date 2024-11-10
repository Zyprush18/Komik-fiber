package requests

type Komik struct{

	Name string `json:"name_comic" validate:"required"`
	Creator string `json:"name_creator" validate:"required"`
	Publisher string `json:"name_publisher" validate:"required"`
	Releases string `json:"release_comic" validate:"required"`
	Image string `json:"image_comic" validate:"required"`
	Decsription string `json:"description" validate:"required"`
}