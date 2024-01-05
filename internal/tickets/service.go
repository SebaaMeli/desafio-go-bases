package tickets

import (
	"strconv"

	"github.com/bootcamp-go/desafio-go-bases/internal/files"
)

const FileName = "../../tickets.csv"

func GetTicketsFromFile() ([]Ticket, error) {
	file, err := files.OpenFile(FileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	content, err := files.ReadCSV(file)

	if err != nil {
		return nil, err
	}

	tickets := NewTicketsSlice(content)

	return tickets, nil
}

func NewTicketsSlice(content [][]string) []Ticket {
	tickets := []Ticket{}

	for _, line := range content {
		tickets = append(tickets, NewTicket(line))
	}

	return tickets
}

func NewTicket(line []string) Ticket {
	ID, _ := strconv.Atoi(line[0])
	Price, _ := strconv.Atoi(line[5])

	return Ticket{
		ID:                 ID,
		Name:               line[1],
		email:              line[2],
		DestinationCountry: line[3],
		FlightTime:         line[4],
		Price:              Price,
	}
}
