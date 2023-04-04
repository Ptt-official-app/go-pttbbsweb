package fav

import (
	"encoding/binary"
	"io"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	pttbbsfav "github.com/Ptt-official-app/go-pttbbs/ptt/fav"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
	pttbbstypes "github.com/Ptt-official-app/go-pttbbs/types"
)

// Fav
//
// It's with it's own serialize method and does not directly copy by struct.
// We can add MTime in FavRaw.
// The content of FavFolder
type Fav struct {
	FavNum int
	Depth  int

	Root   *Fav
	Parent *Fav

	NBoards  int16         /* number of the boards */
	NLines   int8          /* number of the lines */
	NFolders int8          /* number of the folders */
	LineID   pttbbsfav.Lid /* current max line id */
	FolderID pttbbsfav.Fid /* current max folder id */

	Favh []*FavType
}

func NewFav(parent *Fav, root *Fav, depth int) (f *Fav, err error) {
	if depth >= pttbbsfav.FAV_MAXDEPTH {
		return nil, ErrTooMuchDepth
	}
	f = &Fav{
		Parent: parent,
		Root:   root,
		Depth:  depth,
	}
	if root == nil {
		f.Root = f
	}

	return f, nil
}

func (f *Fav) getDataNumber() int {
	return int(f.NBoards) + int(f.NLines) + int(f.NFolders)
}

func (f *Fav) WriteFavrec(file io.Writer) (err error) {
	if f == nil {
		return nil
	}

	if f.Depth == 0 {
		err = binary.Write(file, binary.LittleEndian, FAV_VERSION)
		if err != nil {
			return err
		}
	}

	err = binary.Write(file, binary.LittleEndian, &f.NBoards)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, &f.NLines)
	if err != nil {
		return err
	}
	err = binary.Write(file, binary.LittleEndian, &f.NFolders)
	if err != nil {
		return err
	}
	total := f.getDataNumber()

	for i := 0; i < total; i++ {
		ft := f.Favh[i]
		err := binary.Write(file, binary.LittleEndian, &ft.TheType)
		if err != nil {
			return err
		}
		err = binary.Write(file, binary.LittleEndian, &ft.Attr)
		if err != nil {
			return err
		}

		switch ft.TheType {
		case pttbbsfav.FAVT_FOLDER:
			favFolder := ft.CastFolder()
			if favFolder == nil {
				return ErrInvalidFavFolder
			}
			err = binary.Write(file, binary.LittleEndian, &favFolder.Fid)
			if err != nil {
				return err
			}
			titleBytes := types.Utf8ToBig5(favFolder.Title)
			title := &ptttype.BoardTitle_t{}
			copy(title[:], titleBytes)
			err = binary.Write(file, binary.LittleEndian, title)
			if err != nil {
				return err
			}
		case pttbbsfav.FAVT_BOARD:
			favBoard := ft.CastBoard()
			if favBoard == nil {
				return ErrInvalidFavBoard
			}
			err = pttbbstypes.BinWrite(file, favBoard, ft.TheType.GetTypeSize())
			if err != nil {
				return err
			}
		case pttbbsfav.FAVT_LINE:
			favLine := ft.CastLine()
			if favLine == nil {
				return ErrInvalidFavLine
			}
			err = pttbbstypes.BinWrite(file, favLine, ft.TheType.GetTypeSize())
			if err != nil {
				return err
			}
		}
	}

	for i := 0; i < total; i++ {
		ft := f.Favh[i]
		if ft == nil {
			continue
		}
		if ft.TheType == pttbbsfav.FAVT_FOLDER {
			favFolder := ft.CastFolder()
			if favFolder == nil {
				return ErrInvalidFavFolder
			}
			err := favFolder.ThisFolder.WriteFavrec(file)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ReadFavrec
//
// We need to:
// 1. read version because calling ReadFavrec
// 2. run root.SetFavTypeFavIdx(0)
//
// because ReadFavrec is a recursive-call,
// and the above 2 happens only in the 1st-call.
func ReadFavrec(file io.ReadSeeker, parent *Fav, root *Fav, depth int) (favrec *Fav, err error) {
	// https://github.com/ptt/pttbbs/blob/master/mbbsd/fav.c
	favrec, err = NewFav(parent, root, depth)
	if err != nil {
		return nil, err
	}
	err = binary.Read(file, binary.LittleEndian, &favrec.NBoards)
	if err != nil {
		return nil, ErrInvalidFavRecord
	}

	err = binary.Read(file, binary.LittleEndian, &favrec.NLines)
	if err != nil {
		return nil, ErrInvalidFavRecord
	}

	err = binary.Read(file, binary.LittleEndian, &favrec.NFolders)
	if err != nil {
		return nil, ErrInvalidFavRecord
	}

	nFavh := favrec.getDataNumber()
	favrec.LineID = 0
	favrec.FolderID = 0
	favrec.Favh = make([]*FavType, nFavh)

	for i := 0; i < nFavh; i++ {
		ft := &FavType{}
		favrec.IncreaseFavNum()

		favrec.Favh[i] = ft

		err = binary.Read(file, binary.LittleEndian, &ft.TheType)
		if err != nil {
			return nil, ErrInvalidFavType
		}
		if !ft.TheType.IsValidFavType() {
			return nil, ErrInvalidFavType
		}

		err = binary.Read(file, binary.LittleEndian, &ft.Attr)
		if err != nil {
			return nil, ErrInvalidFavType
		}

		switch ft.TheType {
		case pttbbsfav.FAVT_FOLDER:
			favFolder := ft.CastFolder()
			if favFolder == nil {
				return nil, ErrInvalidFavFolder
			}
			err = binary.Read(file, binary.LittleEndian, &favFolder.Fid)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			title := &ptttype.BoardTitle_t{}
			err = binary.Read(file, binary.LittleEndian, title)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			favFolder.Title = types.Big5ToUtf8(pttbbstypes.CstrToBytes(title[:]))
		case pttbbsfav.FAVT_BOARD:
			favBoard := ft.CastBoard()
			if favBoard == nil {
				return nil, ErrInvalidFavBoard
			}
			err = pttbbstypes.BinRead(file, favBoard, ft.TheType.GetTypeSize())
			if err != nil {
				return nil, ErrInvalidFavBoard
			}
		case pttbbsfav.FAVT_LINE:
			favLine := ft.CastLine()
			if favLine == nil {
				return nil, ErrInvalidFavLine
			}
			err = pttbbstypes.BinRead(file, favLine, ft.TheType.GetTypeSize())
			if err != nil {
				return nil, ErrInvalidFavLine
			}
		}
	}

	for i := 0; i < nFavh; i++ {
		ft := favrec.Favh[i]
		if ft == nil {
			continue
		}

		switch ft.TheType {
		case pttbbsfav.FAVT_FOLDER:
			favFolder := ft.CastFolder()
			if favFolder == nil {
				return nil, ErrInvalidFavFolder
			}
			p, err := ReadFavrec(file, favrec, favrec.Root, favrec.Depth+1)
			if err != nil {
				return nil, ErrInvalidFavFolder
			}
			favFolder.ThisFolder = p
			favrec.FolderID++
			favFolder.Fid = favrec.FolderID
		case pttbbsfav.FAVT_LINE:
			favLine := ft.CastLine()
			if favLine == nil {
				return nil, ErrInvalidFavLine
			}
			favrec.LineID++
			favLine.Lid = favrec.LineID
		}
	}

	return favrec, nil
}

func (f *Fav) IncreaseFavNum() {
	f.FavNum++
	if f.Parent != nil {
		f.Parent.IncreaseFavNum()
	}
}

func (f *Fav) DecreaseFavNum() {
	f.FavNum--
	if f.Parent != nil {
		f.Parent.DecreaseFavNum()
	}
}

func (f *Fav) AddBoard(bid ptttype.Bid) (favType *FavType, err error) {
	// check whether already added
	favBoard, err := f.GetBoard(bid)
	if err != nil {
		return nil, err
	}

	if favBoard != nil {
		return favBoard, nil
	}

	if f.isMaxSize() {
		return nil, ErrTooManyFavs
	}

	fp := &FavBoard{
		Bid: bid,
	}

	return f.PreAppend(pttbbsfav.FAVT_BOARD, fp)
}

func (f *Fav) GetBoard(bid ptttype.Bid) (favType *FavType, err error) {
	if !bid.IsValid() {
		return nil, ptttype.ErrInvalidBid
	}

	return f.GetFavItem(int(bid), pttbbsfav.FAVT_BOARD), nil
}

func (f *Fav) AddLine() (favType *FavType, err error) {
	if f.isMaxSize() {
		return nil, ErrTooManyFavs
	}

	if f.NLines >= MAX_LINE {
		return nil, ErrTooManyLines
	}

	fp := &FavLine{}
	return f.PreAppend(pttbbsfav.FAVT_LINE, fp)
}

func (f *Fav) AddFolder() (favType *FavType, err error) {
	if f.isMaxSize() {
		return nil, ErrTooManyFavs
	}

	if f.NLines >= MAX_FOLDER {
		return nil, ErrTooManyFolders
	}

	newFav, err := NewFav(f, f.Root, f.Depth+1)
	if err != nil {
		return nil, err
	}

	fp := &FavFolder{
		ThisFolder: newFav,
	}
	return f.PreAppend(pttbbsfav.FAVT_FOLDER, fp)
}

func (f *Fav) GetFavItem(theID int, theType pttbbsfav.FavT) (favType *FavType) {
	for _, each := range f.Favh {
		if each.TheType != theType {
			continue
		}

		eachID := each.GetID()
		if eachID == theID {
			return each
		}
	}

	return nil
}

// PreAppend
//
// https://github.com/ptt/pttbbs/blob/master/mbbsd/fav.c#L804
// Although it is named PreAppend, actually it appends to DataTail
func (f *Fav) PreAppend(theType pttbbsfav.FavT, fp interface{}) (favType *FavType, err error) {
	favt := &FavType{
		TheType: theType,
		Attr:    pttbbsfav.FAVH_FAV,
		Fp:      fp,
	}

	f.Favh = append(f.Favh, favt)
	f.Increase(theType, fp)

	return favt, nil
}

func (f *Fav) Increase(theType pttbbsfav.FavT, fp interface{}) {
	switch theType {
	case pttbbsfav.FAVT_BOARD:
		f.NBoards++
	case pttbbsfav.FAVT_LINE:
		f.NLines++
		f.LineID++
		ft, _ := fp.(*FavLine)
		ft.Lid = f.LineID
	case pttbbsfav.FAVT_FOLDER:
		f.NFolders++
		f.FolderID++
		ft, _ := fp.(*FavFolder)
		ft.Fid = f.FolderID
	}
	f.IncreaseFavNum()
}

func (f *Fav) isMaxSize() bool {
	return f.Root.FavNum >= MAX_FAV
}

func (f *Fav) CleanParentAndRoot() {
	f.Parent = nil
	f.Root = nil
	for _, each := range f.Favh {
		if each.TheType != pttbbsfav.FAVT_FOLDER {
			continue
		}

		ft, _ := each.Fp.(*FavFolder)
		ft.ThisFolder.CleanParentAndRoot()
	}
}

func (f *Fav) SetFavTypeFavIdx(startFavIdx int) (newFavIdx int) {
	newFavIdx = startFavIdx
	for idx, each := range f.Favh {
		each.FavIdx = startFavIdx + idx
	}
	newFavIdx += len(f.Favh)

	for _, each := range f.Favh {
		if each.TheType != pttbbsfav.FAVT_FOLDER {
			continue
		}

		folder := each.CastFolder()
		newFavIdx = folder.ThisFolder.SetFavTypeFavIdx(newFavIdx)
	}

	return newFavIdx
}
