package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pealan/golang-rest-api/model"
	"github.com/pealan/golang-rest-api/repository"
)

type DeviceHandler struct {
	deviceRepository repository.DeviceRepository
}

func DeviceHandlerInit(deviceRepository *repository.DeviceRepository) *DeviceHandler {
	return &DeviceHandler{
		deviceRepository: *deviceRepository,
	}
}

// @Summary		Add Device
// @Description	Adds a new Device
// @Tags			Device
// @Accept			json
// @Produce		    json
// @Param			device  body  model.PartialDevice true	"Device that will be added"
// @Success		201	{object}    model.Device{}
// @Failure		400	{object}	ErrorResponse{}
// @Failure		500	{object}	ErrorResponse{}
// @Router			/device [post]
func (d DeviceHandler) AddDevice(c *gin.Context) {
	log.Println("Starting AddDevice")

	var request model.PartialDevice
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, &ErrorResponse{RawError: err.Error()})
		return
	}

	if request == (model.PartialDevice{}) {
		c.JSON(http.StatusBadRequest, &ErrorResponse{Message: "Request does not contain any expected fields"})
		return
	}

	// Finally fill new device fields
	var newDevice model.Device
	if request.Name != nil {
		newDevice.Name = *request.Name
	}

	if request.Brand != nil {
		newDevice.Brand = *request.Brand
	}

	newDevice, err = d.deviceRepository.Save(&newDevice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{RawError: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newDevice)
}

// @Summary		Get Devices
// @Description	Gets a list of all registered Devices. The user has also the option to filter by brand. Empty lists can be returned.
// @Tags			Device
// @Produce		    json
// @Param			brand  query  string  false	"Filter results by this brand"
// @Success		200	{object}    []model.Device
// @Failure		500	{object}	ErrorResponse{}
// @Router			/device [get]
func (d DeviceHandler) GetAllDevices(c *gin.Context) {
	log.Println("Starting GetAllDevices")

	// TODO: Throw 400 error if request is malformed
	values := c.Request.URL.Query()

	var devices []model.Device
	var err error
	if values.Has("brand") {
		devices, err = d.deviceRepository.FindDeviceByBrand(values.Get("brand"))
	} else {
		devices, err = d.deviceRepository.FindAllDevices()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{RawError: err.Error()})
		return
	}

	c.JSON(http.StatusOK, devices)
}

// @Summary		Get Device by ID
// @Description	Gets Device by their ID
// @Tags			Device
// @Produce		    json
// @Param			id  path  int  true	"Device's ID"
// @Success		200	{object}    model.Device{}
// @Failure		404	{object}	ErrorResponse{}
// @Failure		500	{object}	ErrorResponse{}
// @Router			/device/{id} [get]
func (d DeviceHandler) GetDeviceById(c *gin.Context) {
	log.Println("Starting GetDeviceById")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{RawError: err.Error()})
		return
	}

	device, err := d.deviceRepository.FindDeviceById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{RawError: err.Error()})
		return
	}

	if device == nil {
		c.JSON(http.StatusNotFound, &ErrorResponse{Message: "Device with ID not found"})
		return
	}

	c.JSON(http.StatusOK, *device)
}

// @Summary		Update Device
// @Description	Updates the Device with the given ID. Full or partial forms are accepted
// @Tags			Device
// @Accept			json
// @Produce		    json
// @Param			id  path  int  true	"Device's ID"
// @Param			device  body  model.PartialDevice  true	"Device's new attributes."
// @Success		200	{object}    model.Device{}
// @Failure		400	{object}	ErrorResponse{}
// @Failure		404	{object}	ErrorResponse{}
// @Failure		500	{object}	ErrorResponse{}
// @Router			/device/{id} [put]
func (d DeviceHandler) UpdateDevice(c *gin.Context) {
	log.Println("Starting UpdateDevice")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{RawError: err.Error()})
		return
	}

	device, err := d.deviceRepository.FindDeviceById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{RawError: err.Error()})
		return
	}

	if device == nil {
		c.JSON(http.StatusNotFound, &ErrorResponse{Message: "Device with ID not found"})
		return
	}

	var request model.PartialDevice
	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, &ErrorResponse{RawError: err.Error()})
		return
	}

	if request == (model.PartialDevice{}) {
		c.JSON(http.StatusBadRequest, &ErrorResponse{Message: "Request does not contain any expected fields"})
		return
	}

	// Finally update device fields
	if request.Name != nil {
		device.Name = *request.Name
	}

	if request.Brand != nil {
		device.Brand = *request.Brand
	}

	updatedDevice, err := d.deviceRepository.Save(device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{RawError: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedDevice)
}

// @Summary		Delete Device
// @Description	Deletes the Device with the given ID
// @Tags			Device
// @Produce		    json
// @Param			id  path  int  true	"Device's ID"
// @Success		204
// @Failure		404	{object}	ErrorResponse{}
// @Failure		500	{object}	ErrorResponse{}
// @Router			/device/{id} [delete]
func (d DeviceHandler) DeleteDevice(c *gin.Context) {
	log.Println("Starting GetDeviceById")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{RawError: err.Error()})
		return
	}

	found, err := d.deviceRepository.DeleteDeviceById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrorResponse{RawError: err.Error()})
		return
	}

	if !found {
		c.JSON(http.StatusNotFound, &ErrorResponse{Message: "Device with ID not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
