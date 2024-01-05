package tickets_test

import (
	"testing"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets(t *testing.T) {
	t.Run("should return 15 when destination is Argentina", func(t *testing.T) {
		destination := "Argentina"
		expected := 15
		total, _ := tickets.GetTotalTickets(destination)

		require.Equal(t, expected, total)

	})

	t.Run("should return 45 when destination is Brazil", func(t *testing.T) {
		destination := "Brazil"
		expected := 45
		total, _ := tickets.GetTotalTickets(destination)

		require.Equal(t, expected, total)
	})
}

func TestGetCountByPeriod(t *testing.T) {
	t.Run("should return 256 when period is morning", func(t *testing.T) {
		period := "morning"
		expected := 256
		total, _ := tickets.GetCountByPeriod(period)

		require.Equal(t, expected, total)

	})

	t.Run("should return 289 when period is afternoon", func(t *testing.T) {
		period := "afternoon"
		expected := 289
		total, _ := tickets.GetCountByPeriod(period)

		require.Equal(t, expected, total)
	})
	t.Run("should return error when period is invalid", func(t *testing.T) {
		period := "invalid"
		_, err := tickets.GetCountByPeriod(period)

		require.Error(t, err)
	})
}

func TestAverageDestination(t *testing.T) {
	t.Run("should return 1.5 when destination is Argentina", func(t *testing.T) {
		destination := "Argentina"
		total, _ := tickets.GetTotalTickets("")
		expected := 1.5

		averageDestination, _ := tickets.AverageDestination(destination, total)

		require.Equal(t, expected, averageDestination)

	})

	t.Run("should return 4.5 when destination is Brazil", func(t *testing.T) {
		destination := "Brazil"
		total, _ := tickets.GetTotalTickets("")
		expected := 4.5

		averageDestination, _ := tickets.AverageDestination(destination, total)

		require.Equal(t, expected, averageDestination)
	})
}
