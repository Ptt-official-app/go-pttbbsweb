package fav

import (
	"testing"

	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
)

func TestFavType_GetID(t *testing.T) {
	ft0 := &FavType{0, pttbbsfav.FAVT_BOARD, pttbbsfav.FAVH_FAV, &FavBoard{2, 0, 0}}
	ft1 := &FavType{0, pttbbsfav.FAVT_LINE, pttbbsfav.FAVH_FAV, &FavLine{3}}
	ft2 := &FavType{0, pttbbsfav.FAVT_FOLDER, pttbbsfav.FAVH_FAV, &FavFolder{Fid: 4}}
	tests := []struct {
		name          string
		ft            *FavType
		expectedTheID int
	}{
		// TODO: Add test cases.
		{
			ft:            ft0,
			expectedTheID: 2,
		},
		{
			ft:            ft1,
			expectedTheID: 3,
		},
		{
			ft:            ft2,
			expectedTheID: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft := tt.ft
			if gotTheID := ft.GetID(); gotTheID != tt.expectedTheID {
				t.Errorf("FavType.GetID() = %v, want %v", gotTheID, tt.expectedTheID)
			}
		})
	}
}
