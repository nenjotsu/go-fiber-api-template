package usecase

import (
	"errors"
	"fmt"
	"go-fiber-api-template/internals/entity"
	"go-fiber-api-template/internals/repository"
	"strings"
	"time"
)

type ForexfactoryUsecase interface {
	GetForexFactory(timeString string) ([]*entity.ForexFactory, error)
	UpsertForexFactory(requestBody []*entity.CreateForexFactoryRequest) ([]*entity.ForexFactory, error)
}

type forexfactoryUsecase struct {
	repo repository.ForexFactoryRepository
}

func NewForexfactoryUsecase(repo repository.ForexFactoryRepository) ForexfactoryUsecase {
	return &forexfactoryUsecase{repo: repo}
}

func switchDayMonth(input string) string {

	// Split the date and time components
	parts := strings.Split(input, " ")
	date := parts[0]
	timePart := parts[1]

	// Split the date into components (MM/DD/YYYY)
	dateParts := strings.Split(date, "/")
	month := dateParts[0]
	day := dateParts[1]
	year := dateParts[2]

	// Switch month and day
	switchedDate := fmt.Sprintf("%s/%s/%s", day, month, year)

	// Recombine with the time part
	result := fmt.Sprintf("%s %s", switchedDate, timePart)
	return result
}

func parseTimeToManila(timeText string) (*time.Time, error) {
	layout := "01/02/2006 15:04"
	manilaTime, err := time.Parse(layout, timeText)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, errors.New("error parsing date")
	}
	return &manilaTime, nil
}
func parseTimeToUtc(timeText string) (*time.Time, error) {
	layout := "01/02/2006 15:04"
	manilaTZ, err := time.LoadLocation("Asia/Manila")
	if err != nil {
		fmt.Println("Error loading Manila timezone:", err)
		return nil, errors.New("error loading timezone")
	}

	// Parse the input date in Manila timezone
	manilaTime, err := time.ParseInLocation(layout, timeText, manilaTZ)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, errors.New("error parsing date")
	}

	utcTime := manilaTime.UTC()
	return &utcTime, nil
}

func (u *forexfactoryUsecase) GetForexFactory(timeString string) ([]*entity.ForexFactory, error) {
	forexFactoryList, err := u.repo.GetForexFactory(timeString)
	if err != nil {
		return nil, err
	}
	if forexFactoryList == nil {
		return nil, errors.New("forexFactoryList is nil")
	}

	return forexFactoryList, nil
}
func (u *forexfactoryUsecase) UpsertForexFactory(requestBody []*entity.CreateForexFactoryRequest) ([]*entity.ForexFactory, error) {
	list := make([]*entity.ForexFactory, 0)
	for _, forexfactory := range requestBody {
		existing, err := u.repo.FindByTime(forexfactory.Time)
		if err != nil {
			fmt.Printf("Error finding forex factory: %v", err)
		}
		validTime := switchDayMonth(forexfactory.Time)
		manilaTime, err := parseTimeToManila(validTime)
		if err != nil {
			fmt.Printf("Error finding forex factory: %v", err)
		}

		utcTime, err := parseTimeToUtc(validTime)
		if err != nil {
			fmt.Printf("Error finding forex factory: %v", err)
		}

		if existing == nil {
			newForexFactory := &entity.ForexFactory{
				Currency:       forexfactory.Currency,
				Event:          forexfactory.Event,
				Impact:         forexfactory.Impact,
				Time:           forexfactory.Time,
				DateTimeUtc:    *utcTime,
				DateTimeManila: *manilaTime,
			}
			list = append(list, newForexFactory)
			err = u.repo.Create(newForexFactory)
			if err != nil {
				fmt.Printf("Error creating forex factory: %v", err)
			}
		}

		if existing != nil {
			existing.Impact = forexfactory.Impact
			existing.Time = forexfactory.Time
			existing.DateTimeUtc = *utcTime
			existing.DateTimeManila = *manilaTime
			err = u.repo.UpdateByTime(existing)
			if err != nil {
				fmt.Printf("Error updating forex factory: %v", err)
			}
		}
	}

	return list, nil
}
