package bio_card

import "github.com/mieubrisse/teact/teact/components"

type BioCard interface {
	components.Component

	SetName(name string) BioCard
	SetAge(age int) BioCard
}
