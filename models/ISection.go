package models

type SectionProvider interface {
	GetSteps(sectionName string)
}
