package main
import (
  "errors"
  "fmt"
  "sync"
  )

var VehicleType int
//vehicle type enum::
const (
  BIKE VehicleType = iota
  CAR
  TRUCK
)

type Floor struct{
  FloorId          int
  TotalSpaces      int
  AvailableSpace   int
}

type Space struct{
  SpaceId          int
  FloorId          int
  IsAvailable      bool
  VehicleType      VehicleType
}

type BookingInformation struct{
  BookingId        int
  UserId           string
  SpaceId          int
  VehicleType      VehicleType
}

type ParkingLotImplement struct{
   Floor    map([int]*Floor)
   Space    map([int]*Space)
   Booking  map([int]*BookingInformation
   NextSpace  int
   NextBooking  int
   mu       sync.Mutex

}


type ParkingLot interface{
   AddMoreFloor(floor Floor) (int, int)
  BookSpace(userId string, vehicleType Vehicle) (Space, error)
  GetBookedInformation(bookingId int)(BookingInformation, error)
  GetAvailableSpace(vehicle VehicleType) []Space
}

                
type (p ParkingLot) NewParking(){
  Floor:    make(map[int]*Floor)
  Space:    make(map[int]*Space)
  Booking:  make(map[int]*BookingInformation)
}                


func (p ParkingLotImplement) AddMoreFloor(Floor floor) (int, int){
    p.mu.Lock()
    defer p.mu.Unlock()

  p.Floors[floor.FloorId]=&floor

  for i:=0;i<floor.TotalsSpace(floor Floor)
}
}
                
