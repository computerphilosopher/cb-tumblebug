package mcir

import (
	"fmt"
	"net/http"

	"github.com/cloud-barista/cb-tumblebug/src/core/common"
	"github.com/cloud-barista/cb-tumblebug/src/core/mcir"
	"github.com/labstack/echo/v4"
)

// RestPostImage godoc
// @Summary Register image
// @Description Register image
// @Tags Image
// @Accept  json
// @Produce  json
// @Param registeringMethod query string true "registerWithInfo or registerWithId"
// @Param nsId path string true "Namespace ID"
// @Param imageInfo body mcir.TbImageInfo false "Details for an image object"
// @Param imageId body mcir.TbImageReq false "name, connectionName and cspImageId"
// @Success 200 {object} mcir.TbImageInfo
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/image [post]
func RestPostImage(c echo.Context) error {

	nsId := c.Param("nsId")

	action := c.QueryParam("action")
	fmt.Println("[POST Image requested action: " + action)
	/*
		if action == "create" {
			fmt.Println("[Creating Image]")
			content, _ := createImage(nsId, u)
			return c.JSON(http.StatusCreated, content)

		} else */
	if action == "registerWithInfo" {
		fmt.Println("[Registering Image with info]")
		u := &mcir.TbImageInfo{}
		if err := c.Bind(u); err != nil {
			return err
		}
		content, err := mcir.RegisterImageWithInfo(nsId, u)
		if err != nil {
			common.CBLog.Error(err)
			mapA := map[string]string{"message": err.Error()}
			return c.JSON(http.StatusFailedDependency, &mapA)
		}
		return c.JSON(http.StatusCreated, content)
	} else if action == "registerWithId" {
		fmt.Println("[Registering Image with ID]")
		u := &mcir.TbImageReq{}
		if err := c.Bind(u); err != nil {
			return err
		}
		//content, responseCode, body, err := RegisterImageWithId(nsId, u)
		content, err := mcir.RegisterImageWithId(nsId, u)
		if err != nil {
			common.CBLog.Error(err)
			//fmt.Println("body: ", string(body))
			//return c.JSONBlob(responseCode, body)
			mapA := map[string]string{"message": err.Error()}
			return c.JSON(http.StatusFailedDependency, &mapA)
		}
		return c.JSON(http.StatusCreated, content)
	} else {
		mapA := map[string]string{"message": "You must specify: action=registerWithInfo or action=registerWithId"}
		return c.JSON(http.StatusFailedDependency, &mapA)
	}

}

/* function RestPutImage not yet implemented
// RestPutImage godoc
// @Summary Update image
// @Description Update image
// @Tags Image
// @Accept  json
// @Produce  json
// @Param imageInfo body mcir.TbImageInfo true "Details for an image object"
// @Success 200 {object} mcir.TbImageInfo
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/image/{imageId} [put]
*/
func RestPutImage(c echo.Context) error {
	//nsId := c.Param("nsId")

	return nil
}

// Request structure for RestLookupImage
type RestLookupImageRequest struct {
	ConnectionName string `json:"connectionName"`
}

// RestLookupImage godoc
// @Summary Lookup image
// @Description Lookup image
// @Tags Image
// @Accept  json
// @Produce  json
// @Param connectionName body RestLookupImageRequest true "Imageify connectionName"
// @Param imageName path string true "Image name"
// @Success 200 {object} mcir.SpiderImageInfo
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /lookupImage/{imageName} [get]
func RestLookupImage(c echo.Context) error {

	u := &RestLookupImageRequest{}
	if err := c.Bind(u); err != nil {
		return err
	}

	imageName := c.Param("imageName")
	fmt.Println("[Lookup image]" + imageName)
	content, err := mcir.LookupImage(u.ConnectionName, imageName)
	if err != nil {
		common.CBLog.Error(err)
		return c.JSONBlob(http.StatusFailedDependency, []byte(err.Error()))
	}

	return c.JSON(http.StatusOK, &content)

}

// RestLookupImageList godoc
// @Summary Lookup image list
// @Description Lookup image list
// @Tags Image
// @Accept  json
// @Produce  json
// @Param connectionName body RestLookupImageRequest true "Imageify connectionName"
// @Success 200 {object} mcir.SpiderImageList
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /lookupImage [get]
func RestLookupImageList(c echo.Context) error {

	//type JsonTemplate struct {
	//	ConnectionName string
	//}

	u := &RestLookupImageRequest{}
	if err := c.Bind(u); err != nil {
		return err
	}

	fmt.Println("[Get Region List]")
	content, err := mcir.LookupImageList(u.ConnectionName)
	if err != nil {
		common.CBLog.Error(err)
		return c.JSONBlob(http.StatusFailedDependency, []byte(err.Error()))
	}

	return c.JSON(http.StatusOK, &content)

}

// RestFetchImages godoc
// @Summary Fetch images
// @Description Fetch images
// @Tags Image
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Success 200 {object} common.SimpleMsg
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/fetchImages [post]
func RestFetchImages(c echo.Context) error {

	nsId := c.Param("nsId")

	connConfigCount, imageCount, err := mcir.FetchImages(nsId)
	if err != nil {
		common.CBLog.Error(err)
		mapA := map[string]string{
			"message": err.Error()}
		return c.JSON(http.StatusFailedDependency, &mapA)
	}

	mapA := map[string]string{
		"message": "Fetched " + fmt.Sprint(imageCount) + " images (from " + fmt.Sprint(connConfigCount) + " connConfigs)"}
	return c.JSON(http.StatusCreated, &mapA) //content)
}

// RestGetImage godoc
// @Summary Get image
// @Description Get image
// @Tags Image
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Param imageId path string true "Image ID"
// @Success 200 {object} mcir.TbImageInfo
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/image/{imageId} [get]
func RestGetImage(c echo.Context) error {
	// Obsolete function. This is just for Swagger.
	return nil
}

// Response structure for RestGetAllImage
type RestGetAllImageResponse struct {
	Image []mcir.TbImageInfo `json:"image"`
}

// RestGetAllImage godoc
// @Summary List all images
// @Description List all images
// @Tags Image
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Success 200 {object} RestGetAllImageResponse
// @Failure 404 {object} common.SimpleMsg
// @Failure 500 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/image [get]
func RestGetAllImage(c echo.Context) error {
	// Obsolete function. This is just for Swagger.
	return nil
}

// RestDelImage godoc
// @Summary Delete image
// @Description Delete image
// @Tags Image
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Param imageId path string true "Image ID"
// @Success 200 {object} common.SimpleMsg
// @Failure 404 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/image/{imageId} [delete]
func RestDelImage(c echo.Context) error {
	// Obsolete function. This is just for Swagger.
	return nil
}

// RestDelAllImage godoc
// @Summary Delete all images
// @Description Delete all images
// @Tags Image
// @Accept  json
// @Produce  json
// @Param nsId path string true "Namespace ID"
// @Success 200 {object} common.SimpleMsg
// @Failure 404 {object} common.SimpleMsg
// @Router /ns/{nsId}/resources/image [delete]
func RestDelAllImage(c echo.Context) error {
	// Obsolete function. This is just for Swagger.
	return nil
}
