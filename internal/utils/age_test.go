package utils

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name string
		dob  time.Time
		want int
	}{
		{
			name: "birthday already passed this year",
			dob:  time.Date(now.Year()-30, now.Month()-1, now.Day(), 0, 0, 0, 0, now.Location()),
			want: 30,
		},
		{
			name: "birthday not yet occurred this year",
			dob:  time.Date(now.Year()-30, now.Month()+1, now.Day(), 0, 0, 0, 0, now.Location()),
			want: 29,
		},
		{
			name: "birthday today",
			dob:  time.Date(now.Year()-30, now.Month(), now.Day(), 0, 0, 0, 0, now.Location()),
			want: 30,
		},
		{
			name: "leap year dob",
			dob:  time.Date(2004, time.February, 29, 0, 0, 0, 0, now.Location()),
			want: func() int {
				age := now.Year() - 2004
				birthdayThisYear := time.Date(now.Year(), time.February, 29, 0, 0, 0, 0, now.Location())
				if now.Before(birthdayThisYear) {
					age--
				}
				if age < 0 {
					return 0
				}
				return age
			}(),
		},
		{
			name: "future dob should not return negative age",
			dob:  now.AddDate(1, 0, 0),
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateAge(tt.dob)
			if got != tt.want {
				t.Fatalf("CalculateAge(%v) = %d, want %d", tt.dob, got, tt.want)
			}
		})
	}
}
