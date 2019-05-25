package core

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	// 初始化A0IManager
	aoiMgr := NewAOIManager(0, 250, 5, 0, 25, 5)

	// 打印AIOManaget
	fmt.Println(aoiMgr)
}

func TestGetSurroundGridsByGid(t *testing.T) {
	// 初始化A0IManager
	aoiMgr := NewAOIManager(0, 250, 5, 0, 25, 5)

	for gid, _ := range aoiMgr.grids {
		// 得到当前gid的周边九宫格信息
		grids := aoiMgr.GetSurroundGridsByGid(gid)
		fmt.Println("Gid: ", gid, ", gids len = ", len(grids))
		gIDs := make([]int, 0, len(grids))
		for _, grid := range grids {
			gIDs = append(gIDs, grid.GID)
		}
		fmt.Println("Surounding grid IDs are: ", gIDs)
	}
}
