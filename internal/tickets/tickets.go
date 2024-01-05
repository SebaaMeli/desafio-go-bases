package tickets

import (
	"errors"
	"strconv"
	"strings"
)

type Ticket struct {
	ID                 int
	Name               string
	email              string
	DestinationCountry string
	FlightTime         string
	Price              int
}

const (
	earlyMorning = "earlyMorning"
	morning      = "morning"
	afternoon    = "afternoon"
	night        = "night"
)

var errInvalidPeriod = errors.New("invalid period")

func GetFilteredTicketsTotal(filter func(ticket Ticket) bool) (int, error) {
	tickets, err := GetTicketsFromFile()

	if err != nil {
		return 0, err
	}

	total := 0

	for _, ticket := range tickets {
		if filter(ticket) {
			total++
		}
	}

	return total, nil
}

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {
	total, err := GetFilteredTicketsTotal(GetDestinationFilter(destination))

	if err != nil {
		return 0, err
	}

	return total, nil
}

func GetDestinationFilter(destination string) func(ticket Ticket) bool {
	return func(ticket Ticket) bool {
		if destination == "" {
			return true
		}
		return ticket.DestinationCountry == destination
	}
}

// ejemplo 2
func GetCountByPeriod(period string) (int, error) {
	periodFilter, err := GetPeriodFilter(period)

	if err != nil {
		return 0, err
	}

	total, err := GetFilteredTicketsTotal(periodFilter)

	if err != nil {
		return 0, err
	}

	return total, nil
}

func GetPeriodFilter(period string) (func(ticket Ticket) bool, error) {
	if period != earlyMorning && period != morning && period != afternoon && period != night {
		return nil, errInvalidPeriod
	}

	return func(ticket Ticket) bool {
		hour := GetTicketHour(ticket)
		switch period {
		case earlyMorning:
			return hour >= 0 && hour <= 6
		case morning:
			return hour >= 7 && hour <= 12
		case afternoon:
			return hour >= 13 && hour <= 19
		case night:
			return hour >= 20 && hour <= 23
		default:
			return false
		}
	}, nil
}

func GetTicketHour(ticket Ticket) int {
	hour, _ := strconv.Atoi(strings.Split(ticket.FlightTime, ":")[0])
	return hour
}

// ejemplo 3
func AverageDestination(destination string, total int) (float64, error) {
	totalDestinationsTickets, err := GetFilteredTicketsTotal(GetDestinationFilter(destination))
	if err != nil {
		return 0, err
	}

	return float64(totalDestinationsTickets) / float64(total) * 100, nil
}
