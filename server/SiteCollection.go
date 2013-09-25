package server

import (
	"errors"
)

type SiteCollection struct {
	array []*Site
}

func (this *SiteCollection) Push(site *Site) {
	this.array = append(this.array, site)
}

func (this *SiteCollection) Get(index int) (*Site, error) {
	if len(this.array) <= index || index < 0 {
		return nil, errors.New("Argument out of range")
	}
	return this.array[index], nil
}

func (this *SiteCollection) Count() int {
	return len(this.array)
}

func (this *SiteCollection) Insert(index int, site *Site) error {
	err := this.move(1, index)
	this.array[index] = site
	return err
}

func (this *SiteCollection) move(count int, index int) error {
	if len(this.array) <= index || count <= 0 || index < 0 {
		return errors.New("Argument out of range")
	}

	originLength := len(this.array)

	for i := 0; i < count; i++ {
		this.array = append(this.array, nil)
	}

	for i := originLength - 1; i >= index; i-- {
		this.array[i+count] = this.array[i]
	}

	for i := 0; i < index; i++ {
		this.array[i] = nil
	}
	return nil
}
