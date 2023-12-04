package handler

import (
	"fmt"
	"github.com/drakenchef/RIP/internal/app/ds"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strconv"
)

//	func (h *Handler) PlanetsList(ctx *gin.Context) {
//		searchQuery := ctx.Query("search")
//		if searchQuery == "" {
//			planets, err := h.Repository.PlanetsList()
//			if err != nil {
//				ctx.JSON(http.StatusInternalServerError, gin.H{
//					"error": err.Error(),
//				})
//				return
//			}
//			ctx.JSON(http.StatusOK, gin.H{
//				"Planets":   planets,
//				"Flight_id": user_request_id,
//			})
//		} else {
//
//			filteredPlanets, err := h.Repository.SearchPlanet(searchQuery)
//			if err != nil {
//				// обработка ошибки
//			}
//			ctx.JSON(http.StatusOK, gin.H{
//				"Planets":   filteredPlanets,
//				"Flight_id": user_request_id,
//			})
//
//		}
//	}

func (h *Handler) PlanetsList(ctx *gin.Context) {
	userID := 1
	searchQuery := ctx.Query("search")
	if searchQuery == "" {
		planets, err := h.Repository.PlanetsList()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		// Получаем id заявки пользователя
		userRequestID, err := h.Repository.GetUserRequestID(userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Planets":   planets,
			"Flight_id": userRequestID,
		})
	} else {
		filteredPlanets, err := h.Repository.SearchPlanet(searchQuery)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		// Получаем id заявки пользователя
		userRequestID, err := h.Repository.GetUserRequestID(userID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Planets":   filteredPlanets,
			"Flight_id": userRequestID,
		})
	}
}

//func (h *Handler) PlanetsList(ctx *gin.Context) {
//	searchQuery := ctx.Query("search")
//	if searchQuery == "" {
//		planets, err := h.Repository.PlanetsList()
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		// Получаем id заявки пользователя
//		userRequestID, err := h.Repository.GetUserRequestID(1)
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		ctx.JSON(http.StatusOK, gin.H{
//			"Planets":   planets,
//			"Flight_id": userRequestID,
//		})
//	} else {
//		filteredPlanets, err := h.Repository.SearchPlanet(searchQuery)
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//		// Получаем id заявки пользователя
//		userRequestID, err := h.Repository.GetUserRequestID(1)
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"error": err.Error(),
//			})
//		}
//		ctx.JSON(http.StatusOK, gin.H{
//			"Planets":   filteredPlanets,
//			"Flight_id": userRequestID,
//		})
//	}
//}

func (h *Handler) PlanetById(ctx *gin.Context) {
	id := ctx.Param("id")
	idint, _ := strconv.Atoi(id)
	planets, err := h.Repository.PlanetById(idint)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Planets": planets,
	})
	//var request struct {
	//	ID uint `json:"id"`
	//}
	//if err := ctx.BindJSON(&request); err != nil {
	//	h.errorHandler(ctx, http.StatusBadRequest, err)
	//	return
	//}
	//if request.ID == 0 {
	//	h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
	//	return
	//}
	//if planet, err := h.Repository.PlanetById(request.ID); err != nil {
	//	h.errorHandler(ctx, http.StatusInternalServerError, err)
	//	return
	//} else {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"Planet": planet,
	//	})
	//}
}

func (h *Handler) DeletePlanet(ctx *gin.Context) {
	var request struct {
		ID uint `json:"id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if request.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.DeletePlanet(request.ID); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "Planet_id", request.ID)
}

func (h *Handler) AddPlanet(ctx *gin.Context) {

	planetName := ctx.Request.FormValue("name")
	description := ctx.Request.FormValue("description")
	radius := ctx.Request.FormValue("radius")
	distance := ctx.Request.FormValue("distance")
	gravity := ctx.Request.FormValue("gravity")
	types := ctx.Request.FormValue("type")
	radiusfloat, _ := strconv.ParseFloat(radius, 64)
	distancefloat, _ := strconv.ParseFloat(distance, 64)
	gravityfloat, _ := strconv.ParseFloat(gravity, 64)

	newPlanet := ds.Planet{
		Name:        planetName,
		IsDelete:    false,
		Description: description,
		Radius:      radiusfloat,
		Distance:    distancefloat,
		Gravity:     gravityfloat,
		Type:        types,
	}
	file, header, _ := ctx.Request.FormFile("image")
	if errorCode, errCreate := h.createPlanet(&newPlanet); errCreate != nil {
		h.errorHandler(ctx, errorCode, errCreate)
	}
	if file != nil && header.Size != 0 && header != nil {
		newImageURL, errCode, errDB := h.createImagePlanet(&file, header, fmt.Sprintf("%d", newPlanet.ID))
		if errDB != nil {
			h.errorHandler(ctx, errCode, errDB)
			return
		}
		newPlanet.Image = newImageURL
	}
	ctx.Redirect(http.StatusFound, "/Planets")
}

func (h *Handler) createPlanet(planet *ds.Planet) (int, error) {
	if planet.ID != 0 {
		return http.StatusBadRequest, idMustBeEmpty
	}
	if planet.Name == "" {
		return http.StatusBadRequest, planetCannotBeEmpty
	}
	if err := h.Repository.AddPlanet(planet); err != nil {
		return http.StatusBadRequest, err
	}
	return 0, nil
}

//func (h *Handler) UpdatePlanet(ctx *gin.Context) {
//	var updatedPlanet ds.Planet
//	if err := ctx.BindJSON(&updatedPlanet); err != nil {
//		h.errorHandler(ctx, http.StatusBadRequest, err)
//		return
//	}
//	if updatedPlanet.ID == 0 {
//		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
//		return
//	}
//	if err := h.Repository.UpdatePlanet(&updatedPlanet); err != nil {
//		h.errorHandler(ctx, http.StatusBadRequest, err)
//		return
//	}
//
//	h.successHandler(ctx, "updated_planet", gin.H{
//		"id":          updatedPlanet.ID,
//		"name":        updatedPlanet.Name,
//		"description": updatedPlanet.Description,
//		"radius":      updatedPlanet.Radius,
//		"distance":    updatedPlanet.Distance,
//		"gravity":     updatedPlanet.Gravity,
//		"image":       updatedPlanet.Image,
//		"type":        updatedPlanet.Type,
//		"is_delete":   updatedPlanet.IsDelete,
//	})
//}

func (h *Handler) UpdatePlanet(ctx *gin.Context) {
	planetId := ctx.Param("id")
	//planetId := ctx.Request.FormValue("id")
	planetName := ctx.Request.FormValue("name")
	description := ctx.Request.FormValue("description")
	radius := ctx.Request.FormValue("radius")
	distance := ctx.Request.FormValue("distance")
	gravity := ctx.Request.FormValue("gravity")
	types := ctx.Request.FormValue("type")
	planetIduint, _ := strconv.Atoi(planetId)
	radiusfloat, _ := strconv.ParseFloat(radius, 64)
	distancefloat, _ := strconv.ParseFloat(distance, 64)
	gravityfloat, _ := strconv.ParseFloat(gravity, 64)

	newPlanet := ds.Planet{
		ID:          uint(planetIduint),
		Name:        planetName,
		IsDelete:    false,
		Description: description,
		Radius:      radiusfloat,
		Distance:    distancefloat,
		Gravity:     gravityfloat,
		Type:        types,
	}
	file, header, _ := ctx.Request.FormFile("image")
	if errorCode, errCreate := h.updatePlanet(&newPlanet); errCreate != nil {
		h.errorHandler(ctx, errorCode, errCreate)
	}
	if file != nil && header.Size != 0 && header != nil {
		newImageURL, errCode, errDB := h.createImagePlanet(&file, header, fmt.Sprintf("%d", newPlanet.ID))
		if errDB != nil {
			h.errorHandler(ctx, errCode, errDB)
			return
		}
		newPlanet.Image = newImageURL
	}
	ctx.Redirect(http.StatusFound, "/Planets")

}

// asd
func (h *Handler) updatePlanet(planet *ds.Planet) (int, error) {
	if planet.ID == 0 {
		return http.StatusBadRequest, idNotFound
	}
	if err := h.Repository.UpdatePlanet(planet); err != nil {
		return http.StatusBadRequest, err
	}
	return 0, nil
}

func (h *Handler) AddImage(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	planetID := ctx.Request.FormValue("id")

	if planetID == "" {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if header == nil || header.Size == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, headerNotFound)
		return
	}
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	defer func(file multipart.File) {
		errLol := file.Close()
		if errLol != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, errLol)
			return
		}
	}(file)

	// Upload the image to minio server.
	newImageURL, errorCode, errImage := h.createImagePlanet(&file, header, planetID)
	if errImage != nil {
		h.errorHandler(ctx, errorCode, errImage)
		return
	}

	h.successAddHandler(ctx, "image_url", newImageURL)
}

func (h *Handler) createImagePlanet(
	file *multipart.File,
	header *multipart.FileHeader,
	planetID string,
) (string, int, error) {
	newImageURL, errMinio := h.createImageInMinio(file, header)
	if errMinio != nil {
		return "", http.StatusInternalServerError, errMinio
	}
	if err := h.Repository.UpdatePlanetImage(planetID, newImageURL); err != nil {
		return "", http.StatusInternalServerError, err
	}
	return newImageURL, 0, nil
}
