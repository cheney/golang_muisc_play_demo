package library

import "errors"

type MusicEntry struct {
	Id string
	Name string
	Artist string
	Genre string
	Source string
	Type string
}

type MusicManager struct {
	Musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry,0)}
}
func (m *MusicManager) Len() int{
	return len(m.Musics)
}
func (m *MusicManager) Get(index int)(music *MusicEntry, err error) {
	if index <= 0 || index > len(m.Musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.Musics[index - 1], nil
}
func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.Musics) == 0 {
		return nil
	}
	for _, m := range m.Musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}
func (m *MusicManager) Add(music *MusicEntry) {
	m.Musics = append(m.Musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index <= 0 ||index > len(m.Musics) {
		return nil
	}

	if len(m.Musics) == 0 {
		return nil
	}

	removedMusic := &m.Musics[index -1]

	if  index == 1{
		m.Musics = m.Musics[1:]
	}else if index == len(m.Musics) {
		m.Musics = m.Musics[:index]
	}else if index < len(m.Musics) && index > 1 {
		m.Musics = append(m.Musics[:index],m.Musics[index+1:]...)
	}

	return removedMusic
}