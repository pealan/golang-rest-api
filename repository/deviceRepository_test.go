package repository

import (
	"log"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/pealan/golang-rest-api/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

var defaultCreationTime time.Time = time.Now().UTC()

var devices = []model.Device{
	{
		ID:           1,
		Name:         "iPhone 6",
		Brand:        "Apple",
		CreationTime: defaultCreationTime,
	},
	{
		ID:           2,
		Name:         "iPhone 10",
		Brand:        "Apple",
		CreationTime: defaultCreationTime,
	},
	{
		ID:           3,
		Name:         "S21",
		Brand:        "Samsung",
		CreationTime: defaultCreationTime,
	},
	{
		ID:           4,
		Name:         "Pocophone",
		Brand:        "Xiaomi",
		CreationTime: defaultCreationTime,
	},
}

func setup() *DeviceRepository {
	return &DeviceRepository{db: TestDB.Debug().Begin()}
}

func TestMain(m *testing.M) {
	// database file name
	dbName := "database_test.db"

	// remove old database, if it's there
	exec.Command("rm", "-f", dbName)

	// open and create a new database
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	TestDB = db

	err = TestDB.AutoMigrate(&model.Device{})
	if err != nil {
		log.Fatal(err)
	}

	// add mock data
	err = TestDB.Create(devices).Error
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestFindAllDevices_Success(t *testing.T) {
	repo := setup()
	d, err := repo.FindAllDevices()
	assert.Nil(t, err)
	assert.Len(t, d, len(devices))
}

func TestFindDeviceById_Success(t *testing.T) {
	repo := setup()
	d, err := repo.FindDeviceById(1)
	assert.Nil(t, err)
	assert.Equal(t, *d, devices[0])
}

func TestFindDeviceByBrand_Success(t *testing.T) {
	repo := setup()
	ds, err := repo.FindDeviceByBrand("Apple")
	assert.Nil(t, err)
	assert.Len(t, ds, 2)
	assert.Contains(t, ds, devices[0])
	assert.Contains(t, ds, devices[1])

	ds, err = repo.FindDeviceByBrand("Samsung")
	assert.Nil(t, err)
	assert.Len(t, ds, 1)
	assert.Contains(t, ds, devices[2])

	ds, err = repo.FindDeviceByBrand("Xiaomi")
	assert.Nil(t, err)
	assert.Len(t, ds, 1)
	assert.Contains(t, ds, devices[3])
}

func TestFindDeviceById_FailNotFound(t *testing.T) {
	repo := setup()
	d, err := repo.FindDeviceById(1000)
	// No error, but no device either
	assert.Nil(t, err)
	assert.Nil(t, d)
}

func TestFindDeviceByBrand_FailNotFound(t *testing.T) {
	repo := setup()
	d, err := repo.FindDeviceByBrand("xxxxx")
	assert.Nil(t, err)
	assert.Len(t, d, 0)

	d, err = repo.FindDeviceByBrand("")
	assert.Nil(t, err)
	assert.Len(t, d, 0)
}

func TestSaveDevice_Success(t *testing.T) {
	repo := setup()
	defer repo.Rollback()

	newDevice := model.Device{
		Name:  "iPhone 15",
		Brand: "Apple",
	}

	savedDevice, err := repo.Save(&newDevice)
	assert.Nil(t, err)

	expected := savedDevice

	actual, err := repo.FindDeviceById(savedDevice.ID)
	assert.Nil(t, err)
	assert.Equal(t, expected, *actual)
}

func TestSaveDevice_Fail_EmptyDevice(t *testing.T) {
	repo := setup()
	defer repo.Rollback()

	newDevice := model.Device{}
	_, err := repo.Save(&newDevice)

	//TODO Improve empty device error
	assert.NotNil(t, err)
}

func TestDeleteDeviceById_Success(t *testing.T) {
	repo := setup()
	defer repo.Rollback()

	wasDeleted, err := repo.DeleteDeviceById(1)
	assert.Nil(t, err)
	assert.True(t, wasDeleted)

	d, err := repo.FindDeviceById(1)
	assert.Nil(t, err)
	assert.Nil(t, d)
}

func TestDeleteDeviceById_Fail_NotFound(t *testing.T) {
	repo := setup()
	defer repo.Rollback()

	wasDeleted, err := repo.DeleteDeviceById(1000)
	assert.Nil(t, err)
	assert.False(t, wasDeleted)
}
