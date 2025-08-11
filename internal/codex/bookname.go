package codex

//go:generate enumer -type BookName -trimprefix BookName
type BookName byte

const (
	BookNameUnspecified BookName = iota
	BookNameGen
	BookNameExod
	BookNameLev
	BookNameNum
	BookNameDeut
	BookNameJosh
	BookNameJudg
	BookNameRuth
	BookName1Sam
	BookName2Sam
	BookName1Kgs
	BookName2Kgs
	BookName1Chr
	BookName2Chr
	BookNameEzra
	BookNameNeh
	BookNameEsth
	BookNameJob
	BookNamePs
	BookNameProv
	BookNameEccl
	BookNameSong
	BookNameIsa
	BookNameJer
	BookNameLam
	BookNameEzek
	BookNameDan
	BookNameHos
	BookNameJoel
	BookNameAmos
	BookNameObad
	BookNameJonah
	BookNameMic
	BookNameNah
	BookNameHab
	BookNameZeph
	BookNameHag
	BookNameZech
	BookNameMal
	BookNameMatt
	BookNameMark
	BookNameLuke
	BookNameJohn
	BookNameActs
	BookNameRom
	BookName1Cor
	BookName2Cor
	BookNameGal
	BookNameEph
	BookNamePhil
	BookNameCol
	BookName1Thess
	BookName2Thess
	BookName1Tim
	BookName2Tim
	BookNameTitus
	BookNamePhlm
	BookNameHeb
	BookNameJas
	BookName1Pet
	BookName2Pet
	BookName1John
	BookName2John
	BookName3John
	BookNameJude
	BookNameRev
)
