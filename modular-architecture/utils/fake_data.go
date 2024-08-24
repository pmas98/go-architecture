package fakedata

import (
	"fmt"
	"time"

	"github.com/pmas98/go-architecture/modular-architecture/internal/domain"
	"gorm.io/gorm"
)

func PopulateDb(db *gorm.DB) {
	// Create fake data with proper time conversion
	books := []domain.Book{
		{Title: "The Way of Kings", Author: "Brandon Sanderson", PublishedAt: parseDate("2010-08-31")},
		{Title: "The Name of the Wind", Author: "Patrick Rothfuss", PublishedAt: parseDate("2007-03-27")},
		{Title: "Dune", Author: "Frank Herbert", PublishedAt: parseDate("1965-08-01")},
		{Title: "Mistborn: The Final Empire", Author: "Brandon Sanderson", PublishedAt: parseDate("2006-07-01")},
		{Title: "The Hobbit", Author: "J.R.R. Tolkien", PublishedAt: parseDate("1937-09-21")},
		{Title: "Ender's Game", Author: "Orson Scott Card", PublishedAt: parseDate("1985-01-15")},
		{Title: "Neuromancer", Author: "William Gibson", PublishedAt: parseDate("1984-07-01")},
		{Title: "Foundation", Author: "Isaac Asimov", PublishedAt: parseDate("1951-06-01")},
		{Title: "Hyperion", Author: "Dan Simmons", PublishedAt: parseDate("1989-05-01")},
		{Title: "The Expanse: Leviathan Wakes", Author: "James S.A. Corey", PublishedAt: parseDate("2011-06-02")},
	}

	// Insert data into the database
	for _, book := range books {
		result := db.Create(&book)
		if result.Error != nil {
			fmt.Printf("Error inserting book %s: %v\n", book.Title, result.Error)
		}
	}

	fmt.Println("Fake data has been inserted into the books table.")
}

func parseDate(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Printf("Error parsing date %s: %v\n", dateStr, err)
		return time.Time{}
	}
	return t
}
