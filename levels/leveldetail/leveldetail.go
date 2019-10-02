package leveldetail

// Tile structure
type Tile struct {
	ID        int
	Name      string
	Solid     TileSolidity
	Character string
}

func GetTile(i int) Tile {

	switch int(i) {

	case 0:
		return tileAir()
	case 1:
		return tileWall()

	default:
		return tileError()
	}
}

func tileError() Tile {
	return Tile{-1, "ERR", TileWalkable, "X"}
}

func tileAir() Tile {
	return Tile{0, "AIR", TileWalkable, " "}
}

func tileWall() Tile {
	return Tile{1, "WALL", TileSolid, "="}
}
