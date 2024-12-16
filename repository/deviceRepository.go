package repository

import (
	"errors"
	"time"

	"github.com/pealan/golang-rest-api/model"
	"gorm.io/gorm"
)

type DeviceRepository struct {
	db *gorm.DB
}

// Initialize Device database repository
func DeviceRepositoryInit(db *gorm.DB) *DeviceRepository {
	db.AutoMigrate(&model.Device{})

	return &DeviceRepository{db: db}
}

// Returns a list with all devices in the database. Empty lists are a valid response
func (d *DeviceRepository) FindAllDevices() ([]model.Device, error) {
	var devices []model.Device

	err := d.db.Find(&devices).Error
	if err != nil {
		return nil, err
	}

	return devices, nil
}

// Returns device from database with the given id. If no device with that id is found
// returns nil object with nil error.
func (d *DeviceRepository) FindDeviceById(id int) (*model.Device, error) {
	device := model.Device{
		ID: id,
	}

	err := d.db.First(&device).Error
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &device, nil
}

// Returns device from database with the given brand.
func (d *DeviceRepository) FindDeviceByBrand(brand string) ([]model.Device, error) {
	var devices []model.Device

	err := d.db.Where("brand = ?", brand).Find(&devices).Error
	if err != nil {
		return nil, err
	}

	return devices, nil
}

// Saves device to the database. Fills the CreationTime attribute with the current time
func (d *DeviceRepository) Save(device *model.Device) (model.Device, error) {
	if device == nil || *device == (model.Device{}) {
		return model.Device{}, errors.New("cannot process nil/empty device")
	}

	device.CreationTime = time.Now().UTC()

	err := d.db.Save(device).Error
	if err != nil {
		return model.Device{}, err
	}

	return *device, nil
}

// Deletes device with given ID from database. Returns a flag indicating if this device was
// successfully deleted in the process. If no device is found with the ID, no error is returned
// and the flag is set to false
func (d *DeviceRepository) DeleteDeviceById(id int) (bool, error) {
	db := d.db.Delete(&model.Device{}, id)

	if db.Error != nil {
		return false, db.Error
	}

	if db.RowsAffected < 1 {
		return false, nil
	}

	return true, nil
}

// Rollback transaction
func (d *DeviceRepository) Rollback() error {
	err := d.db.Rollback().Error
	if err != nil {
		return err
	}

	return nil
}
