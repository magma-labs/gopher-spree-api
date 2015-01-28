package repositories

import (
  "os"
  "reflect"
  "testing"
)

func TestAssetRepo(t *testing.T) {
  os.Setenv(dbUrlEnvName, "dbname=spree_dev sslmode=disable")
  os.Setenv(dbEngineEnvName, "postgres")

  err := InitDB()

  if err != nil {
    t.Error("An error has ocurred", err)
  }

  if spree_db == nil {
    t.Error("Database helper not initialized")
  }

  defer spree_db.Close()

  assetRepo := NewAssetRepo()

  assets := assetRepo.FindByViewableIds([]int64{11})

  na := len(assets)

  if na < 1 {
    t.Error("Invalid number of assets: %d", na)
  }

  temp := reflect.ValueOf(*assets[0]).Type().String()

  if temp != "models.Asset" {
    t.Error("Invalid type", t)
  }
}
