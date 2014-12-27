package library

import (
	"testing"
)

func TestOps(t *testing.T) {

	mm := NewMusicManager()

	if mm == nil {
		t.Error("NewMusicManager failed.")
	}

	if len(mm.musics) !=0 {
		t.Error("NewMusicManager failed, not empty.")
	}
	m0 := &MusicEntry{"1","My Heart Will Go On","Celion Dion","POP","http://qbox.me/32145","MP3"}

	mm.Add(m0)

	if len(mm.musics) != 1 {
		t.Error("MusicManager.Add() failed.");
	}
	
	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find() failed.")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist ||
		m.Name != m0.Name || m.Genre != m0.Genre ||
		m.Source != m0.Source || m.Type != m0.Type{
		t.Error("MusicManager.Find() failed. Found item mismatch.")
	}

	m, err := mm.Get(1)

	if m == nil {
		t.Error("Musicmanger.Get() failed.",err)
	}

	m = mm.Remove(1)
	if m == nil || len(mm.musics) !=0 {
		t.Error("MusicManager.Remove() failed.",err)
	}
}